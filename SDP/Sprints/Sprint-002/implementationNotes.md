# Sprint-002 Implementation Notes

## SPS-002

Status: active — Study candidate awaiting verification and independent review

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

### Research And Synthesis

- Delegated separate read-only research passes for GitHub CLI/Go,
  Toolkit-installer/ownership, and security/platform/transaction behavior; the
  Master inspected the cited primary evidence rather than treating agent
  summaries as proof.
- Inspected GitHub CLI 2.96.0 source commit
  `b300f2ec7ec9dc9addc39b2ad88c54097ded7ca0`, Go 1.26.5 and its official
  vulnerability record, GitHub REST/release/provenance documentation, platform
  filesystem authorities, and immutable Toolkit objects at `bc110bb5...`.
- Authored `STU-001` with exactly 18 required sections, 21 evidence records, 14
  findings, eight recommendations, 17 Requirements candidates, ten Architecture
  decision inputs, and seven cross-repository findings.
- Resolved every Study identifier in `Relations.yaml` and recorded analytical
  progress for `MAN-Q-001` through `MAN-Q-007` without marking any future success
  criterion delivered or any recommendation approved.
- Left the release-event-only Ledger unchanged; no release transition occurred.

### Work Remaining

- Create exact Study verification evidence and bind it to the committed
  candidate.
- Commit and push the candidate, obtain fresh independent review, remediate all
  blocking/high/medium findings, and obtain exact-head follow-up review.
- Open a draft PR and stop for Steering Group assessment.
