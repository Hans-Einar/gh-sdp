# Design — Thin process adapter

DES-GHS-001 implements ARC-GHS-001 and REQ-GHS-001–004.

Read SDP_RELEASE, SDP_TEST_KEY, SDP_CACHE_DIR and SDP_OFFLINE for canonical
bootstrap configuration. These are explicit development inputs; there is no
implicit production key or invented latest release. Bootstrap returns the fixed
verified compatible native executable path. Go os/exec receives that path and
all client arguments unchanged, inherits cwd/environment and connects all three
standard streams. Return the child's exit code. Local configuration
errors are written to stderr and return exit 2; bootstrap or process-start
errors return exit 4. Saved apply/resume operations force verified-cache-only
bootstrap. Ordinary child exit codes are forwarded unchanged.

The client does not consume installation flags. In particular upgrade --manifest
... --plan-output ... and upgrade --apply ... reach the same SDPTool parser as
direct invocation. Signed artifact and saved-plan semantics belong upstream.
