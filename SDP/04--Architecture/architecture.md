# Architecture — GIP thin client

Status: selected by PLAN-SDP-0003, GIP-3-M2; implementation SPS-003.

## Ownership

ARC-GHS-001: gh-sdp is a process adapter. Canonical SDPTool owns release selection,
trust verification, caching, compatibility, installation planning and execution.
The client imports its shared bootstrap package without local policy forks.
The fixed local executable is invoked directly through Go os/exec, with no shell
or PATH resolution for the child. Arguments and streams are passed unchanged.

This explicitly supersedes STU-REC-002 / STU-ADI-001's historical proposal for a
portable apply engine inside gh-sdp. STU-001 remains an accepted historical Study,
and its source evidence and original recommendations are preserved.

## Scope and limits

No production trust keys/catalog or release asset publication are selected.
Explicit local/HTTPS descriptors and non-production public-key files support the
verified development workflow. Linux is the initial native evidence platform.
Windows and macOS require native verification before support claims.

Requirements: [REQ-GHS-001–004](../03--Requirements/requirements.md).
