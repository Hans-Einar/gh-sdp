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
prerelease at preparation. SPS-007 selects v0.2.0 paired with SDP v2.0.0 because
the delegated default becomes human-readable; machine consumers must request
--json. Earlier release evidence remains historical. The selected candidate
requires its own package checks, human/JSON delegation verification and fresh review.

Follow the Slice's selected publication source: SPS-007 authorizes the exact
reviewed commit on sdp/release-0.2.0, without a main merge. Where a later Slice
explicitly selects a merge, use its reviewed merged commit. Keep publication
identities null until real objects exist. Once the paired production descriptor
is public and the package passes default verification with a fresh cache and
production trust, annotate the selected client version tag on the reviewed clean
source commit and publish its exact asset/checksum/manifest. Verify tag target
and GitHub release remotely, append actual release events, and reconcile
manifests in a small follow-up commit. Do not overwrite a published tag or asset.

SPS-007 does not select global extension installation or live XFMD changes.
The owner will manually upgrade XFMD; repository publication evidence does not
claim that adoption has occurred.
