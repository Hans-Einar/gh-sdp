# VER-SPS-003 — GIP-3-M2 client evidence

Status: passed local verification; fresh independent review pending.
Slice: SPS-003. Release gate: false. Date: 2026-09-27.
Host: Linux amd64; Go go1.27.1; GitHub CLI 2.97.0.
Canonical bootstrap source: 266179e87452272e9f513f0d8f7acb3cdf9e3051.

Client candidate `02e3db23c1b40894d26fbb57db0fb77bfb2c364f` was rebuilt from a
clean tree and passed all checks below. [Exact command/output log](evidence/GIP-3-M2-tests.txt)
records candidate, toolchains, package hashes and results. No native
Windows/macOS execution, production trust or release support is claimed.

## Reproducible checks

```sh
export GOTOOLCHAIN=local
export PATH=/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin:$PATH
go build -o /tmp/gip-gh-sdp .
go test -race ./...
go vet ./...
GH_SDP_BINARY=/tmp/gip-gh-sdp SDPTOOL_BINARY=/tmp/gip3m1-package/sdptool go test -race -count=1 -v ./...
GH_SDP_VIA_GH=true GH_SDP_BINARY=/tmp/gip-gh-sdp SDPTOOL_BINARY=/tmp/gip3m1-package/sdptool go test -race -count=1 -v ./...
git diff --check
```

The paired SDPTool package was built by the coordinating GIP-3-M1 assignment.
It reports `c279d8a2367ec4ea4db3c891a3687352c1434256-dirty` (precommit
GIP-3-M1 source), with SHA256
`550220f27849ff01f15f6ac3e7838a2c372b6ea75443aa693983cedc8844607e`.
This is binary-hash-bound paired evidence, not a claim that its embedded build
revision equals the bootstrap module commit. Final GIP-4 pairing must rerun
against its actual final engine candidate. Client SHA256 is
`1a4c08327ba490efb1297e344a997761168987201f624f191815c9d157a11f9a`.
The integration test creates test keys, signatures, descriptors, plans and
projects in temporary directories. It never updates gh-sdp's installed Toolkit
or any live consumer. The optional gh route uses isolated GH_CONFIG_DIR/XDG_DATA_HOME and
installs only the temporary local extension.

## Results

- Unit/race: exact argv elements (spaces, empty argument, metacharacters), binary
  stdin, stdout/stderr, cwd/environment and child exit 23 pass; invalid boolean
  exits 2, bootstrap/start errors exit 4, saved apply/resume selects offline.
- Packaged candidate: exact direct/client preview JSON and saved-plan bytes
  match. Both entry points apply the same root-bound plan after restoring the
  disposable baseline. AGENTS.md becomes managed, original project instructions
  remain byte-identical in AGENTS-project.md. Repeated preview has zero actions.
- Saved apply succeeds after removing the original descriptor; the wrapper
  obtains its engine from verified cache and SDPTool uses signed plan facts.
- Wrong descriptor protocol, a signed digest-valid binary with incompatible
  protocol response and offline corrupt engine cache all return exit 4 without
  child operation output.
- The same packaged checks pass through actual `gh sdp` (GitHub CLI 2.97.0).
- YAML/NDJSON parse and current-coordinate checks pass; managed AGENTS.md,
  installed Toolkit manifest and append-only ledger match origin/main bytes.
- `go vet ./...` and `git diff --check` pass.

Signature, wrong-key, download bounds, immutable cache, offline-miss and timeout
unit evidence belongs to the pinned upstream bootstrap package. This client
record does not relabel those tests as independently repeated local evidence.
Power loss, signals/cancellation and native non-Linux execution remain outside
this bounded evidence. Independent review must inspect the actual paired source
candidate and the log before Slice closure.
