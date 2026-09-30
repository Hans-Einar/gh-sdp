# gh-sdp

A thin GitHub CLI extension that obtains a verified SDPTool executable and
forwards arguments, streams and exit status. SDPTool owns installation policy,
preview, apply, recovery and project preservation.

The latest published client is v0.1.2, paired with signed Go-only SDP v1.0.0.
Its immutable bootstrap, clean package and public production default were
independently verified. Install with `gh extension install Hans-Einar/gh-sdp`,
or update with `gh extension upgrade sdp`; no local Go or PowerShell is needed.
Native verification covers Linux amd64 only; Windows/macOS support requires native tests.

## Build and use

Use Go 1.26 or later (verification uses Go 1.27.1):

```sh
go build -o gh-sdp .
gh extension install .
```

Install the published Linux amd64 GitHub CLI extension with:

```sh
gh extension install Hans-Einar/gh-sdp
```

The v0.2.0 candidate selects the exact SDP v2.0.0 release descriptor.
This candidate is not yet published.
Without `SDP_RELEASE`, the client uses that default for both verification and the
SDPTool child. An explicit `SDP_RELEASE` overrides both selections. No key option
is needed for the production release: public trust is compiled into the shared
bootstrap. The private signing key is never part of the client or package.

The paired v2.0.0 engine discovers project sources and defaults to human-readable
output. Machine consumers must explicitly request `--json`, including for
`--version`. For example, from a project checkout:

```sh
gh sdp discover
gh sdp discover --json
gh sdp --version --json
```

Preview an existing manual installation with:

```sh
gh sdp upgrade --manifest /absolute/path/to/adoption.json --plan-output /absolute/path/to/plan.json
gh sdp upgrade --apply /absolute/path/to/plan.json
```

An adoption manifest must describe that project's actual current inventory; the
client does not generate or guess it. `gh sdp install --plan-output ...` previews
a clean installation. The default descriptor requires a published paired SDP
release and network access on the first call, then uses the verified cache.

For an explicit nonproduction signed test distribution instead:

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
`SDP_OFFLINE=true` to require cached distribution inputs. Only lowercase `true`
and `false` are accepted so the child receives the same selection. The explicit test key
is required again for cached test distributions. It is a base64 Ed25519 public
key; its selection never means a production release is trusted or supported.
The signed descriptor's `.sig` and assets follow the canonical SDPTool contract.

All command arguments reach SDPTool unchanged. Invalid local configuration exits
2, bootstrap/start failures exit 4, and ordinary child exit codes pass through.
No shell or PATH lookup selects the engine. `gh sdp --version` reports the child
SDPTool identity; it does not claim that the client has that version. The client
release identity is recorded independently in the package manifest; the client
forwards the engine's version output unchanged.

## Package a release candidate

The script requires Bash, Git, Python 3, and Go. Use a clean checkout and an output
directory outside it; existing output files and local module replacements are rejected.
No downloads or publication are performed by the script except Go's normal immutable
module retrieval during compilation.

```sh
SDP_GO=/absolute/path/to/go scripts/package.sh 0.2.0 /absolute/path/to/package
```

Outputs are `gh-sdp-linux-amd64` (the asset naming recognized by GitHub CLI),
`checksums.txt`, and `gh-sdp.manifest.json`. The manifest records the exact clean
source commit, platform, byte count, SHA-256 and Go dependency/build identities.
Builds disable CGO and use `-trimpath`; the same clean source and Go toolchain
produce the same executable. Publication requires the separate
[release gate](docs/Release-Lifecycle.md).

## Verification

```sh
go test -race ./...
GH_SDP_PACKAGE_DIR=/absolute/path/to/package go test -race -run TestPackagedAssetDiscovery -v ./...
GH_SDP_BINARY=/absolute/path/to/gh-sdp SDPTOOL_BINARY=/absolute/path/to/sdptool go test -race -v ./...
GH_SDP_VIA_GH=true GH_SDP_BINARY=/absolute/path/to/gh-sdp SDPTOOL_BINARY=/absolute/path/to/sdptool go test -race -v ./...
```

The asset-discovery test checks the real package using GitHub CLI's `linux-amd64`
suffix rule, then verifies its manifest and checksum; it requires `GH_SDP_PACKAGE_DIR`.
The packaged-candidate test is skipped unless both binary paths are supplied.
It creates disposable signed fixtures and projects and compares direct and
client human/JSON version, discovery, preview and diagnostic output, then
preview/apply.
The optional GitHub CLI route installs the local extension only into a temporary
isolated `GH_CONFIG_DIR` and `XDG_DATA_HOME`.

[Current Slice and authority](SDP/Sprints/Sprint-007/ScrumIterations.md) ·
[Latest published verification](SDP/Verification/VER-SPS-006.md) ·
[Candidate dependency provenance and verification status](SDP/Sprints/Sprint-007/implementationNotes.md)
