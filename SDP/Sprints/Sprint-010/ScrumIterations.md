# Sprint-010 — Linux publisher and SDP2.2.1 default

Status: active
Iteration: SPI-010
Slice: SPS-010
Authorization: owner2026-11-11 explicitly requests publication and installation on
Linux, Beinerud and the laptop, followed by Linux receipt-based project upgrades.
BranchPolicy: current sdp/linux-publisher-trust; no main merge selected.
CommitPolicy: preparation, exact-candidate verification, publication reconciliation.

## Contract

Release gh-sdp0.2.3 with immutable canonical bootstrap35bfdf8, retaining both
publishers and selecting signed SDP2.2.1. This completes a backward-compatible
trust/default patch; no wire, adapter, project runtime or notification redesign.
Completed FIX-0.2.3-001 remains historical bridge evidence, not reopened work.
Root owns SDP2.2.1 publication and consumer project rollout. Preserve dirty work.

## Milestones

P1: immutable dependency, release metadata, race/vet/module verification.
P2: clean portable exact-source package, real signed default and packaged tests,
fresh independent final release review.
P3: annotated tag/release, downloaded hash verification, installed identity and
default on all three hosts; evidence-backed reconciliation.

Publish only after the signed SDP2.2.1 descriptor is publicly downloadable and
verifies through this exact client. Do not rewrite existing tags/assets.
