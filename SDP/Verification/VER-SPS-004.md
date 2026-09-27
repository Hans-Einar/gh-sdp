# VER-SPS-004 — First native release gate

Status: passed; preparation and actual publication verified
Slice: SPS-004
Release: REL-0.1.0
Platform: Linux amd64
Product candidate: 2618e3857a9cf1a7165090ec5209727593fa3221
Bootstrap: v0.0.0-20260927075424-659e543e89d7 (remote, no replacement)

## Candidate evidence

Worker ran Go 1.27.1 race tests and vet successfully. The shared default and
explicit SDP_RELEASE override agree between bootstrap and the actual delegated
child; original argv, streams, cwd, exit and parent environment are preserved.

Two clean detached candidate packages produced identical binary SHA-256:
a76438e039d71fb9bfaec061ed91b9419c55ae2bf2158454e3684bad570a7874.
The [package manifest](evidence/SPS-004-candidate-manifest.json) records source,
module and toolchain. Dirty source and occupied output rejection passed.

The actual isolated GitHub CLI integration passed in 3.153 seconds against
GIP4 engine 07335b0, SHA-256
0d984b548f47b6ecb3b5049a7b2c24268822cf1032fdaeb373bb7574745fc18e.
It checks signed-fixture direct/client preview parity, apply, preserved project
content, repeat no-change, incompatible protocol, corrupt cache and FIFO rejection.
This is fixture evidence; it is not a production descriptor or live XFMD test.

## Commands

```sh
GOTOOLCHAIN=local go test -race -count=1 ./...
GOTOOLCHAIN=local go vet ./...
SDP_GO=/absolute/path/to/go scripts/package.sh 0.1.0 /tmp/fresh-package
GH_SDP_VIA_GH=true GH_SDP_BINARY=/tmp/gh-sdp-r1-final-package/gh-sdp_linux_amd64 SDPTOOL_BINARY=/tmp/gip4-exact-package/sdptool GOTOOLCHAIN=local go test -race -count=1 -v ./...
```

Independent [REV-SPS-004-001](../CodeReview/REV-SPS-004-001.md) reproduced race/vet,
module verification, exact package hash and actual isolated gh integration. It also
verified a local production-signed descriptor without SDP_TEST_KEY, yielding 63
install actions with signed provenance and no project writes. All three draft
documentation findings are resolved; exact preparation 4660b6b is approved.

The
final merged release commit is rebuilt cleanly and its identity/checksum recorded
at publication. Public descriptor availability and actual GitHub extension
installation must be verified then; neither is inferred from these fixtures.

## Remote asset discovery correction

Pre-publication inspection found that the initial underscore-named asset was not
discoverable by GitHub CLI. The local symlink integration did not cover remote
asset selection. Worker candidate c62f5fc changes the asset to
`gh-sdp-linux-amd64`, with no change to the client runtime. Independent
[REV-SPS-004-002](../CodeReview/REV-SPS-004-002.md) approved this correction and
reproduced package SHA-256
`c98257e4ddc688773805fdf6a14e9fec1e23bc424018b1b59a46af1363f800d6`.
Its actual-package discovery test passes and rejects the historical underscore
package. Race tests and vet pass. The [corrected manifest](evidence/SPS-004-assetfix-manifest.json)
is separate from the retained historical package record. Actual GitHub remote
installation remains a required post-publication check.

## Actual release evidence

Final clean merged release source: `8ad0fc2906fc52bd4ee4c214e801a5872a4b0faa`.
[Published package manifest](evidence/SPS-004-published-manifest.json).
Binary SHA-256: `9b34b6775137bafd2c864c17bb9930a5af45024fed4167e291ce9a0f62fcd6d3`.
`TestPackagedAssetDiscovery` passed against the final package. Annotated tag object
`559888833fbe0bcd172f7c9b65c48496b9a821aa` resolves to that source; remote Git
inspection confirms it. [GitHub Release](https://github.com/Hans-Einar/gh-sdp/releases/tag/v0.1.0) is public, not draft or prerelease,
published at `2026-09-27T08:14:45Z`. GitHub's uploaded asset digest equals the
local binary SHA-256.

Actual final client `--version` with an empty fresh cache and neither SDP_RELEASE
nor SDP_TEST_KEY fetched the public default signed descriptor and returned engine
version 0.2.0 / revision 738e6c882daed85591311f18248dab2a48ce2076. This verifies
public production trust and delegation. Global GitHub extension installation and
live XFMD adoption are root-coordinated follow-up evidence, not assumed here.

## Actual remote extension installation

Root coordination installed `Hans-Einar/gh-sdp` through GitHub CLI. The client
coordinator then independently observed `gh extension list`: gh-sdp v0.1.0 and
the existing gh-tree v0.3.14 remain present. With SDP_RELEASE and SDP_TEST_KEY
unset, `gh sdp --version` returned SDPTool 0.2.0 / revision
738e6c882daed85591311f18248dab2a48ce2076. This closes the remote asset-selection
check missed by the historical local symlink test. Live XFMD mutation evidence
remains in the root SDP workstream.
