# VER-SPS-009 — paired SDP 2.2 default

Status: passed; published, globally upgraded and independently reconciled.
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

## Actual publication and global upgrade

Annotated v0.2.2 targets 9776f8a6b26247383b25e12317e69f06ddd1f4e2, tag object
4ca1ed80b97b888a10bac0f55e34ce5311926a2d. GitHub published at
2026-10-10T01:44:05Z: https://github.com/Hans-Einar/gh-sdp/releases/tag/v0.2.2.
Downloaded release assets match the approved package and checksum. The explicitly
authorized `gh extension upgrade sdp` completed from 0.2.1 to 0.2.2; extension list
confirms v0.2.2, and actual global `gh sdp --version --json` with all four development
overrides removed returns SDP2.2.0 at ef8741c. No project SDP migration was run.
Final native smoke and independent publication reconciliation are pending.

## Actual released-default native Run

The globally installed client discovered widget-lab and invoked its revision-bound
run arguments. A real Fyne window opened from the declared make run; native X11
Run input displayed SDUI → SDL Run → GoRun → SDUI: widget lab. Normal window close
returned exit0. The lab's installed process receipt stayed byte-identical. Native
result and screenshot live in evidence/SPS-009 and the lab RSP2 verification at
local commit b7e15d5. This does not claim every widget was retested or XFMD UI Run.

REV-SPS-009-004 independently approves actual publication reconciliation and the
bounded native/global-default observations. All SPS-009 work is complete.
