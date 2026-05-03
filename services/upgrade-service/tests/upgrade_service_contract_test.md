# Upgrade Service Contract

## Required Endpoints

- `POST /v1/upgrade/check-client`
- `POST /v1/upgrade/check-service`
- `POST /v1/releases`
- `POST /v1/deployments`
- `POST /v1/switches`
- `POST /v1/rollbacks`
- `GET /v1/targets/active`
- `GET /healthz`

## Upgrade Rule

- build gap `<= 0`: no upgrade
- build gap `1..3`: optional upgrade
- build gap `> 3`: forced upgrade
