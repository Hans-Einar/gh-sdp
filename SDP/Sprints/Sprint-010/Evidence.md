# SPS-010 exact release evidence

Client03e45710db3ad845d05b26b19c286bd069911294 selects SDP2.2.1 and retains
both publisher anchors. Portable Linux amd64/v1 binary SHA256:
5631dc2e849da5598606394319451be94f6023e0eb7e9b4edc6c82e1647e39c5.

Race tests, vet, module verification and actual packaged-client tests pass.
Linux evidence: /home/warloc/.local/state/sdp-rollout-2.2.1.
Use final-client-packaged-tests.log and final-client-local-signed-version.json;
production engine assets are producer-final/, client assets client/.

Independent review REV-SPS-010 by review_release_rollout found no issues.
The producer review requires exact6f33 CI success. Before client publication,
verify the actual public signed default through this exact binary. Publishing
and installing the client does not itself upgrade project SDP directories.

Owner authorizes publication and installation on Linux, Beinerud and labtop.
Backups exist on each host. Publication and installed identities remain pending.
