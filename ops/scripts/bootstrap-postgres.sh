#!/usr/bin/env bash
set -euo pipefail

POSTGRES_APP_VERSION="${POSTGRES_APP_VERSION:-2.9.4}"
POSTGRES_MAJOR="${POSTGRES_MAJOR:-16}"
POSTGRES_DMG_URL="${POSTGRES_DMG_URL:-https://github.com/PostgresApp/PostgresApp/releases/download/v${POSTGRES_APP_VERSION}/Postgres-${POSTGRES_APP_VERSION}-${POSTGRES_MAJOR}.dmg}"
POSTGRES_APP_DIR="${POSTGRES_APP_DIR:-$HOME/Applications/Postgres.app}"
POSTGRES_BIN_DIR="$POSTGRES_APP_DIR/Contents/Versions/latest/bin"
DOWNLOAD_PATH="${DOWNLOAD_PATH:-/tmp/Postgres-${POSTGRES_APP_VERSION}-${POSTGRES_MAJOR}.dmg}"

mkdir -p "$HOME/Applications"

if [ ! -d "$POSTGRES_APP_DIR" ]; then
  echo "Downloading Postgres.app from $POSTGRES_DMG_URL"
  curl -L "$POSTGRES_DMG_URL" -o "$DOWNLOAD_PATH"
  MOUNT_POINT="$(hdiutil attach "$DOWNLOAD_PATH" -nobrowse | awk '/\/Volumes\// {print $NF; exit}')"
  cp -R "$MOUNT_POINT/Postgres.app" "$POSTGRES_APP_DIR"
  hdiutil detach "$MOUNT_POINT" -quiet
fi

if ! grep -q 'POSTGRES_APP_ROOT' "$HOME/.zprofile" 2>/dev/null; then
  cat >>"$HOME/.zprofile" <<'EOF'
export POSTGRES_APP_ROOT="$HOME/Applications/Postgres.app/Contents/Versions/latest"
export PATH="$POSTGRES_APP_ROOT/bin:$PATH"
EOF
fi

if ! grep -q 'POSTGRES_APP_ROOT' "$HOME/.zshrc" 2>/dev/null; then
  cat >>"$HOME/.zshrc" <<'EOF'
export POSTGRES_APP_ROOT="$HOME/Applications/Postgres.app/Contents/Versions/latest"
export PATH="$POSTGRES_APP_ROOT/bin:$PATH"
EOF
fi

echo "PostgreSQL binaries available at: $POSTGRES_BIN_DIR"
"$POSTGRES_BIN_DIR/psql" --version
