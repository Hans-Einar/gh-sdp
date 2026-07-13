# gh-sdp Project Stable Identifiers

Status: active project instruction
Scope: Mandate, Study, and operating records

The canonical SDP Toolkit at source commit
`bc110bb5fd60009ba67015cf640ad6ddbfe1b04b` defines no identifier grammar for a
Mandate or its outcomes, boundaries, assumptions, questions, or success
criteria. This project-owned instruction therefore establishes a local Phase 1
convention before those identifiers are used. It is not an upstream SDP format.

| Kind | Format | Example |
| --- | --- | --- |
| Mandate | `MAN-###` | `MAN-001` |
| Mandate outcome | `MAN-OUT-###` | `MAN-OUT-001` |
| Mandate boundary | `MAN-BND-###` | `MAN-BND-001` |
| Mandate success criterion | `MAN-SC-###` | `MAN-SC-001` |
| Mandate assumption | `MAN-ASM-###` | `MAN-ASM-001` |
| Mandate question | `MAN-Q-###` | `MAN-Q-001` |
| Study | `STU-###` | `STU-001` |
| Study evidence source | `STU-EVD-###` | `STU-EVD-001` |
| Study finding | `STU-FND-###` | `STU-FND-001` |
| Study recommendation | `STU-REC-###` | `STU-REC-001` |
| Study requirement candidate | `STU-RQC-###` | `STU-RQC-001` |
| Study architecture-decision input | `STU-ADI-###` | `STU-ADI-001` |
| Study cross-repository finding | `STU-XRF-###` | `STU-XRF-001` |
| Sprint | `Sprint-###` | `Sprint-001` |
| Sprint iteration | `SPI-###` | `SPI-001` |
| Sprint slice | `SPS-###` | `SPS-001` |
| Slice verification | `VER-<slice-id>` | `VER-SPS-001` |
| Slice review | `REV-<slice-id>-###` | `REV-SPS-001-001` |

Numbers are three decimal digits, are unique within their kind, and are never
reassigned to a different meaning. The authoritative definition remains in the
Mandate document. `SDP/Traceability/Relations.yaml` resolves the Mandate,
outcome, boundary, success-criterion, assumption, and question IDs to that
document; other records refer to the exact ID. Status is recorded separately
and is not encoded in an identifier.

Study evidence, finding, recommendation, requirement-candidate,
architecture-decision-input, and cross-repository-finding identifiers were
added before Phase 2 used them. They identify project-local records and do not
make a finding normative, a candidate approved, or a recommendation an
Architecture decision.

The canonical release and Fix formats remain authoritative and are not
redefined here. A later authoritative Toolkit format for the project-local kinds
above may require a documented migration while preserving historical links.
