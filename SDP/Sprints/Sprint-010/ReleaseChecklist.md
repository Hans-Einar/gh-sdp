# gh-sdp 0.2.3 release checklist

Adapted from the instantiated Toolkit ReleaseChecklist. Exact candidate:
03e45710db3ad845d05b26b19c286bd069911294. Previous client:0.2.2.
Owner authorization:2026-11-11 publication and three-host installation.

## Scope and content

- [x] Select backward-compatible client0.2.3 independently from engine2.2.1.
- [x] Enumerate trust/default changes and preserve thin-client/offline boundaries.
- [x] Freeze canonical SDP/RELEASE-NOTES.md0.2.3 section; retain old release notes.
- [x] Check exact clean-source package, immutable dependency and actual packaged tests.
- N/A: producer release-log generation. This legacy client workflow uses its
  canonical release-note section directly as the public release body.

## Distribution and review

- [x] Portable Linux amd64/v1 binary, manifest and checksum built from exact source.
- [x] Race tests, vet, module verification and actual packaged integration pass.
- [x] Independent REV-SPS-010 has no findings; producer exact6f33 CI passed.
- [x] Exact client verifies the real public signed2.2.1 default with a fresh cache
  before client publication. See upstream public-default.json evidence.
- N/A: this wrapper does not distribute Toolkit templates, payload inventory,
  predecessor contracts or signed SDP descriptors. Those checklist obligations
  belong to upstream R221 and its final production-signed evidence. Private
  publisher keys are never part of the client package.
- N/A: no client GitHub CI workflow exists; native release tests and independent
  review are the client gates. Do not describe upstream CI as client CI.

## Publication and deployment

- [x] Annotated v0.2.3 tag targets the reviewed source; owner publication authorized.
- [x] Public binary, manifest and checksum downloads match reviewed bytes.
- [x] Reconcile actual source/tag/publication metadata and append release events.
- [x] Linux, Beinerud and labtop extension hashes and public engine identities pass.
- [x] Existing extensions backed up before installation.
- [x] Distinguish extension installation from project SDP upgrades. Root R221
  records three successful consumer upgrades, four blocked histories and one
  excluded legacy installation; this client slice does not claim full-fleet success.

Evidence: Evidence.md, Publication.json and Host-Installations.json.
