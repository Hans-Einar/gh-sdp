# Sprint-006 — Template distribution default patch

Status: active
Iteration: SPI-006
Slice: SPS-006
Authorization: owner explicitly authorized commit, integration, release, extension installation and XFMD upgrade on 2026-09-28; root coordination assigns this repository only.
BranchPolicy: sdp/release-0.1.2 from clean main 9ef2736; independently reviewed PR merge to main, then publication reconciliation on a follow-up branch.
CommitPolicy: contract, bounded Worker implementation, verified review evidence, publication reconciliation.
Release: gh-sdp v0.1.2 paired with SDP v0.2.2, subject to root confirming its final selection; Linux amd64 only.

## Contract

Select the canonical remote immutable bootstrap revision supplied by root that
uses the signed SDP v0.2.2 descriptor by default. This distributes the upstream
current template update through the existing thin client. REQ-GHS-001–005,
ARC-GHS-001 and DES-GHS-001 govern preserved behavior. Arguments, environment,
streams, cwd, child exit status and explicit descriptor overrides remain intact.
No client API, installation policy, supported platform, installed Toolkit or
managed skill migration is selected. Root owns global extension installation and
XFMD adoption. No unrelated repository changes are authorized.

PATCH is appropriate because only the default upstream distribution selection
changes within the existing client API. Client and Toolkit identities remain
separate. Historical release records and append-only ledger bytes are preserved.

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
3. P3: after root confirms public SDP v0.2.2, exercise the packaged default with
   a fresh cache and production trust; merge reviewed preparation, package the
   clean merged commit, publish annotated v0.1.2 and exact assets, inspect remote
   objects, independently review reconciliation, and stop on clean main.

Installed release/versioning skills and docs/Release-Lifecycle.md govern release.
Publication identities remain null until actual tag and GitHub Release exist.
The owner authorization covers this release; no further approval is required.

## Progress

Contract selected. P1 awaits root's immutable remotely available bootstrap SHA;
P2 and P3 are pending. No publication or verification is claimed.
