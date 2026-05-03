#!/usr/bin/env bash
set -euo pipefail

echo "Start AppFactory shared services foundation"
echo "Expected services:"
echo "  - postgres"
echo "  - redis"
echo "  - account-service"
echo "  - upgrade-service"
echo "  - service-manager"
echo
echo "Native mode examples:"
echo "  sh ops/scripts/build-services.sh"
echo "  cd services/account-service && ./bin/account-service"
echo "  cd services/upgrade-service && ./bin/upgrade-service"
echo "  cd services/service-manager && ./bin/service-manager"
echo
echo "Compose mode example:"
echo "  docker compose -f ops/compose/docker-compose.yml up --build"
