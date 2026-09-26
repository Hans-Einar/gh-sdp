# gh-sdp

A thin GitHub CLI extension that obtains a verified SDPTool executable and
forwards arguments, streams and exit status. SDPTool owns installation policy,
preview, apply, recovery and project preservation.

This is an unreleased development candidate. No production signing key, release
catalog or published binary is selected. Native verification currently covers
Linux amd64 only; Windows/macOS support requires native tests.

## Build and use

Use Go 1.26 or later (verification uses Go 1.27.1):

```sh
go build -o gh-sdp .
gh extension install .
```

Configure an explicit signed test distribution supplied by the SDPTool workstream:

```sh
export SDP_RELEASE=/absolute/path/to/release.json
export SDP_TEST_KEY=/absolute/path/to/nonproduction-public-key.txt
# Optional private cache directory; otherwise the user's SDPTool cache is used.
export SDP_CACHE_DIR=/absolute/path/to/private-cache

gh sdp upgrade --manifest /absolute/path/to/adoption.json --plan-output /absolute/path/to/plan.json
gh sdp upgrade --apply /absolute/path/to/plan.json
```

Run from the project directory. Preview requires a fresh external plan path;
apply uses that exact root-bound plan. Saved apply/resume forces offline bootstrap
from the verified cache populated during preview. For other commands, set
`SDP_OFFLINE=true` to require cached distribution inputs. The explicit test key
is required again for cached test distributions. It is a base64 Ed25519 public
key; its selection never means a production release is trusted or supported.
The signed descriptor's `.sig` and assets follow the canonical SDPTool contract.

All command arguments reach SDPTool unchanged. Invalid local configuration exits
2, bootstrap/start failures exit 4, and ordinary child exit codes pass through.
No shell or PATH lookup selects the engine. `gh sdp --version` reports the child
SDPTool identity; it does not claim that the client has that version. The client
release remains independently proposed as 0.1.0 and unreleased.

## Verification

```sh
go test -race ./...
GH_SDP_BINARY=/absolute/path/to/gh-sdp SDPTOOL_BINARY=/absolute/path/to/sdptool go test -race -v ./...
GH_SDP_VIA_GH=true GH_SDP_BINARY=/absolute/path/to/gh-sdp SDPTOOL_BINARY=/absolute/path/to/sdptool go test -race -v ./...
```

The packaged-candidate test is skipped unless both binary paths are supplied.
It creates disposable signed fixtures and projects and compares direct and
client preview/apply. The optional GitHub CLI route installs the local extension
only into a temporary isolated `GH_CONFIG_DIR`.

[Current Slice and authority](SDP/Sprints/Sprint-003/ScrumIterations.md) ·
[Verification](SDP/Verification/VER-SPS-003.md) ·
[Canonical dependency provenance](SDP/Sprints/Sprint-003/implementationNotes.md)
