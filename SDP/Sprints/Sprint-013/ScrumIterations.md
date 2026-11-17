# SPS-013 — Shared SDP 2.3.2 release

SPI-013 / Sprint-013. State: completed. Revision: 1.

Owner explicitly authorizes publication so `gh extension upgrade sdp` obtains the
preview/source-reading output-limit correction (128 files / 32 MiB becomes
512 files / 128 MiB). Canonical root coordinates the signed SDPTool 2.3.2 release,
independent combined review, publication and final reconciliation. The extension
republishes that exact engine without runtime, packager or dependency changes.

BranchPolicy: `sdp/release-2.3.2` from `54f830a`.
CommitPolicy: milestone commits with SPS-013 and concrete delivery.
Previous actual extension/canonical release: 2.3.1. Select patch 2.3.2 for this
compatible limit correction; shared version identity from SPS-011 is retained.

Acceptance: focused packaging tests/vet; clean exact distribution commit;
production-signed canonical descriptor and unchanged engine bytes; actual package
asset-discovery/provenance test; fresh independent combined review; canonical
publication before extension publication; public downloaded asset verification
and isolated extension upgrade qualification. Linux amd64 only. Reuse applicable
SPS-011/SPS-012 behavior evidence for unchanged distribution code.

No main merge, replacement of the owner's installed extension, consumer Framework
upgrade or new platform qualification is selected. Published history is immutable.
The assigned Worker prepares source, packaging, isolated qualification and
reconciliation; root owns public actions.

Preflight: extension worktree was clean at 54f830a; only this extension worktree
exists; no extension record/code overlap or SPS-013 identity collision was found.
Canonical root owns the active Session and cross-repository concurrency record.

Loaded routines: managed AGENTS/project overrides, Framework/manifests, installed
Worker, release, versioning and traceability skills; current source SDP entrypoint,
Worker and release guidance. No installed Framework migration is performed.

Outcome: public 2.3.2, exact package/public download checks, isolated normal upgrade
and clean install, public-default Framework qualification and fresh independent
combined review passed. See Evidence.md and evidence/. Root owns canonical
Session closeout; no further extension work is selected.
