# Durable execution

## Sub-features

- `durable run` with a persistent state directory.
- `durable status` inspection.
- `durable resume` after process interruption.
- Optional committed event and control files.

## How to get to it (user POV)

Run `ompig-slim durable run --state DIR --cwd REPO --model PROVIDER/MODEL
--request-id ID PROMPT`, inspect it with `durable status`, and resume it with
`durable resume --state DIR`.

## Driving it with the repository smoke harness

Run `python3 scripts/smoke-test-durable.py ./bin/ompig-slim` or `make smoke`.
The scenario drives the CLI boundary and checks persisted state across process
lifetimes.

## Gotchas

Use a fresh temporary state directory and a deterministic faux provider. Never
reuse a production durable state directory during verification.
