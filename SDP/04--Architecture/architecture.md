# Architecture — Canonical engine distribution

ARC-GHS-001 is superseded for SPS-011 by ARC-GHS-002, selected through owner
Activity0004 / canonical PLAN-SDP-0029 C2. Historical release clients used runtime
bootstrap/delegation. The current extension distribution ships the exact signed
canonical SDPTool executable renamed `gh-sdp-linux-amd64`.

SDPTool owns all runtime command behavior, setup registration, capabilities,
framework descriptor verification and project preservation. gh-sdp owns only
build-time verification, republishing and distribution provenance. No executable
launcher here selects/downloads/starts a second runtime. Shared signature code
is reused through the immutable bootstrap module only in `cmd/package`.

The build-time verifier checks canonical signature, platform asset bounds/hash,
protocol and engine version/revision against signed source identity. It copies
verified bytes unchanged and retains the signed descriptor and verification
facts. Distribution source identity differs from engine source identity; the
manifest records both. Version2.3.0 is shared across executable entrypoints.

Setup registers the actual executing engine through canonical behavior. gh owns
replacement of its binary extension file. The stable registered link therefore
observes an update immediately without a bootstrap call. Native consumers use
that public path offline. Framework release selection cannot change the
executing engine. Tool upgrades do not mutate installed project framework data.
Canonical managed helper packaging remains canonical SDP responsibility; exact
republishing preserves it without another extension-owned helper/update path.

Linux amd64 remains the native boundary. Real isolated gh release install/upgrade
checks complement descriptor/package tests and preserved project workflow checks.
