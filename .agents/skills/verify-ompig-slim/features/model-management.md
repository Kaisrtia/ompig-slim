# Model management

## Sub-features

- Fullscreen `/model` manager and Roles view.
- Persistent cycle-order edits.
- Session-only Alt+P role selection.
- Confirmed custom-role deletion and reference cleanup.

## How to get to it (user POV)

Start `ompig-slim` interactively, type `/model`, open Roles, edit the cycle,
return to the main screen, then use Alt+P and type a role such as `@review`.

## Driving it with the repository PTY harness

Run `.agents/skills/verify-ompig-slim/helpers/model-management.sh`. The helper
uses `scripts/smoke-test-model-menu.py` with the faux provider and drives the
same keyboard actions through a PTY. It waits for visible readiness and menu
markers rather than sleeping for a fixed startup delay.

## Gotchas

The journey needs a built `bin/ompig-slim`. It uses a temporary profile and
offline mode; it must never write the operator's real `~/.pi` settings.
