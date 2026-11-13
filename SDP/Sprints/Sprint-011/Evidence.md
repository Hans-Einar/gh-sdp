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

## Interrupted update discovery

Actual gh2.97.0 with explicit asset size/digest metadata preserves the previous
registered engine on HTTP503. A deliberately truncated HTTP200 body was accepted,
reported upgraded and installed partial bytes. Direct invocation failed closed;
no older runtime was selected. Both `extension upgrade sdp --force` and
`extension install URL --force` then reported already up to date. In isolation,
removing the extension, expiring/removing its HTTP cache and reinstalling recovered
the exact new executable at the same registered pathname. The executable test
`TestActualGHInterruptedUpdateRecovery` retains this behavior and recovery check.
This is a known upstream gh limitation, not a claim of atomic rollback.

[Official gh2.97.0 download implementation](https://github.com/cli/cli/blob/v2.97.0/pkg/cmd/extension/http.go#L65-L104)
ignores release-asset digest/size metadata and opens the destination for truncation
before copying. It checks io.Copy errors; observed truncated success may involve
its HTTP cache transport. [Upgrade selection](https://github.com/cli/cli/blob/v2.97.0/pkg/cmd/extension/manager.go#L481-L510)
checks update availability before the binary path and does not make force reinstall
a current tag. No gh source changes or runtime fallback were introduced here.

Clean source bbc0c98 was packaged through the actual script against an explicit
signed shell fixture at /tmp/gh-sdp-c2-fixture-package-bbc0c98; package/2 descriptor,
source, canonical asset byte/hash and signature verification passed. This remains
test-signed harness qualification, not a canonical or production artifact.
The build script also inspects its compiled verifier's clean VCS identity and
rejects replacements, disabling ambient Go workspaces/flags before execution so
its pinned build-time trust dependency cannot be replaced by unrelated workspace
configuration. Only the verifier is compiled here; the engine is copied.
