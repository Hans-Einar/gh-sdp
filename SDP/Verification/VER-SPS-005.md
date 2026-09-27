# VER-SPS-005 — Default engine patch release gate

Status: passed; preparation and actual publication verified
Slice: SPS-005
Release: REL-0.1.1
Platform: Linux amd64
Product candidate: 50c2fbe5eff4c17a6116c171797a248d49d9e05e
Bootstrap: v0.0.0-20260927094008-c1901337510e (remote, no replace)

## Candidate evidence

Worker implementation notes record Go 1.27.1 race tests, vet, module verification
and exact v0.2.1 default URL regression results. Both unset and empty SDP_RELEASE
select the literal expected URL; the child receives the same selection and
explicit overrides are preserved. Master inspected the diff and actual package.

Master built clean exact product commit with:

```sh
SDP_GO=/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go scripts/package.sh 0.1.1 /tmp/gh-sdp-sps005-candidate
GH_SDP_PACKAGE_DIR=/tmp/gh-sdp-sps005-candidate GOTOOLCHAIN=local /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -count=1 -run TestPackagedAssetDiscovery -v ./...
```

Package and asset-discovery test passed. Binary SHA-256:
`76bf4f15f61a3defa8787c7cb55a955d669ec67a778499f1cfdba9cee7589b70`.
[Candidate manifest](evidence/SPS-005-candidate-manifest.json) records the clean
source and exact dependency. Independent preparation approval is recorded below.

At candidate preparation, final merged packaging and public production-trust
verification remained pending. Their actual results are recorded below. Global
extension installation and XFMD adoption remain root-coordinated.

## Independent preparation review

[REV-SPS-005-001](../CodeReview/REV-SPS-005-001.md) approves P2 and independently
reproduces the package SHA-256 and complete manifest byte-for-byte from detached
50c2fbe. Race tests, vet, module verification, fresh-cache remote module retrieval,
actual asset discovery and packaged FIFO rejection pass. Upstream module content
diff contains only the default descriptor change. No implementation findings remain.

## Actual release evidence

PR #8 merged to clean source `8cbef9693e353cb04dc6d98021aca454b9078e3f` with
an identical tree to approved preparation `306a993`. Master rebuilt that source
using Go 1.27.1 and ran race tests with both GH_SDP_PACKAGE_DIR and GH_SDP_BINARY
set to the final package, then vet. Package discovery and FIFO rejection ran and
passed; the opt-in direct-engine fixture integration was not rerun or claimed.

[Published package manifest](evidence/SPS-005-published-manifest.json).
Binary SHA-256: `57b1d7dc54adfe6ec15d8e113e6e8943f95741e912e27f10a2afc49e43b3c8f9`.
Annotated tag object `a84d05dddaa080c6582e0ad676f8990dbe448bfa` resolves to the
release source; remote tag inspection matches. The [GitHub Release](https://github.com/Hans-Einar/gh-sdp/releases/tag/v0.1.1)
is public, final and published at `2026-09-27T10:05:48Z`. GitHub's asset digest
matches the local binary. All three assets were downloaded through GitHub CLI and
compared byte-for-byte with the local package successfully.

Both candidate and final merged package were executed with fresh private caches
and SDP_RELEASE, SDP_TEST_KEY and SDP_OFFLINE unset. Default production trust
returned SDPTool `0.2.1`, revision `4bacfce05f92f0dab9680456e297214727b54bc6`.
The upstream release was independently observed public and final before client
publication. Root coordination owns actual global extension installation and
live XFMD verification; this record does not infer those outcomes.
