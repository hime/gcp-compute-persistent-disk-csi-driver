---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/node.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go)

# CSI NodeExpandVolume

## RPC Definition

```protobuf
rpc NodeExpandVolume (NodeExpandVolumeRequest) returns (NodeExpandVolumeResponse) {}
```

## Purpose

This operation is called by the Kubelet to expand the filesystem on the node after the underlying block device has been expanded by `ControllerExpandVolume`.

*   **Trigger:** Kubelet, after `ControllerExpandVolume` has completed and indicated node expansion is required.
*   **Action:** Resizes the filesystem on the device to occupy the new space.

## Parameters

*   `volume_id`: The ID of the volume. (Required)
*   `volume_path`: The path where the volume is mounted (or the block device file exists). (Required)
*   `capacity_range`: The new desired capacity. (Optional)
*   `volume_capability`: The capability of the volume. (Optional)

## Key Logic Flow

1.  **Validate Arguments:** Checks for `volume_id` and `volume_path`.
2.  **Get Device Path:** Determines the block device path for the `volume_id`.
3.  **Handle Block/Readonly:** If the volume capability is Block or ReadOnly, this is a no-op.
4.  **Resize FS:** Calls the `resizefs` utility to expand the filesystem on the `devicePath` to fill the available space. The mount point is specified by `volume_path`.
5.  **Verify Size:** Checks if the block device size is at least the requested size.
6.  **Return Response:** Returns the new capacity.

```mermaid
graph TD
    A[Kubelet: ControllerExpandVolume Done] --> B[CSI Plugin: NodeExpandVolume];
    B --> C{Validate Args};
    C --> D[Get Device Path];
    D --> E{Block or ReadOnly?};
    E -- Yes --> Z[Return Success (No-op)];
    E -- No --> F[Resize Filesystem];
    F --> G{Verify Device Size};
    G -- OK --> Z;
    G -- Too Small --> X[Return Error];
    C -- Fail --> X;
    D -- Fail --> X;
    F -- Fail --> X;
```

## Error Handling

*   `InvalidArgument`: Missing or invalid arguments.
*   `Internal`: Errors during device path retrieval, filesystem resize, or size verification.

## Return Values

*   `NodeExpandVolumeResponse`: Contains the `capacity_bytes` reflecting the new size.

---

[← README.md](./README.md)
