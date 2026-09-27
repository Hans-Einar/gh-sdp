# Design analysis — Thin delegation

The owner-selected GIP design places the portable engine inside SDPTool and
leaves gh-sdp as bootstrap/delegation. This avoids the competing installation
implementations considered by STU-001. The selected implementation reuses the
canonical bootstrap package directly with an exact dependency identity; it does
not copy release, trust, compatibility, cache or installation policy.

Native OS evidence remains necessary: Go cross-compilation alone cannot prove
process, filesystem, permissions, locking or recovery behavior on another OS.
