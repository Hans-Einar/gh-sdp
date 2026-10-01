# VER-SPS-008 — SDP 2.1 default patch

Status: native checks passed; independent release approval pending
Slice: SPS-008
Release: REL-0.2.1
Exact source: 315ec1de9dbe9aee990baae47a1ac2a7a07a0702
Platform: Linux amd64

## Exact source and native gate

Clean source built with Go 1.27.1 using
`env GOMAXPROCS=2 SDP_GO=/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go scripts/package.sh 0.2.1 /tmp/gh-sdp-sps008-candidate`.
The manifest records vcs.modified=false, no replacement, immutable bootstrap
v0.0.0-20261001092620-93517ad98cd1 and client binary SHA-256
230440b0bbeb5af33f98b91bd227b26f78e3deefba9e23dee1e1508e75061dca.
P1 dependency provenance, race tests and vet are in implementationNotes.md.
Master reran vet against exact source, exiting zero.

The full exact-package suite exited zero, with no skipped tests:

```sh
env -u SDP_RELEASE -u SDP_TEST_KEY -u SDP_CACHE_DIR -u SDP_OFFLINE GOMAXPROCS=2 GOTOOLCHAIN=local GH_SDP_VIA_GH=true GH_SDP_PACKAGE_DIR=/tmp/gh-sdp-sps008-candidate GH_SDP_BINARY=/tmp/gh-sdp-sps008-candidate/gh-sdp-linux-amd64 SDPTOOL_BINARY=/tmp/sdp-21-package/sdptool /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -p 2 -count=1 -v ./...
```

The paired engine SHA-256 is
691c021751cefb35c7f92c902545cd6c5d71ccadfcde07ae507a6ce026c8a823,
from reviewed upstream source 93517ad98cd188c0debeb1d0f3d36d123c6e4a3b.
SPS-008-candidate-tests.txt records automatic discovery, human/JSON output and
diagnostic equality, preview/apply, preservation/repetition, incompatible/corrupt
cache, default/override selection, package discovery and FIFO rejection. gh routing
uses temporary configuration/data. The owner's global extension was untouched.

## Public production default

Root published SDP v2.1.0 at 2026-10-01T13:24:49Z. Independently downloaded
public descriptor SHA-256 is
d8436398ff6ed1cf74b7754106792075203dc3745f91d19c56450951b7302dde.
The exact client ran with all inherited SDP_/GH_SDP_ controls removed,
SDP_OFFLINE=false and new empty cache /tmp/gh-sdp-sps008-public-cache-2d4i6184.
No test key or release override was used. --version, --version --json and
discover --json all exited zero with empty stderr. Checked-in public outputs
identify SDPTool 2.1.0, exact upstream revision and sdptool/0.2. Discovery exposes
sessions capability as absent in this repository, consistent with its sources.
This read-only probe did not install or update project files.

## CI and record validation

GitHub API reports zero gh-sdp workflows. Per the bounded contract, local native
release checks apply and no client CI result is claimed. Exact upstream source
has successful contracts and go-installation-linux jobs:
https://github.com/Hans-Einar/SDP/actions/runs/36842826943
That upstream CI is separate from the client evidence above.

All SDP YAML and ledger lines parse. Project manifest, new release record and
new release events validate against current Toolkit JSON schemas read-only.
The established historical current-index extension limitation remains outside
this patch; no process migration or complete current-schema conformity is claimed.

Independent source/package/public-default review and actual client publication
remain pending. No main merge, global extension upgrade or XFMD change is claimed.
