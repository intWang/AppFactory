# Services Foundation TODO

## Immediate Follow-Ups

- Install and verify Docker with Compose support on the local machine.
- Define Redis usage boundaries to avoid turning cache into system-of-record.
- Replace in-memory account storage with PostgreSQL-backed repositories.
- Replace in-memory upgrade storage with PostgreSQL-backed repositories.
- Add service-manager state persistence and local process execution wiring.

## Next Technical Slice

- Add Docker Compose for PostgreSQL, Redis, and the three services.
- Add runnable Compose verification once Docker support exists.
- Add shared config loader with real YAML parsing and environment override support.
- Add account-service password hashing and session issuance.
- Add upgrade-service mutation handlers for releases, deployments, switches, and rollbacks.
- Add service-manager process execution and restart handling.

## Account-Service Follow-Ups

- Add password hashing and session-token strategy.
- Add provider registry interface for Google, Apple, WeChat, and other common social providers.
- Add provider mock adapters for local development.
- Add future web account portal requirements as a PM-led product flow.

## Upgrade-Service Follow-Ups

- Add client version-check API request and response schema.
- Add service-target query and switch API schema.
- Add explicit build-number governance separate from display-version strings.

## Service-Manager Follow-Ups

- Add local process start / stop / restart commands.
- Add profile switching and target visibility.
- Add local hot-switch contract and lock handling.

## Future Product Work

- Define PM draft for shared web account portal.
- Run UD on the portal after PM draft is approved.
- Produce PM final, then AM / SD / QA for the portal product.

## Environment Blockers

- Go is installed locally at `$HOME/develop/go`, but Docker support is still missing.
- `docker` and `docker compose` are not currently available on this machine.
