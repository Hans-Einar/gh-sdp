# SPS-007 P1 implementation notes

Authorization: Master delegated the bounded P1 implementation on 2026-10-01.
Baseline: contract commit 0917212 on sdp/release-0.2.0. Client target is v0.2.0;
paired Toolkit target is v2.0.0.
Status: P1 implemented and locally verified. Exact clean release packaging and
fresh independent review remain P2 work. No publication or merge is claimed.

## Dependency provenance

Root supplied source commit `247fb7aba6282e4eac751b0dcd65a710477c6bbb`.
The canonical remote branch `refs/heads/sdp/session-0003/discovery` resolved to
that exact SHA. Go retrieved the immutable nested bootstrap module as
`v0.0.0-20260930230248-247fb7aba628`; tidy removed the superseded checksum.
Module content checksum: `h1:/BCJuxrLk6ckbp9Mdth/WObQvd4Vtgd+DDqyUtwaFao=`.
The module list includes only this client and the canonical bootstrap, with no
local or remote replacement.

## Changes

The exact default regression requires
`https://github.com/Hans-Einar/SDP/releases/download/v2.0.0/sdp-release.json`
for unset and empty SDP_RELEASE in bootstrap configuration and the actual child
environment. The explicit descriptor override remains covered.

The production adapter needs no change. Packaged integration now compares actual
client and direct engine version/discovery/preview output in human and explicit
JSON modes, checks the JSON version protocol, verifies automatic discovery without
a navigation registration plus the sdptool/0.2 inventory and inline navigation,
and compares diagnostic streams and child exit
status in both modes. Existing signed preview/apply, saved operation offline,
preservation, incompatible protocol and corrupt cache checks remain present.

README, current architecture/design and implementation guidance, release lifecycle
and selected unreleased notes describe upstream discovery, the human output default
and the --json migration. Historical release notes remain unchanged. The visible
compatibility change selects client v0.2.0 independently of Toolkit v2.0.0.
Publication guidance follows the selected reviewed branch without a main merge.
Global extension installation and live XFMD changes are not selected.

## Verification

Executed in `/home/warloc/git/gh-sdp` using Go 1.27.1 on Linux amd64, with
GOMAXPROCS=2 and test/vet package concurrency bounded to two. All commands exited 0:

```sh
git ls-remote https://github.com/Hans-Einar/SDP.git refs/heads/sdp/session-0003/discovery
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go get github.com/Hans-Einar/SDP/SDPTool/bootstrap@247fb7aba6282e4eac751b0dcd65a710477c6bbb
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go mod tidy
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go build -o /tmp/gh-sdp-sps007-worker .
env GOMAXPROCS=2 GH_SDP_BINARY=/tmp/gh-sdp-sps007-worker SDPTOOL_BINARY=/tmp/sdptool-session3 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -p 2 -count=1 -v ./...
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go vet -p 2 ./...
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go mod verify
env GOMAXPROCS=2 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go list -m -json all
git diff --check
```

Race tests passed for the literal v2.0.0 default in bootstrap and the child,
explicit overrides, argument/stream/environment/cwd/exit preservation, failure
boundaries, and the real client/engine integration described above. FIFO rejection
also passed. Module verification reported `all modules verified`.

The supplied engine was an initial upstream development build; the local client
was built from the Worker candidate before commit. Their SHA-256 values were:

- client: `02c7915da5c4a13f32ee56957cd51ee8c6367fee512b7d0ec9b2b5c574e304d9`
- engine: `4383e8b8539eb23aeba5b8333c6cdb4b99f827dc039014cba67fd38794747f8c`

This is integration rehearsal evidence, not clean release-package evidence.
TestPackagedAssetDiscovery skipped because no release package directory was
supplied. Master owns P2 exact clean package verification, final upstream package
pairing, isolated GitHub CLI routing, explicit legacy release compatibility and
fresh independent review. Fresh-cache production default verification awaits the
public SDP v2.0.0 descriptor in P3. Worker stops after the bounded P1 commit.
