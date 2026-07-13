# VER-SPS-001 — SPS-001 Documentation Verification

Status: Master validation passed; independent Reviewer confirmation pending
Slice: SPS-001
Mandate: MAN-001
Release context: REL-0.1.0 (gh-sdp 0.1.0, proposed and unreleased)
Validation time: 2026-07-13T09:43:35Z
Candidate basis: branch codex/phase-1-sdp-bootstrap-mandate, initial commit
2dcc18423a10497678eb19ace8fd9668efda9551, plus the current working tree

This is Slice evidence, not a product-release gate. It claims no final commit,
pull request, review approval, tag, GitHub Release, or published capability.

## Environment

- OS: Microsoft Windows 10.0.26200
- PowerShell: 5.1.26100.8655
- Git: 2.43.0.windows.1
- Python: 3.11.5
- GitHub CLI: 2.96.0
- Target repository: C:\Users\hanse\GIT\gh-sdp
- Canonical SDP clone: C:\Users\hanse\GIT\SDP

## Authoritative Source Evidence

Commands:

~~~powershell
git -C C:\Users\hanse\GIT\SDP fetch --prune origin
git -C C:\Users\hanse\GIT\SDP pull --ff-only origin main
git -C C:\Users\hanse\GIT\SDP rev-parse HEAD
git -C C:\Users\hanse\GIT\SDP status --short --branch
git -C C:\Users\hanse\GIT\SDP tag --list
gh release list --repo Hans-Einar/SDP --limit 20
gh run view 29207275200 --repo Hans-Einar/SDP --json databaseId,headSha,headBranch,event,status,conclusion,url,workflowName,jobs
~~~

Results:

- The clone fast-forwarded from b47c191 to clean origin/main at
  bc110bb5fd60009ba67015cf640ad6ddbfe1b04b.
- SDP.manifest.yaml and canonical documentation identify Toolkit 0.2.0 as
  unreleased, with null tag and release commit.
- git tag --list and gh release list returned no tags or GitHub Releases.
- GitHub Actions run 29207275200 at the exact source commit passed both
  contracts and installer jobs.

Canonical validation commands, run from C:\Users\hanse\GIT\SDP:

~~~powershell
python -m pip install -r Toolkit\tests\requirements.txt
python Toolkit\scripts\validate_sdp.py
python -m unittest discover -s Toolkit\tests -p "test_*.py" -v
.\Toolkit\tests\Install-SDP.Tests.ps1
~~~

Results: canonical validation passed; all 9 Python tests passed; installer
fixture tests passed. Generated Python cache directories were removed afterward,
and the canonical clone was clean. These commands validate the Toolkit source,
not this consuming project.

## Bootstrap Evidence

Commands:

~~~powershell
& 'C:\Users\hanse\GIT\SDP\Toolkit\scripts\Install-SDP.ps1' -ProjectRoot 'C:\Users\hanse\GIT\gh-sdp' -InitializeProjectStructure -Preview
& 'C:\Users\hanse\GIT\SDP\Toolkit\scripts\Install-SDP.ps1' -ProjectRoot 'C:\Users\hanse\GIT\gh-sdp' -InitializeProjectStructure
~~~

Results:

- Preview proposed 46 actions, applied none, and left the worktree unchanged.
- Apply reported Toolkit 0.2.0, 41 applied files, one preserved file, three
  unchanged files, and no release or tag.
- The installed manifest records source commit bc110bb5fd60009ba67015cf640ad6ddbfe1b04b.
- The initializer's Toolkit-specific REL-0.2.0 seed and copied Toolkit
  traceability facts were replaced with truthful gh-sdp REL-0.1.0 records.

## Consuming-Project Record Validation

The following exact command was run from the target repository. It uses the
canonical schemas and checks cross-record identity, publication truth, stable
IDs, paths, and later-lifecycle state.

~~~powershell
@'
from pathlib import Path
import json
import yaml
from jsonschema import Draft202012Validator, FormatChecker

root = Path(r'C:\Users\hanse\GIT\gh-sdp')
upstream = Path(r'C:\Users\hanse\GIT\SDP')

def load_yaml(path):
    return yaml.safe_load(path.read_text(encoding='utf-8'))

def validate_yaml(label, instance_path, schema_name):
    instance = load_yaml(instance_path)
    schema = json.loads((upstream/'Toolkit/schemas'/schema_name).read_text(encoding='utf-8'))
    errors = sorted(Draft202012Validator(schema, format_checker=FormatChecker()).iter_errors(instance), key=lambda e: list(e.path))
    assert not errors, f'{label}: {[e.message for e in errors]}'
    print(f'PASS schema: {label}')
    return instance

project = validate_yaml('project manifest', root/'SDP/SDP-project.manifest.yaml', 'SDP-project-manifest.schema.json')
installed = validate_yaml('installed Toolkit manifest', root/'SDP/Framework/installed-toolkit.manifest.yaml', 'installed-toolkit-manifest.schema.json')
release = validate_yaml('REL-0.1.0 release record', root/'SDP/Releases/REL-0.1.0.yaml', 'release-record.schema.json')
current = load_yaml(root/'SDP/Traceability/CurrentIndex.yaml')
relations = load_yaml(root/'SDP/Traceability/Relations.yaml')
assert isinstance(current, dict) and isinstance(relations, dict)
print('PASS YAML objects: CurrentIndex and Relations')

event_schema = json.loads((upstream/'Toolkit/schemas/release-event.schema.json').read_text(encoding='utf-8'))
ledger_lines = [line for line in (root/'SDP/Traceability/Ledger.ndjson').read_text(encoding='utf-8').splitlines() if line.strip()]
assert ledger_lines, 'Ledger is empty'
for number, line in enumerate(ledger_lines, 1):
    event = json.loads(line)
    errors = list(Draft202012Validator(event_schema, format_checker=FormatChecker()).iter_errors(event))
    assert not errors, f'ledger {number}: {[e.message for e in errors]}'
print(f'PASS release-event schema: {len(ledger_lines)} ledger line(s)')

assert project['project']['name'] == current['project']['name'] == 'gh-sdp'
assert project['release']['currentVersion'] == '0.0.0'
assert project['release']['nextTargetVersion'] == current['release']['targetVersion'] == release['version'] == '0.1.0'
assert project['release']['state'] == current['release']['state'] == release['state'] == 'unreleased'
assert current['release']['activeReleaseId'] == release['releaseId'] == 'REL-0.1.0'
assert project['development']['sprintId'] == current['active']['sprint'] == 'Sprint-001'
assert project['development']['iterationId'] == current['active']['iteration'] == 'SPI-001'
assert project['development']['sliceId'] == current['active']['slice'] == 'SPS-001'
assert installed['toolkitVersion'] == '0.2.0'
assert installed['sourceCommit'] == 'bc110bb5fd60009ba67015cf640ad6ddbfe1b04b'
assert project['release']['latestTag'] is None and project['release']['latestCommit'] is None
assert release['releasePreparationCommit'] is None and release['gitTag'] is None and release['githubReleaseUrl'] is None
assert relations['requirements'] == {} and relations['designs'] == {}
assert 'REL-0.1.0' in relations['releases'] and not (root/'SDP/Releases/REL-0.2.0.yaml').exists()
print('PASS cross-record identity, lifecycle, and publication truth')

mandate_text = (root/'SDP/01--Mandate/mandate.md').read_text(encoding='utf-8')
for section in ('outcomes','boundaries','successCriteria'):
    for stable_id, record in relations[section].items():
        assert stable_id in mandate_text, stable_id
        assert record['mandate'] == 'MAN-001'
        assert (root/record['path']).is_file(), record['path']
for category in ('mandates','sprints','iterations','slices','verification'):
    for stable_id, record in relations[category].items():
        if 'path' in record:
            assert (root/record['path']).is_file(), f'{stable_id}: {record["path"]}'
for key in ('releaseNotes','manifest','releaseRecord'):
    assert (root/relations['releases']['REL-0.1.0'][key]).is_file(), key
print('PASS stable-ID and relation path resolution')

future_templates = [
    'SDP/02--Study/study.md','SDP/03--Requirements/requirements.md',
    'SDP/04--Architecture/architecture.md','SDP/05--DesignAnalysis/design-analysis.md',
    'SDP/06--Design/design.md','SDP/07--Implementation/implementation-plan.md',
]
for rel in future_templates:
    text = (root/rel).read_text(encoding='utf-8')
    assert 'Status: template' in text, rel
print('PASS later lifecycle remains template-only')
print('PHASE 1 RECORD VALIDATION PASSED')
'@ | python -
~~~

Result: all three YAML records passed their canonical schemas; CurrentIndex and
Relations parsed; the one ledger line passed the release-event schema; identity,
lifecycle, publication, relation, ID, path, and template checks passed.

## Managed Files, Layout, and Scope

The following exact command compared SHA-256 hashes with the pinned source and
checked the skill set, untouched future templates, Git layout, and no-product
boundary.

~~~powershell
@'
from pathlib import Path
import hashlib

root = Path(r'C:\Users\hanse\GIT\gh-sdp')
upstream = Path(r'C:\Users\hanse\GIT\SDP')

def digest(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()

pairs = [(upstream/'Toolkit/payload/project-root/AGENTS.md.template', root/'AGENTS.md')]
for source in sorted((upstream/'Toolkit/skills').glob('*/SKILL.md')):
    pairs.append((source, root/'.codex/skills'/source.parent.name/'SKILL.md'))
framework_source = upstream/'Toolkit/payload/sdp-root/Framework'
for source in sorted(path for path in framework_source.rglob('*') if path.is_file()):
    pairs.append((source, root/'SDP/Framework'/source.relative_to(framework_source)))
for source, target in pairs:
    assert target.is_file(), target
    assert digest(source) == digest(target), f'mismatch: {target}'
print(f'PASS managed SHA-256 equality: {len(pairs)} files')

expected_skills = {p.parent.name for p in (upstream/'Toolkit/skills').glob('*/SKILL.md')}
actual_skills = {p.parent.name for p in (root/'.codex/skills').glob('*/SKILL.md')}
assert actual_skills == expected_skills
print(f'PASS managed skill set: {len(actual_skills)} skills')

future_pairs = {
    'SDP/02--Study/study.md': '02--Study/study.md',
    'SDP/03--Requirements/requirements.md': '03--Requirements/requirements.md',
    'SDP/04--Architecture/architecture.md': '04--Architecture/architecture.md',
    'SDP/05--DesignAnalysis/design-analysis.md': '05--DesignAnalysis/design-analysis.md',
    'SDP/06--Design/design.md': '06--Design/design.md',
    'SDP/07--Implementation/implementation-plan.md': '07--Implementation/implementation-plan.md',
}
for target_rel, source_rel in future_pairs.items():
    assert digest(root/target_rel) == digest(upstream/source_rel), target_rel
print('PASS untouched later-lifecycle template hashes: 6 files')
print('MANAGED AND PLACEHOLDER HASH VALIDATION PASSED')
'@ | python -

$gitRoots = @(Get-ChildItem -LiteralPath 'C:\Users\hanse\GIT\gh-sdp' -Force -Recurse -Directory -Filter .git -ErrorAction Stop | ForEach-Object FullName)
if ($gitRoots.Count -ne 1 -or $gitRoots[0] -ne 'C:\Users\hanse\GIT\gh-sdp\.git') { throw "Unexpected Git roots: $($gitRoots -join ', ')" }

$files = @(rg --files -g '!**/.git/**')
$forbidden = @($files | Where-Object { $_ -match '(^|[\\/])(go\.mod|go\.sum|.*\.go|.*\.exe|.*\.dll|.*\.sh|.*\.bat|.*\.cmd|Dockerfile|Makefile)$' -or $_ -match '^\.github[\\/]' })
if ($forbidden.Count -ne 0) { throw "Forbidden product/build files: $($forbidden -join ', ')" }
if (Test-Path -LiteralPath 'SDP\.sdp-backups') { throw 'Unexpected SDP backup directory' }
~~~

Results: all 16 copied managed files matched the pinned source; all 10 skills
matched; all 6 later-lifecycle documents matched their untouched templates;
exactly one .git existed at repository root; all 36 files were governance
artifacts; no product/build/workflow file or backup artifact existed.

## Installer Repeatability

Commands:

~~~powershell
& 'C:\Users\hanse\GIT\SDP\Toolkit\scripts\Install-SDP.ps1' -ProjectRoot 'C:\Users\hanse\GIT\gh-sdp' -Preview
& 'C:\Users\hanse\GIT\SDP\Toolkit\scripts\Install-SDP.ps1' -ProjectRoot 'C:\Users\hanse\GIT\gh-sdp' -InitializeProjectStructure -Preview
~~~

Results:

- Normal same-version reinstall: Proposed 0, Applied 0, five project-owned files
  preserved, and 20 managed or template files unchanged.
- Repeating the one-time initializer: Proposed 1, namely creation of
  SDP/Releases/REL-0.2.0.yaml; zero actions were applied. This is a canonical
  Toolkit payload/idempotency gap, not a gh-sdp release record. It is mitigated
  by project warnings and preview-first one-time use. Correcting it requires a
  separate authorized change in Hans-Einar/SDP.

## Authoritative Tooling Gaps and Limitations

- Toolkit/scripts/validate_sdp.py validates the Toolkit repository layout. It has
  no consuming-project mode and cannot truthfully be presented as this project's
  validator.
- Canonical schemas cover project, installed, and release manifests plus release
  ledger events, but not Mandate, CurrentIndex, Relations, or generic work and
  review events. The checks above are the strongest available without fabricating
  an authoritative validator.
- Toolkit test dependencies use compatible version ranges rather than a fully
  locked environment.
- Evidence currently applies to the working tree. The Master must rerun checks
  against the committed PR head, and a fresh Reviewer must inspect the actual
  diff and evidence before SPS-001 can close.

## Reviewer Confirmation

Pending. No review disposition is claimed here.
