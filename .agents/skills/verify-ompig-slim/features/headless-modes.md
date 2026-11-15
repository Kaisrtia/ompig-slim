# Headless modes

## Sub-features

- One-shot text output with `-p`.
- JSONL events with `--mode json`.
- RPC request/response loops with `--mode rpc`.

## How to get to it (user POV)

Run `ompig-slim -p "..."` for a single answer, or select JSON/RPC when a
script needs machine-readable events.

## Driving it with the repository smoke harness

Use the repository's `scripts/smoke-test-*` scenarios and `make smoke-built`.
These launch the real binary with isolated profiles and faux/offline providers,
then inspect protocol output and exit status.

## Gotchas

A provider/model fixture is required for deterministic no-network runs. Do not
replace the production mode with direct Go calls; protocol framing and process
lifecycle are part of this feature.
