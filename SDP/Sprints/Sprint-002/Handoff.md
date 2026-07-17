# Sprint-002 Handoff

Active Slice: none
Completed Slice: `SPS-002`
Active Fix: `FIX-0.1.0-001`
Study: `STU-001` (`ACCEPTED_WITH_LOW_FINDINGS`)
Status: completed Phase 2 Study; post-merge acceptance reconciliation active

## Authoritative Entry Points

- Slice contract: `SDP/Sprints/Sprint-002/ScrumIterations.md`
- Study: `SDP/02--Study/study.md`
- Governing Mandate: `SDP/01--Mandate/mandate.md` (`MAN-001`)
- Project manifest: `SDP/SDP-project.manifest.yaml`
- Current state: `SDP/Traceability/CurrentIndex.yaml`
- Relations: `SDP/Traceability/Relations.yaml`
- Ledger: `SDP/Traceability/Ledger.ndjson`
- Verification: `SDP/Verification/VER-SPS-002.md`
- Initial review: `SDP/CodeReview/REV-SPS-002-001.md`
- Follow-up review: `SDP/CodeReview/REV-SPS-002-002.md`
- Final review: `SDP/CodeReview/REV-SPS-002-003.md`
- Active Fix: `SDP/Fixes/FIX-0.1.0-001.yaml`
- Identifier convention: `SDP/Instructions/StableIdentifiers.md`

## Completion State

Branch `codex/phase-2-study` starts at exact merged baseline
`3a3b9ece5db6bde438dfd2b4eba57be344350e85`. Study activation is
`50bade647dc0409a40a2969e3b358d899edfa285`; the initial candidate is
`f1f2aba5392012c13b61a652397c6d41001beffd`; corrected Study-content commit is
`6ca9b2c2d4eafdbf83aad95f7cfcd0c418320741`; and the pushed final-review
evidence candidate is `8f59a31ee6c6ee4264cb301217c5a3489fd4444d`.

`VER-SPS-002` records passed Study/Slice verification. Initial review
`REV-SPS-002-001` returned 0 blocking, 0 high, 2 medium, and 2 low findings;
follow-up review `REV-SPS-002-002` returned 0 blocking, 0 high, 1 medium, and
0 low findings. Those findings were resolved through the corrected Study and
persistent evidence/PR reconciliation.

Fresh final Reviewer `/root/study_reviewer_final` independently inspected exact
local, remote, and draft-PR head `8f59a31...`, observed a clean worktree and
open draft PR #2 with merge state `CLEAN`, and reran the complete Study/Slice
matrix. `REV-SPS-002-003` approved `SPS-002` closure with 0 blocking, 0 high,
0 medium, and 0 low findings.

The reviewed Study recommends Go with all five target pairs (including Linux
arm64), an upstream-owned machine-readable Toolkit installation plan and
conformance fixtures, purpose-built verified Toolkit assets, hostile-input
extraction invariants, and platform-honest journaled recovery. These remain
Study inputs only. The current Toolkit remains unreleased and supplies no
eligible stable asset.

PR #2 merged normally into `main` on 2026-07-15 with exact head
`a9fa6980e633c054645b749f8a3babfefa183bf2` and merge commit
`32613734781bf39f2fce176db2acfb2284dfc92f`. The merge occurred before Steering
acceptance was recorded. On 2026-07-17, Steering accepted `STU-001` with
disposition `ACCEPTED_WITH_LOW_FINDINGS`, limited to the historical Study
evidence foundation inspected against SDP commit
`bc110bb5fd60009ba67015cf640ad6ddbfe1b04b`.

The two Low findings are the missing durable PR #2 evidence for the terminal
exact-head review and the control character plus stale open/draft/unmerged
claims in the PR body. `FIX-0.1.0-001` is active solely to reconcile that
governance and evidence state. Sprint-002, SPI-002, and SPS-002 remain complete;
their coordinates remain cleared and are not reopened.

A separate post-Study dependency assessment of `Hans-Einar/SDP` draft PR #4 at
`bf20832bed618ab240cf87c17517fc31ea721311` returned
`UPSTREAM_REWORK_REQUIRED`. Its exactly three Medium, upstream-owned findings
are insufficient normative `PlanJson` action ordering, insufficient AGENTS
conflict-migration collision/idempotency behavior, and the absent promised
portable shared conformance fixture bundle. PR #4 and installation-manifest
schema v1 are not accepted, and their missing semantics must not be defined
locally.

Mandate success criteria remain undelivered. `REL-0.1.0` and Toolkit 0.2.0
remain proposed/unreleased and unsupported, release review and verification
arrays remain empty, and publication identities remain null. The
release-event-only Ledger is unchanged because its installed schema has no
truthful Steering Study-disposition event.

## Next Authorized Decision

Finish only `FIX-0.1.0-001`: verify the exact reconciliation candidate, obtain
fresh independent review, create the new reconciliation PR evidence, and then
correct the already merged PR
[#2](https://github.com/Hans-Einar/gh-sdp/pull/2) body and attach the authorized
durable reconciliation evidence. No new lifecycle phase is authorized after
the Fix; a separate Steering decision is required.

## Residual Limitations

- No Go toolchain, Linux runtime, or macOS runtime evidence was available.
- Upstream PowerShell fixtures and Windows local-extension and
  prerelease-repository experiments were not rerun for final review.
- URL resolution is point-in-time evidence.
- Toolkit release, plan/payload, provenance, and conformance prerequisites
  identified by the Study remain unreleased or unresolved.
- The three Medium SDP PR #4 contract findings remain upstream-owned and block
  accepting that post-Study dependency or schema v1.
- A tracked record cannot self-reference its containing closure commit; final
  approval is bound to exact pushed evidence candidate `8f59a31...`.

## Stop Boundary

Do not begin Requirements, Architecture, Design Analysis, Design,
implementation planning, product code, tests/fixtures, workflows, builds,
compatibility validation, packaging, tags, or releases. Do not modify
`Hans-Einar/SDP`, accept PR #4/schema v1, or redefine its missing contract
semantics in this repository.
