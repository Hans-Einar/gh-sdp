# Sprint-001 Implementation Notes

## SPS-001

Status: in progress — draft PR exists; review remediation validated and follow-up review pending

### Worker Changes

- Replaced the installer seed in `AGENTS-project.md` with bounded `gh-sdp`
  instructions that preserve the managed role hierarchy and prohibit production
  implementation in this Slice.
- Added the project-owned `SDP/Instructions/StableIdentifiers.md` convention
  because the canonical Toolkit defines no Mandate/outcome/boundary ID grammar.
- Replaced the Mandate template with drafting `MAN-001`, including product
  problem, stakeholders, future command intent, version/repository boundaries,
  preservation, stable-release, cross-platform, security/reliability, success,
  assumptions, constraints, open questions, and the Phase 1 stop boundary.
- Expanded `SDP/Traceability/Relations.yaml` so the Mandate, its major outcomes,
  boundaries and success criteria, active work coordinates, working-tree
  verification, and proposed release target resolve to real records. The release
  remains unreleased, publication fields remain null, and no new ledger event was
  appended.
- Added `SDP/Verification/VER-SPS-001.md` and updated the handoff for the pending
  Master integration and fresh Reviewer pass.
- Remediated the four medium findings in `REV-SPS-001-001`: exact-candidate
  evidence now names commit `381b38ed6fc59936e7e2cb30b877e156ff57914c`
  and draft PR #1, `git diff --check` is recorded, the scope check includes
  hidden tracked/untracked paths, and all four assumption plus seven question
  IDs resolve through `Relations.yaml` and validation.

### Verification

Deterministic checks and their exact commands/results are recorded in
`SDP/Verification/VER-SPS-001.md`. The reviewed commit
`381b38ed6fc59936e7e2cb30b877e156ff57914c` was rerun directly from Git objects,
including `git diff --check` and all 46 tracked paths, so hidden `.codex` content
and the `.github` exclusion were evaluated without relying on the remediation
working tree. The current remediation tree separately passes the expanded
relation check for outcomes, boundaries, success criteria, assumptions, and
questions, along with canonical schemas, paths, managed hashes, untouched
later-lifecycle templates, the one-Git-root invariant, and the no-product
boundary.

Repeating the one-time `-InitializeProjectStructure` preview proposed one change:
re-creating the Toolkit-specific `SDP/Releases/REL-0.2.0.yaml`. It was not applied.
The project warning and agent instruction record this canonical initializer gap.

### Decisions

- `MAN-001` subordinate IDs use an explicitly project-local Phase 1 convention;
  no upstream format is claimed.
- The `gh-sdp` `0.1.0` client target and SDP Toolkit `0.2.0` compatibility target
  are separate proposals and both remain unreleased.
- Go is a preferred direction only; Study and Architecture retain the decision.
- `VER-SPS-001` is linked as Slice evidence but not as accepted release-gate
  verification while the bounded remediation still requires Master integration
  and fresh follow-up review.

### Independent Review Pass 1

Fresh Reviewer `/root/phase1_reviewer` inspected committed candidate
`381b38ed6fc59936e7e2cb30b877e156ff57914c` and draft PR #1. The disposition was
changes required: zero blocking, zero high, four medium, and one low finding.
`SDP/CodeReview/REV-SPS-001-001.md` is the authoritative review record. The four
medium findings are addressed in the current remediation working tree; the low
empty-repository-description finding is unchanged for Steering Group
disposition. A fresh follow-up Reviewer remains required.

### Residual Risks, Gaps, and Discoveries

- The canonical project-manifest schema has no fields for repository URL,
  lifecycle phase, Toolkit compatibility target, or document/stable IDs. These
  facts remain in supported manifest fields plus authoritative project records;
  no unsupported manifest fields were invented.
- The canonical `validate_sdp.py` validates the Toolkit repository's own layout
  and cannot be represented as a consuming-project validator. Consumer checks
  therefore compose the canonical JSON schemas with project-local YAML/NDJSON,
  path/ID, hash, repository-layout, and content-boundary checks.
- The initial reviewed commit and draft PR exist. The current bounded remediation
  is not yet committed or pushed, so the Master must rerun it after integration
  and a fresh Reviewer must inspect the resulting PR head before any completion
  decision.
- Normal same-version installation is idempotent, but the optional initializer is
  not idempotent after correcting its Toolkit-specific release seed. The current
  mitigation is preview-first, one-time use; correcting the reusable installer
  requires a separate authorized change in `Hans-Einar/SDP`.
- Detailed artifact/provenance, compatibility, rollback, alternate-source,
  authentication, offline, and platform filesystem policy remains input to
  Study. No later-lifecycle template was populated.
