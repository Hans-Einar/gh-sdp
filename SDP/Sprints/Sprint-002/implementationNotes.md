# Sprint-002 Implementation Notes

## SPS-002

Status: active — Study research and evidence collection

### Activation

- Fetched `origin` safely and confirmed `origin/main` exactly at required normal
  merge commit `3a3b9ece5db6bde438dfd2b4eba57be344350e85`.
- Confirmed PR #1 merged with approved head
  `38a7ea46e0720616f1538ca182c6a4a7e7849eda`, no merge-tree changes, a clean
  gh-sdp worktree, and exactly one root Git repository.
- Created `codex/phase-2-study` from `origin/main`; the Phase 1 branch was not
  reused.
- Activated Sprint-002 / SPI-002 / SPS-002 and Study STU-001 under MAN-001.
- Left `Ledger.ndjson` unchanged because the installed schema defines release
  events only and no release transition occurred.

### Source-State Boundary

The installed Toolkit baseline and canonical upstream `origin/main` both resolve
to `bc110bb5fd60009ba67015cf640ad6ddbfe1b04b`; no tags or GitHub Releases exist.
The separate local SDP clone contains unrelated uncommitted work on another
branch. Study evidence must therefore use immutable Git objects at the pinned
commit (and explicitly classify any other observation) without modifying or
silently updating the upstream repository.

### Work Remaining

- Complete bounded authoritative research workstreams.
- Synthesize and source `STU-001`.
- Create exact Study verification evidence.
- Commit and push the candidate, obtain fresh independent review, remediate all
  blocking/high/medium findings, and obtain exact-head follow-up review.
- Open a draft PR and stop for Steering Group assessment.
