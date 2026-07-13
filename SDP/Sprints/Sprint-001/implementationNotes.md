# Sprint-001 Implementation Notes

## SPS-001

Status: in progress — Master validation passed; commit, draft PR, and independent review pending

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

### Verification

Deterministic checks and their exact commands/results are recorded in
`SDP/Verification/VER-SPS-001.md`. The Master independently reran the checks at
`2026-07-13T09:43:35Z` against the authored working tree. Project, installed
Toolkit, release, and ledger records passed their canonical schemas; YAML,
cross-record identities, paths, stable IDs, managed hashes, later-lifecycle
template hashes, repository layout, and the no-product boundary passed. Normal
same-version reinstall preview proposed zero changes.

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
  verification while review and final-commit evidence remain pending.

### Residual Risks, Gaps, and Discoveries

- The canonical project-manifest schema has no fields for repository URL,
  lifecycle phase, Toolkit compatibility target, or document/stable IDs. These
  facts remain in supported manifest fields plus authoritative project records;
  no unsupported manifest fields were invented.
- The canonical `validate_sdp.py` validates the Toolkit repository's own layout
  and cannot be represented as a consuming-project validator. Consumer checks
  therefore compose the canonical JSON schemas with project-local YAML/NDJSON,
  path/ID, hash, repository-layout, and content-boundary checks.
- The final commit does not yet exist, so current evidence applies to the stated
  base revision plus the working tree and must be rerun against the committed PR
  head before any completion decision.
- Normal same-version installation is idempotent, but the optional initializer is
  not idempotent after correcting its Toolkit-specific release seed. The current
  mitigation is preview-first, one-time use; correcting the reusable installer
  requires a separate authorized change in `Hans-Einar/SDP`.
- Detailed artifact/provenance, compatibility, rollback, alternate-source,
  authentication, offline, and platform filesystem policy remains input to
  Study. No later-lifecycle template was populated.
