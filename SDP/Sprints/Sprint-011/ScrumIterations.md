# SPS-011 — Shared canonical engine distribution

SPI-011 / Sprint-011 executes bounded PLAN-SDP-0029 C2, ACTIVITY-SDP-0004,
canonical SESSION-SDP-0021 and KB-SDP-063. Owner selects execution and one shared
update path. Branch: sdp/shared-tool-installation. Commit policy: milestone.
Target: extension2.3.0 equals canonical engine2.3.0. State: implementation.

Scope: replace runtime bootstrap with build-time verified exact-byte republishing;
retain production publisher trust, separate source provenance, signed framework
preview/apply and project preservation; actual isolated gh release install/upgrade
and offline direct-engine checks; current architecture/design/docs reconciliation.

Acceptance: signed fixture packaging/negative verification, actual gh update
before first gh sdp call with stable setup alias and byte identity, unchanged
projects, direct version/discovery with absent gh/auth and no server, preserved
canonical installer semantics; final exact signed production artifact checks and
fresh independent review are coordinated by root. C2 owns this repository only.
No main merge, public release or live project upgrade is performed here.

Historical SPS-010 release0.2.3 remains published. Installed local profile retains
native Sprint/Slice records; no shared-profile migration or skill installation
is implied by linking the canonical Activity/Plan.
