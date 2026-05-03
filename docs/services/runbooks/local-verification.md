# Local Services Verification

## Health Checks

Run:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8081/healthz
curl http://localhost:8082/healthz
```

## Contract Checks

Run:

```bash
curl http://localhost:8080/v1/services
curl http://localhost:8081/v1/providers
curl http://localhost:8082/v1/targets/active
curl http://localhost:8082/v1/upgrade/check-client
```

## Expected Upgrade Behavior

- latest build not newer: no upgrade
- latest build within 3 builds: optional upgrade
- latest build more than 3 builds ahead: forced upgrade

## Current Blockers

- Native compile verification is unblocked with local Go.
- Compose verification is still blocked until Docker with Compose is installed.
