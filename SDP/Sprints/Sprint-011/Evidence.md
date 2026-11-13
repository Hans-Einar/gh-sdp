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

At the first source milestone, canonical artifact qualification was pending.
The current signed qualification below supersedes that temporary gap. Fresh
independent review and the remaining cross-repository gates are still pending. Public publication/native RPM checks are not
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


## Superseded signed canonical qualification — source63b7f83

Canonical2.3.0 source63b7f8303f3688d135272913bb437895f11080f2 is signed by the
trusted Linux publisher. Engine SHA256:
`8c7c21fc616b57a4dc4b20eba0fa9cd03418ca31335a00b5d4da02f6d88d3c82`.
Descriptor SHA256:
`5086d386e801b63f7aa1df7d1d45ed10ca3da1b1940d50b28e5edf75057ea9cf`.
Clean distribution6065596 packaged it unchanged at
/tmp/activity4-gh-package-6065596. Retained package manifest records production
signed provenance and separate canonical/distribution source identities. This is
a signed local release candidate; no public publication is claimed.

The previous-engine fixture was rebuilt from the same canonical frozen source
and bundled helper with BuildVersion2.2.99 and actual source BuildRevision, solely
for update testing. Its SHA256 is
`42c6291f105a313e4f4c2544f58f21b2a0488f7b8d5f8ed4e72721a7f71213a4`.
It is not a historical public release and was not byte-patched. No tracked source
files in the frozen canonical release worktree were modified.

Actual gh binary-release install/upgrade and interrupted-transfer recovery pass
against these canonical bytes. Registered target lives inside isolated HOME,
retains its pathname and exposes exact replacement bytes before a subsequent
gh sdp call. Production capability/version JSON is identical across entrypoints.
After fixture server closure and removal of gh/auth from the direct environment,
version, discovery, SDUI check and SDUI preview pass with a fresh legacy-engine
cache. HTTP503 preservation and truncated-transfer failure/recovery limitations
remain as above. OS-level network namespace isolation belongs to root's combined
qualification; this test does not claim that separate condition.

The signed installer gate passes direct and actual gh routes: exact engine byte
identity, human/JSON and diagnostic equality, signed preview/apply preservation,
repeat no-change, and no executing-engine selection through framework descriptors.
Commands: GH_SDP_PACKAGE_DIR, GH_SDP_BINARY, SDPTOOL_BINARY and
SDPTOOL_OLD_BINARY set to the paths above; `go test -race -run
'TestPackaged|TestActualGH' -v .`; repeat TestPackagedCandidates with
GH_SDP_VIA_GH=true. Logs are in SDP/Verification/evidence/SPS-011.

## Original signed predecessor matrix

The adapted verification script executes the current canonical engine against
original production signed descriptors0.2.0,0.2.1,1.0.0,2.0.0,2.1.0,2.2.0 and
2.2.1. Every disposable case installs its predecessor, adds owner project/Session/
SDL bytes, refuses a managed-file conflict, upgrades to signed2.3.0 five-phase,
preserves the owner bytes and produces no actions on repeat upgrade. All seven
pass. The original2.2.1 source is the producer-final production descriptor,
not a newly rebuilt predecessor. Per-case signed receipt, predecessor/target
hashes and summary are retained under SDP/Verification/evidence/SPS-011;
full disposable plans/outputs remain /tmp/activity4-c2-predecessor-matrix.
No live project was changed.


## Superseded signed canonical qualification — sourcedab8a03

Source dab8a03c8a4a36d303ca64cbdc0fe79af94843dd supersedes63b7f83 after the
canonical SDUI CLI output-argument fix. Version remains2.3.0. Its production
signed engine SHA256 is
`115b233f203c27b3764e2b9a9d3b19229542d8c1b8b939b13bacd650e1fb409e`;
signed descriptor SHA256 is
`ee68feaebda7be1eb249648d1fd64986ba570d0bacfcfc06cc3d8c626fc0c3a5`.
The previous candidate's package, logs and receipts remain superseded evidence.

Clean distribution4cd3192 repackaged the replacement unchanged at
/tmp/activity4-gh-package-dab8a03-prequalification. Actual gh install/upgrade,
stable registration and before-next-gh-call exact byte identity pass with the
unchanged2.2.99 test fixture. Direct offline version/discovery and SDUI check/
preview pass with absent gh/auth and fresh cache. The test now explicitly puts
--json immediately after the SDUI operation and checks the preview result and
rendered file. Truncated-download failure/recovery and the documented gh-owned
limitation remain verified. Signed installer equivalence/preservation passes both
direct and actual gh routes. Race tests and vet pass.

All seven original signed predecessor upgrades were repeated against dab8a03,
including original production2.2.1: every preservation, managed-conflict refusal,
signed receipt and repeated no-change assertion passes. New per-case receipt,
source/target hashes, adapted script and logs are retained separately under
SDP/Verification/evidence/SPS-011/dab8a03. Full disposable plans/outputs remain
/tmp/activity4-c2-predecessor-matrix-dab8a03. The final reviewed distribution
source is repackaged after the evidence commit and its exact package/gate paths
are handed to root. No public release, live project mutation or main merge is
claimed; root coordinates the independent/integrated publication gates.


## Current signed canonical qualification — source774e150

Source774e150774d184d0604ae78f274586eeb2bad495 supersedes dab8a03 after completing
the typed canonical SDUI preview service boundary and predecessor-profile test/CI
updates. Runtime wire contract, embedded native payload and framework payload
remain unchanged. Version stays2.3.0. Production signed engine SHA256:
`2cfcf12f8bfb0d6fa8d11391db874f68d9bc9210c85974ff3cecee50f915ced6`;
production signed descriptor SHA256:
`c7ef11e749eb2186ca6fee32e8dba1d96d160095f9e0221e1b70d4bd1b3b50b4`.
Earlier receipts and logs remain superseded candidate evidence.

Clean distribution53a25f8 repackaged the exact signed source at
/tmp/activity4-gh-package-774e150-prequalification. The same focused canonical
gates pass: actual gh binary-release installation/upgrade, stable setup alias,
exact new bytes before first subsequent gh sdp call, direct offline version/
discovery/SDUI check/preview without gh/auth or legacy-engine cache, actual
preview artifact, interrupted-transfer failure/recovery, and production package
signature/source/hash checks. Signed framework preview/apply equivalence and
project preservation pass direct and actual gh routes. No new runtime fallback
or release-publication claim is introduced.

All seven original production signed predecessors were rerun against the new
descriptor:0.2.0,0.2.1,1.0.0,2.0.0,2.1.0,2.2.0,2.2.1. Every case passes original
predecessor installation, target signed upgrade, owner-byte preservation,
managed-conflict refusal and repeat no-change. New receipts, descriptor hashes,
script and focused gate logs are retained under
SDP/Verification/evidence/SPS-011/774e150; full disposable plans remain
/tmp/activity4-c2-predecessor-matrix-774e150. Final clean distribution source and
repackaged exact-source gate log are handed to root after this evidence commit.
Root retains independent/integrated gates, public publication and main-merge
ownership. The previously recorded gh interrupted-download limitation remains.

## Published exact candidate — supersedes local candidate gaps

Public extension source 1753a11ffcbda3d6fd563456ee70db3c77e8d6c4; annotated tag v2.3.0 object
ab40dc7b510e0cd729eb2262e691a3dae06d5724; published2026-11-13T09:32:55Z.
Canonical engine source 774e150774d184d0604ae78f274586eeb2bad495; SHA256 `2cfcf12f8bfb0d6fa8d11391db874f68d9bc9210c85974ff3cecee50f915ced6`.
[Public assets](public-evidence/extension-publication.json) and
[normal upgrade/fresh install](public-evidence/consumer-results.json) retain exact
published identities. Actual unpinned0.2.3 normal upgrade and clean install/setup
pass with a stable registered target and exact public executable bytes, followed
by direct version/discovery/SDUI checks without gh/auth and with OS namespace
network isolation. Projects and live extension remain unchanged.
[Signed public-default Framework checks](public-evidence/framework-results.json)
pass preview/apply, owner-file preservation and repeat no-op without release
or trust overrides. [Independent public gate review](public-evidence/independent-review.md)
approves this canonical/extension public scope. Historical signed/local candidates
and the gh truncated-download limitation above remain unchanged.
