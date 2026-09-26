# SPS-003 implementation notes

Authorization: owner-selected PLAN-SDP-0003, GIP-3-M2 (2026-09-27).
Baseline: gh-sdp main be990cc; branch sdp/install-gip-3.
Status: complete after independent REV-SPS-003-002 approval and Master closeout.

## Canonical source reuse

The only product dependency is the stdlib-only nested Go module
`github.com/Hans-Einar/SDP/SDPTool/bootstrap` at
`v0.0.0-20260926233855-fb79727aaa22`, from canonical SDP commit
`fb79727aaa229e5a4a5f228f458d350cd196a6a8`. go.sum pins module content.
No local replace, vendored policy or sibling-module dependency is shipped.
The parent SDPTool module's sibling development replaces are not inherited.

The module performs descriptor signature, explicit test trust, immutable
selector/digest cache, asset digest/platform selection and executable protocol
probe. The client configures it and invokes the fixed returned path with os/exec.
Saved apply/resume requests force offline bootstrap. Installation flags and
streams pass unchanged to SDPTool, whose direct saved-plan path does not resolve
a new descriptor. Local boolean parsing exits 2; bootstrap/start errors exit 4;
ordinary child exits are preserved.

## Evidence and scope

[VER-SPS-003](../../Verification/VER-SPS-003.md) records unit/race, actual packaged
client and real gh sdp integration checks. The fixture uses an explicit ephemeral
Ed25519 test key, a manual disposable SDP baseline and project-owned AGENTS.md.
Exact preview output/plan equality, direct/client apply, preservation and
no-change repetition pass. Descriptor protocol mismatch, signed digest-valid
binary protocol mismatch and corrupt cache fail closed.

Managed AGENTS.md, installed Toolkit facts and the release-event ledger are
unchanged. The legacy installed ledger schema has no implementation transition
kind; lifecycle state is recorded in current project/Sprint/traceability records
without fabricating a release event. Completed Study records and source evidence
remain historical. Current addenda explicitly supersede only the proposed
client-owned apply engine and earlier assignment boundaries.

No production key, release catalog, binary release, merge, live target update or
native Windows/macOS support is claimed. Test subprocesses are not a power-loss
or cancellation/signal-forwarding guarantee. Native Linux amd64 is the measured
platform; the full GIP-4 adoption scenario belongs to the coordinating SDP plan.

The first GitHub CLI test isolated GH_CONFIG_DIR but revealed that 2.97.0 stores
extensions under XDG_DATA_HOME separately. The test-created dangling extension
symlink was identified by its exact disposable target and removed; gh-tree was
untouched. The corrected harness isolates both locations, and repeat checks
confirm no extension link is left in the user's data directory.

Fresh review REV-SPS-003-001 found a canonical bootstrap FIFO read issue (upstream
fix and repin required), an overly broad test-coverage statement, and contradictory
design error wording. The documentation issues were corrected in a82d432; that initial review
remained open for the dependency fix until REV-SPS-003-002. A subsequent client inspection also restricted
SDP_OFFLINE to lowercase true/false: strconv aliases such as 1/TRUE would otherwise
select offline bootstrap while the child interpreted its unchanged environment as
online. Regression tests now reject those divergent spellings.

The dependency now pins fb79727, containing the canonical pre-open regular-file
guard and native FIFO regression. A client packaged Linux regression also probes
a FIFO descriptor under a two-second bound and requires prompt exit 4. No reader
policy was copied into gh-sdp. Fresh follow-up review REV-SPS-003-002 independently confirmed the fix and
approved the exact paired candidates.
