---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/controller.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go)

# CSI ControllerCreateVolume

## RPC Definition

```protobuf
rpc CreateVolume (CreateVolumeRequest) returns (CreateVolumeResponse) {}
```

## Purpose

This operation is called by the CSI external-provisioner sidecar in response to the creation of a `PersistentVolumeClaim` (PVC) in Kubernetes that requests a StorageClass backed by this CSI driver. The driver's primary responsibility is to dynamically provision a new Google Cloud Persistent Disk (PD) volume according to the specifications in the PVC and StorageClass, and then create a `PersistentVolume` (PV) object in Kubernetes to represent the PD.

*   **Trigger:** Creation of a `PersistentVolumeClaim` object.
*   **Action:** Calls the GCE API to create a new Persistent Disk. This can be a zonal or regional disk, or multiple zonal disks if multi-zone provisioning is enabled.
*   **Kubernetes Outcome:** A `PersistentVolume` object is created and bound to the triggering `PersistentVolumeClaim`.

## Parameters

The behavior of `CreateVolume` is primarily controlled by parameters passed in the `StorageClass.parameters` field in the `CreateVolumeRequest`.

Key parameters include:

*   `type`: Specifies the PD type (e.g., `pd-standard`, `pd-ssd`, `hyperdisk-extreme`).
*   `replication-type`: Set to `regional-pd` for regional disks.
*   `zones`: For regional disks, specifies the replica zones.
*   `disk-encryption-kms-key`: KMS key for encryption.
*   `enable-multi-zone-provisioning`: Enables creating separate zonal disks for each zone in the topology.
*   Storage Pool parameters (e.g., `storage-pool`).
*   Provisioned IOPS/Throughput for Hyperdisk volumes.

## Key Logic Flow

1.  **Parameter Extraction & Validation:** Extracts and validates parameters from the request, including disk type, size, topology, and features like storage pools or multi-zone.
2.  **Capacity Validation:** Ensures the requested capacity is within GCE limits.
3.  **Capability Validation:** Checks if the requested `VolumeCapabilities` (e.g., access modes) are supported.
4.  **Mutable Parameters:** Applies any `mutable_parameters` for dynamic provisioning features on supported disk types (e.g., IOPS, Throughput).
5.  **Content Source Handling:** If `VolumeContentSource` is provided (e.g., for cloning or restoring from a snapshot), validates the source.
6.  **Topology Processing:** Determines the target zone(s) based on `AccessibilityRequirements` (preferred/requisite topologies) and StorageClass parameters.
7.  **Multi-Zone Provisioning:** If enabled, it creates individual zonal disks in each target zone (`createMultiZoneDisk`).
8.  **Single Device Provisioning:** Otherwise, it creates a single zonal or regional disk (`createSingleDeviceDisk`).
    *   Checks if a disk with the same name already exists and is compatible.
    *   Calls the GCE API to create the disk (`InsertDisk`). This handles creation from scratch, from a snapshot, or from another volume (cloning).
9.  **Disk Readiness Check:** Ensures the newly created disk is in a `READY` state.
10. **Response Generation:** Returns a `CreateVolumeResponse` containing the `VolumeId` (GCE disk selfLink) and other details.

```mermaid
graph TD
    A[PVC Creation] --> B{CSI external-provisioner};
    B --> C[Controller.CreateVolume];
    C --> D{Extract & Validate Parameters};
    D --> E{Validate Capabilities & Capacity};
    E --> F{Process Topology / Zones};
    F --> G{Multi-Zone?};
    G -- Yes --> H[Create Multi-Zone Disks];
    G -- No --> I[Create Single Zonal/Regional Disk];
    H --> J{GCE API Calls};
    I --> J;
    J --> K{Check Disk(s) READY};
    K --> L[Create PV Object];
    L --> M[Bind PV to PVC];
    C --> Z[Return CreateVolumeResponse];
```

## Supported Features

*   Dynamic Provisioning of various PD types.
*   Regional PDs.
*   Provisioning from Snapshots.
*   Volume Cloning (provisioning from an existing PV).
*   Customer Managed Encryption Keys (CMEK).
*   Storage Pools.
*   Multi-Zone Provisioning.
*   Hyperdisk (Extreme, Throughput, Balanced, ML) with provisioned IOPS/Throughput.

## Error Handling

*   Returns `InvalidArgument` for invalid parameters, capabilities, or topology.
*   Returns `AlreadyExists` if a compatible disk with the same name already exists.
*   Returns `NotFound` if a source snapshot or volume for cloning doesn't exist.
*   Propagates errors from the GCE API calls.
*   Uses a locking mechanism (`volumeLocks`) to prevent concurrent operations on the same volume ID.

## Return Values

*   `CreateVolumeResponse`: Contains the details of the provisioned volume, including:
    *   `VolumeId`: The unique identifier for the GCE disk (e.g., `projects/my-project/zones/us-central1-a/disks/my-disk`).
    *   `CapacityBytes`: Actual capacity of the created disk.
    *   `AccessibleTopology`: The zones where the disk is accessible.
    *   `VolumeContext`: Additional context for the volume.

---

[← README.md](./README.md)
