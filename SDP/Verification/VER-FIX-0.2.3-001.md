# VER-FIX-0.2.3-001 — Linux publisher bridge

Candidate: clean commit b3258f9d9b592bf6f251a32e69530432f4cc038c. Go1.26.10 Linux amd64.

- `go test -race ./...`: passed, including the new public Linux signer proof.
- `go vet ./...`: passed.
- `go mod verify`: passed; immutable upstream bootstrap
  v0.0.0-20261110231554-643ec77a69a4, no local replacement.
- `git diff --check`: passed.
- Public default remains SDP2.2.0; its original production signature verifies.
- Linux-signed SDP2.2.1 candidate643ec77 boots online/local and offline.
- Installed old gh-sdp0.2.2 rejects the new signer (exit4), as expected.
- Tampered descriptor rejects (exit4); actual production-signed install
  preview/apply succeeds in an isolated empty project.
- Packaged integration through isolated gh routing passes: discovery, human/JSON,
  preview equality, apply/preservation/repetition, incompatible/corrupt-cache refusal.
- Final client built with GOAMD64=v1; SHA256
  `274dfc2e2967a957d85571af39c1a2f89f5d4b95dcdd4fe1a9debca152204f0f`.
  Earlier v3 artifact is superseded. No source change was needed for the rebuild.
- [Independent review](../CodeReview/REV-FIX-0.2.3-001.md) approves bounded preparation.

Artifact directory on Linux: `~/.local/state/sdp-linux-signing/`.
Use `gh-sdp-0.2.3-amd64v1/gh-sdp-linux-amd64`, not the older v3 package.
See evidence/FIX-0.2.3-001 for the retained summaries/package identities.
Source-after-build closeout commits change evidence only.

No tag, release publication, installation or project upgrade is claimed.
