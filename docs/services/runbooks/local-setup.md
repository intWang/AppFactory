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
sh ops/scripts/compose-up.sh
```

4. If you need to rebuild a changed service image from a non-login shell, ensure Docker's bundled binaries are on `PATH`:

```bash
export PATH="$HOME/Applications/Docker.app/Contents/Resources/bin:$PATH"
docker compose -f ops/compose/docker-compose.yml build upgrade-service
docker compose -f ops/compose/docker-compose.yml up -d upgrade-service
```

Compose host ports:

- `service-manager`: `http://localhost:18080`
- `account-service`: `http://localhost:18081`
- `upgrade-service`: `http://localhost:18082`

Example upgrade lifecycle smoke test:

```bash
release=$(curl -sS -X POST http://localhost:18082/v1/releases \
  -H 'Content-Type: application/json' \
  -d '{"product_slug":"shared-client","target_type":"client","version_label":"26.2.20.06","build_number":6,"upgrade_url":"https://example.com/client/26.2.20.06"}')

release_id=$(printf '%s' "$release" | sed -n 's/.*"id":"\([^"]*\)".*/\1/p')

curl -sS -X POST http://localhost:18082/v1/deployments \
  -H 'Content-Type: application/json' \
  -d "{\"target_version_id\":\"$release_id\",\"environment\":\"compose-local\",\"status\":\"deployed\"}"

curl -sS -X POST http://localhost:18082/v1/switches \
  -H 'Content-Type: application/json' \
  -d "{\"product_slug\":\"shared-client\",\"target_type\":\"client\",\"to_version_id\":\"$release_id\",\"operator\":\"service-manager\"}"

curl -sS http://localhost:18082/v1/targets/active

curl -sS -X POST http://localhost:18082/v1/rollbacks \
  -H 'Content-Type: application/json' \
  -d '{"product_slug":"shared-client","target_type":"client","rolled_back_to_version_id":"rv-client-2622004","operator":"qa"}'
```

Service-manager release control endpoints:

- `GET http://localhost:18080/v1/releases/targets`
- `GET http://localhost:18080/v1/releases/operations/current`
- `GET http://localhost:18080/v1/releases/operations/history`
- `GET http://localhost:18080/v1/releases/history`
- `GET http://localhost:18080/v1/deployments/history`
- `GET http://localhost:18080/v1/releases/switches/history`
- `GET http://localhost:18080/v1/releases/rollbacks/history`
- `POST http://localhost:18080/v1/releases/create`
- `POST http://localhost:18080/v1/deployments/create`
- `POST http://localhost:18080/v1/releases/promote`
- `POST http://localhost:18080/v1/releases/switch`
- `POST http://localhost:18080/v1/releases/rollback`

Mutating release actions are guarded by an in-process lock keyed by `product_slug + target_type`.
If another release operation is already running for the same target, the service-manager returns `409 Conflict`.

Service-manager persists release operation snapshots to:

- native mode: `services/service-manager/data/service-manager-operations.json`
- compose mode: PostgreSQL table `service_manager_operations` in database `appfactory`

Example compose verification:

```bash
docker exec compose-postgres-1 psql -U appfactory -d appfactory \
  -c "SELECT operation, product_slug, target_type, status, environment FROM service_manager_operations ORDER BY started_at DESC LIMIT 5;"
```

## Current Blockers

- Go is installed locally at `$HOME/develop/go`.
- Docker Desktop is installed locally at `$HOME/Applications/Docker.app`.
