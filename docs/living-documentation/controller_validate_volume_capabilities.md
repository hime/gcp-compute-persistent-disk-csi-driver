---
reviewed_commit: DUMMY_COMMIT_HASH
last_updated: DUMMY_TIMESTAMP
---
[Sourced from: pkg/gce-pd-csi-driver/controller.go](file:///usr/local/google/home/jaimebz/oss/gcp-compute-persistent-disk-csi-driver/pkg/gce-pd-csi-driver/controller.go)

# CSI ControllerValidateVolumeCapabilities

## RPC Definition

```protobuf
rpc ValidateVolumeCapabilities (ValidateVolumeCapabilitiesRequest) returns (ValidateVolumeCapabilitiesResponse) {}
```

## Purpose

This operation is called by Kubernetes to check if a given volume, already provisioned, supports a specific set of capabilities. This allows Kubernetes to ensure that a volume can meet the requirements of a Pod before binding it.

*   **Trigger:** Kubernetes needs to ascertain if an existing Persistent Volume (PV) can satisfy the access mode and other capability requirements of a PersistentVolumeClaim (PVC).
*   **Action:** The driver fetches the specified Persistent Disk (PD) from the GCE API and checks if its properties and the provided parameters are compatible with the requested capabilities.
*   **Kubernetes Outcome:** Kubernetes uses the response to determine if the PV can be bound to the PVC.

## Parameters

*   `volume_id`: The unique identifier of the volume (PD). (Required)
*   `volume_capabilities`: A list of capabilities to check against the volume. (Required)
*   `parameters`: StorageClass parameters, used to validate consistency with the existing disk.
*   `volume_context`: Context from the PV object.

## Key Logic Flow

1.  **Validate Arguments:** Ensures `volume_id` and `volume_capabilities` are provided.
2.  **Parse Volume ID:** Extracts project and volume key from `volume_id`.
3.  **Repair Volume Key:** Ensures the volume key is fully specified (resolves zone/region if necessary).
4.  **Get Disk:** Fetches the PD details from the GCE API using the volume key.
5.  **Capability Check (`validateVolumeCapabilities`):** Checks if the requested capabilities (e.g., access modes like `SINGLE_NODE_WRITER`, `READ_ONLY_MANY`) are generally supported by the driver for PDs.
6.  **Parameter Consistency Check:** Extracts parameters from the request and validates them against the properties of the fetched PD using `gce.ValidateDiskParameters`.
7.  **Secrets Check:** Ensures no secrets are passed, as they are not used for this check.
8.  **Return Response:**
    *   If all checks pass, returns a `Confirmed` response, echoing back the validated capabilities and parameters.
    *   If any check fails, returns a response with a `Message` explaining the incompatibility.

```mermaid
graph TD
    A[K8s Check] --> B[Controller.ValidateVolumeCapabilities];
    B --> C{Validate Request (volume_id, capabilities)};
    C --> D{Get Disk from GCE API};
    D -- Not Found --> X[Return NotFound Error];
    D -- Found --> E{Check General Capability Support};
    E -- Not Supported --> Y[Return Message: Invalid Capabilities];
    E -- Supported --> F{Check Parameter Consistency with Disk};
    F -- Inconsistent --> Z[Return Message: Parameter Mismatch];
    F -- Consistent --> W[Return Confirmed];
```

## Error Handling

*   `InvalidArgument`: Missing or malformed `volume_id` or `volume_capabilities`.
*   `NotFound`: If the disk specified by `volume_id` doesn't exist in GCE.
*   Returns a response with a `Message` for capability or parameter mismatches, rather than an RPC error, as per CSI spec.

## Return Values

*   `ValidateVolumeCapabilitiesResponse`:
    *   `Confirmed`: Set if all requested capabilities are supported and parameters match.
    *   `Message`: Set if validation fails, providing a human-readable reason.

---

[← README.md](./README.md)
