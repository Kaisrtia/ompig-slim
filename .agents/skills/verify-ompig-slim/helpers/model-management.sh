#!/usr/bin/env bash
# Run the canonical isolated PTY journey and preserve its proof after cleanup.
set -euo pipefail

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../.." && pwd)
evidence_dir="$repo_root/.agents/skills/verify-ompig-slim/evidence"
evidence_file="$evidence_dir/model-management.txt"
mkdir -p "$evidence_dir"
{
  printf 'command: python3 scripts/smoke-test-model-menu.py ./bin/ompig-slim\n'
  printf 'journey: startup -> /model -> Roles -> cycle edit -> Alt+P role picker -> custom-role deletion\n'
  cd "$repo_root"
  python3 scripts/smoke-test-model-menu.py ./bin/ompig-slim
} 2>&1 | tee "$evidence_file"
