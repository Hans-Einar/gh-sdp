# SPS-003 implementation notes

Authorization: owner-selected PLAN-SDP-0003, GIP-3-M2 (2026-09-27).
Baseline: gh-sdp main be990cc; branch sdp/install-gip-3.
Status: implemented and verified locally; independent review and closeout pending.

## Canonical source reuse

The only product dependency is the stdlib-only nested Go module
`github.com/Hans-Einar/SDP/SDPTool/bootstrap` at
`v0.0.0-20260926233103-266179e87452`, from canonical SDP commit
`266179e87452272e9f513f0d8f7acb3cdf9e3051`. go.sum pins module content.
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
