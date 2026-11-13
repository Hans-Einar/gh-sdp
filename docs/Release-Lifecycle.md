# Native release gate

Activity0004 / PLAN-SDP-0029 C2 selects distribution2.3.0, identical to the
canonical engine version. Earlier 0.2.x separately versioned bootstrap clients
remain immutable history. Current work is [SPS-011](../SDP/Sprints/Sprint-011/ScrumIterations.md).

Require a clean reviewed distribution commit, race tests/vet, exact canonical
signed release descriptor and asset, and package checks before publication.
`scripts/package.sh VERSION SIGNED-DESCRIPTOR NEW-OUTPUT` verifies canonical
publisher trust/digest/protocol at build time and checks engine version/revision
against the signed source. It copies that engine unchanged under the GitHub CLI
asset suffix `gh-sdp-linux-amd64`. The manifest distinguishes engine repository,
commit, version and hash from distribution repository/commit. Signed descriptor
bytes and signature must accompany the candidate. Production qualification must
reject test-signed provenance; explicit fixture success is insufficient.

Run actual isolated gh install/upgrade, setup, direct/gh identity and byte
comparison, project preview/apply preservation and offline direct use. Update must
replace the registered engine before another gh sdp invocation. A local-extension
symlink test alone does not qualify the binary-release update path. Retain honest
Linux amd64 scope; cross-compilation does not qualify other platforms.

Root coordinates canonical signing, exact production candidate verification,
fresh independent review and publication order. No main merge or publication is
performed by C2. Canonical2.3.0 must be public and verified before extension2.3.0;
both must be available before successor XFMD removes its private engine. Tag,
release, public installation and owner acceptance remain separate recorded facts.
Never replace an already published asset or label a fixture as production proof.
