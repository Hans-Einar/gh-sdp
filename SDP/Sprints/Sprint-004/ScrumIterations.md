# Sprint-004 — gh-sdp first native release

Status: active
Iteration: SPI-004
Slice: SPS-004
Authorization: owner requested merge, release, global gh extension installation and live XFMD upgrade on 2026-09-27.
BranchPolicy: sdp/release-0.1.0 from merged main 74fec81; PR merge after independent review.
CommitPolicy: implementation, verification/review correction, and publication reconciliation.
Release: gh-sdp v0.1.0, paired with SDP v0.2.0; Linux amd64 only.

## Contract

Prepare the first native release, pin the upstream bootstrap containing production
public trust and the exact shared default descriptor, and propagate that descriptor
to the child when SDP_RELEASE is unset. Preserve explicit overrides, exact argv,
streams, cwd and exits. Package a gh-compatible linux-amd64 executable and checksum.
No install policy is copied into the wrapper; no private signing key enters this repo.

Expected files: main.go/tests, go.mod/go.sum, README, packaging and release gate,
project-owned release/verification/traceability records. Managed instructions and
installed Toolkit facts remain unchanged. Linux is the only supported release host;
Windows/macOS need separate native verification before support is advertised.

## Milestones and gate

1. R1: verified default selection/delegation, remote module pin and reproducible packaging.
2. R2: race tests, vet, clean exact-candidate build and fresh independent review.
3. R3: merge approved release preparation, annotated tag, publish asset/checksum,
   verify remote identities and reconcile release records only after publication.

The project-local gate in docs/Release-Lifecycle.md supplies the previously missing
path referenced by the installed release skill. Root coordinator owns SDP release,
production keys, global extension installation and XFMD upgrade. Publish only after
coordination confirms the paired SDP descriptor is publicly available.
