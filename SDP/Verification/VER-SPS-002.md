# VER-SPS-002 — SPS-002 Study Candidate Verification

Status: candidate verification passed; final independent approval and Slice
closure pending
Slice: SPS-002
Study: STU-001
Mandate: MAN-001
Release context: REL-0.1.0 (`gh-sdp` 0.1.0, proposed and unreleased)
Base commit: 3a3b9ece5db6bde438dfd2b4eba57be344350e85
Activation commit: 50bade647dc0409a40a2969e3b358d899edfa285
Initial Study candidate: f1f2aba5392012c13b61a652397c6d41001beffd
Verified corrected candidate: 6ca9b2c2d4eafdbf83aad95f7cfcd0c418320741
Pull request: https://github.com/Hans-Einar/gh-sdp/pull/2 (open draft)
Verification date: 2026-07-13

This record is exact-candidate Study/Slice evidence. It is not a product test,
release gate, publication record, Steering acceptance, compatibility claim, or
authorization to begin Requirements, Architecture, or implementation. The
candidate checks pass. The repository reconciliation is carried by the commit
containing this record; `REV-SPS-002-002` still requires that commit to be
established as the pushed PR head, the PR body to be updated, and a third fresh
review of that exact head.

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

## Candidate And Pull-Request Identity

Commands:

~~~powershell
git rev-parse HEAD
git rev-parse origin/codex/phase-2-study
git rev-parse origin/main
git status --porcelain=v1
gh pr view 2 --json number,url,state,isDraft,baseRefName,headRefName,headRefOid,title
~~~

Results at verification:

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
  persistent evidence/handoff and external PR metadata reconciliation now in
  progress; it did not identify a new Study-content defect.

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
- A commit cannot truthfully self-reference its own identity. Git and PR
  metadata must therefore establish the exact pushed identity of the commit
  containing this record, and a third fresh reviewer must inspect that head
  before anyone claims zero unresolved medium findings or closes `SPS-002`.

## Disposition

The exact corrected Study candidate `6ca9b2c...` passed the documented candidate
verification matrix. Final independent approval and Slice closure remain
pending. No release, Steering, compatibility, support, or later-lifecycle claim
is made.
