# Services Foundation TODO

## Immediate Follow-Ups

- Define Redis usage boundaries to avoid turning cache into system-of-record.
- Persist in-flight `current` release operations across service-manager restarts instead of only persisting completed history.
- Add explicit operation IDs to HTTP responses so RM/QA can trace one release action across create, promote, switch, and rollback.
- Add release-operation cleanup and retention policy for the `service_manager_operations` table.

## Next Technical Slice

- Add account-service password hashing and session issuance.
- Add provider mock adapters and callback contracts for social-login providers.
- Add service-manager orchestration endpoints for rollback plans and promote-with-validation flows.
- Add structured error codes and failure reasons for release orchestration APIs.
- Add operation-state queries filtered by `product_slug`, `target_type`, and time range.

## Account-Service Follow-Ups

- Add password hashing and session-token strategy.
- Add provider registry interface for Google, Apple, WeChat, and other common social providers.
- Add provider mock adapters for local development.
- Add future web account portal requirements as a PM-led product flow.

## Upgrade-Service Follow-Ups

- Add explicit build-number governance separate from display-version strings.
- Add release-channel awareness and staged rollout strategy.
- Add query filters and pagination for releases, deployments, switch events, and rollback events.
- Add release validation rules before switch or promote, such as required deployment status and environment checks.

## Service-Manager Follow-Ups

- Add profile switching persistence and target visibility snapshots.
- Add durable recovery semantics for interrupted operations after restart.
- Add lock ownership metadata, timeout handling, and forced-unlock policy.
- Add Compose and native-mode parity checks in automated tests.
- Add RM-facing release workflow endpoints and audit exports.

## Future Product Work

- Define PM draft for shared web account portal.
- Run UD on the portal after PM draft is approved.
- Produce PM final, then AM / SD / QA for the portal product.

## Environment Blockers

- `Go` is installed locally at `$HOME/develop/go`.
- Docker Desktop is installed locally at `$HOME/Applications/Docker.app`.
- Non-login shells may need `PATH="$HOME/Applications/Docker.app/Contents/Resources/bin:$PATH"` before invoking `docker`.
