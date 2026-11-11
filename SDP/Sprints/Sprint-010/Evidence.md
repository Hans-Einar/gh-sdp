# SPS-010 exact release evidence

Client03e45710db3ad845d05b26b19c286bd069911294 selects SDP2.2.1 and retains
both publisher anchors. Portable Linux amd64/v1 binary SHA256:
5631dc2e849da5598606394319451be94f6023e0eb7e9b4edc6c82e1647e39c5.

Race tests, vet, module verification and actual packaged-client tests pass.
Linux evidence: /home/warloc/.local/state/sdp-rollout-2.2.1.
Use final-client-packaged-tests.log and final-client-local-signed-version.json;
production engine assets are producer-final/, client assets client/.

Independent review REV-SPS-010 by review_release_rollout found no issues.
The producer exact6f33 CI passed. The actual public signed default was
verified through this exact binary before client publication. Publishing
and installing the client does not itself upgrade project SDP directories.

Owner authorizes publication and installation on Linux, Beinerud and labtop.
Backups exist on each host. Publication and installed identities are recorded below.

## Observed completion

Producer exact CI passed. Fresh-cache public-default test passed before client
publication; downloaded client assets match the reviewed package byte-for-byte.
All three installations passed hash and public engine identity checks. See
Publication.json and Host-Installations.json. VER-SPS-010 passed; independent
REV-SPS-010 release gate has no findings.

Independent postpublication reconciliation by review_release_rollout approved
the exact annotated tags, public release/download evidence, three host identities,
three signed receipts/no-op plans and truthful four-consumer blocker scope.

Reconciliation validation: producer Toolkit and management validators pass.
The modern Toolkit project validator is not a passing gate for this client's
unmigrated legacy process: historical study namespaces, release-note categories
and manifest extensions predate its schema. No broad migration is selected.
New Sprint010 identities and reverse links are registered; release record,
manifest, exact source and publication evidence are checked directly.
