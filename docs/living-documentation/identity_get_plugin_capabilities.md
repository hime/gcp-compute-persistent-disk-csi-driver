---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/identity.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/identity.go)

# CSI IdentityGetPluginCapabilities

## RPC Definition

```protobuf
rpc GetPluginCapabilities (GetPluginCapabilitiesRequest) returns (GetPluginCapabilitiesResponse) {}
```

## Purpose

This operation is called by the CO to get the supported capabilities of the plugin.

*   **Trigger:** On CSI plugin registration/discovery.
*   **Action:** Returns a list of capabilities supported by the plugin.

## Parameters

None

## Key Logic Flow

1.  Returns a list of `csi.PluginCapability` messages, indicating the services and features the driver implements.

## Return Values

*   `GetPluginCapabilitiesResponse`: Contains a list of supported capabilities. Examples:
    *   `CONTROLLER_SERVICE`
    *   `VOLUME_ACCESSIBILITY_CONSTRAINTS`
    *   `VOLUME_EXPANSION` (Online/Offline)
    *   `CREATE_DELETE_SNAPSHOT`
    *   `LIST_SNAPSHOTS`
    *   `SINGLE_NODE_MULTI_WRITER`

---

[← README.md](./README.md)
