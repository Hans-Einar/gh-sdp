# Release Notes

## [Unreleased]

No additional changes selected.

## [0.1.1] — Selected, unreleased

### Fixed

- Select the exact signed SDP v0.2.1 descriptor by default through the immutable
  canonical bootstrap dependency. This delivers the upstream external KanBan
  seed correction once the paired distribution is published.
- Assert the exact default descriptor in both bootstrap configuration and the
  delegated child's environment; explicit descriptor overrides remain unchanged.

### Scope

Linux amd64 only. Client arguments, environment, streams, working directory and
exit behavior remain unchanged. The client version is independent from SDP
v0.2.1; publication, global installation and XFMD adoption require separate
recorded evidence.

## [0.1.0] — 2026-09-27

Release-Date: 2026-09-27

### Added

- First native Linux amd64 GitHub CLI extension, delegating installation,
  upgrade, preview and recovery to the canonical Go SDPTool.
- Compiled publisher public trust and an exact default SDP v0.2.0 signed
  descriptor, shared with the delegated child. Explicit SDP_RELEASE overrides
  remain available; test trust requires an explicit test key.
- Clean source packaging with a source/dependency manifest and SHA-256 checksum.

### Scope

The client version is independent from SDP v0.2.0. Windows and macOS binaries
are not published or supported by this release. SDPTool owns policy and
recovery; the extension does not require PowerShell. Publication and live XFMD
adoption are recorded separately when they actually occur.
