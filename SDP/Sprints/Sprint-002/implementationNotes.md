# Sprint-002 Implementation Notes

## SPS-002

Status: complete — Study/Slice verification passed; final exact-head review
approved SPS-002 closure; Steering accepted the historical Study evidence
foundation with Low findings

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
- Commit `6ca9b2c2d4eafdbf83aad95f7cfcd0c418320741` adds the missing relations,
  records the actual first-candidate/review state, replaces the Apple URL, and
  aligns prerelease behavior with pinned GitHub CLI 2.96.0 source. The Master
  committed and pushed that corrected candidate; local HEAD, remote branch, and
  draft PR #2 all resolved to it with a clean worktree at follow-up review.

### Corrected Candidate Verification And Follow-up Review

- Exact candidate verification at `6ca9b2c...` passed all eight YAML parses,
  three applicable record schemas, the one ledger release-event schema, identity
  and coordinate checks, 18 Study sections, all 77 Study IDs
  (21/14/8/17/10/7), all Mandate questions, relation paths and corrected direct
  evidence edges, 50/50 external links, five unchanged later-lifecycle
  templates, the complete 51-path hidden-aware tree, prohibited-file and Git
  layout checks, and full base-to-candidate `git diff --check`.
- Fresh Reviewer `/root/study_reviewer_followup` inspected exact pushed
  `6ca9b2c...`, confirmed every initial content correction and independently
  repeated the full Study and validation review. `REV-SPS-002-002` returned
  `changes-required` with 0 blocking, 0 high, 1 medium, and 0 low findings.
- The sole follow-up finding was persistent evidence/handoff state: these notes
  and Handoff still described `6ca9b2c...` as uncommitted, the two review records
  and `VER-SPS-002` were absent, the Study/Slice arrays were empty, and the draft
  PR body remained stale. No new Study-content defect was identified.
- `VER-SPS-002`, both earlier review records, updated Study/Slice relations,
  current notes/handoff, and a truthful draft-PR body were committed and pushed
  as exact evidence candidate
  `8f59a31ee6c6ee4264cb301217c5a3489fd4444d`.

### Final Exact-Head Review And Closure

- Fresh Reviewer `/root/study_reviewer_final` independently reviewed exact
  local, remote, and draft-PR head `8f59a31...` against baseline `3a3b9ec...`.
  PR #2 was open and draft with merge state `CLEAN`, and the worktree was clean.
- `REV-SPS-002-003` confirmed every prior finding resolved and reran the full
  matrix: schemas and Ledger, release isolation, all 18 sections and 77 Study
  IDs, all Mandate identifier groups and question statuses, 50/50 URLs, five
  unchanged later-lifecycle templates, 54 regular tracked files, one root Git
  repository, no prohibited artifact, and passing full-PR diff hygiene.
- The final disposition was approved for `SPS-002` closure with 0 blocking,
  0 high, 0 medium, and 0 low findings. `STU-001` is therefore a reviewed
  candidate ready for Steering assessment, not a Steering-accepted Study.
- Sprint-002, SPI-002, and SPS-002 are complete. Active development coordinates
  are cleared. `REL-0.1.0` and Toolkit 0.2.0 remain proposed/unreleased;
  release review/verification arrays, publication identities, and the Ledger
  remain unchanged.

### Post-Merge Steering Reconciliation

- PR #2 merged normally into `main` on 2026-07-15. Its exact head was
  `a9fa6980e633c054645b749f8a3babfefa183bf2`, and its merge commit was
  `32613734781bf39f2fce176db2acfb2284dfc92f`. The merge occurred before the
  Steering acceptance record.
- On 2026-07-17, Steering recorded `ACCEPTED_WITH_LOW_FINDINGS` for `STU-001`,
  limited to the historical Study evidence foundation inspected against SDP
  commit `bc110bb5fd60009ba67015cf640ad6ddbfe1b04b`.
- The two accepted Low findings are the absence of durable PR #2 evidence for
  the terminal exact-head review and the control character plus stale
  open/draft/unmerged claims in the PR #2 body. `FIX-0.1.0-001` is the active
  bounded reconciliation; Sprint-002, SPI-002, and SPS-002 remain complete.
- A separate post-Study dependency assessment of `Hans-Einar/SDP` draft PR #4
  at `bf20832bed618ab240cf87c17517fc31ea721311` returned
  `UPSTREAM_REWORK_REQUIRED` with exactly three Medium findings owned upstream:
  normative `PlanJson` action ordering is insufficiently specified; AGENTS
  conflict-migration collision/idempotency behavior is insufficiently
  specified; and the promised portable shared conformance fixture bundle is
  absent.
- PR #4 and installation-manifest schema v1 are not accepted. No missing
  contract semantic is defined locally, and no downstream compatibility
  validation or later lifecycle phase is authorized.
- `Ledger.ndjson` remains unchanged because the installed canonical schema
  defines release events only and has no truthful event type for a Steering
  Study disposition.
- Exact pushed reconciliation candidate
  `d8d0404faff58338020afd31df6c4d9454776080` passed
  `VER-FIX-0.1.0-001`, including Git/GitHub identity, merge parentage and tree
  equality, structured-record/schema, traceability, protected-state, scope, and
  diff-hygiene checks. Fresh independent review and terminal durable evidence
  remain required before the Fix closes.

### Next Decision And Stop

Complete only `FIX-0.1.0-001` verification, fresh independent review, PR
evidence, and the approved PR #2 metadata correction. After the Fix, wait for a
separate next authorization. Do not reopen the completed Sprint/Iteration/Slice
or begin Requirements, Architecture, Design, implementation, compatibility
validation, build, packaging, workflow, release, or upstream Toolkit work.
