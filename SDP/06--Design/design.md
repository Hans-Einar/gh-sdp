# Design — Thin process adapter

DES-GHS-001 implements ARC-GHS-001 and REQ-GHS-001–005.

Read SDP_RELEASE, SDP_TEST_KEY, SDP_CACHE_DIR and SDP_OFFLINE for canonical
bootstrap configuration. When SDP_RELEASE is empty, use the shared bootstrap
DefaultRelease constant for the exact signed SDP v2.1.0 descriptor selected by
SPS-008. Production public trust belongs to that canonical package; explicit SDP_TEST_KEY keeps test
trust separate. Explicitly pass the selected SDP_RELEASE to the child so bootstrap
and installation planning agree. Bootstrap returns the fixed
verified compatible native executable path. Go os/exec receives that path and
all client arguments unchanged, inherits cwd and other environment entries, and connects all three
standard streams. Return the child's exit code. Local configuration
errors are written to stderr and return exit 2; bootstrap or process-start
errors return exit 4. Saved apply/resume operations force verified-cache-only
bootstrap. Ordinary child exit codes are forwarded unchanged.

The client does not consume installation flags. In particular upgrade --manifest
... --plan-output ... and upgrade --apply ... reach the same SDPTool parser as
direct invocation. Signed artifact and saved-plan semantics belong upstream. The paired engine
defaults to human output, including --version; machine consumers must pass
--json explicitly. The client forwards that argument unchanged. Bootstrap
compatibility probing requests --version --json and falls back for older engines;
probe output is separate from delegated command streams.
