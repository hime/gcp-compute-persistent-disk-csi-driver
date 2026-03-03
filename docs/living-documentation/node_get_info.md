---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/node.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go)

# CSI NodeGetInfo

## RPC Definition

```protobuf
rpc NodeGetInfo (NodeGetInfoRequest) returns (NodeGetInfoResponse) {}
```

## Purpose

This operation is called by the Kubelet to get information about the node that the CSI driver is running on.

*   **Trigger:** Kubelet on startup or reconciliation.
*   **Action:** Returns the node ID, max volumes per node, and accessible topology.

## Parameters

None

## Key Logic Flow

1.  **Get Node ID:** Constructs a unique node ID using project, zone, and instance name from the metadata service (`common.CreateNodeID`).
2.  **Get Topology:** Fetches the zone from the metadata service. If dynamic volume limits are enabled, it also fetches disk type labels from the Kubernetes Node object.
3.  **Get Volume Limits:** Determines the maximum number of volumes that can be attached to the node based on the machine type using `getVolumeLimits`. This includes checking for node label overrides.
4.  **Return Response:** Populates `NodeGetInfoResponse` with the Node ID, MaxVolumesPerNode, and AccessibleTopology.

```mermaid
graph TD
    A[Kubelet] --> B[CSI Plugin: NodeGetInfo];
    B --> C{Get Project, Zone, Name from Metadata};
    C --> D[Construct Node ID];
    B --> E{Get Node Object from K8s API};
    E --> F{Get Topology Labels (Zone, Disk Types)};
    B --> G{Get Machine Type from Metadata};
    G --> H[Determine Volume Limits];
    D --> I[Build Response];
    F --> I;
    H --> I;
    I --> Z[Return NodeGetInfoResponse];
```

## Error Handling

*   Errors from fetching the Node object or getting volume limits are logged but generally not returned, to ensure the driver can register.

## Return Values

*   `NodeGetInfoResponse`:
    *   `node_id`: Unique identifier for the node.
    *   `max_volumes_per_node`: Maximum number of volumes that can be attached.
    *   `accessible_topology`: Topology segments, including the zone and potentially disk type availability.

---

[← README.md](./README.md)
