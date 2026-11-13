# gh-sdp

GitHub CLI distribution of the canonical SDPTool engine. The selected 2.3.0
candidate publishes the **exact signed canonical engine bytes** as
`gh-sdp-linux-amd64`. `gh sdp` executes that engine directly. No runtime bootstrap,
child-engine download, private engine cache or separate client version is involved.
The latest published historical bootstrap client remains 0.2.3 with SDP2.2.1;
this branch does not claim 2.3.0 publication.

Install or update the shared tool through GitHub CLI:

```sh
gh extension install Hans-Einar/gh-sdp
gh extension upgrade sdp
gh sdp setup
```

Explicit setup registers `$HOME/.local/bin/sdptool` as a link to the current
executable, without assuming a GitHub extension directory layout. Put
`$HOME/.local/bin` on PATH for shell use. Native consumers resolve that documented
path directly. Ordinary gh binary extension upgrades replace the same target
pathname, so registered consumers immediately execute the new engine even before
the next `gh sdp` call. Local commands work without gh, credentials or network:

```sh
"$HOME/.local/bin/sdptool" --version --json
"$HOME/.local/bin/sdptool" discover --json
```

SDPTool owns setup, actions, discovery, installation policy, signed framework
selection, preview/apply, recovery and preservation. Framework selection through
`SDP_RELEASE` does not choose a different executing tool. A tool update does not
upgrade project SDP directories. Preview and apply project upgrades explicitly.

```sh
gh sdp upgrade --manifest /absolute/path/to/adoption.json --plan-output /absolute/path/to/plan.json
gh sdp upgrade --apply /absolute/path/to/plan.json
```

Human output is the default; machine consumers request `--json`. Declared
project programs remain project-owned (`gh sdp run --program NAME`). Native
release qualification currently covers Linux amd64. Windows/macOS need native
checks before support claims.

## Package a candidate

Use Bash, Git, Python3 and Go1.26+, a clean checkout and a new output directory
outside it. The script verifies an explicitly selected signed canonical release,
checks engine version/revision against its signed descriptor and republishes its
bytes without compiling another engine:

```sh
scripts/package.sh 2.3.0 /absolute/path/to/sdp-release.json /absolute/path/to/new-package
```

HTTPS descriptors are supported. Only build-time verification uses the immutable
canonical bootstrap dependency. It is never shipped as a launcher. Outputs include
the executable, checksum, package manifest, exact signed descriptor/signature and
verification envelope. The manifest distinguishes canonical engine source from
extension distribution source. Explicit `SDP_TEST_KEY` permits nonproduction
fixtures, visibly marked `test-signed`; these do not qualify production publication.
See the [release gate](docs/Release-Lifecycle.md).

## Verify

```sh
go test -race ./...
go vet ./...
GH_SDP_PACKAGE_DIR=/absolute/path/to/package go test -run TestPackagedAssetDiscovery -v .
GH_SDP_BINARY=/absolute/path/to/gh-sdp-linux-amd64 SDPTOOL_BINARY=/absolute/path/to/sdptool go test -run TestPackagedCandidates -v .
SDPTOOL_OLD_BINARY=/absolute/path/to/previous-registered-engine GH_SDP_BINARY=/absolute/path/to/gh-sdp-linux-amd64 go test -run TestActualGH -v .
```

The last gate runs actual gh binary-extension install/upgrade against an isolated
TLS release fixture, expires gh's HTTP metadata cache between simulated releases,
and verifies stable setup registration, byte equality before the first post-update
gh invocation, and direct offline version/discovery without gh on PATH or auth.
Both engine inputs must support setup. It does not prove public release publication.
The installer gate uses disposable signed framework fixtures and checks direct/gh
output, preview/apply and project preservation. No tests modify the user's extension.

[Current bounded work](SDP/Sprints/Sprint-011/ScrumIterations.md)

An interrupted gh2.97.0 download can leave a partial executable; direct consumers
fail without falling back. Forced update may report the current tag up to date.
The isolated recovery gate verifies extension removal and reinstall after HTTP
cache expiration restores the same registered target. See the
[known gh limitation](SDP/Sprints/Sprint-011/Evidence.md#interrupted-update-discovery).
