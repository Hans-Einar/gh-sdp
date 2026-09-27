# Architecture — GIP thin client

Status: SPS-003 thin client and SPS-004 first release delivered; SPS-005 selects
the default-engine patch release.

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

SPS-005 retains the shared bootstrap production public trust and selects the exact
SDP v0.2.1 descriptor as the default. No mutable latest-release catalog is inferred.
Explicit local/HTTPS descriptors and non-production public-key files remain
available for development. Release packaging targets Linux amd64 only.
Windows and macOS require native verification before support claims.

Requirements: [REQ-GHS-001–005](../03--Requirements/requirements.md).
