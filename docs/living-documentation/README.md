# GCE Persistent Disk CSI Driver - Living Documentation

This directory contains the Living Documentation for the GCE Persistent Disk CSI Driver. These documents are intended to be kept synchronized with the codebase through an automated LLM-based pipeline.

## Key Documents

*   [Deployment Architecture](./architecture.md)
*   [Supported Features](./features.md)
*   [LLM Update Instructions](../llm-prompts/llm_instructions.md)
*   [LLM Bootstrap Instructions](../llm-prompts/llm_instructions_bootstrap.md)

## Driver Capabilities

Based on the CSI specification, this driver implements the following services:

### Controller Service Capabilities

*   `CREATE_DELETE_VOLUME`: [ControllerCreateVolume](./controller_create_volume.md), [ControllerDeleteVolume](./controller_delete_volume.md)
*   `PUBLISH_UNPUBLISH_VOLUME`: [ControllerPublishVolume](./controller_publish_volume.md), [ControllerUnpublishVolume](./controller_unpublish_volume.md)
*   `LIST_VOLUMES`: [ControllerListVolumes](./controller_list_volumes.md)
*   `GET_CAPACITY`: [ControllerGetCapacity](./controller_get_capacity.md)
*   `CREATE_DELETE_SNAPSHOT`: [ControllerCreateSnapshot](./controller_create_snapshot.md), [ControllerDeleteSnapshot](./controller_delete_snapshot.md)
*   `LIST_SNAPSHOTS`: [ControllerListSnapshots](./controller_list_snapshots.md)
*   `CLONE_VOLUME`: Handled within [ControllerCreateVolume](./controller_create_volume.md)
*   `EXPAND_VOLUME`: [ControllerExpandVolume](./controller_expand_volume.md)
*   `GET_VOLUME`: [ControllerGetVolume](./controller_get_volume.md)

### Node Service Capabilities

*   `STAGE_UNSTAGE_VOLUME`: [NodeStageVolume](./node_stage_volume.md), [NodeUnstageVolume](./node_unstage_volume.md)
*   `PUBLISH_UNPUBLISH_VOLUME`: [NodePublishVolume](./node_publish_volume.md), [NodeUnpublishVolume](./node_unpublish_volume.md)
*   `GET_VOLUME_STATS`: [NodeGetVolumeStats](./node_get_volume_stats.md)
*   `EXPAND_VOLUME`: [NodeExpandVolume](./node_expand_volume.md)
*   `GET_INFO`: [NodeGetInfo](./node_get_info.md)
*   `GET_CAPACITY`: [NodeGetCapacity](./node_get_capacity.md)

### Identity Service Capabilities

*   [GetPluginInfo](./identity_get_plugin_info.md)
*   [GetPluginCapabilities](./identity_get_plugin_capabilities.md)
*   [Probe](./identity_probe.md)

> [!NOTE]
> The existence of a link above indicates that a living documentation page for that specific CSI operation is intended. Missing files should be generated.
