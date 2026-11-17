# SPS-013 preparation evidence

Baseline: clean extension commit 54f830a; branch sdp/release-2.3.2. Source changes
are release records only; runtime, packaging and pinned trust dependency are
unchanged. Existing SPS-011/SPS-012 distribution behavior evidence remains
applicable. Production package and public consumer verification passed; final combined
independent review approved the actual publication and consumer evidence.

Focused verification uses Go 1.27.1 linux/amd64 with GOTOOLCHAIN=local,
GOCACHE=/home/warloc/.cache/sdp-workstation/gocache and
GOMODCACHE=/home/warloc/.cache/sdp-workstation/gomodcache:

```sh
go test -race -count=1 -v ./cmd/package
go vet ./...
```

Log: /tmp/sdp-r232-extension-tests.log. Both commands passed. The three packaging
tests include exact engine republishing, rejection of untrusted/tampered/wrong
identity inputs, and Linux publisher trust. These fixture tests are not proof of
production package provenance. YAML records and all NDJSON lines parse; no
previously released notes/records or historical ledger byte was changed.


## Production and public verification

Clean distribution commit c1d86bd148e6ad24da55b03482000df68e3fa3d6 republishes
signed canonical f7a6c9f4793b55de5d0017ae74210135f2a4d365 unchanged. Actual package
asset-discovery/provenance test passed on the production package. Public canonical
and extension assets match signed descriptor 34e2fa36…13efac8 and engine digest
7eab2b30…4b5f5bf. Full identities and observed tag/publication receipts are in
`evidence/package.json` and `evidence/publication.json`. Canonical CI44044114546
passed that exact canonical source before publication.

`evidence/consumers.json` and its executing script/command transcript prove normal
unpinned public 2.3.1-to-2.3.2 `gh extension upgrade sdp` and clean installation.
The prior setup symlink reaches the new engine immediately after upgrade, before
setup runs again. gh/direct identities agree with canonical source and digest;
network-namespace-isolated direct discover/check/preview succeeds with no auth or
gh on PATH. Fixture project bytes and the actual live 2.3.1 extension remain
unchanged. No owner project Framework migration was performed.

`evidence/framework.json`, receipt and executing script prove public-default fresh
Framework preview/apply in a disposable project with no release override or test
trust, signed 2.3.2 receipt, owner-file preservation and repeated upgrade no-op.
Framework command summaries retain stdout hashes; the full transcript remains at
`/tmp/sdp-r232-public-framework/commands.json` (embedded payload omitted here).
Linux amd64 only; no new native interaction or platform claim. Earlier unchanged
behavior evidence is reused. Final independent combined review approved: `evidence/IndependentReview.md`.
