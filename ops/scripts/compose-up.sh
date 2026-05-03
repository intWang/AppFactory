#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
COMPOSE_FILE="$REPO_ROOT/ops/compose/docker-compose.yml"
DOCKER_BIN_DIR="${DOCKER_BIN_DIR:-$HOME/Applications/Docker.app/Contents/Resources/bin}"

export PATH="$DOCKER_BIN_DIR:$PATH"

docker compose -f "$COMPOSE_FILE" up --build -d

for container_id in $(docker compose -f "$COMPOSE_FILE" ps -a -q); do
  status="$(docker inspect "$container_id" --format '{{.State.Status}}')"
  if [ "$status" = "created" ]; then
    docker start "$container_id" >/dev/null
  fi
done

echo "Compose services are available on:"
echo "  service-manager: http://localhost:18080"
echo "  account-service: http://localhost:18081"
echo "  upgrade-service: http://localhost:18082"
