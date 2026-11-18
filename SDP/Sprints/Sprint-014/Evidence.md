# SPS-014 release preparation evidence

Baseline: clean c0b6f0f, new sdp/release-2.4.0 branch. Release records only; runtime,
packager and pinned trust dependency are unchanged. Earlier distribution evidence
remains applicable. Production package and final release review are pending.

Focused checks use Go 1.27.1 linux/amd64, GOTOOLCHAIN=local,
GOCACHE=/home/warloc/.cache/sdp-workstation/gocache and
GOMODCACHE=/home/warloc/.cache/sdp-workstation/gomodcache:

```sh
go test -race -count=1 -v ./cmd/package
go vet ./...
```

Log: /tmp/sdp-r240-extension-tests.log. All three packaging tests and vet passed. Tests cover exact signed-engine
republishing, rejection of untrusted/tampered/wrong-identity inputs, and Linux
publisher trust. Fixture tests alone do not
establish production package provenance. Public descriptor selected after canonical
publication: https://github.com/Hans-Einar/SDP/releases/download/v2.4.0/sdp-release.json.
