# Requirements — Shared engine distribution

Current authority: owner Activity0004 execution, PLAN-SDP-0029 C2 and SPS-011.
This supersedes the runtime-client interpretation of REQ-GHS-001–005; earlier
SPS-003/004 and release evidence remain historical.

- REQ-GHS-001: Verify the canonical signed descriptor, exact asset digest/size,
  protocol and source identity at packaging time. Test trust remains explicit,
  visibly nonproduction, and never qualifies production publication.
- REQ-GHS-002: The extension asset is the exact canonical engine bytes. Direct
  invocation and gh routing expose identical version/revision/capabilities and
  command behavior with an empty legacy engine cache.
- REQ-GHS-003: SDPTool exclusively owns setup, discovery, framework selection,
  installation planning/apply/recovery and preservation. No runtime launcher,
  engine selector/download or duplicate command implementation remains here.
- REQ-GHS-004: Actual isolated gh binary-release install/upgrade preserves the
  registered executable pathname and updates direct consumers before the next
  gh sdp invocation. Direct local operations need no gh, auth or network. Project
  bytes remain unchanged during tool updates; signed framework preview/apply
  retains canonical preservation semantics.
- REQ-GHS-005: Distribution2.3.0 shares engine2.3.0 identity. Package provenance
  separately identifies canonical engine repository/commit and extension
  distribution repository/commit. Support claims remain Linux amd64 until native
  qualification elsewhere. Independent exact-candidate review/public gates remain.
