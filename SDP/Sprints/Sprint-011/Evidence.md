# C2 implementation evidence

Baseline: b177e14. Candidate: baseline plus C2 diff, before first milestone commit.
Linux amd64; gh2.97.0. Native installed Framework remains0.2.0. Local SDP
entrypoint skill is absent; available HEOS SDP entrypoint and project-local Worker
were read without installing/migrating skills.

`go test -race ./...`, `go vet ./...`, `go mod verify` pass. Packaging unit tests
use explicit ephemeral Ed25519 signed fixtures. They cover exact byte copying,
separate engine/distribution commits, visible test provenance, retained descriptor,
existing-output refusal and rejection of untrusted signatures, descriptor/asset
tampering, version mismatch and signed source/engine revision mismatch.

Actual gh binary-release install/ordinary upgrade harness passes with disposable
shell fixtures (not canonical-engine acceptance). TLS GitHub API fixture and
CONNECT proxy exercise the installed gh executable. It installs v2.2.1, registers
setup, expires gh's HTTP metadata cache, exposes v2.3.0, invokes extension upgrade,
then compares alias bytes before any gh sdp invocation. Alias target remains
unchanged. Direct version/discovery succeeds with gh absent from PATH, no auth,
server closed and fresh cache. No user's extension/project is modified.

Canonical signed2.3.0 artifact qualification, canonical setup behavior, signed
project preview/apply gate and fresh independent review are pending root's C1
artifact and cross-repository gates. Public publication/native RPM checks are not
claimed. Network disabled at OS namespace level is a separate integrated gate;
server closure/no gh/no auth here proves no dependency on the fixture server.
