# Linux publisher bridge

FIX-0.2.3-001 prepares an unpublished 0.2.3 client candidate. Published 0.2.2
remains the current release until a separate publication is recorded.

The new immutable bootstrap retains the laptop key and adds the Linux key:
`2c947457cffe015f870168cb332242c49606fdea378a13d875579544c2989c03`.
Its public base64 value is `3c4v3py8MxlsZ0zGp0T2Wtmcl2rXQ4f4WEM/YBShBiw=`.
Only SDP's publisher holds private signing material; gh-sdp requires none to build.

The bridge deliberately retains the existing signed SDP2.2.0 default. After this
client is distributed, a Linux-signed release can be explicitly selected with
SDP_RELEASE. A later client default update is separate. Existing 0.2.2 clients
reject the new signer: SSH access, a GitHub public key or publishing the descriptor
does not update their compiled trust. Never use SDP_TEST_KEY for production rollout.

No released descriptor, tag or asset is replaced. This change does not claim
revocation, automatic trust discovery, multiple signatures or completed client adoption.
See the [verification record](../SDP/Verification/VER-FIX-0.2.3-001.md).
