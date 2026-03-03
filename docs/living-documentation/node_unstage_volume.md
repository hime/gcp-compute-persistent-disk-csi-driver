---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/node.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go)

# CSI NodeUnstageVolume

## RPC Definition

```protobuf
rpc NodeUnstageVolume (NodeUnstageVolumeRequest) returns (NodeUnstageVolumeResponse) {}
```

## Purpose

This operation is called by the Kubelet to unmount a volume from its staging path on the node. This reverses the actions of `NodeStageVolume`.

*   **Trigger:** When the last Pod using the volume on the node is terminated.
*   **Action:** Unmounts the volume from the `staging_target_path` and cleans up any data cache.

## Parameters

*   `volume_id`: The ID of the volume to unstage. (Required)
*   `staging_target_path`: The directory on the node where the volume is staged. (Required)

## Key Logic Flow

1.  **Validate Arguments:** Checks for `volume_id` and `staging_target_path`.
2.  **Acquire Lock:** Locks the `volume_id` to prevent concurrent operations.
3.  **Cleanup Stage Path:** Unmounts the volume from the `staging_target_path` using `cleanupStagePath`.
4.  **Device In Use Check:** If `enableDeviceInUseCheck` is true, verifies the underlying block device is no longer in use using `confirmDeviceUnused`. Handles errors and timeouts.
5.  **Data Cache Cleanup:** If data caching is enabled, cleans up the cache setup for the volume using `cleanupCache`.
6.  **Device Cache Removal:** Removes the volume from the in-memory device path cache.
7.  **Return Response:** Returns an empty `NodeUnstageVolumeResponse` on success.

```mermaid
graph TD
    A[Kubelet: Last Pod Deleted] --> B[CSI Plugin: NodeUnstageVolume];
    B --> C{Validate Args};
    C --> D[Acquire Lock];
    D --> E[Unmount Staging Path];
    E --> F{Device In Use Check?};
    F -- Yes --> G[Confirm Device Unused];
    G -- In Use --> X[Return Error / Timeout];
    G -- Not In Use / Timeout / Ignore --> H{Data Cache Cleanup?};
    F -- No --> H;
    H -- Yes --> I[Cleanup Cache];
    H -- No --> J[Remove from Device Cache];
    I --> J;
    J --> Z[Return Success];
    C -- Fail --> X;
    E -- Fail --> X;
    I -- Fail --> X;
```

## Error Handling

*   `InvalidArgument`: Missing or invalid arguments.
*   `Aborted`: Volume lock could not be acquired.
*   `Internal`: Unmounting failed, device still in use after timeout.
*   `DataLoss`: Data cache cleanup errors.

## Return Values

*   `NodeUnstageVolumeResponse`: An empty response indicating success.

---

[← README.md](./README.md)
