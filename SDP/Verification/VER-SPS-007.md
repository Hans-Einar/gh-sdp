# VER-SPS-007 — Discovery client release gate

Status: pending
Slice: SPS-007
Release: REL-0.2.0
Platform: Linux amd64 only
Product candidate: 336d654c7533a30b340a1366109a41a9307ddf1e
Bootstrap: v0.0.0-20260930230248-247fb7aba628 (root supplied, remotely resolved)

## Required evidence

- Go 1.27.1 race tests, vet and module verification, with GOMAXPROCS=2 and test -p 2.
- Remote immutable module resolution with no local replacement.
- Clean exact-commit package manifest, checksum, asset discovery and FIFO rejection.
- Paired SDPTool/direct and client execution: human default, explicit JSON,
  no-argument source discovery, existing preview/apply and preservation checks.
- Isolated GitHub CLI routing with temporary configuration, without changing the
  user's global extension.
- Fresh independent source/package review.
- Public paired SDP v2.0.0 descriptor and fresh-cache production-trust human/JSON
  execution before client publication.
- Actual annotated tag, release assets and downloaded checksums followed by
  independent reconciliation review.

No checks above are passed by this placeholder. No merge, global installation,
Windows/macOS native support or live XFMD outcome is claimed.

## Initial record validation

The project manifest and new release record validate against the current Toolkit
JSON schemas, and the new version-selection ledger event validates against the
release-event schema. All edited YAML and all NDJSON lines parse successfully.
The current-index schema rejects the pre-existing historicalStudy,
historicalDependencyAssessment, lifecycle and lastCompleted extensions, also
present in the clean baseline. As documented in VER-SPS-006, process migration
is not part of this client release; no complete current-Toolkit schema
compatibility is claimed.
