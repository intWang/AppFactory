# Service Manager Contract

## Required Endpoints

- `GET /v1/services`
- `POST /v1/services/start`
- `POST /v1/services/stop`
- `POST /v1/services/restart`
- `GET /v1/services/health`
- `POST /v1/services/switch-profile`
- `GET /v1/services/targets`
- `GET /healthz`

## Scope

- Local light control plane only
- Native process mode first
- Compose support required for parity
