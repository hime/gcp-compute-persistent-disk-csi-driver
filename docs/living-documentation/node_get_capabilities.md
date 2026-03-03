---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/node.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/node.go)

# CSI NodeGetCapabilities

## RPC Definition

```protobuf
rpc NodeGetCapabilities (NodeGetCapabilitiesRequest) returns (NodeGetCapabilitiesResponse) {}
```

## Purpose

This operation is called by the Kubelet to get the supported capabilities of the node service provided by the CSI driver.

*   **Trigger:** Kubelet on startup or reconciliation.
*   **Action:** Returns a predefined list of capabilities supported by the node plugin.

## Parameters

None

## Key Logic Flow

1.  Returns a static list of `csi.NodeServiceCapability` messages stored in `ns.Driver.nscap`.

## Return Values

*   `NodeGetCapabilitiesResponse`: Contains a list of supported capabilities. Examples include:
    *   `STAGE_UNSTAGE_VOLUME`
    *   `EXPAND_VOLUME`
    *   `GET_VOLUME_STATS`
    *   `SINGLE_NODE_MULTI_WRITER`

---

[← README.md](./README.md)
