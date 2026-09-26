# Sprint-003 — Thin SDPTool delegation

Status: in-progress
Iteration: SPI-003
Slice: SPS-003
Milestone: GIP-3-M2 in [PLAN-SDP-0003](https://github.com/Hans-Einar/SDP/blob/sdp/install-gip-3/SDP/05--Implementation/SDPTool/Installation/Plan.md)
Authorization: owner selected the complete GIP implementation plan on 2026-09-27.
Release context: REL-0.1.0 remains proposed and unreleased.
BranchPolicy: phase; branch sdp/install-gip-3 from gh-sdp main be990cc.
CommitPolicy: milestone; GIP-3-M2 concrete delivery, then evidence/review corrections.

## SPI-003 / SPS-003 contract

Implement a native Go gh-sdp executable that bootstraps a verified compatible
SDPTool binary through the canonical shared bootstrap package, then delegates
all arguments, streams, working directory and child exit status. The same
SDPTool owns release verification, platform/protocol selection and caching;
installation policy, preview, apply and recovery remain entirely in SDPTool.
The current [canonical contract](https://github.com/Hans-Einar/SDP/blob/sdp/install-gip-3/SDP/04--Design/SDPTool/Installation/Contract.md) supersedes the Study's proposed
client-owned apply engine. This is a separate authorization, not retroactive
acceptance of historical upstream PR #4 or Study candidates.

Expected files: Go client and meaningful tests, dependency pin/provenance,
README, current Requirements/Architecture/Design/Implementation documents,
project-owned authority addenda and Sprint/verification/traceability records.
Do not change managed AGENTS.md, installed Toolkit facts, historical event bytes,
completed Sprint records, live XFMD, release workflows, tags or releases.

## Acceptance and stop boundary

- Build/run actual gh-sdp with a signed explicit non-production test descriptor
  and actual SDPTool child. Preview/apply results must match direct SDPTool.
- Unit tests prove exact argv (including spaces/metacharacters), stdin/stdout/
  stderr, cwd and nonzero exit propagation; no shell or PATH engine lookup.
- Canonical bootstrap verifies signed descriptor, digest, platform and protocol;
  exercise wrong-key, incompatible, corrupt cache and offline cases in upstream
  evidence, with actual wrapper rejection evidence for incompatible binaries.
- Linux is the verified execution host. Cross-builds cannot prove native Windows
  or macOS behavior; neither platform is advertised supported.
- Validate YAML/NDJSON syntax, current coordinates and diff whitespace.
- Obtain fresh independent review of the actual paired candidates; Master closes
  records only after evidence and findings are resolved. Owner acceptance,
  production signing keys, a catalog, release and live rollout remain unselected.

Return the bounded milestone and paired source identities to the coordinating
Master. Do not begin another Slice or imply completion before actual review.
