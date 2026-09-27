# Sprint-005 — Default engine patch release

Status: complete
Iteration: SPI-005
Slice: SPS-005
Authorization: owner explicitly authorized merge and release of the external KanBan fix on 2026-09-27; root coordination assigned gh-sdp v0.1.1.
BranchPolicy: sdp/release-0.1.1 from clean main e9316ad; reviewed PR merge to main, publication reconciliation on a follow-up branch.
CommitPolicy: bounded implementation, verification/review evidence, then publication reconciliation.
Release: gh-sdp v0.1.1, paired with SDP v0.2.1; Linux amd64 only.

## Contract

Pin the canonical remote bootstrap revision supplied by root coordination that
selects the signed SDP v0.2.1 descriptor. This makes the default client obtain
the upstream external KanBan seed correction. The client remains a thin adapter;
explicit descriptor overrides, argv, environment, streams, cwd and exit behavior
are preserved. No local module replacement, installed-skill migration, application
feature, global extension installation or XFMD mutation is authorized here.

REQ-GHS-001–005 and DES-GHS-001 govern this patch. PATCH is appropriate because
this corrects the default upstream release selection without changing the client
API or supported platform set. Prior v0.1.0 publication records remain historical.

Expected files: go.mod/go.sum, focused default URL regression in main_test.go,
README and current design/default release guidance, release lifecycle wording,
and project-owned work, release, review, verification and traceability records.
Managed AGENTS.md and installed Toolkit facts remain unchanged.

## Milestones and gate

1. P1: Worker pins the supplied immutable remote bootstrap, adds an exact v0.2.1
   default descriptor regression, updates selected documentation and records tests.
2. P2: Go 1.27.1 race tests, vet, module verification, clean exact-commit package,
   actual package discovery and fresh independent review pass.
3. P3: after root confirms public SDP v0.2.1, test the packaged default using a
   fresh cache and production trust; merge approved preparation, build clean merged
   source, publish annotated v0.1.1 and exact asset/checksum/manifest, inspect actual
   remote objects, reconcile records and stop on clean main.

The installed release skill and docs/Release-Lifecycle.md govern publication.
All publication identities remain null until real objects exist. Root owns global
extension verification and live XFMD adoption. Do not claim either here.

## Preparation disposition

P1 and P2 delivered. Independent REV-SPS-005-001 approves production candidate
50c2fbe5eff4c17a6116c171797a248d49d9e05e and independently reproduces its
package checksum and manifest. Race/vet, exact default URL/delegation, fresh
remote module retrieval, asset discovery and FIFO rejection passed. At preparation, P3 remained pending; actual publication and reconciliation
are recorded below.

## Publication and closeout

P3 delivered on 2026-09-27. PR #8 merged reviewed preparation with an identical
tree into source 8cbef9693e353cb04dc6d98021aca454b9078e3f. Its clean package
passed race/vet, asset discovery and FIFO checks. Both candidate and final package
used fresh caches without SDP_RELEASE, SDP_TEST_KEY or SDP_OFFLINE overrides and
returned public SDP 0.2.1 / revision 4bacfce05f92f0dab9680456e297214727b54bc6.

Annotated v0.1.1 resolves to the clean merged source. The public final GitHub
Release was published at 2026-09-27T10:05:48Z:
https://github.com/Hans-Einar/gh-sdp/releases/tag/v0.1.1.
Downloaded binary, checksum and manifest match the exact local package bytes.
Publication records are reconciled; active work coordinates are cleared.
Root coordination retains global extension installation and live XFMD verification.

Independent REV-SPS-005-002 approves publication reconciliation with no findings.
It verified the remote tag, release, independently downloaded assets and actual
downloaded binary's fresh-cache production default. SPS-005 is complete.
