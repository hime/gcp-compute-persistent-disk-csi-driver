---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/identity.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/identity.go)

# CSI IdentityProbe

## RPC Definition

```protobuf
rpc Probe (ProbeRequest) returns (ProbeResponse) {}
```

## Purpose

This operation is called by the CO to check if the plugin is healthy and responsive.

*   **Trigger:** Periodically by the CO (e.g., Kubelet liveness probe).
*   **Action:** Performs a health check.

## Parameters

None

## Key Logic Flow

1.  Returns a default `ProbeResponse`, indicating readiness. Currently, this is a no-op health check.

## Return Values

*   `ProbeResponse`: An empty response, signifying the plugin is ready.

---

[← README.md](./README.md)
