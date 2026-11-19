/*
Copyright 2026 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gceGCEDriver

import (
	"context"
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-cloud-provider/pkg/cloud/meta"
	csi "github.com/container-storage-interface/spec/lib/go/csi"
	compute "google.golang.org/api/compute/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/gcp-compute-persistent-disk-csi-driver/pkg/constants"
	gce "sigs.k8s.io/gcp-compute-persistent-disk-csi-driver/pkg/gce-cloud-provider/compute"
	"sigs.k8s.io/gcp-compute-persistent-disk-csi-driver/pkg/parameters"
)

func TestCreateVolumeWithVMLocationHint(t *testing.T) {
	const (
		pvcName      = "test-pvc"
		pvcNamespace = "test-ns"
		selectedNode = "gke-cluster-pool-1234-abcd"
	)

	zonalTop := &csi.TopologyRequirement{
		Requisite: []*csi.Topology{{Segments: map[string]string{constants.TopologyKeyZone: zone}}},
		Preferred: []*csi.Topology{{Segments: map[string]string{constants.TopologyKeyZone: zone}}},
	}
	regionalTop := &csi.TopologyRequirement{
		Requisite: []*csi.Topology{
			{Segments: map[string]string{constants.TopologyKeyZone: zone}},
			{Segments: map[string]string{constants.TopologyKeyZone: secondZone}},
		},
		Preferred: []*csi.Topology{
			{Segments: map[string]string{constants.TopologyKeyZone: secondZone}},
			{Segments: map[string]string{constants.TopologyKeyZone: zone}},
		},
	}

	pvcWithSelectedNode := func(nodeName string) *v1.PersistentVolumeClaim {
		annotations := map[string]string{}
		if nodeName != "" {
			annotations[annSelectedNode] = nodeName
		}
		return &v1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:        pvcName,
				Namespace:   pvcNamespace,
				Annotations: annotations,
			},
		}
	}

	pvcParams := func(extra map[string]string) map[string]string {
		p := map[string]string{
			parameters.ParameterKeyPVCName:      pvcName,
			parameters.ParameterKeyPVCNamespace: pvcNamespace,
		}
		for k, v := range extra {
			p[k] = v
		}
		return p
	}

	sourceZonalDisk := gce.CloudDiskFromV1(&compute.Disk{
		Name:     "src-disk",
		SizeGb:   20,
		Status:   "READY",
		Type:     fmt.Sprintf("projects/%s/zones/%s/diskTypes/pd-balanced", project, zone),
		Zone:     zone,
		SelfLink: fmt.Sprintf("projects/%s/zones/%s/disks/src-disk", project, zone),
	})

	testCases := []struct {
		name               string
		nilClient          bool
		k8sObjects         []runtime.Object
		seedDisks          []*gce.CloudDisk
		enableStoragePools bool
		enableMultiZone    bool
		req                *csi.CreateVolumeRequest
		wantDiskKey        *meta.Key
		wantLocationHint   string
		wantNoK8sAPICalls  bool
	}{
		{
			name:       "zonal disk sets VM location hint",
			k8sObjects: []runtime.Object{pvcWithSelectedNode(selectedNode)},
			req: &csi.CreateVolumeRequest{
				Name:                      name,
				CapacityRange:             stdCapRange,
				VolumeCapabilities:        stdVolCaps,
				Parameters:                pvcParams(map[string]string{parameters.ParameterKeyType: "pd-balanced"}),
				AccessibilityRequirements: zonalTop,
			},
			wantDiskKey:      meta.ZonalKey(name, zone),
			wantLocationHint: fmt.Sprintf("project/%s/zone/%s/instance/%s", project, zone, selectedNode),
		},
		{
			name:       "regional disk uses preferred zone for VM location hint",
			k8sObjects: []runtime.Object{pvcWithSelectedNode(selectedNode)},
			req: &csi.CreateVolumeRequest{
				Name:               name,
				CapacityRange:      stdCapRange,
				VolumeCapabilities: stdVolCaps,
				Parameters: pvcParams(map[string]string{
					parameters.ParameterKeyType:            "pd-balanced",
					parameters.ParameterKeyReplicationType: replicationTypeRegionalPD,
				}),
				AccessibilityRequirements: regionalTop,
			},
			wantDiskKey:      meta.RegionalKey(name, region),
			wantLocationHint: fmt.Sprintf("project/%s/zone/%s/instance/%s", project, secondZone, selectedNode),
		},
		{
			name:       "HDHA regional disk sets VM location hint",
			k8sObjects: []runtime.Object{pvcWithSelectedNode(selectedNode)},
			req: &csi.CreateVolumeRequest{
				Name:               name,
				CapacityRange:      stdCapRange,
				VolumeCapabilities: stdVolCaps,
				Parameters: pvcParams(map[string]string{
					parameters.ParameterKeyType: parameters.DiskTypeHdHA,
				}),
				AccessibilityRequirements: regionalTop,
			},
			wantDiskKey:      meta.RegionalKey(name, region),
			wantLocationHint: fmt.Sprintf("project/%s/zone/%s/instance/%s", project, secondZone, selectedNode),
		},
		{
			name:       "zonal to regional clone uses selected node preferred zone, not source disk zone",
			k8sObjects: []runtime.Object{pvcWithSelectedNode(selectedNode)},
			seedDisks:  []*gce.CloudDisk{sourceZonalDisk},
			req: &csi.CreateVolumeRequest{
				Name:               name,
				CapacityRange:      stdCapRange,
				VolumeCapabilities: stdVolCaps,
				Parameters: pvcParams(map[string]string{
					parameters.ParameterKeyType:            "pd-balanced",
					parameters.ParameterKeyReplicationType: replicationTypeRegionalPD,
				}),
				VolumeContentSource: &csi.VolumeContentSource{
					Type: &csi.VolumeContentSource_Volume{
						Volume: &csi.VolumeContentSource_VolumeSource{
							VolumeId: fmt.Sprintf("projects/%s/zones/%s/disks/src-disk", project, zone),
						},
					},
				},
				AccessibilityRequirements: regionalTop,
			},
			wantDiskKey:      meta.RegionalKey(name, region),
			wantLocationHint: fmt.Sprintf("project/%s/zone/%s/instance/%s", project, secondZone, selectedNode),
		},
		{
			name:       "immediate binding (no selected-node annotation) leaves location hint empty",
			k8sObjects: []runtime.Object{pvcWithSelectedNode("")},
			req: &csi.CreateVolumeRequest{
				Name:                      name,
				CapacityRange:             stdCapRange,
				VolumeCapabilities:        stdVolCaps,
				Parameters:                pvcParams(nil),
				AccessibilityRequirements: zonalTop,
			},
			wantDiskKey:      meta.ZonalKey(name, zone),
			wantLocationHint: "",
		},
		{
			name:      "feature disabled (nil kubeClient) leaves location hint empty",
			nilClient: true,
			req: &csi.CreateVolumeRequest{
				Name:                      name,
				CapacityRange:             stdCapRange,
				VolumeCapabilities:        stdVolCaps,
				Parameters:                pvcParams(nil),
				AccessibilityRequirements: zonalTop,
			},
			wantDiskKey:      meta.ZonalKey(name, zone),
			wantLocationHint: "",
		},
		{
			name:       "missing PVC metadata skips lookup and leaves location hint empty",
			k8sObjects: []runtime.Object{pvcWithSelectedNode(selectedNode)},
			req: &csi.CreateVolumeRequest{
				Name:                      name,
				CapacityRange:             stdCapRange,
				VolumeCapabilities:        stdVolCaps,
				Parameters:                map[string]string{parameters.ParameterKeyType: "pd-balanced"},
				AccessibilityRequirements: zonalTop,
			},
			wantDiskKey:       meta.ZonalKey(name, zone),
			wantLocationHint:  "",
			wantNoK8sAPICalls: true,
		},
		{
			name:       "PVC not found fails open without error",
			k8sObjects: nil,
			req: &csi.CreateVolumeRequest{
				Name:                      name,
				CapacityRange:             stdCapRange,
				VolumeCapabilities:        stdVolCaps,
				Parameters:                pvcParams(nil),
				AccessibilityRequirements: zonalTop,
			},
			wantDiskKey:      meta.ZonalKey(name, zone),
			wantLocationHint: "",
		},
		{
			name:       "malformed selected-node containing slash is ignored",
			k8sObjects: []runtime.Object{pvcWithSelectedNode("invalid/node")},
			req: &csi.CreateVolumeRequest{
				Name:                      name,
				CapacityRange:             stdCapRange,
				VolumeCapabilities:        stdVolCaps,
				Parameters:                pvcParams(nil),
				AccessibilityRequirements: zonalTop,
			},
			wantDiskKey:      meta.ZonalKey(name, zone),
			wantLocationHint: "",
		},
		{
			name:               "storage pool disk skips VM location hint and PVC lookup",
			k8sObjects:         []runtime.Object{pvcWithSelectedNode(selectedNode)},
			enableStoragePools: true,
			req: &csi.CreateVolumeRequest{
				Name:               name,
				CapacityRange:      stdCapRange,
				VolumeCapabilities: stdVolCaps,
				Parameters: pvcParams(map[string]string{
					parameters.ParameterKeyType:         parameters.DiskTypeHyperdiskBalanced,
					parameters.ParameterKeyStoragePools: fmt.Sprintf("projects/%s/zones/%s/storagePools/pool-1", project, zone),
				}),
				AccessibilityRequirements: zonalTop,
			},
			wantDiskKey:       meta.ZonalKey(name, zone),
			wantLocationHint:  "",
			wantNoK8sAPICalls: true,
		},
		{
			name:            "multi-zone disk skips VM location hint and PVC lookup",
			k8sObjects:      []runtime.Object{pvcWithSelectedNode(selectedNode)},
			enableMultiZone: true,
			req: &csi.CreateVolumeRequest{
				Name:               name,
				CapacityRange:      stdCapRange,
				VolumeCapabilities: stdVolCaps,
				Parameters: pvcParams(map[string]string{
					parameters.ParameterKeyType:                        "pd-ssd",
					parameters.ParameterKeyEnableMultiZoneProvisioning: "true",
				}),
				AccessibilityRequirements: zonalTop,
			},
			wantDiskKey:       meta.ZonalKey(name, zone),
			wantLocationHint:  "",
			wantNoK8sAPICalls: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fcp, err := gce.CreateFakeCloudProvider(project, zone, tc.seedDisks)
			if err != nil {
				t.Fatalf("Failed to create fake cloud provider: %v", err)
			}

			var fakeClient *fake.Clientset
			var kubeClient kubernetes.Interface
			if !tc.nilClient {
				fakeClient = fake.NewSimpleClientset(tc.k8sObjects...)
				kubeClient = fakeClient
			}

			gceDriver := initGCEDriverWithCloudProvider(t, fcp, &GCEControllerServerArgs{
				KubeClient: kubeClient,
			})
			gceDriver.cs.enableStoragePools = tc.enableStoragePools
			if tc.enableMultiZone {
				gceDriver.cs.multiZoneVolumeHandleConfig = MultiZoneVolumeHandleConfig{
					Enable:    true,
					DiskTypes: []string{"pd-ssd"},
				}
			}

			resp, err := gceDriver.cs.CreateVolume(context.Background(), tc.req)
			if err != nil {
				t.Fatalf("CreateVolume() returned unexpected error: %v", err)
			}
			if resp.GetVolume() == nil {
				t.Fatalf("CreateVolume() returned nil volume")
			}

			disk, err := fcp.GetDisk(context.Background(), project, tc.wantDiskKey)
			if err != nil {
				t.Fatalf("GetDisk(%v) failed: %v", tc.wantDiskKey, err)
			}
			if got := disk.GetLocationHint(); got != tc.wantLocationHint {
				t.Errorf("disk.GetLocationHint() = %q, want %q", got, tc.wantLocationHint)
			}
			if tc.wantNoK8sAPICalls && fakeClient != nil && len(fakeClient.Actions()) != 0 {
				t.Errorf("expected no k8s API calls, got %d: %v", len(fakeClient.Actions()), fakeClient.Actions())
			}
		})
	}
}
