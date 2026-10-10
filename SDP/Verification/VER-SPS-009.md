# VER-SPS-009 — paired SDP 2.2 default

Status: publication gate passed; publication and local upgrade pending.
Source: 9776f8a6b26247383b25e12317e69f06ddd1f4e2; Linux amd64; client 0.2.2.

The clean package manifest and native logs are under evidence/SPS-009. Final
client SHA-256 is f863d847185a7fecee01c724f6e3e54750ae09ffb3d30d0319efe6b079b33ee4.
The full final package race suite passes without skipped cases, including real gh
routing under isolated configuration, default/override, streams/exit, installation,
repetition, preservation, cache corruption and nonregular-input checks. Source vet
and module verification pass; main adapter code is unchanged.

Initial REV-SPS-009-001 required provenance checksum and README authority fixes.
Those were corrected in 9776f8a6b26247383b25e12317e69f06ddd1f4e2; the original report remains intact.
REV-SPS-009-002 independently approves the exact final clean candidate and package,
reproduces its binary, checks hashes/provenance and verifies the public production
default with a new empty cache and no SDP/GH_SDP overrides. It returns SDP2.2.0 at
ef8741c433c7fde67e689f94a74e7afdf938f24f and the original five lab sources plus
runnable widget-lab. Lab file hashes are unchanged. Production signature/hash
matches the published upstream release; no test key or descriptor override is used.
An immediate coordinator probe after publication had transient HTTP404; both later
coordinator and independent fresh-cache probes passed without a trust bypass.

GitHub has no client workflows; the native gate applies. Upstream exact-candidate
CI passed run38013745278. No client publication, global upgrade, native GUI Run or
XFMD UI adoption is inferred from this gate. RSP2 coordinates those next steps.
