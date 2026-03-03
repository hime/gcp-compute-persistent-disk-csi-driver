# CreateVolume Features and Supported Modes

This document details the various features and modes supported by the `CreateVolume` operation in the GCE PD CSI Driver, primarily within the `createVolumeInternal` function.

[Sourced from: pkg/gce-pd-csi-driver/controller.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go)

## Key Features:

1.  **Standard Provisioning:** Creating a new, empty Persistent Disk.
2.  **Volume Cloning:** Creating a new volume from an existing PVC (VolumeContentSource: Volume).
3.  **Snapshot Restore:** Creating a new volume from a Snapshot (VolumeContentSource: Snapshot).
4.  **Topology-Aware Provisioning:** Honoring zone and region constraints specified in `AccessibilityRequirements`.
5.  **Multi-Zone (Regional Disks):** Support for creating and validating Regional Persistent Disks.
6.  **Disk Type Selection:** Defaulting and allowing specification of various PD disk types (pd-standard, pd-ssd, pd-balanced, hyperdisk-*, etc.).
7.  **Dynamic Disk Type Selection:** Automatically choosing a disk type based on requirements when `enableDynamicVolumes` is active.
8.  **Mutable Parameters:** Allowing IOPS and Throughput to be specified and potentially modified for supported disk types (e.g., hyperdisk).
9.  **Capacity Specification:** Handling capacity requests from `CapacityRange`.
10. **Access Modes:** Validating and supporting different access modes (RWO, RWX, ROX).
11. **Storage Pools:** Validating and supporting the use of Storage Pools.

## Validation Checks:

*   Request Parameter Validation
*   Capacity Range Checks
*   Volume Capability Validation (Access Modes, Mount Options)
*   Mutable Parameter Compatibility with Disk Type
*   Multi-Writer Capability
*   Read-Only Mode Requirements
*   Multi-Attachment Restrictions for certain Disk Types
*   Storage Pool Compatibility
*   Topology Constraint Feasibility
*   Multi-Zone Configuration Validity

## Conditional Flows:

*   The provisioning path differs significantly based on whether `VolumeContentSource` is nil, a Volume, or a Snapshot.
*   Zone selection logic depends on topology constraints and whether the disk is zonal or regional.
*   Error handling for various GCE API call failures.
