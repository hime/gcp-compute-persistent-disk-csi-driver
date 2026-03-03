---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/node.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go)

# CSI NodeStageVolume

## RPC Definition

```protobuf
rpc NodeStageVolume (NodeStageVolumeRequest) returns (NodeStageVolumeResponse) {}
```

## Purpose

This operation is called by the Kubelet to prepare a volume for mounting on a specific node. This typically involves formatting the disk (if necessary) and mounting it to a global staging path on the node.

*   **Trigger:** When a Pod first consumes a PVC on a node.
*   **Action:** Locates the attached PD, formats it, and mounts it to the `staging_target_path`.

## Parameters

*   `volume_id`: The ID of the volume to stage. (Required)
*   `staging_target_path`: The directory on the node to mount the volume to. (Required)
*   `volume_capability`: Specifies the access type (Mount/Block) and mode. (Required)
*   `publish_context`: Additional information, can include expected disk size and data cache parameters.
*   `volume_context`: Additional volume attributes, can include partition number.

## Key Logic Flow

1.  **Validate Arguments:** Checks for `volume_id`, `staging_target_path`, and `volume_capability`.
2.  **Acquire Lock:** Locks the `volume_id` to prevent concurrent operations.
3.  **Get Device Path:** Determines the block device path for the given `volume_id` (e.g., `/dev/disk/by-id/google-my-disk`).
4.  **Data Cache Setup:** If data caching is enabled and configured in `publish_context`, sets up the cache layers.
5.  **Check Existing Mount:** If `staging_target_path` is already a mount point, returns success.
6.  **Prepare Staging Path:** Creates the `staging_target_path` directory if it doesn't exist.
7.  **Handle Block Volume:** If `volume_capability` is Block, this is a no-op, returns success.
8.  **Determine Mount Options:** Extracts fstype and mount flags from `volume_capability`.
9.  **Validate Disk Size:** If `constants.ContextDiskSizeGB` is in `publish_context`, compares it against the actual block device size.
10. **Format and Mount:** Calls `ns.formatAndMount` to format (if needed) and mount the `devicePath` to `stagingTargetPath` with the determined options. Includes retry logic with `noload` option for readonly ext4/ext3 mounts.
11. **Resize Filesystem:** If not a readonly mount, resizes the filesystem to match the block device size.
12. **Update Read Ahead:** Configures the block device's read_ahead_kb value if specified in mount options.
13. **BTRFS Tuning:** Applies specific sysfs settings if BTRFS is used and flags are provided.
14. **Return Response:** Returns an empty `NodeStageVolumeResponse` on success.

```mermaid
graph TD
    A[Kubelet: Pod needs Volume] --> B[CSI Plugin: NodeStageVolume];
    B --> C{Validate Args};
    C --> D{Get Device Path};
    D --> E{Data Cache?};
    E -- Yes --> F[Setup Cache];
    E -- No --> G[Check Existing Mount];
    F --> G;
    G -- Yes --> Z[Return Success];
    G -- No --> H{Prepare Staging Dir};
    H --> I{Block Volume?};
    I -- Yes --> Z;
    I -- No --> J{Get Mount Options};
    J --> K{Validate Size (Optional)};
    K --> L[Format & Mount];
    L --> M{Resize FS?};
    M -- Yes --> N[Resize FS];
    M -- No --> N;
    N --> O{Update Read Ahead?};
    O -- Yes --> P[Set Read Ahead];
    O -- No --> P;
    P --> Q{BTRFS Tuning?};
    Q -- Yes --> R[Set BTRFS Sysfs];
    Q -- No --> R;
    R --> Z;
    C -- Fail --> X[Return Error];
    D -- Fail --> X;
    F -- Fail --> X;
    H -- Fail --> X;
    L -- Fail --> X;
    N -- Fail --> X;
    P -- Fail --> X;
    R -- Fail --> X;
```

## Error Handling

*   `InvalidArgument`: Missing or invalid arguments.
*   `Aborted`: Volume lock could not be acquired.
*   `Internal`: Failures in getting device path, mkdir, formatting, mounting, resizing, or sysfs operations.
*   `DataLoss`: Data cache setup errors.

## Return Values

*   `NodeStageVolumeResponse`: An empty response indicating success.

---

[← README.md](./README.md)
