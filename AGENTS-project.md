# gh-sdp Project Instructions

`gh-sdp` is the official GitHub CLI extension project for the future `gh sdp`
command. `Hans-Einar/SDP` is the canonical, independent SDP Toolkit repository;
do not modify that repository without a separate, explicit authorization.

- Follow this repository's project-local `SDP/` records for all work.
- Keep the `gh-sdp` client version distinct from every SDP Toolkit version and
  describe release/support state only from repository evidence.
- Support claims with repository evidence and applicable tests, then require a
  fresh independent review of the actual change.
- In later lifecycle phases, address Windows, Linux, and macOS behavior together
  with preservation, provenance, path-safety, and other security concerns.
- `SPS-001` authorizes governance, Mandate, traceability, and verification
  records only. It authorizes no production implementation, build, packaging,
  or release workflow.
- At installed source commit `bc110bb5fd60009ba67015cf640ad6ddbfe1b04b`,
  `-InitializeProjectStructure` is a one-time bootstrap operation: always preview
  it and do not apply a repeat that proposes the Toolkit-specific
  `SDP/Releases/REL-0.2.0.yaml` in this project.

The managed `AGENTS.md` remains authoritative for role boundaries, Slice/Fix
discipline, review, verification, release gates, and stop conditions.

## Current implementation assignment

Owner authorization on 2026-09-27 selects SDP PLAN-SDP-0003, including GIP-3-M2
as local Sprint-003 / SPI-003 / SPS-003. Its project-owned contract is
`SDP/Sprints/Sprint-003/ScrumIterations.md`. gh-sdp is only a thin bootstrap and
delegation client; SDPTool owns all installation policy. Earlier SPS-001 and
SPS-002 boundaries remain historical and do not prohibit this separately
authorized Slice. Managed AGENTS.md and installed Toolkit facts stay unchanged.
Fresh independent review remains required; no merge or release is selected.

## Authorized release assignment — 2026-09-27

The owner selected merge, release, extension installation and live XFMD upgrade.
SPS-004 in SDP/Sprints/Sprint-004/ScrumIterations.md governs this repository's
release preparation and publication; it supersedes SPS-003's historical no-release
boundary only for this new work. Root coordination owns upstream SDP and live XFMD.

## Authorized default-engine patch — 2026-09-27

The owner explicitly authorized merge and release of the external KanBan fix.
SPS-005 in SDP/Sprints/Sprint-005/ScrumIterations.md selects gh-sdp v0.1.1 with
canonical SDP v0.2.1 as its default engine. This supersedes SPS-004's completed
boundary only for the bounded patch. Root owns upstream publication, global
extension installation and XFMD adoption. Fresh independent review is required.

## Authorized template-default patch — 2026-09-28

The owner authorized commit, integration, release, extension installation and
XFMD upgrade. SPS-006 in SDP/Sprints/Sprint-006/ScrumIterations.md selects client
v0.1.2 paired with owner-selected SDP v1.0.0, which removes public PowerShell
entrypoints. It supersedes SPS-005's completed boundary for this bounded patch.
Root owns upstream publication and XFMD adoption; the client coordinator owns
delegated global extension installation and default verification. Root confirmed the revised Go-only candidate and immutable pin after exact-head
CI and upstream integration. REV-SPS-006-002 approves P2; P3 verifies the public
production descriptor and publishes the exact merged client. Independent review
of final publication reconciliation remains required.

## Authorized discovery release — 2026-10-01

Session 0003 / upstream PLAN-SDP-0014 selects local SPS-007 in
SDP/Sprints/Sprint-007/ScrumIterations.md. Prepare and release client v0.2.0
paired with SDP v2.0.0 through the root-supplied immutable bootstrap. The human
default output change requires explicit --json for machine consumers. Publication
from the reviewed branch is authorized after the paired descriptor is public;
main merge, global extension installation and live XFMD changes are not selected.
Root owns upstream release; the owner will manually upgrade XFMD.

## Authorized SDP 2.1 default patch — 2026-10-01

Owner explicitly authorized publication of prepared SDPTool 2.1.0 and paired
gh-sdp. SPS-008 in SDP/Sprints/Sprint-008/ScrumIterations.md selects client v0.2.1
using immutable upstream bootstrap 93517ad98cd188c0debeb1d0f3d36d123c6e4a3b.
Publish the exact reviewed branch after the upstream descriptor is public.
No main merge, global extension upgrade/install or XFMD change is authorized.
Keep the owner's existing extension intact for their manual gh notification test.
