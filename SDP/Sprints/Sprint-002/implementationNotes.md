# Sprint-002 Implementation Notes

## SPS-002

Status: active — initial review changes required; bounded rework applied,
verification and exact-head follow-up review pending

### Activation

- Fetched `origin` safely and confirmed `origin/main` exactly at required normal
  merge commit `3a3b9ece5db6bde438dfd2b4eba57be344350e85`.
- Confirmed PR #1 merged with approved head
  `38a7ea46e0720616f1538ca182c6a4a7e7849eda`, no merge-tree changes, a clean
  gh-sdp worktree, and exactly one root Git repository.
- Created `codex/phase-2-study` from `origin/main`; the Phase 1 branch was not
  reused.
- Activated Sprint-002 / SPI-002 / SPS-002 and Study STU-001 under MAN-001 in
  commit `50bade647dc0409a40a2969e3b358d899edfa285`.
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

### Candidate And Initial Review

- Committed the first complete Study candidate as
  `f1f2aba5392012c13b61a652397c6d41001beffd`, pushed it to
  `origin/codex/phase-2-study`, and opened draft PR
  [#2](https://github.com/Hans-Einar/gh-sdp/pull/2) against `main`.
- Fresh Reviewer `/root/study_reviewer_initial` returned `changes-required`
  with 0 blocking, 0 high, 2 medium, and 2 low findings. The findings covered
  three omitted direct evidence links, stale candidate/handoff state, one dead
  Apple evidence URL, and inaccurate prerelease-only binary-install wording.
- This bounded rework adds the missing relations, records the actual candidate
  and review state, replaces the Apple URL, and aligns the prerelease behavior
  with pinned GitHub CLI 2.96.0 source. The corrected exact candidate has not
  yet been committed or pushed by the Master; verification and fresh follow-up
  review remain pending.

### Work Remaining

- Have the Master commit and push this bounded correction, then create exact
  Study verification evidence and persistent review records bound to the
  corrected candidate.
- Obtain fresh exact-head follow-up review with zero unresolved blocking, high,
  or medium findings and update draft PR #2 truthfully.
- Stop with the PR still draft for Steering Group assessment.
