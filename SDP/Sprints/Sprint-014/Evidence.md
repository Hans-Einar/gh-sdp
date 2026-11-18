# SPS-014 release preparation evidence

Baseline: clean c0b6f0f, new sdp/release-2.4.0 branch. Release records only; runtime,
packager and pinned trust dependency are unchanged. Earlier distribution evidence
remains applicable. Production package, public consumer qualification and final independent
combined release review passed.

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


## Production/public qualification

Clean distribution f7763acf27900c0fb51fd28d3f2b6edc8a40a5f7 republishes signed
canonical 52888aaf96c657c0d6656b104c9594e06f99d46f unchanged. Actual package
asset-discovery test passed, verifying production signature/envelope and exact
canonical bytes. Public downloaded assets match recorded API hashes. See
`evidence/package.json` and `evidence/publication.json` for full engine, descriptor,
source and tag identities. Exact canonical CI44287504558 passed before publication.

Normal unpinned public 2.3.2-to-2.4.0 extension upgrade and clean installation pass
in isolated homes. Prior setup registration reaches the new engine immediately
before setup runs again. gh/direct identities and bytes match; disconnected direct
version/discovery/check/preview succeeds without gh or credentials. Fixture project
bytes and live owner extension are unchanged. Retained scripts/results/transcript
are in evidence/; native DocumentHost qualification remains canonical-root-owned.

Public-default fresh Framework preview/apply obtains signed 2.4.0 without release
or test-trust overrides, preserves an owner file and repeats upgrade as no-op in
a disposable project. Framework transcript summaries retain stdout hashes; full
payload output remains `/tmp/sdp-r240-public-framework/commands.json`.

Final independent combined review: `evidence/IndependentReview.md`, approved for
R240/SPS-014 closeout. Linux amd64 only; no main merge, owner extension replacement
or real project Framework migration. Earlier unchanged distribution evidence is
reused. Published notes and historical record bytes remain unchanged.
