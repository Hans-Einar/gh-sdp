# Native release gate

A release candidate must have completed preparation milestones in its bounded
release Slice,
passing Go race tests and vet, a clean exact-commit Linux amd64 package with SHA-256,
and a fresh independent review of the same production source. The upstream bootstrap
must be a remotely resolvable immutable module version, with no local replace.
The default signed SDP descriptor must be public before the gh-sdp release is published.

The first public client API is v0.1.0 (previous currentVersion 0.0.0 is an unreleased
placeholder). Platform scope is Linux amd64. Package naming is gh-sdp-linux-amd64,
which GitHub CLI discovers as a binary extension asset. The package script records
source commit and refuses tracked or untracked dirt. A checksum accompanies the binary.

Freeze selected notes into the selected client version and set release state to
prerelease at preparation. SPS-006 selects v0.1.2 paired with Go-only SDP v1.0.0.
The earlier v0.2.2 candidate's package and review evidence remain historical;
the selected v1.0.0 candidate requires its own package checks and fresh review.
Keep publication identities null until real objects exist. After owner-authorized PR
merge, tag the clean merged source commit with an annotated tag for the selected client
version and publish its exact asset/checksum/manifest. Verify tag target and GitHub
release remotely, append actual release events, and reconcile manifests in a small follow-up commit. Do not overwrite
a published tag or asset. The root coordinator installs the extension from GitHub and
checks actual default delegation before live XFMD adoption.
