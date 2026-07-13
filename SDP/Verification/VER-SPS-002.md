# VER-SPS-002 — SPS-002 Study Candidate Verification

Status: passed — independently confirmed for SPS-002 closure
Slice: SPS-002
Study: STU-001
Mandate: MAN-001
Release context: REL-0.1.0 (`gh-sdp` 0.1.0, proposed and unreleased)
Base commit: 3a3b9ece5db6bde438dfd2b4eba57be344350e85
Activation commit: 50bade647dc0409a40a2969e3b358d899edfa285
Initial Study candidate: f1f2aba5392012c13b61a652397c6d41001beffd
Verified Study-content commit: 6ca9b2c2d4eafdbf83aad95f7cfcd0c418320741
Final reviewed evidence candidate: 8f59a31ee6c6ee4264cb301217c5a3489fd4444d
Reviewer confirmation: REV-SPS-002-003
Pull request: https://github.com/Hans-Einar/gh-sdp/pull/2 (open draft)
Verification date: 2026-07-13

This record is passed Study/Slice evidence. Its authored deterministic matrix is
bound to Study-content commit `6ca9b2c...`; fresh final Reviewer
`/root/study_reviewer_final` independently reran the final matrix against exact
pushed evidence candidate `8f59a31...` and approved `SPS-002` closure in
`REV-SPS-002-003`. A tracked record cannot truthfully contain the identity of
the commit that contains itself, so no closure-metadata self-reference is
claimed. This is not a product test, release gate, publication record, Steering
acceptance, compatibility claim, or authorization to begin Requirements,
Architecture, or implementation.

## Environment

- OS: Microsoft Windows 10.0.26200, x64
- PowerShell: 5.1.26100.8655
- Git: 2.43.0.windows.1
- Python: 3.11.5
- GitHub CLI: 2.96.0
- Go: not installed
- Target repository: `C:\Users\hanse\GIT\gh-sdp`
- Canonical SDP clone used only through immutable Git objects:
  `C:\Users\hanse\GIT\SDP`
- Installed/canonical Toolkit source pin:
  `bc110bb5fd60009ba67015cf640ad6ddbfe1b04b`

## Study-Content Candidate And Pull-Request Identity

Commands:

~~~powershell
git rev-parse HEAD
git rev-parse origin/codex/phase-2-study
git rev-parse origin/main
git status --porcelain=v1
gh pr view 2 --json number,url,state,isDraft,baseRefName,headRefName,headRefOid,title
~~~

Results at authored Study-content verification:

- `HEAD` and `origin/codex/phase-2-study` both resolved to exact corrected
  candidate `6ca9b2c2d4eafdbf83aad95f7cfcd0c418320741`.
- `origin/main` resolved to required baseline
  `3a3b9ece5db6bde438dfd2b4eba57be344350e85`.
- The worktree was clean.
- PR #2 was open and draft, based on `main`, sourced from
  `codex/phase-2-study`, and reported exact head `6ca9b2c...`.

## Structured Records And Identity

The verifier loaded exact candidate blobs with `git show
6ca9b2c...:<path>`. Canonical schemas were loaded from immutable Toolkit Git
objects with `git -C C:\Users\hanse\GIT\SDP show
bc110bb5...:Toolkit/schemas/<schema>`.

Checks and results:

- Parsed all eight YAML files: the installed manifest, three Framework
  templates, release record, project manifest, CurrentIndex, and Relations.
- Applied the canonical project-manifest, installed-manifest, and release-record
  schemas successfully.
- Parsed the one non-empty `Ledger.ndjson` line and applied the canonical
  release-event schema successfully.
- Confirmed project, release, lifecycle, Sprint/Iteration/Slice, installed
  Toolkit, and publication identities agree. Client 0.1.0 and Toolkit 0.2.0
  remain separate and unreleased; tag/release publication identities remain
  null.
- Confirmed `Ledger.ndjson` is unchanged from the baseline and the
  `REL-0.1.0` review/verification arrays remain empty. Study verification is not
  represented as a release gate.

Result: passed.

## Study Shape, IDs, Relations, And Mandate Coverage

Checks and results:

- Confirmed numbered sections 1 through 18 exactly once and in order.
- Confirmed exact stable-ID counts and bidirectional Study/Relations membership:
  21 evidence records, 14 findings, eight recommendations, 17 Requirements
  candidates, ten Architecture decision inputs, and seven cross-repository
  findings (77 Study IDs total).
- Confirmed all MAN-Q-001 through MAN-Q-007 links. MAN-Q-005 remains
  `answered-for-study`; the other six remain `partially-answered-by-study`.
  These are analytical statuses, not approved Requirements or Architecture.
- Confirmed corrected direct evidence edges: `STU-FND-005` to `STU-EVD-021`,
  `STU-XRF-001` to `STU-EVD-009`, and `STU-XRF-007` to `STU-EVD-009`.
- Resolved every path recorded by Relations against an actual candidate blob.

Result: passed.

## Source Links

The exact candidate Study was scanned for unique HTTPS links. The check issued
bounded HTTP requests with a non-secret verifier user agent and treated 2xx/3xx
responses as resolvable.

Result: 50 of 50 unique Study URLs resolved. This is point-in-time reachability,
not proof that mutable documentation will remain unchanged.

## Scope, Templates, Layout, And Diff Hygiene

Commands and deterministic checks:

~~~powershell
git diff --check 3a3b9ece5db6bde438dfd2b4eba57be344350e85 6ca9b2c2d4eafdbf83aad95f7cfcd0c418320741
git ls-tree -r 6ca9b2c2d4eafdbf83aad95f7cfcd0c418320741
Get-ChildItem -LiteralPath C:\Users\hanse\GIT\gh-sdp -Force -Recurse -Directory -Filter .git
~~~

Results:

- The five later-lifecycle templates under Requirements, Architecture, Design
  Analysis, Design, and Implementation were byte-identical to the exact Phase-2
  baseline.
- The exact candidate contained 51 tracked paths, including hidden `.codex`
  content.
- No `.github` workflow, product source, Go module, executable/library, build or
  packaging file, shell/batch implementation, archive, backup tree, symlink, or
  submodule was present.
- Exactly one `.git` directory existed, at the target repository root.
- Full base-to-candidate `git diff --check` emitted no output.

Result: passed.

## Independent Review Evidence

- `REV-SPS-002-001` reviewed initial candidate `f1f2aba...` and returned
  changes required: 0 blocking, 0 high, 2 medium, 2 low.
- Corrected candidate `6ca9b2c...` addressed those four findings.
- `REV-SPS-002-002` independently reviewed exact pushed candidate `6ca9b2c...`,
  confirmed the content corrections and validation matrix, and returned changes
  required: 0 blocking, 0 high, 1 medium, 0 low. Its sole finding is the
  persistent evidence/handoff and external PR metadata reconciliation; it did
  not identify a new Study-content defect.
- `REV-SPS-002-003` independently reviewed exact pushed evidence candidate
  `8f59a31...`, confirmed all prior findings resolved, reran the complete final
  matrix, and approved `SPS-002` closure with 0 blocking, 0 high, 0 medium, and
  0 low findings.

## Final Exact-Head Confirmation

Fresh Reviewer `/root/study_reviewer_final` confirmed local HEAD,
`origin/codex/phase-2-study`, and open draft PR #2 all resolved to
`8f59a31ee6c6ee4264cb301217c5a3489fd4444d`, the worktree was clean, and GitHub
reported merge state `CLEAN`.

Against that exact candidate, the Reviewer independently confirmed:

- four intentional base-to-head commits and 13 changed SDP-only files;
- 54 tracked regular files, one root `.git`, no symlink, submodule, executable,
  ignored artifact, backup, or nested repository, and no product, workflow,
  build, package, release, or later-lifecycle addition;
- five byte-identical later-lifecycle templates and passing full-PR
  `git diff --check`;
- eight parsed YAML files, the three applicable canonical record schemas, and
  the sole non-empty Ledger event schema; byte-identical Ledger history, empty
  release review/verification arrays, and null publication identities;
- all 18 Study sections, all 77 Study IDs in the exact 21/14/8/17/10/7 counts,
  bidirectional relations, all paths, all Mandate identifier groups, and the
  stated question statuses;
- sufficient Study treatment of alternatives A through D, security,
  provenance, platform behavior, recovery, and the stop boundary; and
- 50 of 50 URL resolutions, live canonical SDP main at `bc110bb5...` with no
  tags/releases, pinned GitHub CLI source `b300f2e...`, and a truthful draft PR
  body.

Result: passed and independently confirmed for `SPS-002` closure.

## Limitations And Residual Work

- No Go toolchain was installed, and no product build is authorized in this
  Study Slice.
- No Linux or macOS runtime evidence was produced.
- The upstream PowerShell installer fixtures were not rerun. The separate SDP
  worktree had unrelated changes, so canonical evidence and schemas were read
  from immutable Git objects at the installed pin.
- Windows local-extension and prerelease repository experiments remain empirical
  work for a later authorized phase; the prerelease conclusion here is bound to
  pinned GitHub CLI 2.96.0 source.
- External links are mutable and were checked only at this verification point.
- Final exact-head review did not rerun the upstream PowerShell fixtures or the
  Windows local-extension and prerelease-repository experiments.
- A commit cannot truthfully self-reference its own identity. Final review is
  bound to exact pushed evidence candidate `8f59a31...`; no later closure-record
  commit identity is invented here.

## Disposition

The exact Study-content commit `6ca9b2c...` passed the documented verification
matrix, and fresh independent review confirmed the complete final evidence
candidate `8f59a31...`. `SPS-002` may close and `STU-001` may proceed as a
reviewed candidate to Steering assessment. No release, Steering acceptance,
compatibility, support, or later-lifecycle claim is made.
