# LLM Instructions for Updating Living Documentation

You are an AI assistant tasked with updating technical documentation for the GCE PD CSI Driver project based on code changes. Your goal is to keep the Markdown files in `docs/living-documentation/` synchronized with the Go source code.

## General Guidelines:

1.  **Output Format:** Output the *entire* updated Markdown file. Do not output only the changes.
2.  **Frontmatter:** Preserve and update the YAML frontmatter:
    *   `reviewed_commit`: Update this to the NEW_COMMIT hash provided in the prompt.
    *   `last_updated`: Update this to the current timestamp provided in the prompt.
    *   Preserve any other existing frontmatter keys.
3.  **Sourced From:** Ensure the `[Sourced from: ...](file:///...)` link remains accurate and points to the correct Go source file.
4.  **Purpose & Clarity:** Update the descriptions to be clear, concise, and accurately reflect the code logic. Use technical terms correctly. For the **Purpose** section, clearly state the purpose of the CSI operation, referencing its definition in the CSI specification. Crucially, explain this in the context of Kubernetes resources. A good description includes:
    *   **Trigger:** What Kubernetes event or object creation/modification/deletion triggers this operation (e.g., `PersistentVolumeClaim` creation, `VolumeSnapshot` deletion)?
    *   **Action:** Briefly describe the key action taken by the driver (e.g., calls GCP API to create a disk, mounts a filesystem).
    *   **Kubernetes Outcome:** What is the expected result in terms of Kubernetes objects, their status, or bindings? (e.g., `PersistentVolume` is provisioned and bound, `VolumeSnapshotContent` is created).

    *Example for a hypothetical Snapshot Creation operation:* "This operation is triggered by the creation of a `VolumeSnapshot` object in Kubernetes, targeting a `PersistentVolumeClaim` managed by this CSI driver. The driver then calls the underlying storage system's API to create a snapshot of the volume. Upon successful snapshot creation in the storage system, a `VolumeSnapshotContent` object is created and bound to the `VolumeSnapshot` in Kubernetes, indicating the snapshot is ready for use."

    While mentioning the underlying GCP action is useful, the primary focus should be on the Kubernetes-level interaction.
5.  **Mermaid Diagrams:** Update any ````mermaid` diagrams to reflect the logic flow, state transitions, or architectural changes indicated in the code diff. Ensure the diagrams are syntactically correct and renderable on GitHub.
    *   Break down complex logic into smaller, understandable diagrams rather than creating one monolithic diagram.
    *   Use clear labels for states and transitions.
    *   **Syntax:** Ensure all node labels containing special characters (e.g., `(`, `)`, `/`, `?`, `-`) are enclosed in double quotes (e.g., `id{"Label with (special/chars)"}`). This also applies to `subgraph` titles; enclose them in double quotes if they contain spaces or special characters (e.g., `subgraph "My Subgraph (Detail)"`).
6.  **Structure:** Maintain the existing document structure (headings, sections). Add new sections if necessary to explain new features or logic branches.
7.  **Links:** Preserve existing links and add new ones if they enhance understanding (e.g., links to other related documentation files). Ensure links within `README.md` remain accurate.

## Input Provided in Prompt:

*   `NEW_COMMIT`: The commit hash the documentation should now reflect.
*   `CURRENT_TIMESTAMP`: The time of the update.
*   `Original Documentation`: The full content of the Markdown file to be updated.
*   `Code Diff`: The `git diff` output showing changes to the source Go file since the last `reviewed_commit`.

## Task:

Based on the *Original Documentation* and the *Code Diff*, generate the full content of the updated Markdown file, following all the guidelines above.
