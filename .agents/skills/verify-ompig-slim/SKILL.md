---
name: verify-ompig-slim
description: Verify ompig-slim's Go TUI and headless user flows with isolated PTY smoke runs; use after changes to interactive menus, model roles, startup, or durable execution.
---

# Verify ompig-slim

## Surface

The primary user surface is the native Go terminal UI. The binary also exposes
headless text/JSON/RPC modes and durable execution. This skill uses the existing
PTY smoke harness rather than internal setters or test-only APIs.

## Launch

Build the binary from the repository root:

```sh
make build
```

For the model-management journey, the helper launches an isolated PTY process
with `PIG_TEST_FAUX=1`, `PI_OFFLINE=1`, and a temporary `PIG_HOME`/`PI_HOME`.
No provider login or network access is required. The harness waits for the
observable `interactive-ready` startup marker before driving `/model` and the
compact picker. It terminates the exact child process it started and removes
its temporary profile.

## Doctor

Run this read-only check before driving anything:

```sh
./bin/ompig-slim --version && ./bin/ompig-slim --help >/dev/null
```

If either command fails, rebuild with `make build`. Do not drive a binary from
another checkout. The smoke helper also requires Python 3 and the repository's
`python3 scripts/smoke-test-model-menu.py` harness.

## Drive

Run the model-management journey and save its proof:

```sh
.agents/skills/verify-ompig-slim/helpers/model-management.sh
```

The journey exercises the real user path: startup, `/model`, the Roles view,
cycle-order editing, returning to the main screen, Alt+P session picker role
selection, reopening Roles, and confirmed custom-role deletion. It asserts the
persisted settings after each user action while the application is running.

For the broader repository smoke suite:

```sh
make smoke
```

That suite covers durable execution, model menus, logout, widgets, todo
arguments, codemode, subagents, web tools, and agent state. Use the focused
helper first for model-management changes; use `make smoke` for cross-cutting
runtime changes.

## Evidence

The helper writes `evidence/model-management.txt` under this skill directory.
The file contains the command, environment-independent PASS summary, and the
observable user actions that were asserted. A valid proof must show both the
input path and resulting state: the existing harness checks settings writes,
role selection, cycle updates, and deletion cleanup, not merely the final
terminal frame. Preserve the evidence file when cleaning up.

## Cleanup

The helper's `finally` path terminates only the PTY process it started and
removes its temporary profile. It does not kill by process name. If a run is
interrupted, inspect for a child of the helper and terminate that exact PID;
never kill all `ompig-slim` processes. Scratch profiles are created below the
system temporary directory and are removed by `TemporaryDirectory`.

## Helpers

`helpers/model-management.sh` is executable and runs the repository's canonical
PTY scenario, capturing its output in the evidence path. Invoke it from the
repository root as shown above.

## Feature map

See `features/README.md` for the maintained user-facing journey index:

- `features/model-management.md` — interactive model roles and picker.
- `features/headless-modes.md` — text, JSON, and RPC invocation surfaces.
- `features/durable-execution.md` — resumable unattended execution.

Update the map when routes, commands, prompts, or observable side effects
change. Reach for `/maintain-verification-skill` when the application changes.
