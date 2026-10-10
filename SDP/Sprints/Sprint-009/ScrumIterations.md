# Sprint-009 — runnable-program default release

Status: active
Iteration: SPI-009
Slice: SPS-009
Authorization: owner Session0010 T011 (2026-10-10) requests gh-sdp publication and local gh extension upgrade.
BranchPolicy: sdp/release-0.2.2 from clean 1bc3327; publish exact reviewed branch, no main merge needed.
CommitPolicy: preparation, exact-candidate verification/review, publication reconciliation.
Release: client 0.2.2 paired with signed SDP 2.2.0; Linux amd64.

## Contract and scope

Update the immutable canonical bootstrap to the exact upstream release candidate
whose compiled DefaultRelease selects public SDP 2.2.0. Do not add a local replace.
The client remains a thin adapter under REQ-GHS-001–005, ARC-GHS-001 and DES-GHS-001.
Default update is a client patch; forwarded additive application discovery/run
belongs to upstream SDPTool. Preserve argv, environment, stream, cwd, exit,
human/JSON, explicit descriptor override and offline saved-operation behavior.
No new client notifier, platform or XFMD code change. Managed Framework untouched.

Owner authorization covers publication after native release gates, followed by
gh extension upgrade sdp and actual default discovery/run. The root coordinator
owns upstream SDP publication. A project SDP-folder upgrade is not selected.

## Milestones

P1: remotely resolvable immutable bootstrap, regression/guidance/version updates;
race, vet and module provenance pass.
P2: clean exact package, full packaged tests and actual gh routing, fresh independent
review; public paired descriptor/default production verification before publication.
P3: annotated tag/release, downloaded asset hashes, global extension upgrade,
production default and lab discovery/native Run; truthful reconciliation and review.

Current result: preparation active; no package, review or publication success claimed.
