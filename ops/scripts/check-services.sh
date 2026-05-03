#!/usr/bin/env bash
set -euo pipefail

echo "Check shared services foundation endpoints"
echo "  curl http://localhost:8080/healthz"
echo "  curl http://localhost:8081/healthz"
echo "  curl http://localhost:8082/healthz"
echo "  curl http://localhost:8082/v1/upgrade/check-client"
