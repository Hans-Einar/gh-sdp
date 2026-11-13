#!/usr/bin/env bash
# Verify and republish the exact canonical engine; do not compile another engine.
set -euo pipefail
if [[ $# -ne 3 || ! $1 =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo 'usage: scripts/package.sh VERSION SIGNED-DESCRIPTOR NEW-OUTPUT-DIRECTORY' >&2
  exit 2
fi
repo=$(git -C "$(dirname "$0")/.." rev-parse --show-toplevel)
cd "$repo"
if [[ -n $(git status --porcelain --untracked-files=all) ]]; then
  echo 'package: source checkout must be clean (including untracked files)' >&2
  exit 2
fi
output=$(python3 - "$3" "$repo" <<'PY'
import pathlib, sys
output, repo = map(lambda p: pathlib.Path(p).resolve(), sys.argv[1:])
if output == repo or repo in output.parents:
    raise SystemExit('package: output must be outside the source checkout')
if output.exists():
    raise SystemExit('package: output directory already exists')
print(output)
PY
)
commit=$(git rev-parse HEAD)
go_tool=${SDP_GO:-go}
staging=$(mktemp -d)
trap 'rm -rf "$staging"' EXIT
# Compile only the build-time verifier. Disable ambient workspaces/flags so an
# unrelated checkout or -modfile cannot replace the pinned trust implementation.
GOENV=off GOWORK=off GOFLAGS= GOTOOLCHAIN=local "$go_tool" build -trimpath -buildvcs=true -o "$staging/packager" ./cmd/package
"$go_tool" version -m "$staging/packager" > "$staging/build-info.txt"
python3 - "$staging/build-info.txt" "$commit" <<'PYBUILD'
import pathlib,sys
info=pathlib.Path(sys.argv[1]).read_text()
if f'vcs.revision={sys.argv[2]}' not in info or 'vcs.modified=false' not in info or '=>' in info:
    raise SystemExit('package: verifier must identify clean source without replacements')
PYBUILD
"$staging/packager" "$1" "$2" "$output" "$commit"
cp "$staging/build-info.txt" "$output/gh-sdp.packager-build.txt"
if [[ -n $(git status --porcelain --untracked-files=all) || $(git rev-parse HEAD) != "$commit" ]]; then
  rm -rf -- "$output"
  echo 'package: source changed during packaging' >&2
  exit 2
fi
cat "$output/checksums.txt"
