---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/identity.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/identity.go)

# CSI IdentityGetPluginInfo

## RPC Definition

```protobuf
rpc GetPluginInfo (GetPluginInfoRequest) returns (GetPluginInfoResponse) {}
```

## Purpose

This operation is called by the CO to get identifying information about the plugin.

*   **Trigger:** On CSI plugin registration.
*   **Action:** Returns the plugin's name and version.

## Parameters

None

## Key Logic Flow

1.  Returns a static response containing the driver name and version.

## Return Values

*   `GetPluginInfoResponse`:
    *   `name`: The name of the driver (e.g., `pd.csi.storage.gke.io`).
    *   `vendor_version`: The version of the driver.

---

[← README.md](./README.md)
