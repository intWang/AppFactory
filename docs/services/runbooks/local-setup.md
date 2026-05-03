# Local Services Setup

## Required Tooling

- Go toolchain
- Docker with Compose support
- PostgreSQL access for local mode or Compose
- Redis access for local mode or Compose

## Native Process Mode

1. Run `ops/scripts/bootstrap-postgres.sh`.
2. Run `ops/scripts/init-local-postgres.sh`.
3. Run `ops/scripts/apply-postgres-migrations.sh`.
4. Run `ops/scripts/build-services.sh`.
5. Start `service-manager` with `services/service-manager/bin/service-manager`.
6. Let `service-manager` launch `account-service` and `upgrade-service` from their `bin/` directories.

The PostgreSQL bootstrap and migration scripts are intentionally idempotent:

- `init-local-postgres.sh` reuses an already running local instance.
- `apply-postgres-migrations.sh` records applied files in `schema_migrations` and skips them on later runs.

## Compose Mode

1. Run `ops/scripts/bootstrap-docker-desktop.sh`.
2. Start Docker Desktop once so the daemon is ready.
3. Run:

```bash
docker compose -f ops/compose/docker-compose.yml up --build
```

## Current Blockers

- Go is installed locally at `$HOME/develop/go`.
- `docker compose` is not installed on this machine yet.
