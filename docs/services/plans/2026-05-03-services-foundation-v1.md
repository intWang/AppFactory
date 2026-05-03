# Services Foundation V1 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development or superpowers:executing-plans to implement this plan task-by-task when delegation is appropriate.

**Goal:** Build the first shared backend foundation for AppFactory, including account-service, upgrade-service, service-manager, shared Go support packages, local config files, and verification assets.

**Architecture:** Three Go services share a root `services/` workspace. `account-service` owns account-domain truth. `upgrade-service` owns version and deployment governance. `service-manager` owns local service control and health aggregation. Shared libraries provide logging, config, health, and runtime helpers. PostgreSQL and Redis are modeled now, with local runtime integration prepared even if the toolchain is not yet installed.

**Tech Stack:** Go, HTTP JSON APIs, PostgreSQL, Redis, Docker Compose, Markdown runbooks, shell scripts

**Known Local Blockers:** Docker / Docker Compose are not currently installed on this machine. Go can be installed and verified locally, so implementation should proceed through scaffold and compile validation even if full containerized runtime verification remains deferred.

---

## File Structure

### New directories

- Create: `services/account-service/cmd/account-service`
- Create: `services/account-service/internal/httpapi`
- Create: `services/account-service/internal/domain`
- Create: `services/account-service/internal/storage`
- Create: `services/account-service/configs`
- Create: `services/account-service/migrations`
- Create: `services/account-service/tests`
- Create: `services/upgrade-service/cmd/upgrade-service`
- Create: `services/upgrade-service/internal/httpapi`
- Create: `services/upgrade-service/internal/domain`
- Create: `services/upgrade-service/internal/storage`
- Create: `services/upgrade-service/configs`
- Create: `services/upgrade-service/migrations`
- Create: `services/upgrade-service/tests`
- Create: `services/service-manager/cmd/service-manager`
- Create: `services/service-manager/internal/httpapi`
- Create: `services/service-manager/internal/domain`
- Create: `services/service-manager/internal/runtime`
- Create: `services/service-manager/configs`
- Create: `services/service-manager/tests`
- Create: `services/shared/go/config`
- Create: `services/shared/go/health`
- Create: `services/shared/go/httpx`
- Create: `services/shared/go/logging`
- Create: `services/shared/go/runtime`
- Create: `services/shared/proto`
- Create: `services/shared/configs`
- Create: `ops/compose`
- Create: `ops/scripts`
- Create: `docs/services/runbooks`

### New files

- Create: `services/go.work`
- Create: `services/account-service/go.mod`
- Create: `services/account-service/cmd/account-service/main.go`
- Create: `services/account-service/internal/httpapi/router.go`
- Create: `services/account-service/internal/domain/account.go`
- Create: `services/account-service/internal/storage/memory_store.go`
- Create: `services/account-service/configs/local.yaml`
- Create: `services/account-service/tests/account_service_contract_test.md`
- Create: `services/upgrade-service/go.mod`
- Create: `services/upgrade-service/cmd/upgrade-service/main.go`
- Create: `services/upgrade-service/internal/httpapi/router.go`
- Create: `services/upgrade-service/internal/domain/release.go`
- Create: `services/upgrade-service/internal/storage/memory_store.go`
- Create: `services/upgrade-service/configs/local.yaml`
- Create: `services/upgrade-service/tests/upgrade_service_contract_test.md`
- Create: `services/service-manager/go.mod`
- Create: `services/service-manager/cmd/service-manager/main.go`
- Create: `services/service-manager/internal/httpapi/router.go`
- Create: `services/service-manager/internal/domain/service.go`
- Create: `services/service-manager/internal/runtime/registry.go`
- Create: `services/service-manager/configs/local.yaml`
- Create: `services/service-manager/tests/service_manager_contract_test.md`
- Create: `services/shared/go/config/config.go`
- Create: `services/shared/go/health/health.go`
- Create: `services/shared/go/httpx/json.go`
- Create: `services/shared/go/logging/logger.go`
- Create: `services/shared/go/runtime/runtime.go`
- Create: `ops/compose/docker-compose.yml`
- Create: `ops/scripts/start-services.sh`
- Create: `ops/scripts/check-services.sh`
- Create: `docs/services/runbooks/local-setup.md`
- Create: `docs/services/runbooks/local-verification.md`

## Task 1: Scaffold the Services Workspace

**Files:**
- Create all directories listed under `services/`, `ops/`, and `docs/services/runbooks`
- Create: `services/go.work`

- [ ] **Step 1: Write the failing directory contract**

Record this target structure in notes or a temporary verification command:

```text
services/account-service
services/upgrade-service
services/service-manager
services/shared/go
ops/compose
ops/scripts
docs/services/runbooks
```

- [ ] **Step 2: Verify the structure does not exist yet**

Run:

```bash
find services ops docs/services/runbooks -maxdepth 2 -type d 2>/dev/null | sort
```

Expected: some or all of the target directories are missing.

- [ ] **Step 3: Create the minimal workspace tree**

Create the directory structure and add a minimal `services/go.work` file that references:

```text
./account-service
./upgrade-service
./service-manager
```

- [ ] **Step 4: Verify the structure now exists**

Run:

```bash
find services ops docs/services/runbooks -maxdepth 2 -type d | sort
```

- [ ] **Step 5: Commit**

```bash
git add services ops docs/services/runbooks
git commit -m "chore: scaffold shared services workspace"
```

## Task 2: Scaffold Shared Go Support Packages

**Files:**
- Create: `services/shared/go/config/config.go`
- Create: `services/shared/go/health/health.go`
- Create: `services/shared/go/httpx/json.go`
- Create: `services/shared/go/logging/logger.go`
- Create: `services/shared/go/runtime/runtime.go`

- [ ] **Step 1: Write failing contract notes**

Each package must expose one minimal type or function:

- `config`: app config struct
- `health`: health snapshot struct
- `httpx`: JSON writer helper
- `logging`: logger interface or wrapper
- `runtime`: service metadata struct

- [ ] **Step 2: Verify files are absent**

Run:

```bash
find services/shared/go -maxdepth 2 -type f 2>/dev/null | sort
```

- [ ] **Step 3: Create minimal package stubs**

Each file should compile in principle and remain intentionally small.

- [ ] **Step 4: Static verify package declarations**

Run:

```bash
grep -R "^package " -n services/shared/go
```

- [ ] **Step 5: Commit**

```bash
git add services/shared/go
git commit -m "feat: add shared services support packages"
```

## Task 3: Scaffold account-service

**Files:**
- Create: `services/account-service/go.mod`
- Create: `services/account-service/cmd/account-service/main.go`
- Create: `services/account-service/internal/httpapi/router.go`
- Create: `services/account-service/internal/domain/account.go`
- Create: `services/account-service/internal/storage/memory_store.go`
- Create: `services/account-service/configs/local.yaml`
- Create: `services/account-service/tests/account_service_contract_test.md`

- [ ] **Step 1: Write failing API contract**

Document that first-slice endpoints must include:

```text
POST /v1/accounts/register
POST /v1/accounts/login
POST /v1/accounts/logout
GET /v1/accounts/me
GET /v1/providers
GET /healthz
```

- [ ] **Step 2: Verify files do not exist yet**

Run:

```bash
find services/account-service -maxdepth 4 -type f 2>/dev/null | sort
```

- [ ] **Step 3: Create minimal service scaffold**

Add a small `net/http` router with stub handlers and a simple in-memory store placeholder.

- [ ] **Step 4: Verify contract text is present**

Run:

```bash
grep -R "/v1/accounts/register\\|/healthz" -n services/account-service
```

- [ ] **Step 5: Commit**

```bash
git add services/account-service
git commit -m "feat: scaffold account service"
```

## Task 4: Scaffold upgrade-service

**Files:**
- Create: `services/upgrade-service/go.mod`
- Create: `services/upgrade-service/cmd/upgrade-service/main.go`
- Create: `services/upgrade-service/internal/httpapi/router.go`
- Create: `services/upgrade-service/internal/domain/release.go`
- Create: `services/upgrade-service/internal/storage/memory_store.go`
- Create: `services/upgrade-service/configs/local.yaml`
- Create: `services/upgrade-service/tests/upgrade_service_contract_test.md`

- [ ] **Step 1: Write failing API contract**

Document that first-slice endpoints must include:

```text
POST /v1/upgrade/check-client
POST /v1/upgrade/check-service
POST /v1/releases
POST /v1/deployments
POST /v1/switches
POST /v1/rollbacks
GET /v1/targets/active
GET /healthz
```

- [ ] **Step 2: Verify files do not exist yet**

Run:

```bash
find services/upgrade-service -maxdepth 4 -type f 2>/dev/null | sort
```

- [ ] **Step 3: Create minimal service scaffold**

Include the shared 3-build forced-upgrade rule in domain-level logic or contract notes.

- [ ] **Step 4: Verify contract text is present**

Run:

```bash
grep -R "check-client\\|forced" -n services/upgrade-service
```

- [ ] **Step 5: Commit**

```bash
git add services/upgrade-service
git commit -m "feat: scaffold upgrade service"
```

## Task 5: Scaffold service-manager

**Files:**
- Create: `services/service-manager/go.mod`
- Create: `services/service-manager/cmd/service-manager/main.go`
- Create: `services/service-manager/internal/httpapi/router.go`
- Create: `services/service-manager/internal/domain/service.go`
- Create: `services/service-manager/internal/runtime/registry.go`
- Create: `services/service-manager/configs/local.yaml`
- Create: `services/service-manager/tests/service_manager_contract_test.md`

- [ ] **Step 1: Write failing API contract**

Document that first-slice endpoints must include:

```text
GET /v1/services
POST /v1/services/start
POST /v1/services/stop
POST /v1/services/restart
GET /v1/services/health
POST /v1/services/switch-profile
GET /v1/services/targets
GET /healthz
```

- [ ] **Step 2: Verify files do not exist yet**

Run:

```bash
find services/service-manager -maxdepth 4 -type f 2>/dev/null | sort
```

- [ ] **Step 3: Create minimal service scaffold**

Provide a service-registry placeholder and a manager router that can report static service status.

- [ ] **Step 4: Verify contract text is present**

Run:

```bash
grep -R "switch-profile\\|/healthz" -n services/service-manager
```

- [ ] **Step 5: Commit**

```bash
git add services/service-manager
git commit -m "feat: scaffold service manager"
```

## Task 6: Add Local Config and Ops Assets

**Files:**
- Create: `ops/compose/docker-compose.yml`
- Create: `ops/scripts/start-services.sh`
- Create: `ops/scripts/check-services.sh`
- Create: `docs/services/runbooks/local-setup.md`
- Create: `docs/services/runbooks/local-verification.md`

- [ ] **Step 1: Write failing runtime expectations**

Document expected local dependencies:

```text
postgres
redis
account-service
upgrade-service
service-manager
```

- [ ] **Step 2: Verify the assets are absent**

Run:

```bash
find ops docs/services/runbooks -maxdepth 3 -type f 2>/dev/null | sort
```

- [ ] **Step 3: Create compose, scripts, and runbooks**

The scripts may remain non-executable text if the local toolchain is missing, but commands and expected behavior must be concrete.

- [ ] **Step 4: Verify references are present**

Run:

```bash
grep -R "postgres\\|redis\\|service-manager" -n ops docs/services/runbooks
```

- [ ] **Step 5: Commit**

```bash
git add ops docs/services/runbooks
git commit -m "chore: add local service ops assets"
```

## Task 7: Final Verification and Blocker Capture

**Files:**
- Modify if needed: `todo/services-foundation-todo.md`
- Modify if needed: `docs/services/runbooks/local-setup.md`

- [ ] **Step 1: Run static repository verification**

Run:

```bash
find services ops docs/services todo -maxdepth 4 -type f | sort
```

- [ ] **Step 2: Run text-level contract verification**

Run:

```bash
grep -R "/healthz\\|check-client\\|accounts/register\\|switch-profile" -n services
```

- [ ] **Step 3: Capture environment blockers with evidence**

Run:

```bash
go version
docker compose version
```

Expected: `go version` should succeed once Go is provisioned. `docker compose version` may still fail until Docker support is installed. Record that explicitly.

- [ ] **Step 4: Commit**

```bash
git add services ops docs/services todo
git commit -m "docs: record services foundation blockers"
```
