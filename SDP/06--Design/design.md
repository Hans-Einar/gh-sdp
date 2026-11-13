# Design — Verified byte republishing

DES-GHS-002 implements ARC-GHS-002 and REQ-GHS-001–005 for SPS-011.
`scripts/package.sh VERSION DESCRIPTOR NEW-OUTPUT` requires clean distribution
source and external fresh output. `cmd/package` uses a disposable private cache
for canonical build-time verification. It rejects wrong platform, invalid version
or commit identity, signature/asset/protocol failure, engine version/revision
mismatch and existing output. Failed packaging leaves no partial new directory.

The executable is copied unchanged, preserving the canonical embedded helper and
source build identity. Outputs include SHA256/checksum, package/2 manifest,
unchanged descriptor, detached signature and verification envelope. Engine source
and distribution source remain distinct. Test trust is visible as test-signed.
The temporary cache is removed and is not part of the installed runtime.

No runtime main is built from this repository. gh invokes the canonical engine
in place. Canonical setup performs explicit registration and conflict checks;
version/discovery never install or select another engine. Signed framework
selection still serves project preview/apply, with policy and errors owned by
SDPTool. Distribution replacement changes no project files.
