#!/usr/bin/env bash
set -euo pipefail

DOCKER_DMG_URL="${DOCKER_DMG_URL:-https://desktop.docker.com/mac/main/amd64/Docker.dmg}"
DOWNLOAD_PATH="${DOWNLOAD_PATH:-/tmp/Docker.dmg}"
APP_DEST="${APP_DEST:-/Applications/Docker.app}"

echo "Downloading Docker Desktop from $DOCKER_DMG_URL"
curl -L "$DOCKER_DMG_URL" -o "$DOWNLOAD_PATH"
MOUNT_POINT="$(hdiutil attach "$DOWNLOAD_PATH" -nobrowse | awk '/\/Volumes\// {print $NF; exit}')"
sudo cp -R "$MOUNT_POINT/Docker.app" "$APP_DEST"
hdiutil detach "$MOUNT_POINT" -quiet

echo "Docker Desktop copied to $APP_DEST"
echo "Open Docker.app once to complete first-run setup, then run: docker compose version"
