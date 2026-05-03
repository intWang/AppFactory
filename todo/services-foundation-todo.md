# Services Foundation TODO

## Immediate Follow-Ups

- Install and verify Docker with Compose support on the local machine.
- Add runbooks for local startup, troubleshooting, and service switching.
- Define concrete PostgreSQL schemas and migrations for account-service.
- Define concrete PostgreSQL schemas and migrations for upgrade-service.
- Define Redis usage boundaries to avoid turning cache into system-of-record.

## Next Technical Slice

- Add shared Go workspace files and module boundaries.
- Add minimal HTTP servers and `/healthz` endpoints for all three services.
- Add shared config loader, logger, and runtime helpers under `services/shared/go`.
- Add service-manager process registry and health aggregation.
- Add Docker Compose for PostgreSQL, Redis, and the three services.

## Account-Service Follow-Ups

- Add password hashing and session-token strategy.
- Add provider registry interface for Google, Apple, WeChat, and other common social providers.
- Add provider mock adapters for local development.
- Add future web account portal requirements as a PM-led product flow.

## Upgrade-Service Follow-Ups

- Add deployment record model and migration.
- Add active-target switching model and rollback model.
- Add client version-check API request and response schema.
- Add service-target query and switch API schema.
- Add explicit build-number governance separate from display-version strings.

## Service-Manager Follow-Ups

- Add local process start / stop / restart commands.
- Add health snapshot aggregation and status summaries.
- Add profile switching and target visibility.
- Add local hot-switch contract and lock handling.

## Future Product Work

- Define PM draft for shared web account portal.
- Run UD on the portal after PM draft is approved.
- Produce PM final, then AM / SD / QA for the portal product.

## Environment Blockers

- Go is installed locally at `$HOME/develop/go`, but Docker support is still missing.
- `docker` and `docker compose` are not currently available on this machine.
