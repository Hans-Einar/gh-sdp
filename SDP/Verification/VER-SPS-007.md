# VER-SPS-007 — Discovery client release gate

Status: initial P2 preparation independently approved; final upstream pairing and P3 pending
Slice: SPS-007
Release: REL-0.2.0
Platform: Linux amd64 only
Product candidate: 336d654c7533a30b340a1366109a41a9307ddf1e
Bootstrap: v0.0.0-20260930230248-247fb7aba628 (root supplied, remotely resolved)

## Required evidence

- Go 1.27.1 race tests, vet and module verification, with GOMAXPROCS=2 and test -p 2.
- Remote immutable module resolution with no local replacement.
- Clean exact-commit package manifest, checksum, asset discovery and FIFO rejection.
- Paired SDPTool/direct and client execution: human default, explicit JSON,
  no-argument source discovery, existing preview/apply and preservation checks.
- Isolated GitHub CLI routing with temporary configuration, without changing the
  user's global extension.
- Fresh independent source/package review.
- Public paired SDP v2.0.0 descriptor and fresh-cache production-trust human/JSON
  execution before client publication.
- Actual annotated tag, release assets and downloaded checksums followed by
  independent reconciliation review.

No checks above are passed by this placeholder. No merge, global installation,
Windows/macOS native support or live XFMD outcome is claimed.

## Initial record validation

The project manifest and new release record validate against the current Toolkit
JSON schemas, and the new version-selection ledger event validates against the
release-event schema. All edited YAML and all NDJSON lines parse successfully.
The current-index schema rejects the pre-existing historicalStudy,
historicalDependencyAssessment, lifecycle and lastCompleted extensions, also
present in the clean baseline. As documented in VER-SPS-006, process migration
is not part of this client release; no complete current-Toolkit schema
compatibility is claimed.

## Clean exact candidate checks

Release source candidate: 72b4e048a90841a1994e9846bb4db0d0ad9b922e.
Go 1.27.1 built the clean checkout with scripts/package.sh 0.2.0 into
/tmp/gh-sdp-sps007-candidate. The checked-in candidate manifest records Linux
amd64, no replacement, vcs.modified=false, size 10043233 and SHA-256
9c7471817285647e3ca12c7ef0618171bd09021e43a9a8392a3779c85b0ba72b.

Commands exited zero:

```sh
env GOMAXPROCS=2 SDP_GO=/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go scripts/package.sh 0.2.0 /tmp/gh-sdp-sps007-candidate
env GOMAXPROCS=2 GOTOOLCHAIN=local GH_SDP_PACKAGE_DIR=/tmp/gh-sdp-sps007-candidate GH_SDP_BINARY=/tmp/gh-sdp-sps007-candidate/gh-sdp-linux-amd64 SDPTOOL_BINARY=/tmp/sdptool-session3 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -p 2 -count=1 -v ./...
env GOMAXPROCS=2 GOTOOLCHAIN=local GH_SDP_VIA_GH=true GH_SDP_PACKAGE_DIR=/tmp/gh-sdp-sps007-candidate GH_SDP_BINARY=/tmp/gh-sdp-sps007-candidate/gh-sdp-linux-amd64 SDPTOOL_BINARY=/tmp/sdptool-session3 /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -p 2 -count=1 -run TestPackagedCandidates -v ./...
env GOMAXPROCS=2 GOTOOLCHAIN=local /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go vet -p 2 ./...
```

The candidate-tests and gh-routing logs record actual packaged discovery, FIFO,
automatic source discovery without navigation.json, sdptool/0.2 inventory and
inline navigation, human/JSON output, diagnostics, preview/apply, preservation,
repetition, incompatible protocol and corrupt cache checks. GitHub CLI used
isolated temporary configuration/data and did not change the user's installation.
The initial paired engine SHA-256 is
4383e8b8539eb23aeba5b8333c6cdb4b99f827dc039014cba67fd38794747f8c;
it reports 2.0.0-dev and revision unknown, so these checks are rehearsal evidence.
Final upstream release pairing is still required.

The same clean client passed a new production-trust cache with explicit
SDP_RELEASE=https://github.com/Hans-Einar/SDP/releases/download/v1.0.0/sdp-release.json,
SDP_OFFLINE=false and no SDP_TEST_KEY. Its --version returned upstream 1.0.0 at
fede327d6f3af35fe7aad323e2c485a27134d20d, recorded in SPS-007-legacy-version.json.
This confirms legacy explicit selection and the bootstrap compatibility fallback;
it does not establish new default production availability.

## Independent preparation disposition

REV-SPS-007-001 independently approves the reviewed preparation source and exact
client package with no findings. The Reviewer reran all race tests, isolated
GitHub CLI routing, vet, module verification, pinned bootstrap tests and a new
production-trust legacy override. Final immutable upstream release pairing and
public SDP v2.0.0 default verification remain pending; releaseGate remains false.

## Authoritative upstream package pairing

The unchanged exact client package above passed the full race suite with
SDPTOOL_BINARY=/tmp/sdp-session3-release-package/sdptool, then passed the isolated
GitHub CLI TestPackagedCandidates route against the same final binary. Commands
match the initial package checks with only the supplied engine path changed.
Actual outputs are retained as SPS-007-final-pair.txt and
SPS-007-final-gh-routing.txt. No tests were skipped.

Final engine SHA-256:
fa0c4a9f387a7752fcc5928ed267c4de13c54f50d80d709cbedc1b9cae42325c.
A separate fresh-cache production-trust invocation used the signed local
/tmp/sdp-session3-release-package/sdp-release.json with no SDP_TEST_KEY and
SDP_OFFLINE=false. The client --version --json returned version 2.0.0,
schema sdptool/0.2 and revision 6f9af514c99076d9067b3fbbfc93193a2f986070;
SPS-007-final-signed-version.json retains the response. This establishes the
actual immutable signed pair locally. It does not establish public descriptor
availability, which remains a P3 requirement before client publication.
