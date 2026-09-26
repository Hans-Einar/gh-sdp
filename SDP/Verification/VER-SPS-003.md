# VER-SPS-003 — GIP-3-M2 client evidence

Status: passed local verification; fresh independent review pending.
Slice: SPS-003. Release gate: false. Date: 2026-09-27.
Host: Linux amd64; Go go1.27.1; GitHub CLI 2.97.0.
Canonical bootstrap source: fb79727aaa229e5a4a5f228f458d350cd196a6a8.

Client candidate `fca8480b1619f325fbbbd7b9d7224cbf865dbcdf` was rebuilt from a
clean tree and passed all checks below. [Final command/output log](evidence/GIP-3-M2-final-tests.txt)
records candidate, toolchains, package hashes and results. No native
Windows/macOS execution, production trust or release support is claimed.

## Reproducible checks

```sh
export GOTOOLCHAIN=local
export PATH=/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin:$PATH
go build -o /tmp/gip-gh-sdp .
go test -race ./...
go vet ./...
GH_SDP_BINARY=/tmp/gip-gh-sdp SDPTOOL_BINARY=/tmp/gip3-exact-package/sdptool go test -race -count=1 -v ./...
GH_SDP_VIA_GH=true GH_SDP_BINARY=/tmp/gip-gh-sdp SDPTOOL_BINARY=/tmp/gip3-exact-package/sdptool go test -race -count=1 -v ./...
git diff --check
```

The paired SDPTool package was built from a detached clean canonical worktree at
`fb79727aaa229e5a4a5f228f458d350cd196a6a8`; its embedded revision agrees.
SDPTool SHA256: `b80504b1a1c89f586a02b938e7a7f514fc415b8b14ff005d3c35693076293007`.
Client SHA256: `f7fa971d950591512919c65ede5d009d65cce2d2d4c25c33b91ecfd640af50f8`.
The [initial precommit pairing log](evidence/GIP-3-M2-tests.txt) remains historical
and is superseded by this exact clean pairing; it is not relabeled as clean.
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
- Packaged Linux FIFO regression returns exit 4 promptly within its two-second
  guard, proving the repinned canonical reader addresses REV-SPS-003-001 R1.
- Invalid SDP_OFFLINE spellings such as 1/TRUE are rejected, preserving identical
  offline selection between the bootstrap and unchanged child environment.
- `go vet ./...` and `git diff --check` pass.

Signature, wrong-key, download bounds, immutable cache, offline-miss and timeout
behavior is owned by the pinned upstream bootstrap package. This client record
does not assert that each behavior has a corresponding upstream unit test, nor
claim independent local verification beyond the explicitly listed checks.
Power loss, signals/cancellation and native non-Linux execution remain outside
this bounded evidence. Independent review must inspect the actual paired source
candidate and the log before Slice closure.
