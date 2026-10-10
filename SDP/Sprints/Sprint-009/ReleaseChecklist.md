# gh-sdp 0.2.2 release checklist

Copy this checklist into the selected release's work area. Record version,
previous release, exact commit, owner publication authority, evidence paths and
actual outcomes. Unchecked means pending; use N/A only with a reason. A generated
release log is not proof of publication. The project release contract governs.

## Scope and content

- [x] Select product version from compatibility impact. Keep language/profile,
  process/management schema and client versions distinct.
- [x] Enumerate included work and remaining exclusions; validate selected model,
  tests, review dispositions and links on the actual candidate.
- [x] Update canonical RELEASE-NOTES.md Unreleased entries with stable work IDs.
  Freeze selected entries into a dated version section; preserve released bytes.
- [x] Run `sdptool release-log --version X.Y.Z --output Releases/X.Y.Z.md`.
  Check with the same command plus --check. Use --all --output Releases --check
  to check every recorded version. Generator extracts reviewed notes, not Git guesses.

## Installation distribution

N/A to this thin client: payload inventory, signed installer descriptor, original
predecessor migrations and extracted Toolkit archive are owned by SDP2.2.0 under
MAINT-SDP-0015. Its independently approved exact source is ef8741c, published before
this client gate. Native client package/default/isolated-gh tests apply here.

## Publish and consumer readiness

- [x] Exact candidate CI and required independent review pass; working tree clean;
  tag absent; merge/publication explicitly authorized.
- [x] Publish annotated tag and GitHub Release with generated log, signed descriptor,
  executable and checksums; verify downloaded assets against exact release.
- [x] Reconcile actual tag/commit/time and append truthful release events only after
  success. Do not rewrite a released log or silently correct released notes.
- [x] Verify gh-sdp selection: its compiled default is pinned. Either publish/update
  the client default or document an explicit SDP_RELEASE selecting this release.
- [x] Provide preview/apply commands and expected receipt. Owner-requested manual
  project upgrade remains manual; never confuse updating the gh extension with
  upgrading the project's SDP directory.

Consuming projects reuse the scope/content/publication checks; mark SDP publisher
inventory/signing steps N/A if they do not distribute SDP itself. Keep their own
application/package/installer acceptance checks alongside this list.

Evidence: VER-SPS-009, REV-SPS-009-002 and evidence/SPS-009. The client retains its
installed older native lifecycle. Per-version generated Toolkit logs are upstream;
client release notes are the exact frozen section from SDP/RELEASE-NOTES.md.
No live project upgrade was requested or applied. Actual client publication,
extension upgrade and one native Run route are complete. Independent publication
reconciliation remains pending and is separate from source/package approval.

Final independent publication reconciliation approved in REV-SPS-009-004.
