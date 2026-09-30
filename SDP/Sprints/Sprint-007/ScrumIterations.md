# Sprint-007 — SDP discovery and human output default

Status: in-progress
Iteration: SPI-007
Slice: SPS-007
Authorization: owner selected Session 0003 on 2026-10-01, implementing automatic source discovery and releasing SDP plus the required gh-sdp update for manual XFMD adoption.
Upstream coordination: SDP PLAN-SDP-0014 / Session 0003.
BranchPolicy: sdp/release-0.2.0 from clean main be00296ab88590aea2eff757f1f76148449beb7c; publish only the reviewed branch commit. No main merge is authorized.
CommitPolicy: contract, bounded Worker implementation, exact-candidate evidence/review, publication reconciliation.
Release: gh-sdp v0.2.0 paired with SDP v2.0.0; Linux amd64 only.

## Contract

REQ-GHS-001–005, ARC-GHS-001 and DES-GHS-001 continue to govern the thin
adapter. Select the immutable remote bootstrap revision supplied by root with
SDP v2.0.0 as its exact default. Upstream owns source discovery, installation
policy and human default output. Bootstrap uses machine-readable version probing
with compatibility fallback; the client continues to forward arguments and
streams unchanged. Explicit --json is the migration path for JSON consumers.

The pre-1.0 MINOR increment reflects a visible default output compatibility
change from the default delegated SDPTool. This is not merely a default version
patch. Client and Toolkit versions remain independent. Arguments, environment,
streams, cwd, exit status, signed production trust, explicit descriptor overrides
and saved operation offline behavior remain preserved.

Expected Worker files: go.mod/go.sum, main_test.go and necessary compatibility
integration tests, README.md, current architecture/design default guidance,
docs/Release-Lifecycle.md, selected release notes, implementationNotes.md.
Master owns manifest, traceability, release, verification and review records.
No managed skills/framework migration, new platform support, main merge, global
extension installation or live XFMD change is selected.

## Milestones and gate

1. P1: Worker pins root-supplied remote bootstrap, updates exact default regression
   and current usage guidance. Go race tests, vet, module verification and remote
   no-replacement provenance pass.
2. P2: build clean exact-commit Linux amd64 package and checksum; verify actual
   package discovery/FIFO rejection and explicit JSON delegation. Fresh independent
   Reviewer inspects the source, package evidence and release records.
3. P3: after root confirms public SDP v2.0.0, run packaged default with fresh cache
   and production trust; confirm human default and explicit JSON. Publish annotated
   v0.2.0 and exact assets from the reviewed branch commit, inspect remote objects,
   reconcile actual records, independently review closeout and stop.

Publication is owner-authorized but waits for upstream production availability
and actual gate completion. Publication identities remain null until they exist.
The owner's manual XFMD upgrade is outside this repository assignment.

## Progress

P1 delivered in 336d654c7533a30b340a1366109a41a9307ddf1e using remote bootstrap
v0.0.0-20260930230248-247fb7aba628. Race tests with the initial paired engine,
vet and module provenance passed as recorded in implementationNotes.md.
P2 exact clean package and independent review are pending. P3 awaits the public
paired descriptor. No client publication or integration has occurred.

P2 initial clean package checks and independent REV-SPS-007-001 are approved.
The exact package source is 72b4e048a90841a1994e9846bb4db0d0ad9b922e.
Final immutable upstream package pairing is pending; development-engine rehearsal
evidence is not substituted for that gate. P3 still awaits the public descriptor.

P2 final d304261 immutable pair and the public production default gate now pass.
REV-SPS-007-002 approves exact client source 72b4e048 for authorized publication.
P3 actual tag/release and independent reconciliation remain pending.

## Actual publication

Annotated v0.2.0 now resolves to exact reviewed source
72b4e048a90841a1994e9846bb4db0d0ad9b922e. GitHub Release was published at
2026-09-30T23:24:45Z; downloaded checksums and manifest match the reviewed
package. Main remains unchanged; no global installation or XFMD mutation
occurred. Independent publication reconciliation remains pending.
