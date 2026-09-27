#!/usr/bin/env bash
# Build a clean, exact-source Linux amd64 GitHub CLI extension asset.
set -euo pipefail
if [[ $# -ne 2 || ! $1 =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "usage: scripts/package.sh VERSION OUTPUT-DIRECTORY" >&2
  exit 2
fi
version=$1
repo=$(git -C "$(dirname "$0")/.." rev-parse --show-toplevel)
cd "$repo"
if [[ -n $(git status --porcelain --untracked-files=all) ]]; then
  echo "package: source checkout must be clean (including untracked files)" >&2
  exit 2
fi
output=$(python3 - "$2" "$repo" <<'PY'
import pathlib, sys
output, repo = map(lambda p: pathlib.Path(p).resolve(), sys.argv[1:])
if output == repo or repo in output.parents:
    raise SystemExit('package: output must be outside the source checkout')
print(output)
PY
)
go_tool=${SDP_GO:-go}
commit=$(git rev-parse HEAD)
mkdir -p "$output"
for name in gh-sdp_linux_amd64 checksums.txt gh-sdp.manifest.json; do
  if [[ -e "$output/$name" || -L "$output/$name" ]]; then
    echo "package: output already exists: $output/$name" >&2
    exit 2
  fi
done
staging=$(mktemp -d "$output/.package.XXXXXXXX")
trap 'rm -rf "$staging"' EXIT
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOTOOLCHAIN=local "$go_tool" build -trimpath -buildvcs=true -o "$staging/gh-sdp_linux_amd64" .
"$go_tool" version -m "$staging/gh-sdp_linux_amd64" > "$staging/build-info.txt"
python3 - "$staging" "$version" "$commit" <<'PY'
import hashlib, json, pathlib, sys
root, version, commit = pathlib.Path(sys.argv[1]), sys.argv[2], sys.argv[3]
info = (root / 'build-info.txt').read_text()
if f'vcs.revision={commit}' not in info or 'vcs.modified=false' not in info:
    raise SystemExit('package: binary does not identify the clean source commit')
if '=>' in info:
    raise SystemExit('package: module replacements are not permitted')
name = 'gh-sdp_linux_amd64'
data = (root / name).read_bytes()
digest = hashlib.sha256(data).hexdigest()
manifest = {'schemaVersion': 'gh-sdp-package/1', 'version': version,
            'sourceCommit': commit, 'platform': 'linux/amd64', 'file': name,
            'sha256': digest, 'size': len(data),
            'buildInfo': info.splitlines()[1:]}
(root / 'gh-sdp.manifest.json').write_text(json.dumps(manifest, indent=2) + '\n')
(root / 'checksums.txt').write_text(f'{digest}  {name}\n')
PY
# Recheck source after compilation as well; publication requires the recorded tree.
if [[ -n $(git status --porcelain --untracked-files=all) || $(git rev-parse HEAD) != "$commit" ]]; then
  echo "package: source changed during packaging" >&2
  exit 2
fi
for name in gh-sdp_linux_amd64 checksums.txt gh-sdp.manifest.json; do
  ln "$staging/$name" "$output/$name"
done
cat "$output/checksums.txt"
