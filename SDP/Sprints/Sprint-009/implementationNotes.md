# SPS-009 implementation notes

## P1 — immutable default

Owner Session0010 T011 authorizes publication and global extension upgrade.
The clean upstream release candidate is ef8741c433c7fde67e689f94a74e7afdf938f24f.
GOPROXY=direct resolved canonical bootstrap
v0.0.0-20261010013506-ef8741c433c7, checksum
h1:HiLMErh4PEajKGLK5+CsTYPOexN224k9loHzpKnSZkk=.
No local replacement or checksum bypass is used. The compiled DefaultRelease is
https://github.com/Hans-Einar/SDP/releases/download/v2.2.0/sdp-release.json.
The corresponding exact URL regression and maintained guidance are updated;
main.go and client adapter behavior remain unchanged.

Package, production-default verification, independent review and publication
remain pending. Updating the extension does not upgrade a project's SDP folder.

Initial source race suite, vet and module verification passed. The final upstream
commit only corrects a generated release-log plan reference; bootstrap content
is identical. Final exact-package tests will verify the newly pinned identity.
