# SPS-013 preparation evidence

Baseline: clean extension commit 54f830a; branch sdp/release-2.3.2. Source changes
are release records only; runtime, packaging and pinned trust dependency are
unchanged. Existing SPS-011/SPS-012 distribution behavior evidence remains
applicable. Production package and final combined independent review are pending.

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
released record or historical ledger byte was changed.
