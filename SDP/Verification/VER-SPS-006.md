# VER-SPS-006 — Template distribution default patch gate

Status: previous v0.2.2 candidate approved; v1.0.0 verification pending; publication paused
Slice: SPS-006
Release: REL-0.1.2
Platform: Linux amd64
Product candidate: f7b690633413b0a1f13b49d3e6265bd3494d1995
Bootstrap: v0.0.0-20260927231627-591920b69534 (remote, no replacement)

## Candidate evidence

Worker implementation notes record Go 1.27.1 race tests, vet, module verification
and the literal v0.2.2 default URL regression. Master inspected the actual diff:
no production adapter changes; dependency, literal expectation and current
documentation select the new default while historical release records remain.

Master built the clean exact product candidate using Go 1.27.1 on Linux amd64:

```sh
SDP_GO=/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go scripts/package.sh 0.1.2 /tmp/gh-sdp-sps006-candidate
GH_SDP_PACKAGE_DIR=/tmp/gh-sdp-sps006-candidate GH_SDP_BINARY=/tmp/gh-sdp-sps006-candidate/gh-sdp-linux-amd64 GOTOOLCHAIN=local /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -count=1 -run 'TestPackagedAssetDiscovery|TestPackagedClientRejectsFIFO' -v ./...
```

Both commands exited 0. Actual package discovery and packaged FIFO rejection
passed. SHA-256: `1cce96074c58ee9c43583d0724628294c0c58f51d4d05f0345672d748fb447ce`.
The [candidate manifest](evidence/SPS-006-candidate-manifest.json) identifies
exact clean source, dependency and binary bytes. Independent review, fresh-cache
public production default execution, clean merged packaging and publication
are pending; no Windows/macOS, global installation or XFMD evidence is claimed.

## Independent review and owner pause

REV-SPS-006-001 independently reproduced the clean candidate package bytes,
passed race tests, vet, module verification, discovery and FIFO checks, and
confirmed fresh remote module provenance. P2 is approved for this candidate.
The owner then paused integration/publication pending a revised Go-only upstream
SDP release. P3 and global extension installation remain pending; root must
confirm the revised candidate and bootstrap pin before continuation.

## Revised upstream target

The owner subsequently selected SDP v1.0.0 to remove public PowerShell entrypoints.
The client stays at target v0.1.2. This record and REV-SPS-006-001 prove only the
previous v0.2.2 candidate, not the new target. Root has not supplied its gated
replacement bootstrap SHA. Updated pin, regression and exact-candidate checks
remain pending, as do integration, publication and installation.
