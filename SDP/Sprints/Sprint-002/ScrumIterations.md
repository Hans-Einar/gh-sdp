# Sprint-002 — Phase 2 Evidence-Backed Study

Status: complete
Iteration: SPI-002
Completed Slice: SPS-002
Study: STU-001
Mandate: MAN-001
Release context: REL-0.1.0 (`gh-sdp` 0.1.0, proposed and unreleased)
Installed Toolkit baseline: 0.2.0 source commit
`bc110bb5fd60009ba67015cf640ad6ddbfe1b04b` (unreleased)

## Sprint Goal

Produce a rigorous, evidence-backed Study that resolves the factual and
analytical groundwork needed before Requirements and Architecture can be
written, while making no later-lifecycle or implementation decision.

## SPI-002 — Study Research, Synthesis, and Review

Status: complete

### SPS-002 — Authoritative Phase 2 Study

Status: complete

#### Goal

Research and synthesize `STU-001` under governing Mandate `MAN-001`, validate
the Study and its traceability, obtain fresh independent review of the exact
candidate, remediate every blocking/high/medium finding, and prepare a reviewed
draft PR for Steering Group assessment.

#### Why Now

PR #1 merged at exact baseline
`3a3b9ece5db6bde438dfd2b4eba57be344350e85`, and the Steering Group separately
authorized Study work. Requirements and Architecture need verified facts about
GitHub CLI extensions, Toolkit installation, distribution, authentication,
platforms, security, recovery, provenance, and conformance before they can be
defined responsibly.

#### Scope

- GitHub CLI extension naming, invocation, local development, installation,
  upgrade, removal, precompiled distribution, release assets, and enterprise
  considerations.
- Go suitability and Windows, Linux, and macOS target analysis, including an
  explicit recommendation for Linux arm64.
- GitHub authentication, API, rate-limit, offline, cache, and test-seam inputs.
- The actual pinned SDP Toolkit installer, manifest, payload, ownership,
  migration, backup, preview, idempotence, schema, and failure contract.
- Installer-rule ownership alternatives and drift prevention.
- Toolkit release discovery, selection, compatibility, provenance, integrity,
  archive/path security, filesystem behavior, transaction/recovery, and
  conformance-testing inputs.
- Requirements candidates, Architecture decision inputs, cross-repository
  findings, and explicit treatment of MAN-Q-001 through MAN-Q-007.

#### Non-Goals

- Requirements approval or substantive changes to
  `SDP/03--Requirements/requirements.md`.
- Architecture selection or substantive changes to
  `SDP/04--Architecture/architecture.md`.
- Design Analysis, Design, implementation planning, product code, Go modules,
  handlers, installer/downloader/archive/migration engines, product tests or
  fixtures, workflows, builds, packaging, binaries, tags, or releases.
- Modification of `Hans-Einar/SDP`, creation of upstream issues, or silent
  Toolkit upgrade.
- Claims that `gh-sdp` 0.1.0 or Toolkit 0.2.0 is published or supported.

#### Expected Documents

- `SDP/02--Study/study.md` (`STU-001`).
- `SDP/Sprints/Sprint-002/ScrumIterations.md`, `implementationNotes.md`, and
  `Handoff.md`.
- `SDP/Verification/VER-SPS-002.md` after real verification exists.
- One or more `SDP/CodeReview/REV-SPS-002-###.md` records after real independent
  review exists.
- Truthful updates to `SDP/SDP-project.manifest.yaml`,
  `SDP/Traceability/CurrentIndex.yaml`, and `SDP/Traceability/Relations.yaml`.
- `SDP/Traceability/Ledger.ndjson` only if an event defined by its installed
  release-event schema actually occurs; Study activation is not such an event.

#### Evidence Standard

- Prefer official documentation, official source, immutable repository paths,
  tagged source, schemas, tests, manifests, release records, and actual command
  output.
- Every material verified fact records a source URL or repository path,
  revision/version when applicable, inspection date, what the evidence proves,
  and its limitations.
- Conclusions are explicitly classified as verified external fact, verified
  repository fact, observation, inference, recommendation, assumption, or
  unresolved question.
- Mutable search snippets and research-agent summaries are discovery aids, not
  final proof. The Master inspects cited primary evidence.
- Installed Toolkit baseline, upstream `origin/main`, published tags/releases,
  and proposed/uncommitted state remain distinct.

#### Invariants

- Exactly one Git repository exists at the `gh-sdp` root.
- The branch begins at the required Phase 1 merge baseline and contains only
  Study/governance evidence.
- Later-lifecycle templates remain byte-identical to their Phase 2 baseline.
- `gh-sdp` client, installed Toolkit, selected Toolkit, latest stable Toolkit,
  GitHub CLI, and Sprint/Slice identities remain separate.
- Project-owned content, credentials, paths, remote archives, and interruption
  are treated as preservation and trust concerns.
- No Study recommendation is represented as an approved Requirement or
  Architecture decision.

#### Traceability

- Governing Mandate: MAN-001.
- Study: STU-001.
- Sprint / Iteration / Slice: Sprint-002 / SPI-002 / SPS-002.
- Link all relevant MAN-OUT, MAN-BND, MAN-SC, MAN-ASM, and MAN-Q identifiers.
- Add stable Study evidence, finding, recommendation, requirements-candidate,
  Architecture-input, and cross-repository-finding IDs according to
  `SDP/Instructions/StableIdentifiers.md`.
- Do not mark future Mandate success criteria delivered.

#### Verification

- Parse all YAML and every NDJSON line; apply canonical schemas where they
  exist without claiming schemas for formats the Toolkit does not define.
- Validate CurrentIndex/Relations agreement, stable-ID/path resolution, and
  coverage of MAN-Q-001 through MAN-Q-007.
- Check source/citation completeness and classification of material claims.
- Confirm later-lifecycle templates are unchanged and no production, workflow,
  build, packaging, executable, release, nested-repository, or backup artifact
  exists.
- Enumerate the complete hidden-aware tree, confirm exactly one root Git
  repository, verify external URLs where practical, and run `git diff --check`.
- Bind final evidence to the exact candidate commit.

#### Independent Review

A fresh Reviewer that did not author the Study must inspect the actual
`origin/main...HEAD` diff, authority, evidence/source quality, classifications,
all required Study areas, every MAN-Q question, alternatives, security,
cross-platform coverage, traceability, scope discipline, and absence of later
lifecycle or implementation work. Every blocking, high, and medium finding must
be resolved and the corrected exact head reviewed again.

#### Completion Signal

- `STU-001` addresses all required Study sections with stable IDs and inspectable
  primary evidence.
- Verification passes against the exact pushed candidate.
- Independent review of the corrected exact head has zero unresolved blocking,
  high, or medium findings.
- A draft PR against `main` identifies Phase 2 Study only and remains unmerged
  for Steering Group assessment.

#### Stop Boundary

Stop after reporting the reviewed draft Study PR. Do not populate Requirements,
Architecture, Design Analysis, Design, or implementation-plan templates, and do
not begin implementation, packaging, workflow, or release work.
