# Implementation and release work

Current: [SPS-004](../Sprints/Sprint-004/ScrumIterations.md) prepares the owner-authorized
first native client release and records real publication after independent review.
Root coordination owns paired SDP publication, global extension installation and
live XFMD adoption. Linux amd64 is the bounded support target.

## Delivered GIP-3-M2 foundation


Status: complete for GIP-3-M2; owner authorized PLAN-SDP-0003 on 2026-09-27.

One bounded Sprint-003 / SPI-003 / SPS-003 implements the thin Go client and
its canonical bootstrap dependency. The [Slice contract](../Sprints/Sprint-003/ScrumIterations.md)
defines files, invariants, tests, independent review and stop boundary.
The historical Study phase and completed Slices remain complete.

Branch: sdp/install-gip-3 from refreshed gh-sdp main be990cc.
Milestone commits identify GIP-3-M2 and the concrete delivery. Push for review;
no merge, tag, release or live installation is authorized by this work.

REV-SPS-003-002 approved the exact client/engine pair; the coordinating Master
closed SPS-003 and cleared active work pointers. GIP-4 remains upstream work.
