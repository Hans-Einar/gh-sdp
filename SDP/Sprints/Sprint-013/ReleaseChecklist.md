# SDP 2.3.2 extension release checklist

Authority and scope: SPS-013. Previous release: 2.3.1. Linux amd64.

- [x] Select compatible patch 2.3.2 and preserve released history.
- [x] Prepare notes, prerelease identity, bounded contract and traceability.
- [x] Pass focused packaging race tests and vet; record exact environment.
- [x] Package from clean committed extension source using the exact production-signed
  canonical descriptor; verify signature, source/version identity and unchanged bytes.
- [x] Pass actual package asset-discovery/provenance test.
- [x] Obtain fresh independent combined review on actual candidates/artifacts.
- [x] Root verifies canonical public descriptor and engine before extension publication.
- [x] Root publishes annotated tag and GitHub Release, verifies public asset hashes
  and isolated extension upgrade, then reconciles real identities.

Canonical inventory, descriptor signing, predecessor rehearsal and generated
canonical release log belong to the root canonical release checklist. The
extension does not modify the engine. Runtime/setup/update behavior evidence from
SPS-011/SPS-012 remains applicable to unchanged source; this checklist makes no
new cross-platform claim. Actual user extension replacement and project Framework
upgrades are outside this assignment.

All applicable gates passed; public evidence and final independent combined review
are retained in evidence/. Released tag/source stays c1d86bd; this closeout commit
changes records only. The owner's live extension remains 2.3.1 for manual upgrade.
