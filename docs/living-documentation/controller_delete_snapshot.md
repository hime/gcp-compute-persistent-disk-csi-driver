---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/controller.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go)

# CSI ControllerDeleteSnapshot

## RPC Definition

```protobuf
rpc DeleteSnapshot (DeleteSnapshotRequest) returns (DeleteSnapshotResponse) {}
```

## Purpose

This operation is called by the CSI external-snapshotter sidecar to delete a snapshot of a Persistent Disk (PD) volume.

*   **Trigger:** Deletion of a `VolumeSnapshotContent` object in Kubernetes.
*   **Action:** Calls the GCE API to delete the GCE Snapshot or GCE Image corresponding to the given `snapshot_id`.

## Parameters

*   `snapshot_id`: The unique identifier of the snapshot to be deleted. (Required)

## Key Logic Flow

1.  **Validate Arguments:** Checks if the `snapshot_id` is provided.
2.  **Parse Snapshot ID:** Parses the `snapshot_id` to determine the project, backup type (Snapshot or Image), and resource key using `common.SnapshotIDToProjectKey`.
3.  **Handle Invalid ID:** If the `snapshot_id` is malformed, logs a warning and returns success as per the CSI spec (idempotency).
4.  **Dispatch Deletion:**
    *   If the type is `DiskSnapshotType`, calls GCE API `DeleteSnapshot`.
    *   If the type is `DiskImageType`, calls GCE API `DeleteImage`.
5.  **Return Response:** Returns an empty `DeleteSnapshotResponse` on successful deletion or if the snapshot was not found or the ID was invalid.

```mermaid
graph TD
    A[VolumeSnapshotContent Deleted] --> B{CSI external-snapshotter};
    B --> C[Controller.DeleteSnapshot];
    C --> D{Validate snapshot_id};
    D -- Invalid --> Z[Return Success];
    D -- Valid --> E{Parse snapshot_id (Type & Key)};
    E --> F{Snapshot or Image?};
    F -- Snapshot --> G[GCE API: DeleteSnapshot];
    F -- Image --> H[GCE API: DeleteImage];
    G --> Y{Check Error};
    H --> Y;
    Y -- Error --> X[Return Error];
    Y -- No Error / Not Found --> Z;
    C --> Z[Return DeleteSnapshotResponse];
```

## Error Handling

*   `InvalidArgument`: Missing `snapshot_id` or unknown snapshot type from ID.
*   Returns success if the `snapshot_id` is malformed (idempotency).
*   Propagates GCE API errors, except for `NotFound` which is treated as success.

## Return Values

*   `DeleteSnapshotResponse`: An empty response indicating success.

---

[← README.md](./README.md)
