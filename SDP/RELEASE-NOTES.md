# Release Notes

## [Unreleased]

### Changed

- Select the exact signed SDP v2.0.0 default through the immutable canonical
  bootstrap, including upstream automatic source discovery.
- The paired SDPTool now defaults to human-readable output. Machine consumers
  must pass --json explicitly, including when requesting --version. This visible
  compatibility change selects client v0.2.0 independently of Toolkit v2.0.0.
- Shared bootstrap version probing requests machine output with compatibility
  fallback for older explicitly selected engines.

### Scope

Linux amd64 only. The client continues to forward arguments, streams, working
directory, environment and child exit status. Explicit descriptors and saved
operation offline behavior remain supported. Publication awaits exact-candidate
verification, independent review and the public paired descriptor; no merge,
global extension installation or live XFMD change is selected.

## [0.1.2] — 2026-09-28

Release-Date: 2026-09-28

### Fixed

- Select the exact signed SDP v1.0.0 descriptor by default through the immutable
  canonical bootstrap dependency. This distributes the upstream Go-only release
  and current template update through the existing thin client.
- Assert the exact default descriptor in both bootstrap configuration and the
  delegated child's environment; explicit descriptor overrides remain unchanged.

### Scope

Linux amd64 only. Client arguments, environment, streams, working directory and
exit behavior remain unchanged. The client version is independent from SDP
v1.0.0, whose removal of public PowerShell entrypoints does not change the Go
client API. Publication and global extension installation are verified in
VER-SPS-006; the upstream RP3 record owns live XFMD process-upgrade evidence.

## [0.1.1] — 2026-09-27

Release-Date: 2026-09-27

### Fixed

- Select the exact signed SDP v0.2.1 descriptor by default through the immutable
  canonical bootstrap dependency. This delivers the upstream external KanBan
  seed correction from the published paired distribution.
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
