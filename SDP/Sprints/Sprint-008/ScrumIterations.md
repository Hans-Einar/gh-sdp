# Sprint-008 — SDP 2.1 default patch

Status: complete
Iteration: SPI-008
Slice: SPS-008
Authorization: owner explicitly requested publication of prepared SDPTool 2.1.0 and paired gh-sdp on 2026-10-01.
BranchPolicy: sdp/release-0.2.1 from clean e5c08fb; publish reviewed exact branch source without merging main.
CommitPolicy: contract, bounded Worker implementation, exact-candidate evidence and independent review, actual publication reconciliation.
Release: gh-sdp v0.2.1 paired with SDP v2.1.0; Linux amd64 only.

## Contract

Select immutable canonical bootstrap from upstream source
93517ad98cd188c0debeb1d0f3d36d123c6e4a3b, whose exact DefaultRelease is v2.1.0.
The client remains a thin adapter under REQ-GHS-001–005, ARC-GHS-001 and
DES-GHS-001. This client patch changes its default engine selection without
changing argument, environment, stream, cwd, exit, human/JSON, explicit override
or offline saved-operation contracts. Toolkit and client versions are independent.
GitHub CLI owns extension update notifications; no client notifier is selected.

Expected Worker files: go.mod/go.sum, exact default regression tests, README and
current default architecture/design guidance where applicable, release lifecycle,
release notes and implementation notes. Master owns governance, manifest,
traceability, verification and release records.

No main merge, global extension installation/upgrade, XFMD changes, managed
Framework migration or new platform support is selected. Preserve the owner's
installed extension for their manual update-notification test. No CI workflow
exists in this repository; do not introduce CI infrastructure for this patch.
Use the native local release gate and record upstream exact-source CI separately.

## Milestones and gate

1. P1: Worker pins remotely resolved immutable bootstrap, updates exact default
   regressions and current guidance; race tests, vet and module provenance pass.
2. P2: build a clean exact-source Linux amd64 package, checksum and manifest;
   run complete packaged race integration and isolated gh routing against the
   supplied final upstream engine. Fresh independent review inspects actual evidence.
3. P3: after root confirms public SDP 2.1.0 descriptor, verify packaged human/JSON
   default through a fresh cache and production trust. Publish annotated v0.2.1
   with exact reviewed assets, inspect remote identities/downloads, reconcile
   records, obtain fresh independent closeout review and stop.

Publication is explicitly authorized but awaits actual gate completion. Release
identities remain null until they exist. Root owns upstream publication.

## Chronological progress

P1 delivered in 315ec1de9dbe9aee990baae47a1ac2a7a07a0702. P2 exact clean
package and full native suite, and P3 public production default checks pass as
recorded in VER-SPS-008. Independent review is pending; client is unpublished.

REV-SPS-008-001 approves the exact candidate and public production default.
Publication may proceed under owner authorization; actual objects are pending.

## Actual publication

Annotated v0.2.1 and GitHub Release now exist at approved exact source 315ec1d.
All downloaded assets match the reviewed package. Independent reconciliation
remains pending; owner global extension and XFMD remain untouched.

## Closeout

REV-SPS-008-002 independently approves the actual published release, downloaded
assets and production default with no findings. P1–P3 complete; records reconcile
real publication. Main, installed global extension and XFMD remain untouched.
No further Slice is selected.
