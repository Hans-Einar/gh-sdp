# Linux publisher trust

Current Activity0004/SPS-011 uses the same immutable public trust at build time
to verify canonical engine2.3.0 before exact-byte republishing. No bootstrap is
installed. Package with `scripts/package.sh 2.3.0 SIGNED-DESCRIPTOR NEW-OUTPUT`.
Production requires production-signed provenance; explicit test keys qualify
fixtures only. Canonical publisher owns signing; no private keys are needed here.

## Historical 0.2.3 rollout


The initial FIX-0.2.3-001 bridge was reviewed with SDP2.2.0 as default. Owner
authorization on 2026-11-11 selects SPS-010 publication with SDP2.2.1 as default,
followed by installation on Linux, Beinerud and the laptop. Historical bridge
evidence remains scoped to its own source; the release has its own final gate.

The immutable bootstrap retains the laptop key and adds the Linux key:
`2c947457cffe015f870168cb332242c49606fdea378a13d875579544c2989c03`.
Its public base64 value is `3c4v3py8MxlsZ0zGp0T2Wtmcl2rXQ4f4WEM/YBShBiw=`.
Only SDP's publisher holds private signing material; gh-sdp requires none to build.

Publish and verify signed SDP2.2.1 first, then publish/install gh-sdp0.2.3. Old
0.2.2 clients reject the Linux signer. GitHub/SSH authentication or SDP_RELEASE
cannot introduce trust. Never use SDP_TEST_KEY for production rollout.
The release does not revoke the laptop key, change signature formats, enable
automatic trust discovery or replace existing released artifacts.

The historical0.2.3 packager built a portable Linux amd64 client.
Do not inherit a host-specific GOAMD64 setting. Exact-candidate evidence belongs
to [SPS-010](../SDP/Sprints/Sprint-010/ScrumIterations.md).
Project SDP upgrades are a separate root-bound preview/apply operation.
