# Local Services Setup

## Required Tooling

- Go toolchain
- Docker with Compose support
- PostgreSQL access for local mode or Compose
- Redis access for local mode or Compose

## Native Process Mode

1. Install `go`.
2. Start PostgreSQL and Redis locally.
3. Start `account-service` on `:8081`.
4. Start `upgrade-service` on `:8082`.
5. Start `service-manager` on `:8080`.

## Compose Mode

Run:

```bash
docker compose -f ops/compose/docker-compose.yml up --build
```

## Current Blockers

- Go is installed locally at `$HOME/develop/go`.
- `docker compose` is not installed on this machine yet.
