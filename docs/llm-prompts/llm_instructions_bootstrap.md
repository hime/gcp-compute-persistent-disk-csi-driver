# LLM Instructions for Bootstrapping Living Documentation

You are an AI assistant tasked with autonomously generating the initial technical documentation for the GCE PD CSI Driver project. Your goal is to scan the project's Go source code, identify all implemented Container Storage Interface (CSI) operations, and generate a set of Markdown files in `docs/living-documentation/`, one for each operation, plus an index file.

## Contextual Files & Information:

*   **Official Documentation URLs:** Content from the following URLs will be provided in the prompt to offer context on the driver's features and capabilities. Ensure the documentation you generate is consistent with the information from these sources:
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/gce-pd-csi-driver
    *   https://docs.cloud.google.com/kubernetes-engine/docs/concepts/persistent-volume-attach-limits
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/preexisting-pd
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/install-csi-driver
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/readonlymany-disks
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/ssd-pd
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/regional-pd
    *   https://docs.cloud.google.com/kubernetes-engine/docs/concepts/hyperdisk
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/hyperdisk
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/hyperdisk-storage-pools
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/hyperdisk-ml
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/data-cache
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/custom-boot-disks
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/volume-cloning
    *   https://docs.cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/backup-pd-volume-snapshots

## General Guidelines for Each Operation File:

1.  **Output Format:** Output the *entire* new Markdown file, including the YAML frontmatter.
2.  **Filename:** The filename MUST be derived from the CSI operation name (e.g., `ControllerCreateVolume` -> `controller_create_volume.md`, `NodePublishVolume` -> `node_publish_volume.md`).
3.  **Frontmatter:** Include the following YAML frontmatter:
    *   `reviewed_commit`: Set this to the NEW_COMMIT hash provided in the prompt.
    *   `last_updated`: Set this to the CURRENT_TIMESTAMP provided in the prompt.
4.  **Sourced From:** Add a line after the frontmatter linking to the source file: `[Sourced from: <relative_path_to_source>](file:///<absolute_path_to_source>)`
5.  **Purpose:** Clearly state the purpose of the CSI operation, referencing its definition in the CSI specification. Crucially, explain this in the context of Kubernetes resources. A good description includes:
    *   **Trigger:** What Kubernetes event or object creation/modification/deletion triggers this operation (e.g., `PersistentVolumeClaim` creation, `VolumeSnapshot` deletion)?
    *   **Action:** Briefly describe the key action taken by the driver (e.g., calls GCP API to create a disk, mounts a filesystem).
    *   **Kubernetes Outcome:** What is the expected result in terms of Kubernetes objects, their status, or bindings? (e.g., `PersistentVolume` is provisioned and bound, `VolumeSnapshotContent` is created).
6.  **Clarity and Conciseness:** Describe the key logic, parameters, and return values. Use technical terms correctly.
7.  **Feature Consistency:** Ensure the described functionality aligns with the features outlined in the provided URL contents. Cross-link to relevant public documentation sections if possible.
8.  **Mermaid Diagrams:** Include ````mermaid` diagrams to illustrate logic flow, state transitions, or architecture. Ensure the diagrams are syntactically correct and renderable on GitHub.
9.  **Structure:** Organize the documentation with clear headings. Typical sections for a CSI operation might include:
    *   RPC Definition (e.g., `CreateVolume`)
    *   Purpose
    *   Parameters
    *   Key Logic Flow (with diagrams)
    *   Supported Features (referencing provided URL context)
    *   Error Handling
    *   Return Values
10. **Links:** Link to `README.md` and any other relevant documents within `docs/living-documentation/` using relative paths (e.g., `./README.md`).
11. **Mermaid Syntax:** When creating ````mermaid` diagrams, ensure all node labels containing special characters (e.g., `(`, `)`, `/`, `?`, `-`) are enclosed in double quotes (e.g., `id{"Label with (special/chars)"}`). This is crucial for correct rendering. This also applies to `subgraph` titles; enclose them in double quotes if they contain spaces or special characters (e.g., `subgraph "My Subgraph (Detail)"`).

## Input Provided in Prompt:

*   `NEW_COMMIT`: The current commit hash the documentation should reflect.
*   `CURRENT_TIMESTAMP`: The time of the generation.
*   `URL_CONTENTS`: A section containing the text content extracted from the official documentation URLs listed above.

## Task:

1.  **Code Scan:** Scan the Go source files within the `pkg/gce-pd-csi-driver/` directory of the project to identify all implemented CSI service methods (Controller, Node, Identity).
2.  **Generate Operation Docs:** For each identified CSI operation, create a new Markdown file in `docs/living-documentation/` following all the "General Guidelines for Each Operation File" above.
3.  **Generate README.md:** After generating all individual operation files, create or completely overwrite `docs/living-documentation/README.md`. This file should serve as the main index.
    *   Title: `# GCE Persistent Disk CSI Driver - Living Documentation`
    *   Introduction paragraph.
    *   `## Key Documents`: Include links to `./architecture.md` and `./features.md`.
    *   `## Driver Capabilities`: Introduction.
        *   `### Controller Service Capabilities`: List all *found* Controller CSI operations, linking to their generated `.md` files (e.g., `[ControllerCreateVolume](./controller_create_volume.md)`).
        *   `### Node Service Capabilities`: List all *found* Node CSI operations and their links.
        *   `### Identity Service Capabilities`: List all *found* Identity CSI operations and their links.
    *   Ensure all links use the relative `./filename.md` format.

4.  **Output:** You MUST output the full content of EVERY new or modified Markdown file within `docs/living-documentation/`, including the `README.md` and all the individual operation files.

> [!IMPORTANT]
> Do NOT modify any files within the `docs/llm-prompts/` directory. Your output should only be the content for the files within `docs/living-documentation/`.
