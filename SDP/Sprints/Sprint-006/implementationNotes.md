# SPS-006 P1 implementation notes

Authorization: SPS-006 P1, assigned by the Master on 2026-09-28.
Baseline: contract commit 00a4737 on sdp/release-0.1.2.
Status: P1 implemented and locally verified; fresh review and release gates remain pending.

## Implementation and provenance

The selected canonical DefaultRelease is
`https://github.com/Hans-Einar/SDP/releases/download/v0.2.2/sdp-release.json`.
TestSelectedReleaseReachesChild compares against this literal URL for both unset
and empty SDP_RELEASE. It checks bootstrap configuration and an actual child
process environment, retaining the explicit path-with-spaces override case.

The client pins the canonical remote nested module
`github.com/Hans-Einar/SDP/SDPTool/bootstrap` at
`v0.0.0-20260927231627-591920b69534`. The Master supplied immutable source commit
`591920b69534fa12e587f5f95a55b6bce9b13da4`; `git ls-remote` confirmed that exact
revision on the canonical remote's `refs/heads/sdp/mvp1-source-ui`. Go retrieved
the revision with `go get`; `go mod tidy` removed the superseded checksums.
No local module replacement is used.

Module content checksum:
`h1:W82Nb7A4D+LGsbXtJwyuhBsXlmaYRLA0wFxmwQxgLyo=`.

No production adapter code or API changes are selected. README, current
architecture/design, release instructions and selected 0.1.2 release notes
identify the candidate while preserving prior release history. SDPTool owns the
template distribution; this client selects its paired default release.

## Verification

Executed in `/home/warloc/git/gh-sdp` with Go 1.27.1, Linux amd64
(Linux 6.12.0-211.53.1.el10_2.x86_64). All commands exited 0:

```sh
git ls-remote https://github.com/Hans-Einar/SDP.git refs/heads/sdp/mvp1-source-ui
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go get github.com/Hans-Einar/SDP/SDPTool/bootstrap@591920b69534fa12e587f5f95a55b6bce9b13da4
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go mod tidy
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -count=1 -v ./...
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go vet ./...
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go mod verify
/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go list -m -json all
git diff --check
```

Race tests passed, including exact default delegation for unset/empty SDP_RELEASE,
explicit override handling, argv/streams/cwd/inherited environment/exit behavior
and configuration/failure boundaries. Module verification reported
`all modules verified`; the module listing contained only this project and the
pinned canonical bootstrap, with no Replace.

The optional packaged-candidate, asset-discovery and packaged FIFO tests skipped
because their binary/package environment variables were not supplied.

Clean exact-commit packaging, package discovery, FIFO rejection and fresh
independent review belong to P2. Public production-trust default resolution
belongs to P3 after root confirms SDP v0.2.2 publication. No Windows/macOS
verification, publication, merge, global installation or XFMD adoption is claimed.

Master-owned manifests, release records and traceability are outside this Worker
change. The Worker stops after the bounded verified P1 commit.
