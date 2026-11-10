# VER-FIX-0.2.3-001 — Linux publisher bridge

Candidate: clean commit containing the FIX-0.2.3-001 dependency update (recorded
after commit in the package manifest). Go1.26.10 Linux amd64.

- `go test -race ./...`: passed, including the new public Linux signer proof.
- `go vet ./...`: passed.
- `go mod verify`: passed; immutable upstream bootstrap
  v0.0.0-20261110231554-643ec77a69a4, no local replacement.
- `git diff --check`: passed.
- Public default remains SDP2.2.0. New client package, real old/new signer
  checks, corruption checks and fresh independent review remain pending.

No tag, release publication, installation or project upgrade is claimed.
