# Services Foundation TODO

## Immediate Follow-Ups

- Define Redis usage boundaries to avoid turning cache into system-of-record.
- Add service-manager state persistence beyond process memory.
- Add release-target switching through `service-manager`, not just directly through `upgrade-service`.
- Add deploy/switch audit queries for release history views.

## Next Technical Slice

- Add account-service password hashing and session issuance.
- Add provider mock adapters and callback contracts for social-login providers.
- Add upgrade-service release history, switch history, and rollback history query APIs.
- Add service-manager release orchestration endpoints that call `upgrade-service`.
- Add service-manager lock handling for conflicting start/stop/restart/switch actions.

## Account-Service Follow-Ups

- Add password hashing and session-token strategy.
- Add provider registry interface for Google, Apple, WeChat, and other common social providers.
- Add provider mock adapters for local development.
- Add future web account portal requirements as a PM-led product flow.

## Upgrade-Service Follow-Ups

- Add explicit build-number governance separate from display-version strings.
- Add release-channel awareness and staged rollout strategy.
- Add read APIs for deployments, switch events, and rollback events.

## Service-Manager Follow-Ups

- Add version-switch actions that coordinate with `upgrade-service`.
- Add profile switching persistence and target visibility snapshots.
- Add local hot-switch contract and lock handling.
- Add Compose and native-mode parity checks in automated tests.

## Future Product Work

- Define PM draft for shared web account portal.
- Run UD on the portal after PM draft is approved.
- Produce PM final, then AM / SD / QA for the portal product.

## Environment Blockers

- `Go` is installed locally at `$HOME/develop/go`.
- Docker Desktop is installed locally at `$HOME/Applications/Docker.app`.
- Non-login shells may need `PATH="$HOME/Applications/Docker.app/Contents/Resources/bin:$PATH"` before invoking `docker`.
