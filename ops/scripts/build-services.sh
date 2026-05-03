#!/usr/bin/env bash
set -euo pipefail

GO_BIN="${GO_BIN:-$HOME/develop/go/bin/go}"
REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
SERVICES_ROOT="$REPO_ROOT/services"

mkdir -p "$SERVICES_ROOT/account-service/bin"
mkdir -p "$SERVICES_ROOT/upgrade-service/bin"
mkdir -p "$SERVICES_ROOT/service-manager/bin"

cd "$SERVICES_ROOT"
"$GO_BIN" build -o "$SERVICES_ROOT/account-service/bin/account-service" ./account-service/cmd/account-service
"$GO_BIN" build -o "$SERVICES_ROOT/upgrade-service/bin/upgrade-service" ./upgrade-service/cmd/upgrade-service
"$GO_BIN" build -o "$SERVICES_ROOT/service-manager/bin/service-manager" ./service-manager/cmd/service-manager
