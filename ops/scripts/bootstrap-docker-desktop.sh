#!/usr/bin/env bash
set -euo pipefail

DOCKER_DMG_URL="${DOCKER_DMG_URL:-https://desktop.docker.com/mac/main/amd64/Docker.dmg}"
DOWNLOAD_PATH="${DOWNLOAD_PATH:-/tmp/Docker.dmg}"
APP_DEST="${APP_DEST:-$HOME/Applications/Docker.app}"
DOCKER_APP_ROOT="$APP_DEST/Contents/Resources/bin"
FORCE_DOWNLOAD="${FORCE_DOWNLOAD:-false}"

echo "Downloading Docker Desktop from $DOCKER_DMG_URL"
if [ "$FORCE_DOWNLOAD" = "true" ] || [ ! -f "$DOWNLOAD_PATH" ]; then
  curl -L "$DOCKER_DMG_URL" -o "$DOWNLOAD_PATH"
fi
MOUNT_POINT="$(hdiutil attach "$DOWNLOAD_PATH" -nobrowse | awk '/\/Volumes\// {print $NF; exit}')"
mkdir -p "$(dirname "$APP_DEST")"
rm -rf "$APP_DEST"
ditto "$MOUNT_POINT/Docker.app" "$APP_DEST"
hdiutil detach "$MOUNT_POINT" -quiet

for shell_rc in "$HOME/.zprofile" "$HOME/.zshrc"; do
  touch "$shell_rc"
  if ! grep -q 'Docker.app/Contents/Resources/bin' "$shell_rc"; then
    printf '\nexport PATH="%s:$PATH"\n' "$DOCKER_APP_ROOT" >>"$shell_rc"
  fi
done

echo "Docker Desktop copied to $APP_DEST"
echo "Open Docker.app once to complete first-run setup, then run: docker compose version"
