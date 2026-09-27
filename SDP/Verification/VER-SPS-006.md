# VER-SPS-006 — Template distribution default patch gate

Status: publication and extension installation verified; reconciliation review pending
Slice: SPS-006
Release: REL-0.1.2
Platform: Linux amd64
Product candidate: 6cd4c3e3b16d265a79d2e8879fd964cb1c8ccbd2
Bootstrap: v0.0.0-20260927233033-705df7e3d550 (remote, no replacement)

## Historical v0.2.2 candidate evidence

Historical product candidate: f7b690633413b0a1f13b49d3e6265bd3494d1995.
Historical bootstrap: v0.0.0-20260927231627-591920b69534.

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

## Current revised P1/P2 evidence

Exact product candidate: 6cd4c3e3b16d265a79d2e8879fd964cb1c8ccbd2. Bootstrap:
v0.0.0-20260927233033-705df7e3d550. Independent REV-SPS-006-002 passed full
race tests including packaged discovery/FIFO rejection, vet, module verification
and direct remote dependency retrieval. No adapter or package source changes.
The [revised candidate manifest](evidence/SPS-006-go-only-candidate-manifest.json)
records clean exact-source provenance, Linux amd64, size 10042505 and SHA-256
2a02675a9f61979ae781b8a7bfb855167c6db61d08a4d5e2a38be64a6d9c318b.
Root inspected these records and confirmed the new upstream production candidate.
P3 public default execution, merged packaging and publication remain to be recorded.

## Public production default gate

The reviewed packaged client was executed with a new empty SDP_CACHE_DIR and no
SDP_RELEASE override after upstream publication. It downloaded and verified the
public signed descriptor and executable, then returned SDPTool version 1.0.0
and revision fede327d6f3af35fe7aad323e2c485a27134d20d, exit zero. Public upstream
release: https://github.com/Hans-Einar/SDP/releases/tag/v1.0.0. No local descriptor
or trust override was used. Final client merge/build/publication remain pending.

## Additional process-schema inspection

The current upstream Toolkit project validator was tried as an additional check,
not a selected native release gate. It rejects this project's pre-existing custom
study/traceability extensions and historical release-note Scope headings, also
present on clean main 9ef2736. No compatibility with that validator is claimed.
The new upstreamTargetVersion field is retained only in extensible Relations, not
the strict CurrentIndex release object. Migrating historical process records is
outside this default-engine patch; native client release gates above govern.

## P3 actual publication and installation

PR10 merged to a4988846432c7f7f7b3f722f2786359c9b127372. The clean merged source package
passed actual discovery/FIFO tests with the race detector. Its annotated v0.1.2
tag and GitHub Release were published at 2026-09-27T23:42:46Z:
https://github.com/Hans-Einar/gh-sdp/releases/tag/v0.1.2

Downloaded assets pass checksums.txt and match the published manifest. SHA-256:
14c4824f765e4621260714746ad0c91e66c3167ae75d1dac6b2b76a50311ae41.
The installed extension was upgraded with `gh extension upgrade sdp` from v0.1.1
to v0.1.2. `gh sdp --version` with its ordinary production configuration returns
SDPTool 1.0.0 at fede327d6f3af35fe7aad323e2c485a27134d20d. gh-tree remains installed.
Root owns the separate live XFMD upgrade evidence. Independent publication-record
reconciliation review is pending; historical released records remain unchanged.
