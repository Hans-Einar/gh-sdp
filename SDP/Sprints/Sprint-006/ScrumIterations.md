# Sprint-006 — Template distribution default patch

Status: paused by owner
Iteration: SPI-006
Slice: SPS-006
Authorization: owner explicitly authorized commit, integration, release, extension installation and XFMD upgrade on 2026-09-28; root coordination assigns this repository only.
BranchPolicy: sdp/release-0.1.2 from clean main 9ef2736; independently reviewed PR merge to main, then publication reconciliation on a follow-up branch.
CommitPolicy: contract, bounded Worker implementation, verified review evidence, publication reconciliation.
Release: gh-sdp v0.1.2 paired with owner-selected SDP v1.0.0; Linux amd64 only.

## Contract

Select the canonical remote immutable bootstrap revision supplied by root that
uses the signed SDP v1.0.0 descriptor by default. This distributes the upstream
Go-only release and current template update through the existing thin client. REQ-GHS-001–005,
ARC-GHS-001 and DES-GHS-001 govern preserved behavior. Arguments, environment,
streams, cwd, child exit status and explicit descriptor overrides remain intact.
No client API, installation policy, supported platform, installed Toolkit or
managed skill migration is selected. This coordinator owns the delegated global
extension installation and default-version verification after publication; root
owns upstream publication and XFMD adoption. No unrelated changes are authorized.

PATCH is appropriate because only the default upstream distribution selection
changes within the existing client API. SDP v1.0.0 is independently selected
because it removes public PowerShell entrypoints; the Go client API stays intact.
Client and Toolkit identities remain separate. Historical release records and append-only ledger bytes are preserved.

Expected Worker files: go.mod/go.sum, exact default regression in main_test.go,
README.md, current architecture/design default guidance, docs/Release-Lifecycle.md,
selected release notes and implementationNotes.md. Master owns manifest,
traceability, release and verification records.

## Milestones and gate

1. P1: bounded Worker pins the remotely published bootstrap SHA supplied by root,
   updates the exact default URL regression and current docs; race tests, vet,
   module verification and no-replacement check pass.
2. P2: build clean exact-commit Linux amd64 package with checksum; verify actual
   package discovery and FIFO rejection. Fresh independent Reviewer inspects the
   same source, required records, real evidence and remote module provenance.
3. P3: after root confirms the new Go-only candidate, required pin and public SDP v1.0.0, exercise the packaged default with
   a fresh cache and production trust; merge reviewed preparation, package the
   clean merged commit, publish annotated v0.1.2 and exact assets, inspect remote
   objects, independently review reconciliation, and stop on clean main.

Installed release/versioning skills and docs/Release-Lifecycle.md govern release.
Publication identities remain null until actual tag and GitHub Release exist.
The owner authorization covers this release; no further approval is required.

## Progress

P1 delivered in f7b690633413b0a1f13b49d3e6265bd3494d1995. Root confirmed SDP
v0.2.2 and supplied remotely available bootstrap 591920b69534fa12e587f5f95a55b6bce9b13da4.
P2 clean exact-commit package, asset discovery and FIFO checks passed; independent
REV-SPS-006-001 approved that candidate only. VER-SPS-006 records the actual
historical evidence. The v1.0.0 pin, updated regression/docs and corresponding
P1/P2 verification are now pending. P3 remains paused.

## Owner steering — publication and integration paused

The owner stopped publication and integration on 2026-09-28 and requires removal
of all PowerShell dependencies from the upstream release. Preserve this prepared
client branch. P2 approval applies to product candidate f7b6906 only; do not merge,
tag, publish or upgrade the global extension until root confirms the revised
Go-only upstream candidate and whether its bootstrap requires a new immutable pin.
No client PR, merge, tag, release or global installation occurred in SPS-006.
The installed extension remains v0.1.1. Root still owns XFMD application changes.
Before the pause, root explicitly delegated global extension installation and
fresh-cache default verification to this client coordinator after publication;
that operation is also paused.

## Owner selection — SDP v1.0.0

The owner selected SDP v1.0.0 for removal of the public PowerShell entrypoints.
Root is preparing the Go-only candidate, including the native payload inventory
and removal of legacy build/test runners. Client v0.1.2 remains a backward-compatible
default update. Do not pin a replacement SHA until root supplies the gated new
candidate. Preserve the previous candidate and its evidence as historical;
its v0.2.2 default and P2 review do not prove or approve the v1.0.0 target.
Publication and integration remain paused. Existing owner authorization applies
when root confirms the revised candidate; no new owner approval is required.
