# STU-001 — gh-sdp Phase 2 Study

Status: accepted evidence foundation — `ACCEPTED_WITH_LOW_FINDINGS`
Mandate: MAN-001
Sprint: Sprint-002
Iteration: SPI-002
Slice: SPS-002
Inspection date: 2026-07-13
Steering disposition date: 2026-07-17
Inspected SDP source: `bc110bb5fd60009ba67015cf640ad6ddbfe1b04b`
Authorization boundary: Study only

This Study records evidence, findings, recommendations, candidate requirements,
and Architecture inputs. It approves none of the latter and delivers none of
`MAN-SC-001` through `MAN-SC-007`. Requirements, Architecture, Design Analysis,
Design, implementation planning, product code, packaging, workflows, tags, and
releases remain outside `SPS-002`.

## Steering disposition and post-Study boundary

On 2026-07-17, the Steering Group accepted `STU-001` with disposition
`ACCEPTED_WITH_LOW_FINDINGS` as the Phase 2 Study evidence foundation. The
acceptance applies only to the historical Study performed against pinned SDP
commit `bc110bb5fd60009ba67015cf640ad6ddbfe1b04b`; it does not retroactively
change the inspection date, evidence inventory, findings, recommendations, or
source classifications below.

PR #2 merged normally before acceptance was recorded. Its exact head was
`a9fa6980e633c054645b749f8a3babfefa183bf2`; GitHub created merge commit
`32613734781bf39f2fce176db2acfb2284dfc92f` on 2026-07-15. Sprint-002,
SPI-002, and SPS-002 remain complete and are not reopened. The accepted Low
findings are:

1. the terminal exact-head review claim for `a9fa698...` was not durably
   attached to PR #2 as a PR comment or native review; and
2. the PR #2 body contains a control character before the short head identity
   and stale statements that the PR is open, draft, and unmerged.

A separate post-Study dependency assessment inspected `Hans-Einar/SDP` draft
PR #4 at exact head `bf20832bed618ab240cf87c17517fc31ea721311`. Its disposition
is `UPSTREAM_REWORK_REQUIRED`, with exactly three Medium, upstream-owned
findings:

1. normative `PlanJson` action ordering is insufficiently specified;
2. AGENTS conflict-migration collision/idempotency behavior is insufficiently
   specified; and
3. the promised portable shared conformance fixture bundle is absent.

PR #4 and installation-manifest schema v1 are not accepted. These later facts
are not historical Study evidence, and this repository does not redefine the
missing semantics or work around the findings. The Steering disposition
authorizes no Requirements, Architecture, Design Analysis, Design,
implementation, compatibility validation, packaging, workflow, or release
work. `gh-sdp` 0.1.0 and SDP Toolkit 0.2.0 remain unreleased and unsupported.

`SDP/Traceability/Ledger.ndjson` remains unchanged because the installed
canonical schema defines release events only and has no truthful event type for
a Steering Study disposition.

## 1. Scope and methodology

### Goal, authority, and exclusions

The goal is to establish the factual and analytical basis needed to write later
Requirements and Architecture for the `gh-sdp` GitHub CLI extension. The
governing authority is `MAN-001`, especially `MAN-OUT-001` through
`MAN-OUT-006`, active boundaries `MAN-BND-002` through `MAN-BND-008`, approved
but undelivered success criteria `MAN-SC-001` through `MAN-SC-007`, assumptions
`MAN-ASM-001` through `MAN-ASM-004`, and questions `MAN-Q-001` through
`MAN-Q-007`. The execution coordinates are `Sprint-002` / `SPI-002` /
`SPS-002`. [verified repository fact; STU-EVD-001]

For exact stable-ID resolution, this Study traces to outcomes `MAN-OUT-001`,
`MAN-OUT-002`, `MAN-OUT-003`, `MAN-OUT-004`, `MAN-OUT-005`, and
`MAN-OUT-006`; applicable boundaries `MAN-BND-002`, `MAN-BND-003`,
`MAN-BND-004`, `MAN-BND-005`, `MAN-BND-006`, `MAN-BND-007`, and
`MAN-BND-008`; undelivered success criteria `MAN-SC-001`, `MAN-SC-002`,
`MAN-SC-003`, `MAN-SC-004`, `MAN-SC-005`, `MAN-SC-006`, and `MAN-SC-007`;
and assumptions `MAN-ASM-001`, `MAN-ASM-002`, `MAN-ASM-003`, and
`MAN-ASM-004`. [verified repository fact; STU-EVD-001]

The Study excludes all later-lifecycle approval and implementation. It does not
modify `Hans-Einar/SDP`, define final CLI semantics, select Architecture, create
product tests, or claim either `gh-sdp` 0.1.0 or SDP Toolkit 0.2.0 is released or
supported. [verified repository fact; STU-EVD-001]

### Method

Sources were ranked: immutable repository objects and official source; official
product/platform documentation; repository command evidence; then explicit
inference. Search snippets and research summaries were discovery aids only.
Material claims use these classifications:

- **verified external fact** — directly supported by an authoritative external
  source;
- **verified repository fact** — directly supported by an immutable project or
  Toolkit object, manifest, test, or command result;
- **observation** — directly observed in the inspected environment but not a
  general guarantee;
- **inference** — reasoned consequence of facts, still subject to validation;
- **recommendation** — Study advice for later approval;
- **assumption** — premise not yet verified;
- **unresolved question** — a decision or evidence gap left open.

### Inspected baselines and limitations

- `gh-sdp` began Phase 2 from merge baseline
  `3a3b9ece5db6bde438dfd2b4eba57be344350e85`; Study activation is commit
  `50bade647dc0409a40a2969e3b358d899edfa285`. [verified repository fact;
  STU-EVD-001]
- The installed and current fetched SDP Toolkit source is
  `bc110bb5fd60009ba67015cf640ad6ddbfe1b04b`; no tag or GitHub Release existed
  when inspected. [verified repository fact; STU-EVD-011]
- GitHub CLI 2.96.0, source commit
  `b300f2ec7ec9dc9addc39b2ad88c54097ded7ca0`, is the client behavior baseline.
  [observation and verified external fact; STU-EVD-003, STU-EVD-017]
- The local environment was Windows NT 10.0.26200 AMD64, Git
  2.43.0.windows.1, Python 3.11.5, and GitHub CLI 2.96.0. No `go` executable was
  on `PATH`, so this Study contains no local compilation or cross-execution
  evidence. [observation; STU-EVD-017]
- The separate SDP checkout carried unrelated user branch/state and was not
  treated as a writable clean-room fixture. The Study made no change there and
  read pinned evidence with `git show <revision>:<path>`; current upstream state
  was read from fetched `origin/main` and the API. [observation; STU-EVD-011]
- Live documentation and API behavior may evolve after the inspection date;
  later Requirements must pin supported versions and re-verify them.
  [assumption]

### Evidence inventory

Every entry was inspected on 2026-07-13.

| ID | Source and revision/version | What it proves | Limitations |
| --- | --- | --- | --- |
| STU-EVD-001 | Project paths `SDP/01--Mandate/mandate.md`, `SDP/Sprints/Sprint-002/ScrumIterations.md`, `SDP/Traceability/{CurrentIndex.yaml,Relations.yaml}` at activation commit `50bade647dc0409a40a2969e3b358d899edfa285`; merge `3a3b9ece5db6bde438dfd2b4eba57be344350e85` | Phase 2 authority, boundaries, identities, and active contract | Project-local governance, not product behavior |
| STU-EVD-002 | [GitHub CLI extension manual](https://cli.github.com/manual/gh_extension), [install manual](https://cli.github.com/manual/gh_extension_install), [upgrade manual](https://cli.github.com/manual/gh_extension_upgrade), [remove manual](https://cli.github.com/manual/gh_extension_remove), and [official extension authoring guide](https://docs.github.com/en/github-cli/github-cli/creating-github-cli-extensions), live docs | Naming, argument forwarding, local/remote install, release binaries, asset suffixes, upgrade/remove, Windows `.exe`, GHES URL | Live documentation; some operational detail requires source |
| STU-EVD-003 | GitHub CLI 2.96.0 immutable source: [`manager.go`](https://github.com/cli/cli/blob/b300f2ec7ec9dc9addc39b2ad88c54097ded7ca0/pkg/cmd/extension/manager.go#L91-L133), [install/distribution paths](https://github.com/cli/cli/blob/b300f2ec7ec9dc9addc39b2ad88c54097ded7ca0/pkg/cmd/extension/manager.go#L208-L377), [upgrade/platform paths](https://github.com/cli/cli/blob/b300f2ec7ec9dc9addc39b2ad88c54097ded7ca0/pkg/cmd/extension/manager.go#L449-L841), [`http.go`](https://github.com/cli/cli/blob/b300f2ec7ec9dc9addc39b2ad88c54097ded7ca0/pkg/cmd/extension/http.go#L66-L148), and [Windows pointer](https://github.com/cli/cli/blob/b300f2ec7ec9dc9addc39b2ad88c54097ded7ca0/pkg/cmd/extension/symlink_windows.go#L1-L14) | Exact 2.96.0 execution, binary detection/download, latest release, platform mapping, manifest, local pointer, and upgrade behavior | Implementation evidence for 2.96.0, not a permanent public contract |
| STU-EVD-004 | [`gh help environment`](https://cli.github.com/manual/gh_help_environment), [`gh auth login`](https://cli.github.com/manual/gh_auth_login), [`gh auth token`](https://cli.github.com/manual/gh_auth_token), and [`gh api`](https://cli.github.com/manual/gh_api), live docs | Token precedence, host selection, credential storage, token output, authenticated API and cache option | Does not select the extension's API architecture |
| STU-EVD-005 | GitHub REST [rate limits](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api), [latest release](https://docs.github.com/en/rest/releases/releases?apiVersion=latest), and [release assets](https://docs.github.com/en/rest/releases/assets), current API docs | Public unauthenticated access, 60/hour unauthenticated and 5,000/hour ordinary authenticated limit, latest-release rules, private read permission, asset digest/download | Limits vary by credential/account and secondary limits; API is versioned |
| STU-EVD-006 | Go [source/install and target docs](https://go.dev/doc/install/source) and [porting policy](https://go.dev/wiki/PortingPolicy), live docs inspected against Go 1.26 generation | `GOOS`/`GOARCH`, `CGO_ENABLED=0`, target combinations, first-class ports and glibc caveat | Build support does not prove this future program runs correctly |
| STU-EVD-007 | Go 1.26.5 [`archive/zip`](https://pkg.go.dev/archive/zip@go1.26.5), [`archive/tar`](https://pkg.go.dev/archive/tar@go1.26.5), [`os`](https://pkg.go.dev/os@go1.26.5), [`os/exec`](https://pkg.go.dev/os/exec@go1.26.5), immutable [Go source commit `c19862e…`](https://github.com/golang/go/tree/c19862e5f8415b4f24b189d065ed739517c548ba/src), [traversal-resistant APIs](https://go.dev/blog/osroot), and [GO-2026-4970](https://pkg.go.dev/vuln/GO-2026-4970) | Archive entry types, explicit ZIP path mode, root-scoped operations, process CWD, rename/sync differences, and the patched boundary for an `os.Root` escape vulnerability | Standard library primitives are not a complete extraction or transaction policy; future toolchains must be pinned and re-audited |
| STU-EVD-008 | SDP immutable docs at `bc110…`: [`SDP.manifest.yaml`](https://github.com/Hans-Einar/SDP/blob/bc110bb5fd60009ba67015cf640ad6ddbfe1b04b/SDP.manifest.yaml), [`Toolkit/README.md`](https://github.com/Hans-Einar/SDP/blob/bc110bb5fd60009ba67015cf640ad6ddbfe1b04b/Toolkit/README.md), [`Distribution-And-Upgrades.md`](https://github.com/Hans-Einar/SDP/blob/bc110bb5fd60009ba67015cf640ad6ddbfe1b04b/docs/Distribution-And-Upgrades.md), [`Installer-Migration.md`](https://github.com/Hans-Einar/SDP/blob/bc110bb5fd60009ba67015cf640ad6ddbfe1b04b/docs/Installer-Migration.md), [`Toolkit-Manifest.md`](https://github.com/Hans-Einar/SDP/blob/bc110bb5fd60009ba67015cf640ad6ddbfe1b04b/docs/Toolkit-Manifest.md) | Normative Toolkit 0.2.0 identity, ownership and migration statements | Toolkit is unreleased; prose is not fully machine-readable |
| STU-EVD-009 | SDP immutable [`Toolkit/scripts/Install-SDP.ps1`](https://github.com/Hans-Einar/SDP/blob/bc110bb5fd60009ba67015cf640ad6ddbfe1b04b/Toolkit/scripts/Install-SDP.ps1) at `bc110…` | Actual PowerShell detection, comparison, preservation, backup, preview, copying and reporting behavior | Current implementation behavior is not automatically normative |
| STU-EVD-010 | SDP immutable [`Install-SDP.Tests.ps1`](https://github.com/Hans-Einar/SDP/blob/bc110bb5fd60009ba67015cf640ad6ddbfe1b04b/Toolkit/tests/Install-SDP.Tests.ps1), `Toolkit/payload/`, `Toolkit/schemas/`, and repository lifecycle seeds at `bc110…` | Covered fixtures, actual payload boundary, schemas, and initializer inputs | Tests are Windows/PowerShell fixtures and omit several failure/concurrency/security cases |
| STU-EVD-011 | Fetched `Hans-Einar/SDP` `origin/main` = `bc110…`; `git tag --list`, `gh release list`, and REST tags/releases returned empty on 2026-07-13; [commit](https://github.com/Hans-Einar/SDP/commit/bc110bb5fd60009ba67015cf640ad6ddbfe1b04b) | Installed baseline equals fetched main and no published Toolkit release/tag existed | Point-in-time mutable upstream observation; local checkout was preserved |
| STU-EVD-012 | GitHub [release asset digests](https://docs.github.com/en/rest/releases/assets), [immutable release verification](https://docs.github.com/en/code-security/how-tos/secure-your-supply-chain/secure-your-dependencies/verify-release-integrity), [artifact attestations](https://docs.github.com/en/actions/concepts/security/artifact-attestations), and [signature verification](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification) | Integrity/provenance mechanisms and their boundaries | Mechanisms only help when publication and verification policy actually uses them |
| STU-EVD-013 | GitHub [source archive stability](https://docs.github.com/en/repositories/working-with-files/using-files/downloading-source-code-archives) | Generated archive bytes may change; commit contents are stable; releases are preferred for security | Does not define an SDP payload format |
| STU-EVD-014 | Microsoft [path naming](https://learn.microsoft.com/en-us/windows/win32/fileio/naming-a-file), [long paths](https://learn.microsoft.com/en-us/windows/win32/fileio/maximum-file-path-limitation), [links/junctions](https://learn.microsoft.com/en-us/windows/win32/fileio/hard-links-and-junctions), [reparse points](https://learn.microsoft.com/en-us/windows/win32/fileio/reparse-points), and [`ReplaceFileW`](https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-replacefilew) | Windows roots, reserved names, case default, link types, long-path opt-in, same-volume replacement and partial failure states | Filesystem and policy configuration still vary by machine |
| STU-EVD-015 | POSIX Issue 8 [`rename`](https://pubs.opengroup.org/onlinepubs/9799919799/functions/rename.html) and [durability rationale](https://pubs.opengroup.org/onlinepubs/9799919799/xrat/V4_xbd_chap01.html) | Atomic rename semantics, cross-filesystem failure, and distinction between atomicity and durability | Applies to conforming Unix behavior, not Windows and not whole-tree transactions |
| STU-EVD-016 | Apple [files/directories](https://developer.apple.com/documentation/technologyoverviews/files-and-directories), [APFS FAQ](https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/APFS_Guide/FAQ/FAQ.html), and [`replaceItem`](https://developer.apple.com/documentation/foundation/filemanager/replaceitemat%28_%3Awithitemat%3Abackupitemname%3Aoptions%3A%29) | APFS case/normalization behavior, portable path advice, same-volume replacement | macOS may use non-APFS or case-sensitive volumes |
| STU-EVD-017 | Local commands `gh --version`, `git --version`, `python --version`, `Get-Command go`, and OS environment | Exact Study tool environment and absence of local Go | One Windows workstation only |
| STU-EVD-018 | [Semantic Versioning 2.0.0](https://semver.org/) | Full SemVer syntax and prerelease precedence | Does not define Toolkit compatibility policy |
| STU-EVD-019 | `SDP/Framework/installed-toolkit.manifest.yaml`, `SDP/SDP-project.manifest.yaml`, and `SDP/Releases/REL-0.1.0.yaml` at activation commit | Installed Toolkit and proposed client identities/state | Project state only; not publication evidence |
| STU-EVD-020 | GitHub's immutable [`cli/go-gh` README](https://github.com/cli/go-gh/blob/a89b68a441b259f483266e4cee718d9a8e606722/README.md) at `a89b68a…` | Go API client can follow `gh` token/host conventions and stored OAuth fallback | Current library direction, not an approved dependency |
| STU-EVD-021 | Project `SDP/Verification/VER-SPS-001.md` and `AGENTS-project.md` at merge `3a3b9ec…` | Actual normal-repeat result and repeat-initializer proposal of `SDP/Releases/REL-0.2.0.yaml` | Historical project-specific execution, not a corrected upstream contract |

## 2. Product and version identities

| Identity | Study state | Classification and evidence |
| --- | --- | --- |
| `gh-sdp` client | Current `0.0.0`; proposed target `0.1.0`; unreleased, no release commit/tag/URL | verified repository fact; STU-EVD-019 |
| Installed project Toolkit | Toolkit `0.2.0`, installer `0.2.0`, source `bc110…` | verified repository fact; STU-EVD-019 |
| Selected/available Toolkit release | None selected or published | verified repository fact; STU-EVD-011 |
| Latest stable Toolkit release | None existed at inspection; `0.2.0` was explicitly `unreleased` | verified repository fact; STU-EVD-008, STU-EVD-011 |
| GitHub CLI | 2.96.0, source `b300f2e…` | observation/verified external fact; STU-EVD-003, STU-EVD-017 |
| Development work | `Sprint-002` / `SPI-002` / `SPS-002` / `STU-001` | verified repository fact; STU-EVD-001 |

`gh-sdp` 0.1.0 and SDP Toolkit 0.2.0 are independent versions of separate
products. Neither advances, releases, or proves compatibility for the other.
[verified repository fact; STU-EVD-001, STU-EVD-019]

**STU-FND-001 — identity truth.** Status and version output must model the
client, installed Toolkit, selected Toolkit, latest eligible Toolkit, and each
release state separately. [inference; STU-EVD-011, STU-EVD-019]

## 3. GitHub CLI extension contract

Official documentation requires an extension repository name beginning `gh-`.
For this project `gh-sdp`, the extension name is `sdp`; `gh sdp <args>` executes
the `gh-sdp` extension and forwards `<args>`. A script extension has an
executable named exactly like the repository at repository root; a precompiled
extension instead supplies platform binaries as release assets. An extension
cannot replace a core `gh` command. [verified external fact; STU-EVD-002]

| Operation | Documented behavior | 2.96.0 source-derived detail |
| --- | --- | --- |
| `gh extension install OWNER/gh-sdp` | Install remote repository; release artifacts are tried as binary extension, otherwise repository is cloned as script extension; `--pin` means release tag for binary or commit for script | Latest-release endpoint and recognized asset suffixes determine binary mode; a missing current-platform asset in a binary release fails rather than building source. [STU-EVD-002, STU-EVD-003] |
| `gh extension install .` | Local development pointer/symlink targets a root executable named like the repository; a precompiled binary must first be built there | Windows uses an equivalent pointer file, but 2.96.0 local validation appears to check extensionless `gh-sdp` while released Windows assets require `.exe`; the exact precompiled Windows local flow is unresolved pending an executable test. [STU-EVD-002, STU-EVD-003] |
| `gh extension upgrade sdp` | Upgrade installed client extension; pinned/local extensions are not ordinary upgrade candidates | Binary mode installs the latest release; script mode pulls the repository. [STU-EVD-002, STU-EVD-003] |
| `gh extension remove sdp` | Remove the installed extension | 2.96.0 removes its extension directory; it does not uninstall SDP project content. [STU-EVD-002, STU-EVD-003] |

The 2.96.0 manager connects standard input/output/error and forwards arguments
after the extension name. It does not set `exec.Cmd.Dir`; Go specifies that an
empty `Dir` inherits the caller's current directory. Therefore inheritance of
the `gh` process working directory is a **source-derived fact for 2.96.0**, not
a separately documented permanent extension contract. [verified external fact;
STU-EVD-003, STU-EVD-007]

On Windows a release binary must end in `.exe`. The documented local-install
wording and the 2.96.0 extensionless local source check do not establish whether
repository-root `gh-sdp.exe` alone works with `gh extension install .`.
Therefore the Windows local precompiled workflow is an unresolved question,
not a claimed contract. Directly running a built `gh-sdp.exe` with controlled
arguments/CWD remains a useful debug seam; later work must test local install
against the supported GitHub CLI versions. [verified external fact,
source observation, recommendation; STU-EVD-002, STU-EVD-003]

Remote GitHub Enterprise installation can use a full repository URL. Host and
credentials can be selected with `GH_HOST` and enterprise token variables;
support for alternate hosts must preserve host-qualified repository identity
and never silently redirect to `github.com`. [verified external fact and
recommendation; STU-EVD-002, STU-EVD-004]

**STU-FND-002 — viable extension mechanism.** GitHub CLI provides the required
repository, native-binary, local-development, argument, and platform mechanisms,
so `MAN-ASM-001` remains viable. CWD inheritance and downloader details still
need version-pinned conformance tests. [inference; STU-EVD-002, STU-EVD-003]

## 4. Distribution and release assets

Precompiled assets conventionally use the full executable name plus
`OS-ARCHITECTURE`, for example `gh-sdp-windows-amd64.exe`,
`gh-sdp-linux-amd64`, `gh-sdp-linux-arm64`, `gh-sdp-darwin-amd64`, and
`gh-sdp-darwin-arm64`. All five are recognized by GitHub CLI 2.96.0.
[verified external fact; STU-EVD-002, STU-EVD-003]

GitHub CLI first queries the latest release. Its current source selects an asset
by platform suffix and downloads the API asset body directly. The release asset
structure used by this path contains name and API URL, not the REST `digest`
field, and the installer performs no checksum-file or digest verification.
Thus GitHub CLI transport/download must not be represented as automatic
`gh-sdp` binary integrity verification. [verified external fact for 2.96.0;
STU-EVD-003, STU-EVD-005]

If no release exists, remote installation clones the repository as a script
extension; it does not compile Go. A normal precompiled Go repository with no
root binary will consequently not become runnable through that fallback. If a
release signals binary mode but lacks the user's matching asset, installation
fails. [verified external fact/source-derived; STU-EVD-002, STU-EVD-003]

The latest-release API excludes drafts and prereleases and currently selects by
`created_at`, not highest SemVer. During remote install, GitHub CLI 2.96.0 first
calls `isBinExtension`, whose binary detector inspects the latest stable release
before the requested `--pin` tag. Only after that classification succeeds does
the pinned binary path fetch the requested release tag. A prerelease-only
precompiled repository with no root script therefore receives no latest stable
release, falls through to script detection, and is rejected as not installable;
`--pin <prerelease-tag>` does not rescue it. This is a pinned-source fact, not an
ambiguity. A version-pinned conformance fixture remains useful to detect future
GitHub CLI changes and to characterize repositories that have both a stable
binary release and a pinned prerelease. These rules apply to client-extension
distribution; Toolkit selection needs its own policy. [verified external
fact/source-derived behavior; STU-EVD-002, STU-EVD-003, STU-EVD-005,
STU-EVD-018]

`gh extension upgrade sdp` updates the installed **client**. Future
`gh sdp update` intends to update **SDP Toolkit content in a selected project**.
Neither command may change the other's identity implicitly. [verified
repository fact and inference; STU-EVD-001, STU-EVD-003]

**STU-FND-003 — release prerequisite.** A GitHub Release containing every
claimed target asset is operationally required for dependable normal
precompiled installation. Client release checksums/attestations are a separate
project release concern and are not supplied by `gh extension install`.
[inference; STU-EVD-002, STU-EVD-003, STU-EVD-012]

## 5. Go technology direction

Go can emit one native executable, cross-target with `GOOS`/`GOARCH`, and avoid
a target C toolchain when the dependency graph supports `CGO_ENABLED=0`. Its
standard library covers HTTP, JSON, ZIP/tar parsing, process execution, path and
filesystem primitives, and testing. GitHub officially documents Go extension
scaffolding and `go-gh` offers GitHub CLI-compatible auth/host conventions.
[verified external fact; STU-EVD-002, STU-EVD-006, STU-EVD-007, STU-EVD-020]

Pure-Go feasibility remains conditional: no dependency set or platform adapter
has been selected, and no local build was possible in this Study. Filesystem
security, replacement, locking, and crash consistency still require deliberate
platform abstractions; standard-library availability is not proof of correct
behavior. [observation and inference; STU-EVD-007, STU-EVD-017]

| Target | Go / GitHub CLI evidence | Study assessment |
| --- | --- | --- |
| Windows amd64 | Go target and first-class port; `windows-amd64.exe` recognized | Initial target; needs NTFS/reparse/sharing/locking fixtures. [STU-EVD-003, STU-EVD-006] |
| Linux amd64 | Go target and first-class glibc port; suffix recognized | Initial target; test declared minimum distro/glibc and filesystem matrix. [STU-EVD-003, STU-EVD-006] |
| macOS amd64 | Go first-class port; `darwin-amd64` recognized | Initial target; test APFS case variants and signing/quarantine release concerns later. [STU-EVD-003, STU-EVD-006] |
| macOS arm64 | Go first-class port; `darwin-arm64` recognized | Initial target; native Apple Silicon is necessary for a current macOS promise. [STU-EVD-003, STU-EVD-006] |
| Linux arm64 | Go first-class glibc port; `linux-arm64` recognized | **Include initially**, subject to real arm64 CI/runtime evidence. [STU-EVD-003, STU-EVD-006] |

Alternatives remain credible but weaker against this Mandate: Rust can also
produce native binaries but adds a different ecosystem/toolchain and no
equivalent official `go-gh` reuse; self-contained .NET introduces a larger
runtime/distribution surface; interpreted Bash/PowerShell/Python shifts runtime
availability to users and conflicts with the no-forced-PowerShell constraint.
Those are comparative inferences, not measured artifact-size results. A later
Architecture decision must compare reproducible builds, dependency audit,
binary size, startup, platform APIs, and team operability with prototypes.
[inference; STU-EVD-002, STU-EVD-006]

**STU-REC-001 — Go and targets.** Retain Go as the preferred direction, target a
pure-Go (`CGO_ENABLED=0`) native binary where evidence permits, and include all
five targets above—explicitly Linux arm64—in the initial release matrix. Pin a
patched toolchain and require build plus native runtime tests before support is
claimed. This is a Study recommendation, not Architecture approval.
[recommendation; STU-EVD-006, STU-EVD-007]

## 6. GitHub authentication and API behavior

`GH_TOKEN` takes precedence over `GITHUB_TOKEN` for `github.com`/`ghe.com`;
enterprise-server variables apply to GHES, and `GH_HOST` supplies an otherwise
unknown host. Stored login normally uses the system credential store, with a
documented plaintext fallback. `gh auth token` prints a credential to stdout;
it should not be used merely to move a token between processes when an API
client can follow GitHub CLI auth conventions. [verified external fact and
recommendation; STU-EVD-004]

`gh api` makes authenticated requests and can cache responses. Direct REST
requests for public release data may be unauthenticated; the documented primary
limit is 60 requests/hour per source IP unauthenticated and ordinarily 5,000
requests/hour authenticated. Private release/asset reads require suitable
repository Contents read access. Secondary limits and host policy also apply.
[verified external fact; STU-EVD-004, STU-EVD-005]

| Approach | Benefits | Costs/risks | Evidence still needed |
| --- | --- | --- | --- |
| Spawn `gh api` | Reuses installed CLI host/auth behavior; easy command fake; no token value handled by extension | Process overhead; `gh api` requires auth even where direct public HTTP could work; output/error/version coupling | Exact exit/error/cache/host behavior across supported `gh` versions. [STU-EVD-004, STU-EVD-007] |
| Direct client following `go-gh` conventions | Reuses env/stored OAuth without printing token; direct status/rate headers; in-process test transport | Adds dependency and credential/config code path; must preserve GHES and scope semantics | Pin/audit library, inject transport/auth provider, test env precedence and credential non-disclosure. [STU-EVD-020] |
| Direct unauthenticated HTTP for public default | Works before login and uses public 60/hour budget | Lower shared-IP limit; cannot access private/alternate repos; must add host/TLS/cache policy | Rate/error/redirect tests and explicit upgrade to auth. [STU-EVD-005] |

Local `status` should first read project manifests and report installed facts
without network. Network-derived fields must be optional and label repository,
host, retrieval time, cache age, and stale/offline state. A cache may serve
immutable metadata only after validating its repository/version/digest identity;
staleness must never be hidden. [recommendation; STU-EVD-004, STU-EVD-005,
STU-EVD-019]

Test seams should inject API transport, clock, cache, auth/host resolver, and
subprocess runner; deterministic fixtures and local HTTP servers should cover
authenticated/unauthenticated, rate-limit, redirects, timeout, stale cache,
private/alternate repository, and GHES cases without a live account.
[recommendation]

**STU-FND-004 — auth is reusable but architecture is open.** Reusing GitHub CLI
identity is feasible by subprocess or compatible client. Direct client access
is the Study preference because it avoids token stdout and permits public
unauthenticated fallback, but the final API/auth architecture remains for later
evidence and approval. [inference/recommendation; STU-EVD-004, STU-EVD-005,
STU-EVD-020]

## 7. Actual SDP Toolkit installation contract

The authoritative Toolkit manifest identifies Toolkit/installer 0.2.0,
supported project-manifest schema `1.0`, migration metadata, and an unreleased
state with null tag and release commit. Normative prose identifies
`SDP.manifest.yaml` as the release manifest, classifies ownership, requires
preview, preserves project-owned content, backs up replaced managed files,
stops on unsupported schemas, refuses downgrade, and states same-install
idempotence. [verified repository fact; STU-EVD-008]

The table separates that normative contract from the current PowerShell
implementation at `bc110…`:

| Area | Actual behavior | Classification/evidence |
| --- | --- | --- |
| Fixed records and parsing | Project and installed facts are read only from hardcoded `SDP/SDP-project.manifest.yaml` and `SDP/Framework/installed-toolkit.manifest.yaml`. `Get-YamlScalar` uses a regular expression for selected scalar names; the installer does not parse the whole YAML document or validate it against the published JSON schemas. A missing installed manifest is unconditionally classified as the supported pre-versioning baseline, without first proving what legacy layout is present. Present manifests with missing/malformed required scalars or unsupported schema strings stop before the later write phase. | current implementation, not a complete schema/legacy contract; STU-EVD-009, STU-EVD-010 |
| Toolkit version/downgrade | Reads installed `toolkitVersion`; validates SemVer-shaped text but compares only major/minor/patch, ignoring prerelease precedence; rejects installed core version newer than 0.2.0. The root manifest's `compatibilityRange` has no documented subject or evaluated range grammar in this installer. | current implementation and unresolved contract meaning; STU-EVD-008, STU-EVD-009, STU-EVD-018 |
| Ownership | Managed: root `AGENTS.md`, `.codex/skills/sdp-*/SKILL.md`, `SDP/Framework/*`. Project-owned: `AGENTS-project.md`, reminders, project manifest, release notes, lifecycle, release/review/verification/traceability/instructions records. The script creates/refreshes current managed sources but has no inventory tombstones or deletion pass for files that were managed by an older Toolkit and are absent from the new payload. | normative ownership classes plus current additive implementation; STU-EVD-008, STU-EVD-009 |
| Preservation/force exceptions | Project-owned differing files are preserved, and most same-version differing managed files are preserved unless `-ForceManagedFiles`; upgrade or force refreshes those files. However, root `AGENTS.md` is always refreshed after any migration/preservation copy, and the installed manifest is always regenerated with managed refresh enabled. These are exceptions to broad same-version-preservation prose. `ForceManagedFiles` never authorizes project-owned overwrite. | normative intent with verified implementation exceptions; STU-EVD-008, STU-EVD-009, STU-EVD-010 |
| AGENTS migration | A differing prior `AGENTS.md` is copied to missing `AGENTS-project.md`; if that destination exists, old content is copied to timestamped `AGENTS-project.migration-*.md`; managed `AGENTS.md` is then installed without the ordinary managed-file backup path. | current additive migration implementing preservation with a special replacement path; STU-EVD-008, STU-EVD-009, STU-EVD-010 |
| Backups and recovery | Changed managed files are copied sequentially before their replacement. The default is `SDP/.sdp-backups/<UTC-to-seconds>/`; `-BackupRoot` may resolve to an absolute location outside the project. There is no lock, journal, transaction boundary or rollback. Second-resolution default names plus forced copy create a collision/overwrite risk for concurrent or repeated same-second runs. | current implementation plus direct risk inference; STU-EVD-009 |
| Preview, plan and output | Covered fixtures prove `-Preview` writes nothing. Preview reruns installer control flow and emits `PROPOSED` human text; it does not create a reusable plan object, digest or input binding, and tests do not prove preview/apply action-for-action equivalence. Skills are sorted, but framework and initializer recursive enumeration is not explicitly sorted; the remaining fixed operations impose only partial ordering. Reporting is `Write-Host` human output with category/count summaries and no JSON contract. | current implementation/test fact plus explicit unproven guarantees; STU-EVD-009, STU-EVD-010 |
| Idempotence/timestamp | Normal same-version fixture repeat preserves `toolkitInstalledAt` and covered tree content. This evidence is limited to the tested base-install tree and does not cover changed ambient source identity, obsolete managed files, initializer repeat, concurrency or same-second backups. | normative claim with bounded test evidence; STU-EVD-008, STU-EVD-009, STU-EVD-010, STU-EVD-021 |
| Project/Toolkit path guard | `Resolve-Path`/`GetFullPath` strings are compared lexically with `OrdinalIgnoreCase` for equality or repository-prefix membership; a sibling such as `SDP-Analyzer` is accepted and a lexical child is rejected. The guard does not resolve and compare final symlink/junction destinations, so the tests prove complete string-segment handling, not link-aware containment. | current implementation and bounded test evidence; STU-EVD-009, STU-EVD-010 |
| Source and installed identity | `sourceCommit` is obtained from ambient `git -C <ToolkitRepositoryRoot> rev-parse HEAD`, not from a verified release/payload identity. Installed facts contain that commit and version/capability fields but no source owner/repository/host, release asset/digest, plan/contract identity or payload manifest hash. Root and installed capability sets differ: installed facts omit root-manifest `sdp.release-notes.v1` and `sdp.skill-metadata.v1`, without a documented projection rule. | current implementation and manifest comparison; STU-EVD-008, STU-EVD-009 |
| Base traceability and structure initialization | Base install copies the Toolkit repository's live `Traceability/CurrentIndex.yaml` and `Relations.yaml` as project-owned files when missing, so Toolkit-specific lifecycle state can seed a consumer. `-InitializeProjectStructure` additionally copies project-owned lifecycle/instruction files and `SDP-DOCUMENT-GUIDE.md`, including Toolkit-specific `Releases/REL-0.2.0.yaml`, only when missing. | current implementation/payload behavior; STU-EVD-009, STU-EVD-010, STU-EVD-021 |

The base installation payload contains managed AGENTS/Framework templates and
skills, plus missing project manifest, release notes, empty ledger, and the live
Toolkit repository's traceability seeds. It does **not** install
`Toolkit/scripts`, tests, schemas, root `SDP.manifest.yaml`, upstream `.git`, or
other development-only repository files. The initializer broadens that copy to
canonical lifecycle seeds, including Toolkit-specific
`Releases/REL-0.2.0.yaml`; none is a nested Git repository. [verified repository
fact; STU-EVD-009, STU-EVD-010]

Fixture coverage proves preview zero mutation; empty install and normal repeat;
project-owned preservation; force-managed backup/restore; legacy AGENTS and
requirements preservation; unsupported installed/project schemas; downgrade
rejection; and sibling-versus-child path handling. It does not prove concurrent
invocation, crash recovery, archive safety, hostile links, transactional
rollback, full YAML schema validation by the script, or initializer repeat.
[verified repository fact and observation; STU-EVD-010]

Actual project evidence shows a normal repeat proposed zero content changes,
while repeating the initializer in preview proposed creating
`SDP/Releases/REL-0.2.0.yaml`. This is a known upstream payload/idempotency gap.
It must remain documented and unfixed here; `gh-sdp` must not create that false
release record for a consuming project. [verified repository fact;
STU-EVD-021]

**STU-FND-005 — contract is evidenced but neither complete nor reusable.**
Toolkit 0.2.0 exposes clear ownership/migration intent plus a tested PowerShell
implementation, but no versioned machine-readable plan, ownership inventory,
portable engine, deletion/tombstone model or bound preview/apply contract. Its
regex scalar checks, unconditional missing-manifest legacy assumption,
core-only SemVer comparison, lexical non-link-aware containment, preservation
exceptions, unsorted/sequential human-only actions, ambient `sourceCommit`,
identity/capability gaps and Toolkit-specific consumer seeds are current
implementation behavior—not safe new normative semantics. A portable client
must fail closed rather than infer the missing contract. [verified repository
fact/inference; STU-EVD-008, STU-EVD-009, STU-EVD-010, STU-EVD-021]

## 8. Installer-rule ownership and drift prevention

| Alternative | Source of truth / platforms / PowerShell | Drift, conformance, schema and release coupling | Complexity, security, offline and failure modes | SDP implications |
| --- | --- | --- | --- | --- |
| **A. gh-sdp cross-platform engine + conformance** | Toolkit prose/fixtures remain owner; Go engine supports Windows/Linux/macOS without PowerShell | Medium-high drift: two interpretations; requires fixtures for every supported Toolkit release and explicit capability/schema negotiation | Medium build complexity; strong local control, but security fixes can diverge; offline works with cached payload; mismatch may silently produce a different plan unless fail-closed | Toolkit must publish contract fixtures/payload; gh-sdp must release for rule changes |
| **B. Toolkit machine-readable plan/contract consumed by gh-sdp** | SDP owns versioned ownership, payload inventory, compatibility and plan semantics; gh-sdp supplies portable apply engine | Lowest practical semantic drift if the plan is canonical and fixtures bind producer/consumer; schema evolution and minimum client capability are explicit | Upfront schema/planner work; smaller duplicated policy; offline works from verified asset; malformed/unsupported plan fails before mutation | Requires separately authorized SDP change and release asset; can preserve PowerShell as one consumer |
| **C. PowerShell when present + independent fallback** | PowerShell stays canonical on some hosts; fallback becomes second rule owner | Highest behavioral split and test matrix; fallback drift and host-dependent outcomes; releases coupled to both engines | Low initial Windows work but forced/runtime PowerShell elsewhere, duplicate attack/failure surfaces, inconsistent offline behavior | Conflicts with `MAN-BND-006`; still requires portable fallback and cross-engine fixtures |
| **D. Generate engines from one canonical specification** | SDP owns a generator/spec; generated PowerShell/Go can cover all platforms | Potentially low drift, but generator/spec/tool versions become release contracts; generated output and hand-edits need enforcement | Highest initial and maintenance complexity; generator compromise affects all engines; offline generation/runtime policy must be defined | Major upstream toolchain/schema addition; useful only if declarative plan cannot express required operations |

All alternatives need authenticated/verified release input, unsupported-schema
failure, deterministic fixtures, explicit compatibility, hostile-input defense,
and a way to reproduce the selected Toolkit's exact plan. A cannot truthfully
claim conformance until the Toolkit owns fixtures; C preserves the current
PowerShell dependency and creates host-dependent semantics; D can be revisited
if B's plan cannot express migrations safely. [inference; STU-EVD-005,
STU-EVD-008, STU-EVD-009]

**STU-REC-002 — canonical plan.** Prefer **B**: `Hans-Einar/SDP` should publish a
versioned, machine-readable installation-plan contract, payload inventory,
ownership/capability metadata, and conformance fixtures; a cross-platform
`gh-sdp` engine should validate and apply that contract. A is a possible
time-boxed bridge only when bound to canonical fixtures. Do not select B as
Architecture until the upstream contract exists and a prototype proves it can
represent migration, preview, and preservation. [recommendation; STU-EVD-008,
STU-EVD-010]

**STU-FND-006 — upstream prerequisite.** The preferred alternative requires a
separately authorized change and release in `Hans-Einar/SDP`; this repository
cannot implement or declare Toolkit compatibility before that dependency is
resolved. [inference; STU-EVD-011]

## 9. Toolkit release discovery and selection

The REST latest-release endpoint means “most recent non-draft,
non-prerelease by creation time,” not “highest compatible SemVer.” Therefore a
stable-by-default client should enumerate/validate candidates or verify that the
latest endpoint result is SemVer-valid and compatible; it must not infer
compatibility from recency. Exact `--version` selection should use full SemVer
2.0 precedence, match a release/tag/manifest, and reject downgrade unless the
later approved policy provides an explicit recovery path. [verified external
fact and recommendation; STU-EVD-005, STU-EVD-018]

Prereleases, alternate owner/repository, forks, and private repositories should
be explicit trust modes. The selected host/repository must flow through auth,
cache keys, provenance policy, status output, and audit evidence. Private assets
need Contents read permission; public assets can be fetched unauthenticated.
[verified external fact/recommendation; STU-EVD-004, STU-EVD-005]

Generated source archives have stable commit contents but their compressed
bytes can change and they are not purpose-built installation payloads. Prefer a
versioned Toolkit release asset containing only the installation payload and
machine-readable contract. Validate release tag, asset identity, manifest
Toolkit version/release state/capabilities/schema, compatibility declaration,
and verified digest before extraction. [verified external fact/recommendation;
STU-EVD-008, STU-EVD-012, STU-EVD-013]

There is no current stable Toolkit release to select. Normal install/update must
therefore fail truthfully with an actionable “no eligible stable release” state;
using `bc110…` for development must remain an explicit unreleased source choice,
not a default. [verified repository fact/recommendation; STU-EVD-011]

Downloads should use a uniquely created temporary location, stream with size
limits and hash computation, close/sync before validation, and publish to cache
only after complete verification. On timeout, non-success status, short body,
redirect-policy failure, digest mismatch, or cancellation: stop before project
mutation, delete untrusted partials, and retain only a non-secret diagnostic.
Cache entries need host/repo/version/asset/digest keys and age; corrupt or
unverifiable entries are misses. [recommendation]

### Intended option inputs, not final semantics

| Option | Study input and unresolved semantics |
| --- | --- |
| `--project-root <path>` | Explicit root; otherwise CWD. Resolve without escaping trust boundary, reject Toolkit-repository nesting, define symlink/root identity and non-existent-root policy. [STU-EVD-009] |
| `--version <semver>` | Exact Toolkit release selection using full SemVer; prerelease opt-in and downgrade/recovery interaction remain Requirements choices. [STU-EVD-018] |
| `--repository <owner/repo>` | Explicit alternate trust source; hostname syntax, allowlist, private auth, provenance policy and cache namespace remain open. [STU-EVD-004, STU-EVD-005] |
| `--force-managed-files` | Directionally maps to restoring changed managed files with backup; scope must come from selected Toolkit contract, never client hard-coding. [STU-EVD-008, STU-EVD-009] |
| `--initialize-project-structure` | One-time additive bootstrap with preview; must not reproduce the known upstream `REL-0.2.0` seed gap. [STU-EVD-021] |
| `--json` | Machine-readable, versioned result/plan/error model; exact schema and stdout/stderr rules are unresolved. |
| `--verbose` | Diagnostic detail without credentials or sensitive headers/paths; interaction with JSON and redaction remains unresolved. |

**STU-REC-003 — release selection.** Use an eligible stable, purpose-built,
immutable Toolkit asset by default; allow exact SemVer pins and other repositories
only explicitly; fail closed on absent release, incompatible capability/schema,
identity mismatch, or unverifiable input. This remains candidate policy.
[recommendation; STU-EVD-005, STU-EVD-011, STU-EVD-012, STU-EVD-013]

## 10. Provenance and integrity

| Property | Mechanism and boundary |
| --- | --- |
| Transport security | HTTPS protects a connection to the authenticated host and data in transit. It does not by itself prove which source commit/workflow produced an asset or that the publisher intended a particular digest. [inference; STU-EVD-005] |
| Integrity | SHA-256 checksum, REST asset `digest`, or manifest-declared file hashes detect byte changes only when the expected value is obtained through a trusted, identity-bound channel. [verified external fact/inference; STU-EVD-005, STU-EVD-012] |
| Authenticity | A verified signed tag identifies a signer under a key/identity policy; release immutability prevents later tag/asset movement. Neither proves code safety. [verified external fact; STU-EVD-012] |
| Provenance | GitHub artifact attestations cryptographically link an artifact to repository/workflow/commit and provide SLSA v1 Build Level 2; public attestations use Sigstore transparency. They must be verified and explicitly do not guarantee the artifact is secure. [verified external fact; STU-EVD-012] |

GitHub exposes a `sha256:` digest for release assets and can verify immutable
release assets, but GitHub's generated source ZIP/tar cannot be verified with
`gh release verify-asset`. A standalone checksum file stored beside a mutable
asset does not add authenticity unless the checksum is signed/attested or
anchored in an immutable trusted manifest. [verified external fact and inference;
STU-EVD-012, STU-EVD-013]

**STU-REC-004 — layered verification.** Require immutable purpose-built assets;
verify asset SHA-256 against API and a manifest-declared payload hash; verify
manifest/release/tag/version agreement; and prefer a verified GitHub attestation
bound to the expected repository, workflow, and commit. Define trusted signer,
repository and workflow policy later. Any mismatch, truncation, duplicate
manifest, missing digest, or failed attestation must stop before extraction.
[recommendation; STU-EVD-005, STU-EVD-012]

**STU-FND-007 — provenance gap.** SDP Toolkit 0.2.0 has no published asset,
checksum, immutable release, signed release tag, or attestation to verify, so no
cryptographic provenance claim or support declaration is currently possible.
[verified repository fact; STU-EVD-011]

## 11. Archive and path security

Archive names and selected project paths must be treated as attacker-controlled.
Lexical `Clean`/prefix checks are insufficient when links or concurrent renames
exist. Go's root-scoped APIs improve traversal resistance, but
`GO-2026-4970` allowed a trailing-slash final-symlink escape before Go 1.25.12
and in Go 1.26.0–1.26.4; a patched toolchain (1.26.5 in that line) plus a
regression fixture is required. [verified external fact; STU-EVD-007]

Go ZIP readers report non-local/backslash names under
`GODEBUG=zipinsecurepath=0`, with documentation noting that default behavior may
change later. Security policy must therefore validate every name explicitly and
must not depend on process environment. Tar additionally represents hard links,
symlinks, devices, FIFOs, sparse/extended entries, creating a larger rejection
surface. [verified external fact; STU-EVD-007]

### Testable later-requirement invariants

1. Parse archive names as `/`-separated protocol paths before OS conversion;
   reject empty, NUL/invalid UTF-8, absolute, rooted, `..`, backslash or mixed
   separators, Windows drive/UNC/device namespaces, trailing-space/dot and
   reserved Windows components. [recommendation; STU-EVD-007, STU-EVD-014]
2. Normalize only under a documented Unicode policy and reject duplicate,
   case-fold, and normalization-equivalent destination collisions before any
   write; preserve original legal names only after the portable-collision pass.
   [recommendation; STU-EVD-014, STU-EVD-016]
3. Accept only declared regular files/directories. Reject archive symlinks,
   junction/reparse semantics, hard links, devices, FIFOs, sparse/unknown entry
   types, alternate data-stream names, and unexpected executable/permission
   bits. [recommendation; STU-EVD-007, STU-EVD-014]
4. Inspect the complete central directory/header stream before extraction;
   reject duplicate paths, undeclared payload files, impossible sizes, excessive
   entry count, per-file size, total uncompressed size, compression ratio, depth,
   and path length. Enforce streamed byte limits even when headers lie.
   [recommendation]
5. Extract only into a newly created trusted staging root. Use traversal-resistant
   root/handle-relative operations, refuse existing link/reparse components, and
   never form a trusted destination by string prefix. Revalidate selected
   project identity immediately before apply. [recommendation; STU-EVD-007,
   STU-EVD-014]
6. No archive operation may write outside staging; no project apply may write
   outside the resolved project root or explicit backup root. Cleanup must not
   follow links. Tests must race link/rename replacement to exercise TOCTOU
   defenses. [recommendation; STU-EVD-007]

ZIP is the narrower preferred candidate for a purpose-built payload because it
avoids tar's native link/device entry variety, but ZIP is not intrinsically
safe. An even narrower manifest plus individual verified files could avoid
general archive semantics at higher request/packaging cost. Final format remains
an Architecture decision. [inference/recommendation; STU-EVD-007]

**STU-FND-008 — extraction is a security boundary.** Safe extraction requires
format-independent policy, complete preflight, explicit limits and collisions,
patched root-scoped writes, and hostile filesystem tests; library defaults alone
do not satisfy `MAN-SC-006`. [inference; STU-EVD-007, STU-EVD-014,
STU-EVD-016]

## 12. Transaction, backup and recovery

The current PowerShell installer detects schema/version before mutation and
backs up replaced managed files, but then creates/replaces files sequentially.
It has no complete immutable plan, plan digest, lock, staging transaction,
rollback journal, crash recovery protocol, or tree-wide commit. [verified
repository fact; STU-EVD-009]

**Recommended operation protocol (not approved design):**

1. Resolve/validate project, selected Toolkit, verified payload and compatibility;
   acquire a project-scoped exclusive lock with owner/time metadata.
2. Read a consistent inventory and build the entire deterministically ordered
   action plan (`create`, `replace-managed`, `preserve`, `backup`, `warning`,
   `error`) before mutation; hash canonical plan plus inputs.
3. Preview emits that exact plan/identity/hash and performs zero project, backup,
   cache or lock-persistent mutation. Apply recomputes and rejects changed inputs
   rather than applying a stale preview.
4. Stage new bytes and journal on the destination filesystem; validate hashes,
   sizes, permissions, free space and parent/link state. Sync staged files where
   the promised durability level requires it.
5. Create backups before each destructive replacement, recording original
   identity/hash/metadata and every committed action in a durable append-only
   journal. Never overwrite an earlier recovery set.
6. Commit in deterministic dependency order with platform replacement adapters.
   On failure/interruption, stop, journal exact state, attempt bounded reverse
   restoration, and report `unchanged`, `applied`, `rolled-back`, or
   `recovery-required` truthfully.
7. Verify resulting inventory/manifest, sync applicable parent directories,
   mark journal complete, release lock, then clean trusted staging. Retain or
   prune backups only under explicit policy.

[recommendation; STU-EVD-014, STU-EVD-015]

POSIX same-filesystem `rename` provides atomic directory-entry replacement, but
atomic does not imply durable; durability may require syncing the new file and
directory. Cross-filesystem rename fails. Go documents that `os.Rename` is not
atomic even within one directory on non-Unix platforms. Windows `ReplaceFileW`
requires the same volume, can preserve metadata/create a backup, and documents
partial failure states. Apple likewise recommends destination-volume temporary
storage for replacement. [verified external fact; STU-EVD-007, STU-EVD-014,
STU-EVD-015, STU-EVD-016]

| Platform | Practical claim ceiling |
| --- | --- |
| Linux/Unix | Per-file same-filesystem atomic replacement can be offered after file/directory sync policy and tests; a multi-file tree is not one atomic transaction. [STU-EVD-015] |
| macOS | Same-volume per-item safe replacement is feasible; APFS capabilities help but non-APFS volumes and multi-item updates require journal/recovery. [STU-EVD-016] |
| Windows | Use a Windows replacement adapter and handle sharing/AV interference plus documented partial states; do not claim generic `os.Rename` atomicity. [STU-EVD-007, STU-EVD-014] |

Retries are safe only before mutation or for operations proved idempotent from
journal state; blind retry after an ambiguous replacement can destroy recovery
evidence. Lock acquisition must distinguish active owner, stale/crashed owner,
and explicit recovery. Exit/output must include plan identity, phase, changed
paths/counts, backup/journal locations, rollback result, and next safe action,
while redacting credentials and sensitive headers. [recommendation]

**STU-REC-005 — transactional honesty.** Require deterministic preflight,
preview/apply equivalence, same-filesystem staging, backups, durable journal,
project lock, platform replacement adapters and explicit recovery state. Promise
best-effort rollback with evidenced outcomes, never a universal atomic
whole-tree transaction. [recommendation; STU-EVD-009, STU-EVD-014,
STU-EVD-015, STU-EVD-016]

**STU-FND-009 — current recovery gap.** Toolkit 0.2.0's backup-first sequential
installer supplies useful preservation behavior but not the transaction,
concurrency, or crash guarantees needed for a future cross-platform support
claim. [verified repository fact/inference; STU-EVD-009, STU-EVD-010]

## 13. Cross-platform filesystem behavior

| Dimension | Windows | Linux / POSIX target | macOS | Proposed portable handling |
| --- | --- | --- | --- | --- |
| Paths, separators, roots | Drive-relative, drive-rooted, single-rooted, UNC and device namespaces differ; `\` separates components and many characters/names are reserved. [verified external fact; STU-EVD-014] | `/` root/components and POSIX rename rules apply; mounted filesystems can differ. [verified external fact; STU-EVD-015] | `/` paths; Apple advises URL/component APIs and assuming case sensitivity. [verified external fact; STU-EVD-016] | Parse archive paths as a platform-neutral protocol; reject all rooted/drive/UNC/backslash forms before OS conversion. |
| Case and preservation | Normally case-insensitive/case-preserving, but per-directory behavior can vary. [verified external fact; STU-EVD-014] | Common supported filesystems are case-sensitive, but this Study has not established one universal Linux-filesystem rule. [observation/limitation] | APFS is case-insensitive by default but configurable case-sensitive and preserves case. [verified external fact; STU-EVD-016] | Reject portable case collisions before staging; test actual target volume capabilities. |
| Unicode | Windows namespace/reserved-character rules apply. [verified external fact; STU-EVD-014] | Filesystem byte/name policy varies; no universal Unicode normalization assumption is safe. [inference] | APFS accepts valid UTF-8, preserves normalization, and modern APFS is normalization-insensitive. [verified external fact; STU-EVD-016] | Require valid UTF-8, defined normalization comparison and collision rejection without silently renaming. |
| Links | NTFS hard links, junctions, symlinks and other reparse points require explicit detection. [verified external fact; STU-EVD-014] | Symlinks/hard links can redirect access; handle-relative traversal is required under hostile-local-filesystem threat. [verified external fact/inference; STU-EVD-007] | Unix links plus mount/filesystem variation apply. [inference; STU-EVD-016] | Reject archive links and existing destination link/reparse traversal; use root-scoped handles. |
| Permissions/executable bit | Go's portable chmod surface uses only owner-writable meaning for Windows. [verified external fact; STU-EVD-007] | Unix mode bits are meaningful. [verified external fact; STU-EVD-007] | Unix mode bits are meaningful, subject to volume/mount policy. [inference] | Payload manifest allowlists mode intent; never trust archive mode; do not infer Windows ACL equivalence. |
| Locking | Sharing modes and external scanners can deny replacement; exact lock primitive/policy is unresolved. [risk/observation] | Advisory-lock behavior and filesystems vary; exact primitive is unresolved. [unresolved question] | Unix/APFS and external processes vary; exact primitive is unresolved. [unresolved question] | Platform lock adapter plus process metadata, contention/stale-lock fixtures, and no unsupported universal guarantee. |
| Rename/replacement | Go says non-Unix `Rename` is not atomic; `ReplaceFileW` is same-volume and has partial failure states. [verified external fact; STU-EVD-007, STU-EVD-014] | Same-filesystem POSIX rename is atomic; atomicity is not durability and cross-filesystem can fail. [verified external fact; STU-EVD-015] | Same-volume replacement/safe-save is available; non-APFS volumes remain possible. [verified external fact; STU-EVD-016] | Stage on destination volume; platform adapter; journal every committed replacement. |
| Long paths/reserved names | Long-path behavior requires application/system conditions; device names and trailing dot/space remain hazards. [verified external fact; STU-EVD-014] | Limits vary by filesystem and mount. [unresolved question] | Limits vary by filesystem. [unresolved question] | Enforce conservative manifest/path/depth limits and test boundary values per target. |
| Temporary directories | Default Windows temp search is environment/system based and not guaranteed accessible. [verified external fact; STU-EVD-007] | `/tmp`/`TMPDIR` default is not necessarily the destination filesystem. [verified external fact; STU-EVD-007] | OS temporary storage may be a different volume; Apple provides a destination replacement directory. [verified external fact; STU-EVD-016] | Create untrusted download temp separately; create apply staging on the destination volume. |
| Line endings/hidden files | Windows attributes and text tooling can differ. | Dot names conventionally hide entries. | Dot names/Finder metadata can differ. | Treat payload as exact bytes, never translate line endings, enumerate hidden entries, and reject undeclared metadata. [recommendation] |
| Antivirus/indexer interference | Sharing/rename denial has been observed as an operational risk on Windows, but this Study has no controlled measurement. [observation/risk] | Indexers/security tools may race operations; not measured. [assumption] | Spotlight/security tools may race operations; not measured. [assumption] | Preserve journal state, return actionable sharing/busy diagnostics, and test injected transient failures rather than claim universal behavior. |

**STU-FND-010 — platform abstraction is mandatory.** Path policy can be shared,
but root handling, reparse/link inspection, locking, replacement, durability and
permission mapping need target-specific adapters and real filesystem tests.
Portable support cannot be inferred from cross-compilation. [inference;
STU-EVD-006, STU-EVD-007, STU-EVD-014, STU-EVD-015, STU-EVD-016]

## 14. Conformance and testing strategy inputs

Later testing should be layered and deterministic:

| Layer | Required coverage |
| --- | --- |
| Pure domain tests | Full SemVer including prerelease/build metadata; command/flag parsing; client/Toolkit identity; compatibility/capability selection; deterministic ordering; JSON schema; documented exit-state mapping |
| Project-selection fixtures | Current directory and explicit root; missing/file/non-empty roots; Windows and Unix syntax; drive/UNC/mixed separators; real/symlinked roots; sibling `SDP` and `SDP-Analyzer`; project inside Toolkit rejection |
| Release/API fixtures | Latest stable and exact pinned versions; prerelease-only precompiled rejection plus stable-latest/pinned-prerelease behavior; no releases/assets; authenticated and unauthenticated public access; private/alternate/GHES host; rate-limit, timeout, redirect, stale cache, corrupt cache and offline |
| Payload/security corpus | Safe ZIP; `../`, absolute, drive/UNC/backslash/mixed paths; duplicate/case/Unicode collisions; symlink/junction/hard-link/device/FIFO; existing destination link; file-count/size/ratio/depth limits; truncated/corrupt/partial download; digest/manifest/version mismatch; `os.Root` trailing-slash regression |
| Canonical installer fixtures | Empty/non-empty projects; legacy AGENTS/lifecycle migration; managed/project-owned preservation; force backup; normal-repeat idempotence; downgrade and unsupported schemas; initializer repeat and `REL-0.2.0` exclusion; preview zero mutation |
| Apply/recovery fault injection | Failure at every backup, journal, write, sync, replace, verify and cleanup step; interruption/cancellation; lock contention/stale lock; rollback success/failure; disk-full/permission/sharing failures; retry from each journal state |
| Native matrix | Windows amd64, Linux amd64/arm64, macOS amd64/arm64; case-sensitive and insensitive volumes where available; long paths; executable-bit/permission mapping; link and rename semantics |
| CLI acceptance | All intended command intents; CWD propagation; local extension ambiguity on Windows; stdout/stderr separation; JSON validity with verbose diagnostics; redaction; terminal/nonterminal and documented exit codes |

Every item above is a recommendation for later tests, not a test implemented by
this Study. It covers command parsing, CWD/explicit root, discovery/pins/auth,
safe extraction, manifest agreement, empty/non-empty/legacy projects,
preservation/backups/idempotence/downgrade/schema/preview, platform paths,
siblings, JSON/exit codes, cleanup/network/partial download and failed apply as
required by `SPS-002`. [recommendation]

Use injected clock, filesystem/root operations, platform replacement/lock,
random/ID source, cache, HTTP transport, auth/host resolver, subprocess runner,
signal/cancellation and output writers. Use local HTTP/TLS servers instead of
live GitHub for protocol cases. Keep malicious archives and failure schedules as
reviewed test data; produce golden canonical plans only from the Toolkit-owned
contract, with normalized paths/timestamps and a plan-schema version.
[recommendation]

Conformance ownership should be split: `Hans-Einar/SDP` owns versioned contract,
payload and golden input/expected-plan fixtures for each Toolkit release;
`gh-sdp` proves its parser/planner/apply engine against those fixtures on every
supported OS/architecture. Cross-implementation differential tests may compare
PowerShell and the portable consumer where both support the same contract, but
PowerShell output alone must not become an undocumented golden oracle.
[recommendation; STU-EVD-008, STU-EVD-010]

**STU-REC-006 — conformance gate.** No Toolkit version or platform should be
declared supported until immutable upstream fixtures, hostile-input corpus,
native target execution and failure/recovery tests pass against the exact
client/Toolkit pair. [recommendation; STU-EVD-006, STU-EVD-010]

**STU-FND-011 — current coverage is insufficient for product support.** The
pinned PowerShell fixtures establish valuable ownership/migration behavior but
not portable API, archive, concurrency, transaction or native matrix
conformance. [verified repository fact/inference; STU-EVD-010]

## 15. Requirements candidates

These are Phase 3 inputs only. “Open” names the unresolved choice; none is an
approved Requirement.

| ID | Candidate requirement | Mandate / evidence | Open choice / affected stakeholders |
| --- | --- | --- | --- |
| STU-RQC-001 | Every output shall distinguish client, installed Toolkit, selected Toolkit, latest eligible Toolkit, and real release state. | MAN-OUT-003, MAN-SC-003; STU-EVD-011, STU-EVD-019 | Output schema; maintainers, automation, reviewers |
| STU-RQC-002 | Commands shall default to CWD and accept an explicit project root while resolving one stable root identity and rejecting Toolkit-repository nesting/escape. | MAN-OUT-001/002, MAN-BND-007; STU-EVD-003, STU-EVD-009 | Symlink/nonexistent-root semantics; project owners |
| STU-RQC-003 | Client install shall publish precompiled assets for every supported target using GitHub CLI naming and keep `gh extension upgrade` separate from Toolkit update. | MAN-OUT-001/003/004; STU-EVD-002, STU-EVD-003 | Windows local-install ambiguity; users/release operators |
| STU-RQC-004 | Toolkit selection shall default to an eligible stable release, permit exact full-SemVer pins explicitly, and fail on no eligible release or incompatible schema/capability. | MAN-OUT-003/005; STU-EVD-005, STU-EVD-011, STU-EVD-018 | Prerelease/downgrade policy; Toolkit/client maintainers |
| STU-RQC-005 | Alternate/private/enterprise repositories shall be explicit and host-qualified throughout auth, cache, provenance and output. | MAN-OUT-003/006; STU-EVD-004, STU-EVD-005 | Supported host/repository scope; enterprise operators/security |
| STU-RQC-006 | Local status shall function without network; remote/cache fields shall identify source, retrieval time, age and stale/offline state. | MAN-OUT-001/003/006; STU-EVD-004, STU-EVD-019 | Cache TTL/refresh policy; users/automation |
| STU-RQC-007 | Credentials, auth headers and tokens shall never be printed, persisted in project/cache, or exposed to child processes unnecessarily. | MAN-OUT-006, MAN-SC-006; STU-EVD-004, STU-EVD-020 | Auth architecture; security/private-repo users |
| STU-RQC-008 | Before extraction, a Toolkit asset shall pass repository/release/tag/manifest/version/capability, SHA-256 and configured provenance checks. | MAN-OUT-003/005/006, MAN-SC-006; STU-EVD-005, STU-EVD-012 | Required signer/attestation policy; release/security reviewers |
| STU-RQC-009 | Extraction shall enforce the six hostile-path/archive invariants in section 11, explicit limits, a patched toolchain and zero writes outside trusted staging. | MAN-OUT-002/006, MAN-SC-006; STU-EVD-007, STU-EVD-014, STU-EVD-016 | ZIP versus non-archive format; project owners/security |
| STU-RQC-010 | The complete deterministic action plan shall exist before mutation; preview shall make zero mutation and apply shall reject changed inputs from the previewed plan. | MAN-OUT-002, MAN-SC-002; STU-EVD-009 | Plan serialization/hash/UI; users/automation |
| STU-RQC-011 | Project-owned files shall never be overwritten; managed replacement/force shall follow the selected Toolkit contract and create verified backups first. | MAN-OUT-002/005, MAN-BND-004; STU-EVD-008, STU-EVD-009 | Canonical machine-readable contract; project/Toolkit owners |
| STU-RQC-012 | Apply shall use a project lock, same-volume staging, durable journal, platform replacement adapter and truthful rollback/recovery state. | MAN-OUT-002/006, MAN-SC-002/006; STU-EVD-014, STU-EVD-015, STU-EVD-016 | Durability/retention/lock policy; operators/security |
| STU-RQC-013 | Unsupported/malformed installed or project schemas, downgrade, and manifest/version mismatch shall fail before project mutation. | MAN-OUT-002/003/005; STU-EVD-008, STU-EVD-009 | Explicit recovery/downgrade mode; project owners |
| STU-RQC-014 | Windows amd64, Linux amd64/arm64, and macOS amd64/arm64 shall pass native conformance before support; Linux/macOS shall require no PowerShell. | MAN-OUT-004, MAN-BND-006, MAN-SC-004; STU-EVD-003, STU-EVD-006 | Minimum OS/distro/filesystem; all users/release reviewers |
| STU-RQC-015 | Human and JSON output shall be deterministic, versioned, redacted and separate diagnostics from machine stdout; exits shall encode useful final state. | MAN-OUT-001/006; Study sections 6, 9, 12 | Exact JSON/exit/accessibility contract; automation/users |
| STU-RQC-016 | Every supported Toolkit release shall provide Toolkit-owned plan/payload/conformance fixtures; the exact client/Toolkit pair shall pass them and the security/recovery matrix. | MAN-OUT-005/006, MAN-SC-005/006/007; STU-EVD-008, STU-EVD-010 | Upstream format and ownership; both repositories/reviewers |
| STU-RQC-017 | Initialization shall be explicit, preview-first and additive, and shall never create an upstream Toolkit release record in a consuming project. | MAN-OUT-002/003; STU-EVD-021 | Corrected upstream seed/plan; project and Toolkit owners |

## 16. Architecture decision inputs

| ID | Later decision | Study recommendation | Remaining evidence before approval |
| --- | --- | --- | --- |
| STU-ADI-001 | Installer-rule owner | Toolkit-owned machine-readable plan consumed by portable engine (B) | Upstream schema/prototype expresses migration, ownership and preview; A–D review. [STU-REC-002] |
| STU-ADI-002 | API/auth approach | Prefer direct client following `gh` conventions, with optional unauthenticated public access; retain subprocess seam | Pinned library audit; GHES/private/env/stored-token tests; process alternative measurements. [STU-EVD-004, STU-EVD-020] |
| STU-ADI-003 | Release artifact format | Purpose-built immutable Toolkit asset, preferably narrow ZIP or manifest/files | Upstream payload proposal; request/size comparison; GitHub release prototype. [STU-REC-003, STU-REC-004] |
| STU-ADI-004 | Safe-extraction boundary | Verify and fully preflight in isolated staging, then apply only declared regular files with root-scoped operations | Patched Go prototype, hostile corpus and platform link-race tests. [STU-EVD-007] |
| STU-ADI-005 | Planning/apply boundary | Canonical immutable plan and plan hash; apply recomputes/revalidates | Machine-readable plan schema and preview UX/JSON experiment. [STU-REC-005] |
| STU-ADI-006 | Transaction model | Per-file same-volume replacement, backups, durable journal, lock and explicit recovery; no whole-tree atomic claim | Fault-injection prototype and native durability/replacement results. [STU-EVD-014, STU-EVD-015, STU-EVD-016] |
| STU-ADI-007 | Compatibility discovery | Release manifest declares client range/capabilities/schemas; stable default plus exact pins | Upstream compatibility model and version-skew/migration policy. [STU-EVD-008, STU-EVD-018] |
| STU-ADI-008 | Output model | Shared typed result feeding human and versioned JSON, with stdout/stderr/redaction rules | Requirements for accessibility, interactivity, fields and exit codes. |
| STU-ADI-009 | Test seams | Inject transport/auth/cache/clock/filesystem/platform/process/signal/output and use local servers/golden plans | Ownership of fixtures and representative native runners. [STU-REC-006] |
| STU-ADI-010 | Platform abstraction | Shared path/plan policy plus Windows and Unix/macOS lock/replacement/durability adapters | Minimum-platform decision and real NTFS/APFS/Linux filesystem experiments. [STU-FND-010] |

No row is an Architecture decision. The Architecture owner must record the
chosen alternative only after Requirements are approved and the remaining
evidence exists. [boundary; MAN-BND-008]

## 17. Cross-repository findings

The following are proposed inputs for a separately authorized
`Hans-Einar/SDP` effort. No upstream file or issue was changed/created.

| ID | Finding / requested upstream capability | Evidence and consequence |
| --- | --- | --- |
| STU-XRF-001 | Publish a purpose-built Toolkit release payload excluding repository/dev/test/tooling and Toolkit-specific consuming-project state seeds. | Current base install can seed live Toolkit `CurrentIndex`/`Relations`, the initializer copies `REL-0.2.0`, and release distribution is absent. STU-EVD-009, STU-EVD-010, STU-EVD-011, STU-EVD-021 |
| STU-XRF-002 | Own a versioned machine-readable ownership and canonical installation-plan format, including deterministic actions, deletions/tombstones, migration, backup and preview/apply binding. | Current truth is prose plus partially ordered PowerShell branches with no reusable plan, so a portable client would otherwise duplicate rules and could not identify obsolete managed files. STU-EVD-008, STU-EVD-009 |
| STU-XRF-003 | Publish immutable assets with manifest/file SHA-256 and verified attestation/signature policy. | Toolkit 0.2.0 has no release provenance objects. STU-EVD-011, STU-EVD-012 |
| STU-XRF-004 | Declare supported project/installed schemas; compatibility-range subject and grammar; capability projection; source repository/host; release asset/digest; payload/plan/contract identity; and migration prerequisites in the release contract and installed facts. | Root and installed capability sets differ without a projection rule; `compatibilityRange` meaning is undefined; ambient Git HEAD is recorded without origin, asset, digest or contract identity. STU-EVD-008, STU-EVD-009 |
| STU-XRF-005 | Own platform-neutral conformance inputs and golden expected plans for empty, non-empty, legacy, preservation, force, downgrade, unsupported schema and initializer scenarios. | Current fixtures are PowerShell behavior tests and omit portable plan output. STU-EVD-010 |
| STU-XRF-006 | Provide a consuming-project validator for full manifest schemas, evidenced legacy-layout detection, link-aware root identity, ownership/inventory and plan/result state without requiring the whole Toolkit source checkout. | Regex scalar extraction, unconditional missing-manifest legacy classification and lexical path comparison do not constitute that consumer contract. STU-EVD-009, STU-EVD-010 |
| STU-XRF-007 | Correct or explicitly contract Toolkit-specific state seeding, including base `CurrentIndex`/`Relations` and repeat-initializer `SDP/Releases/REL-0.2.0.yaml`. | Source inspection establishes base live-traceability copying; actual project preview reproduces the false release-record proposal. STU-EVD-009, STU-EVD-021 |

**STU-REC-007 — upstream coordination gate.** Obtain separate authority and SDP
maintainer agreement on `STU-XRF-001` through `STU-XRF-007` before an
implementation Slice claims canonical compatibility. [recommendation;
MAN-BND-002, MAN-ASM-002]

## 18. Open questions and conclusions

### MAN-001 question status

| Question | Status | Study answer/evidence | Decision owner and unresolved part |
| --- | --- | --- | --- |
| MAN-Q-001 — provenance/integrity | **Partially answered** | Layer immutable purpose asset, SHA-256/API+manifest hash, version agreement, signed/immutable release and verified attestation; no current Toolkit objects exist. STU-EVD-011, STU-EVD-012; STU-REC-004 | SDP maintainers and release/security reviewers must publish and approve signer/workflow/provenance policy. |
| MAN-Q-002 — compatibility/migration/rollback/skew | **Partially answered** | Canonical plan/capability contract, full SemVer, fail-before-mutation, journal/backups/recovery are candidate model. STU-EVD-008, STU-EVD-009, STU-REC-002, STU-REC-005 | Requirements/Architecture plus SDP maintainers must define compatibility ranges, migration graph, downgrade/recovery and retention. |
| MAN-Q-003 — offline/alternate/fork/prerelease/private | **Partially answered** | Stable canonical default; all other trust sources explicit/host-qualified; local status offline and cache labeled. STU-EVD-004, STU-EVD-005, STU-REC-003 | Steering/Requirements decide supported modes, syntax, cache/offline install and private/GHES scope. |
| MAN-Q-004 — auth/API/cache/rate limit | **Partially answered** | Reuse `gh` conventions; subprocess, direct compatible client and unauthenticated public approaches compared; local status avoids network. STU-EVD-004, STU-EVD-005, STU-EVD-020 | Architecture selects approach after private/GHES/rate/cache/redaction tests; Requirements set offline/stale behavior. |
| MAN-Q-005 — Go direction | **Answered for Study** | Go remains preferred; pure Go is conditional; initial matrix should include Windows amd64, Linux amd64/arm64, macOS amd64/arm64. STU-EVD-002, STU-EVD-006, STU-REC-001 | Architecture must still approve/replace after prototypes, dependency audit, binary/operational measurements and native tests. |
| MAN-Q-006 — flags/output/exits/interaction/accessibility | **Partially answered** | Section 9 analyzes intended flags; typed human/JSON, stdout/stderr, redaction and useful final-state exits are candidates. STU-RQC-002/004/005/011/015/017 | Requirements owner defines exact syntax, precedence, schema, exit codes, prompts/noninteractive mode and accessibility. Windows local install remains ambiguous; the 2.96.0 prerelease-only precompiled pin rejection is source-derived but still warrants version-pinned conformance. |
| MAN-Q-007 — filesystem/archive/lock/backup/recovery | **Partially answered** | Sections 11–13 define hostile-input invariants, platform claim ceilings and transaction/recovery inputs. STU-EVD-007, STU-EVD-014, STU-EVD-015, STU-EVD-016; STU-REC-005 | Requirements/Architecture/security reviewers set exact limits, lock/replacement/durability/backup policy after native fault/race tests. |

### Conclusions

- **STU-FND-012 — feasibility.** The extension and Go mechanisms support the
  Mandate direction, including Linux arm64, but distribution, Windows local
  development, and native filesystem behavior require pinned conformance.
  [inference; STU-EVD-002, STU-EVD-003, STU-EVD-006]
- **STU-FND-013 — principal blocker.** The canonical Toolkit is unreleased and
  lacks a purpose-built asset, reusable plan/ownership contract and provenance;
  therefore compatibility cannot yet be implemented or claimed truthfully.
  [verified repository fact/inference; STU-EVD-008, STU-EVD-011]
- **STU-FND-014 — safety ceiling.** The current installer establishes valuable
  preservation semantics but not archive, concurrency, cross-platform
  transaction or crash-recovery guarantees. [verified repository fact/inference;
  STU-EVD-009, STU-EVD-010]
- **STU-REC-008 — historical Study disposition at Slice closure.** Submit this
  reviewed candidate for
  Steering assessment now that `VER-SPS-002` and fresh exact-head review
  `REV-SPS-002-003` have passed. Do not treat those process gates as Steering
  acceptance or authorize Requirements, Architecture, or implementation from
  the Study alone. Resolve upstream contract/provenance prerequisites before a
  future implementation Slice. [recommendation]

Traceability is `MAN-001` → `STU-001` → `Sprint-002` / `SPI-002` / `SPS-002`,
with evidence `STU-EVD-001`–`021`, findings `STU-FND-001`–`014`,
recommendations `STU-REC-001`–`008`, requirement candidates
`STU-RQC-001`–`017`, Architecture inputs `STU-ADI-001`–`010`, and upstream
findings `STU-XRF-001`–`007`. It covers all Mandate outcomes, preserves active
boundaries and assumptions, and does not mark any success criterion delivered.
`VER-SPS-002` records passed Study/Slice verification, and
`REV-SPS-002-001` through `REV-SPS-002-003` record the real independent review
and remediation chain. Final Reviewer `/root/study_reviewer_final` approved
exact pushed evidence candidate `8f59a31ee6c6ee4264cb301217c5a3489fd4444d`
for `SPS-002` closure with zero blocking, high, medium, or low findings. At
Slice closure, Steering assessment was the next authorized decision and
`STU-001` was not yet Steering accepted. The 2026-07-17 disposition recorded
above now accepts only that historical evidence foundation; no Mandate success
criterion is marked delivered and no later lifecycle phase is authorized.
[verified repository-process boundary; STU-EVD-001]
