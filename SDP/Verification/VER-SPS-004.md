# VER-SPS-004 — First native release gate

Status: preparation passed; independent REV-SPS-004-001 approved
Slice: SPS-004
Release: REL-0.1.0
Platform: Linux amd64
Product candidate: 2618e3857a9cf1a7165090ec5209727593fa3221
Bootstrap: v0.0.0-20260927075424-659e543e89d7 (remote, no replacement)

## Candidate evidence

Worker ran Go 1.27.1 race tests and vet successfully. The shared default and
explicit SDP_RELEASE override agree between bootstrap and the actual delegated
child; original argv, streams, cwd, exit and parent environment are preserved.

Two clean detached candidate packages produced identical binary SHA-256:
a76438e039d71fb9bfaec061ed91b9419c55ae2bf2158454e3684bad570a7874.
The [package manifest](evidence/SPS-004-candidate-manifest.json) records source,
module and toolchain. Dirty source and occupied output rejection passed.

The actual isolated GitHub CLI integration passed in 3.153 seconds against
GIP4 engine 07335b0, SHA-256
0d984b548f47b6ecb3b5049a7b2c24268822cf1032fdaeb373bb7574745fc18e.
It checks signed-fixture direct/client preview parity, apply, preserved project
content, repeat no-change, incompatible protocol, corrupt cache and FIFO rejection.
This is fixture evidence; it is not a production descriptor or live XFMD test.

## Commands

```sh
GOTOOLCHAIN=local go test -race -count=1 ./...
GOTOOLCHAIN=local go vet ./...
SDP_GO=/absolute/path/to/go scripts/package.sh 0.1.0 /tmp/fresh-package
GH_SDP_VIA_GH=true GH_SDP_BINARY=/tmp/gh-sdp-r1-final-package/gh-sdp_linux_amd64 SDPTOOL_BINARY=/tmp/gip4-exact-package/sdptool GOTOOLCHAIN=local go test -race -count=1 -v ./...
```

Independent [REV-SPS-004-001](../CodeReview/REV-SPS-004-001.md) reproduced race/vet,
module verification, exact package hash and actual isolated gh integration. It also
verified a local production-signed descriptor without SDP_TEST_KEY, yielding 63
install actions with signed provenance and no project writes. All three draft
documentation findings are resolved; exact preparation 4660b6b is approved.

The
final merged release commit is rebuilt cleanly and its identity/checksum recorded
at publication. Public descriptor availability and actual GitHub extension
installation must be verified then; neither is inferred from these fixtures.
