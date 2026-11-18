# SPS-014 — Shared SDP 2.4.0 DocumentHost release

SPI-014 / Sprint-014. State: completed. Revision: 1.

The owner authorizes paired SDPTool/gh-sdp 2.4.0 publication of the completed
DocumentHost launcher. Canonical implementation/model independent review covers
735c865, closed at 8f1fb9a. Root coordinates canonical signing, exact-candidate
combined release review, both publications and final reconciliation.

The extension republishes the exact production-signed canonical engine unchanged.
Minor 2.4.0 selects additive native DocumentHost panes through explicit `--context`
choice and SVG data, with reload preserving the last valid view. Legacy SDUI 0.2
behavior remains supported. Domain callbacks and an XFMD context picker are outside
this release; no extension runtime, packager or dependency change is selected.

BranchPolicy: `sdp/release-2.4.0` from released-state commit `c0b6f0f`.
CommitPolicy: milestone commits naming SPS-014 and concrete delivery.
Previous actual canonical/extension release: 2.3.2. Shared version identity and
single update path from SPS-011 remain authoritative.

Acceptance: focused packaging race tests/vet; clean exact distribution commit;
production signature, descriptor/source/version identity and same-byte engine
qualification; actual package asset-discovery/provenance test; fresh independent
combined release review; canonical public descriptor before extension publication;
public asset hashes, isolated normal extension upgrade/install and setup/direct
offline identity checks. Root owns canonical runtime/native/predecessor evidence.
Reuse existing unchanged distribution behavior evidence. Linux amd64 only.

No main merge, user extension replacement, consumer Framework upgrade or platform
expansion is selected. This Worker prepares records, packages and reconciliation; root owns public
actions and canonical closeout. Published history must remain immutable.

Preflight: extension checkout clean at c0b6f0f; single worktree; no SPS-014 or release
identity collision and no overlapping extension edits found. Scoped verdict: green.
Root maintains cross-repository concurrency and the active canonical Session.

Loaded routines: managed AGENTS/project overrides, reminders, Framework/manifests,
installed Worker/release/versioning/traceability; current SDP entrypoint, Worker
and release. Installed Framework remains unchanged.

Outcome: public 2.4.0 signed-engine distribution passed actual package/public hash
checks, isolated normal 2.3.2 upgrade and clean install, setup/offline identity,
public-default Framework qualification and final independent combined review.
See evidence/. Root owns canonical Session closeout; no next extension work selected.
