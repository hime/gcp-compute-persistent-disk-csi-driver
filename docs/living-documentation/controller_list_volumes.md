---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/controller.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go)

# CSI ControllerListVolumes

## RPC Definition

```protobuf
rpc ListVolumes (ListVolumesRequest) returns (ListVolumesResponse) {}
```

## Purpose

This operation allows Kubernetes to get a list of all Persistent Disk (PD) volumes managed by the driver within the accessible scope (typically the project). It's used for various purposes, including PV garbage collection and dynamic provisioning checks.

*   **Trigger:** Called periodically by Kubernetes and potentially by administrative tools.
*   **Action:** Fetches a list of disks from the GCE API. It can also optionally list instances to determine which nodes each disk is attached to.
*   **Kubernetes Outcome:** Provides a list of volumes and their attachment status.

## Parameters

*   `max_entries`: The maximum number of volume entries to return in a single response. Pagination is supported.
*   `starting_token`: A token provided by a previous `ListVolumesResponse` to get the next page of results.

## Key Logic Flow

1.  **Validate Request:** Checks if `max_entries` is non-negative.
2.  **Handle Pagination:** If `starting_token` is provided, finds the offset from the cached results.
3.  **Fetch Volumes (`listVolumeEntries`):** (Only if no `starting_token` or cache is empty)
    *   Calls GCE API `ListDisks` to get all disks in the project.
    *   Optionally, calls GCE API `ListInstances` to get all VM instances.
    *   Processes results using `disksAndInstancesToVolumeEntries`:
        *   Iterates through disks, creating a map of `volumeId` to the list of node IDs it's attached to (from `disk.Users`).
        *   Handles Multi-Zone volumes by creating an additional entry with the multi-zone volume ID format.
        *   If instances are listed, it augments the attachment information by checking `instance.Disks`.
        *   Compiles a list of `ListVolumesResponse_Entry` objects.
    *   Caches the full list of entries and initializes the pagination token map.
4.  ** paginate Results:** Extracts the appropriate slice of volume entries based on `max_entries` and the calculated offset.
5.  **Generate Next Token:** If there are more results, generates a new unique `starting_token` for the next call.
6.  **Return Response:** Returns `ListVolumesResponse` with the current page of entries and the `next_token`.

```mermaid
graph TD
    A[K8s or Admin] --> B[Controller.ListVolumes];
    B --> C{Validate max_entries};
    C --> D{Starting Token?};
    D -- Yes --> E{Lookup Offset in Cache};
    E -- Not Found --> X[Return Error];
    D -- No --> F[listVolumeEntries];
    F --> G[GCE API: ListDisks];
    G --> H[GCE API: ListInstances (Optional)];
    H --> I[disksAndInstancesToVolumeEntries];
    I --> J{Map volumeId to nodeIds};
    J --> K{Cache Results & Init Tokens};
    E --> L{Calculate Page Slice};
    K --> L;
    L --> M{Generate next_token?};
    M --> Z[Return ListVolumesResponse];
```

## Return Values

*   `ListVolumesResponse`:
    *   `entries`: A list of volumes, each including:
        *   `Volume`: Contains the `volume_id`.
        *   `Status`: Contains `published_node_ids`.
    *   `next_token`: Token for fetching the next page of results, if any.

---

[← README.md](./README.md)
