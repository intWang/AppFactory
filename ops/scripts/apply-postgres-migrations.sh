#!/usr/bin/env bash
set -euo pipefail

POSTGRES_APP_DIR="${POSTGRES_APP_DIR:-$HOME/Applications/Postgres.app}"
POSTGRES_BIN_DIR="${POSTGRES_BIN_DIR:-$POSTGRES_APP_DIR/Contents/Versions/latest/bin}"
PGPORT="${PGPORT:-5432}"
PGUSER="${PGUSER:-appfactory}"
PGDATABASE="${PGDATABASE:-appfactory}"
REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"

"$POSTGRES_BIN_DIR/psql" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" \
  -c "CREATE TABLE IF NOT EXISTS schema_migrations (filename TEXT PRIMARY KEY, applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW());" >/dev/null

is_applied() {
  local filename="$1"
  "$POSTGRES_BIN_DIR/psql" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" -tAc \
    "SELECT 1 FROM schema_migrations WHERE filename = '$filename'" | grep -q 1
}

run_dir() {
  local migration_dir="$1"
  for file in "$migration_dir"/*.sql; do
    [ -f "$file" ] || continue
    local filename
    filename="${file#$REPO_ROOT/}"
    if is_applied "$filename"; then
      echo "Skipping migration already applied: $file"
      continue
    fi
    echo "Applying migration: $file"
    "$POSTGRES_BIN_DIR/psql" -v ON_ERROR_STOP=1 -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" -f "$file"
    "$POSTGRES_BIN_DIR/psql" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" \
      -c "INSERT INTO schema_migrations (filename) VALUES ('$filename');" >/dev/null
  done
}

run_dir "$REPO_ROOT/services/account-service/migrations"
run_dir "$REPO_ROOT/services/upgrade-service/migrations"
run_dir "$REPO_ROOT/services/service-manager/migrations"
