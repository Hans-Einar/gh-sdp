# SPS-008 P1 implementation notes

Authorization: Master delegated bounded P1 on 2026-10-01.
Contract baseline: 94351791843b66a3e41fce5694d0db2022dddf0e on
sdp/release-0.2.1. Master committed contract/traceability corrections and
prerelease preparation through e2190dc during this work; Worker files were separate.
Client target is v0.2.1, paired with Toolkit v2.1.0.
Status: P1 implemented and locally verified. Clean exact-source packaging,
packaged integration, independent review and publication remain Master-owned
P2/P3 work. No publication, merge or installation is claimed.

## Dependency provenance

The public upstream annotated v2.1.0 tag peeled to the assigned source
93517ad98cd188c0debeb1d0f3d36d123c6e4a3b. Direct remote Go resolution recorded
that exact Origin.Hash, subdirectory SDPTool/bootstrap and refs/tags/v2.1.0.
The immutable module is v0.0.0-20261001092620-93517ad98cd1, with content checksum
h1:lw9R5Np2sa9xv7IFVsVPSQpHCBkg5FHKtZQuai47ml4=.
Its DefaultRelease is exactly
https://github.com/Hans-Einar/SDP/releases/download/v2.1.0/sdp-release.json.
The module graph contains only this client and canonical bootstrap; no replacement
is present. Tidy removed the superseded module checksums.

The initial default-proxy go get failed by reporting that the parent SDPTool
module did not contain the bootstrap package. GOPROXY=direct successfully
resolved and retrieved the exact nested module; no local source replacement,
checksum bypass or trust override was used.

## Changes

The existing regression now requires the literal v2.1.0 descriptor in bootstrap
configuration and the actual child environment for unset/empty SDP_RELEASE;
explicit overrides remain covered. Production client logic is unchanged.
README, current architecture/design and implementation guidance, release lifecycle
and selected v0.2.1 notes describe the new default and current Slice. Historical
release claims remain historical and the selected patch remains unpublished.
No CI infrastructure, global extension operation or XFMD change was performed.

## Verification

Executed from /home/warloc/git/gh-sdp with Go 1.27.1 on Linux amd64.
The following successful commands establish remote provenance and P1 verification:

```sh
git ls-remote https://github.com/Hans-Einar/SDP.git refs/tags/v2.1.0 'refs/tags/v2.1.0^{}'
env GOMAXPROCS=2 GOPROXY=direct /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go list -m -json github.com/Hans-Einar/SDP/SDPTool/bootstrap@93517ad98cd188c0debeb1d0f3d36d123c6e4a3b
env GOMAXPROCS=2 GOPROXY=direct /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go get github.com/Hans-Einar/SDP/SDPTool/bootstrap@93517ad98cd188c0debeb1d0f3d36d123c6e4a3b
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go mod tidy
env -u SDP_RELEASE -u SDP_TEST_KEY -u SDP_OFFLINE -u SDP_CACHE_DIR -u GH_SDP_BINARY -u SDPTOOL_BINARY -u GH_SDP_PACKAGE_DIR -u GH_SDP_VIA_GH GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -p 2 -count=1 -v ./...
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go vet -p 2 ./...
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go mod verify
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go list -m -json all
git diff --check
```

All listed commands exited 0. Race tests passed exact argument/stream/environment/
exit delegation, configuration and failure boundaries, and default/override
selection. Module verification reported all modules verified. Tests removed
inherited SDP and packaged-test controls; no production client cache was used.
TestPackagedCandidates, TestPackagedAssetDiscovery and TestPackagedClientRejectsFIFO
skipped because no client package was supplied in P1. These are required P2 checks,
not claimed passes. Master owns clean package pairing with the supplied final
engine, isolated gh routing, public fresh-cache defaults and independent review.
The supplied engine's SHA-256 was independently checked as
691c021751cefb35c7f92c902545cd6c5d71ccadfcde07ae507a6ce026c8a823;
no engine execution is claimed by this checksum check.

Worker stops after the bounded P1 commit.

## P2/P3 Master reconciliation

VER-SPS-008 and REV-SPS-008-001 establish the clean exact source 315ec1d
package, complete native suite and production public default. Annotated client
v0.2.1 was published at 2026-10-01T13:35:16Z; downloaded assets match.
Independent publication closeout remains pending. No main/global/XFMD mutation occurred.

REV-SPS-008-002 now independently approves publication reconciliation. SPS-008
is complete; no next work is selected.
