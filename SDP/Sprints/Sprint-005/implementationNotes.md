# SPS-005 P1 implementation notes

Authorization: SPS-005 P1, assigned by the Master on 2026-09-27.
Baseline: contract commit d43290eb8653d07b5e80f574fcfb518a29e49e70 on
sdp/release-0.1.1.
Status: P1 implemented and locally verified; fresh review and release gates remain pending.

## Implementation and provenance

The client pins the canonical remote nested module
`github.com/Hans-Einar/SDP/SDPTool/bootstrap` at
`v0.0.0-20260927094008-c1901337510e`. The immutable source commit is
`c1901337510e13ecf4f7fda64443729c816f7f42`, resolved from the actual remote
`refs/heads/sdp/release-r2` using `git ls-remote https://github.com/Hans-Einar/SDP.git`.
Go retrieved this remote revision with `go get` and `go mod tidy` removed the
superseded checksums. No local module replacement is used.

Module content checksum:
`h1:iGb4QLoIYpnXeRuW0nXOrqNQqoccADoj3Gd0muvQWiY=`.
The canonical DefaultRelease selects
`https://github.com/Hans-Einar/SDP/releases/download/v0.2.1/sdp-release.json`.
The strengthened TestSelectedReleaseReachesChild compares against this literal
URL for both unset and empty SDP_RELEASE. It checks bootstrap configuration and
an actual child process environment, retaining the explicit path-with-spaces
override case. The expectation no longer derives from the imported constant,
so an unintended default change cannot silently change both sides of the test.

No production adapter code or API changes were needed. README, current
architecture/design, generic release instructions and selected 0.1.1 release
notes now identify the candidate accurately while preserving v0.1.0 history.
The upstream engine owns the external KanBan correction; this client change
selects its paired release rather than copying installation policy.

## Verification

Executed in `/home/warloc/git/gh-sdp` with Go 1.27.1, Linux amd64
(Linux 6.12.0-211.53.1.el10_2.x86_64). All commands exited 0:

```sh
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -count=1 -v ./...
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go vet ./...
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go mod verify
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go list -m -json all
git diff --check
```

Race tests passed, including unset/empty/override descriptor delegation, exact
argv/streams/cwd/inherited environment/exit behavior and configuration/failure
boundaries. Module verification reported `all modules verified`; module listing
contained only this project and the pinned canonical bootstrap, with no Replace.

The optional packaged-candidate, asset-discovery and packaged FIFO tests skipped
because their binary/package environment variables were not supplied. These
checks, clean exact-commit packaging and fresh independent review belong to P2.
Public production-trust default resolution belongs to P3 after root confirms
SDP v0.2.1 publication. This evidence does not claim those gates, Windows/macOS
verification, publication, a merge, global installation or XFMD adoption.

Master-owned manifests, release records and traceability are outside this Worker
change. The Worker stops after the bounded P1 commit.
