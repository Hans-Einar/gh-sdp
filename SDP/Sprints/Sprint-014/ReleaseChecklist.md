# SDP 2.4.0 extension release checklist

Authority/scope: SPS-014. Prior release: 2.3.2. Qualified platform: Linux amd64.

- [x] Select additive minor 2.4.0; preserve legacy 0.2 and immutable release history.
- [x] Prepare notes, bounded contract, prerelease manifest and traceability.
- [x] Pass focused packaging race tests and vet; retain environment/results.
- [ ] Package from clean committed extension source using production-signed canonical
  descriptor; verify signature, source/version identity and identical engine bytes.
- [ ] Pass actual package asset-discovery/provenance test.
- [ ] Obtain fresh independent combined release review on actual candidates/artifacts.
- [ ] Root verifies canonical public descriptor before extension publication.
- [ ] Root publishes extension tag/release and verifies downloaded assets, isolated
  normal upgrade/install and setup/offline identity; reconcile actual identities.

Canonical inventory/signing, native DocumentHost qualification, supported Framework
predecessor rehearsals, exact-candidate CI and generated canonical release log are
root-owned gates. Existing SPS-011 through SPS-013 behavior evidence remains
applicable to unchanged extension code. No real user extension or project upgrade.
