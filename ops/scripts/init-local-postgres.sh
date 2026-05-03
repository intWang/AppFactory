#!/usr/bin/env bash
set -euo pipefail

POSTGRES_APP_DIR="${POSTGRES_APP_DIR:-$HOME/Applications/Postgres.app}"
POSTGRES_BIN_DIR="${POSTGRES_BIN_DIR:-$POSTGRES_APP_DIR/Contents/Versions/latest/bin}"
PGDATA="${PGDATA:-$HOME/.appfactory/postgres-data}"
PGLOG="${PGLOG:-$HOME/.appfactory/postgres.log}"
PGPORT="${PGPORT:-5432}"
PGUSER="${PGUSER:-appfactory}"
PGDATABASE="${PGDATABASE:-appfactory}"
PGPASSWORD="${PGPASSWORD:-appfactory}"

mkdir -p "$(dirname "$PGDATA")"

if [ ! -f "$PGDATA/PG_VERSION" ]; then
  LANG="${LANG:-en_US.UTF-8}" LC_ALL="${LC_ALL:-en_US.UTF-8}" \
    "$POSTGRES_BIN_DIR/initdb" -D "$PGDATA" -U "$PGUSER" --auth=trust --locale=C --encoding=UTF8
fi

if ! "$POSTGRES_BIN_DIR/pg_isready" -h localhost -p "$PGPORT" -U "$PGUSER" >/dev/null 2>&1; then
  "$POSTGRES_BIN_DIR/pg_ctl" -D "$PGDATA" -l "$PGLOG" -o "-p $PGPORT" start
fi

if ! "$POSTGRES_BIN_DIR/psql" -p "$PGPORT" -U "$PGUSER" -d postgres -tAc "SELECT 1 FROM pg_database WHERE datname='$PGDATABASE'" | grep -q 1; then
  "$POSTGRES_BIN_DIR/createdb" -p "$PGPORT" -U "$PGUSER" "$PGDATABASE"
fi

"$POSTGRES_BIN_DIR/psql" -p "$PGPORT" -U "$PGUSER" -d postgres -c "ALTER USER $PGUSER WITH PASSWORD '$PGPASSWORD';" >/dev/null

echo "Local PostgreSQL is running on port $PGPORT"
