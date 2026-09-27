# VER-SPS-005 — Default engine patch release gate

Status: candidate checks passed; independent review and publication pending
Slice: SPS-005
Release: REL-0.1.1
Platform: Linux amd64
Product candidate: 50c2fbe5eff4c17a6116c171797a248d49d9e05e
Bootstrap: v0.0.0-20260927094008-c1901337510e (remote, no replace)

## Candidate evidence

Worker implementation notes record Go 1.27.1 race tests, vet, module verification
and exact v0.2.1 default URL regression results. Both unset and empty SDP_RELEASE
select the literal expected URL; the child receives the same selection and
explicit overrides are preserved. Master inspected the diff and actual package.

Master built clean exact product commit with:

```sh
SDP_GO=/home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go scripts/package.sh 0.1.1 /tmp/gh-sdp-sps005-candidate
GH_SDP_PACKAGE_DIR=/tmp/gh-sdp-sps005-candidate GOTOOLCHAIN=local /home/warloc/.local/share/sdp-toolchains/go1.27.1/go/bin/go test -race -count=1 -run TestPackagedAssetDiscovery -v ./...
```

Package and asset-discovery test passed. Binary SHA-256:
`76bf4f15f61a3defa8787c7cb55a955d669ec67a778499f1cfdba9cee7589b70`.
[Candidate manifest](evidence/SPS-005-candidate-manifest.json) records the clean
source and exact dependency. Independent review must confirm preparation.

The final merged source needs a fresh clean package. Default production trust
with a fresh cache is required after the paired SDP v0.2.1 release is public.
No publication, global extension installation or XFMD adoption is claimed here.
