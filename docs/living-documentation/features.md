# GCE PD CSI Driver - Supported Features

This document lists the key features supported by the driver, gathered from public documentation. This page is intended to be a reference for other living documentation pages.

## Implemented CSI Operations

This section lists the Container Storage Interface (CSI) operations implemented by the driver, linking to their in-depth documentation and source code.

### Identity Service

*   **GetPluginInfo:** [Docs](./identity_get_plugin_info.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/identity.go#L34)
*   **GetPluginCapabilities:** [Docs](./identity_get_plugin_capabilities.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/identity.go#L45)
*   **Probe:** [Docs](./identity_probe.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/identity.go#L80)

### Controller Service

*   **CreateVolume:** [Docs](./controller_create_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L314)
*   **DeleteVolume:** [Docs](./controller_delete_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L911)
*   **ControllerPublishVolume:** [Docs](./controller_publish_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1011)
*   **ControllerUnpublishVolume:** [Docs](./controller_unpublish_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1249)
*   **ValidateVolumeCapabilities:** (Part of other operations) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1338)
*   **ListVolumes:** [Docs](./controller_list_volumes.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1466)
*   **GetCapacity:** [Docs](./controller_get_capacity.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1603)
*   **ControllerGetVolume:** [Docs](./controller_get_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1440)
*   **CreateSnapshot:** [Docs](./controller_create_snapshot.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1616)
*   **DeleteSnapshot:** [Docs](./controller_delete_snapshot.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1880)
*   **ListSnapshots:** [Docs](./controller_list_snapshots.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1914)
*   **ControllerExpandVolume:** [Docs](./controller_expand_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go#L1965)

### Node Service

*   **NodeStageVolume:** [Docs](./node_stage_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go#L358)
*   **NodeUnstageVolume:** [Docs](./node_unstage_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go#L659)
*   **NodePublishVolume:** [Docs](./node_publish_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go#L202)
*   **NodeUnpublishVolume:** [Docs](./node_unpublish_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go#L335)
*   **NodeGetVolumeStats:** [Docs](./node_get_volume_stats.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go#L779)
*   **NodeExpandVolume:** [Docs](./node_expand_volume.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go#L836)
*   **NodeGetInfo:** [Docs](./node_get_info.md) | [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go#L741)
*   **NodeGetCapabilities:** [Code](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go#L720)
*   **NodeGetCapacity:** Not Implemented

## Core Features

*   **Dynamic Provisioning:** Volumes are dynamically provisioned using Kubernetes `StorageClass` objects. Different disk types and parameters can be specified within the `StorageClass`.
*   **CSI Specification:** Implements the Container Storage Interface (CSI) spec for volume lifecycle management.

## Disk Types & Performance

*   **Supported Disk Types:**
    *   Standard Persistent Disk (`pd-standard`)
    *   Balanced Persistent Disk (`pd-balanced`)
    *   SSD Persistent Disk (`pd-ssd` / `premium-rwo` StorageClass)
    *   Hyperdisk Balanced
    *   Hyperdisk Extreme
    *   Hyperdisk Throughput
*   **Regional Persistent Disks:** Provides high availability by replicating data across two zones in a region. Configured via `replication-type: regional-pd` in the `StorageClass` parameters and constrained using `allowedTopologies`.
*   **Dynamic IOPS/Throughput Modification:** For Hyperdisk volumes, IOPS and throughput can be dynamically modified using the `VolumeAttributesClass` CRD (Requires GKE 1.34+).

## Data Management

*   **Volume Snapshots:** Supports creation and restoration from Volume Snapshots.
*   **Volume Cloning:** Supports creating new volumes by cloning existing PVCs.
*   **Pre-existing Disks:** Supports using existing GCP Persistent Disks. This can be done by manually creating `PersistentVolume` and `PersistentVolumeClaim` objects, or by using `StatefulSet` volume claim templates with pre-bound PVs.

## Access Modes

*   **ReadWriteOnce (RWO):** The default, volume can be mounted as read-write by a single node.
*   **ReadOnlyMany (ROX):** Volumes can be mounted read-only by multiple nodes. This is typically achieved by:
    *   Restoring from a Volume Snapshot into a ROX PVC.
    *   Cloning a volume into a ROX PVC.
    *   Using a pre-existing disk, ensuring the data is populated, and setting `readOnly: true` in the `spec.csi` section of the `PersistentVolume`.

## Filesystems

*   **Default:** `ext4` for Linux.
*   **Other Supported:** `xfs` is supported on Linux node images that include it.
*   **Windows:** `NTFS` is required for Windows Server node pools.
    *   Filesystem type can be specified in the `StorageClass` or in the `spec.csi.fsType` field of a manual `PersistentVolume`.

## Security

*   **Customer-Managed Encryption Keys (CMEK):** Supported, allowing users to control the keys encrypting their data.

## Topology & Scheduling

*   **Zone Awareness:** Used with Regional Persistent Disks to specify replication zones via `allowedTopologies` in the `StorageClass`.
*   **Attach Limits:** The driver reports per-node volume attachment limits to the Kubernetes scheduler to prevent overloading nodes. These limits depend on the machine type and the types of disks being attached (Persistent Disk vs. Hyperdisk). Limits can be overridden using the node label `node-restriction.kubernetes.io/gke-volume-attach-limit-override`.

## Other

*   **Volume Expansion:** Supports resizing PersistentVolumeClaims.
*   **Windows Support:** The driver can be used with Windows node pools.
