# Design analysis — Shared executable

Owner Activity0004 selects one tool update owner: GitHub CLI replaces its actual
extension executable, which is also the directly registered standalone engine.
A second bootstrap would retain the competing runtime lookup/update state the
migration removes. Exact signed-byte republishing preserves canonical module and
helper build composition without duplicating canonical mains or imports here.

Build-time reuse of the immutable canonical bootstrap verifier retains publisher
trust, path/bounds checks, digest and protocol validation. Additional packaging
checks bind signed release/source to the actual executable version/revision.
The extension distribution manifest records its own source separately.

Native update qualification uses actual gh release installation rather than a
local extension symlink, because only the former exercises executable replacement.
An isolated TLS API fixture permits reproducible release transitions; public
signed-release verification remains a separate root-coordinated gate.
