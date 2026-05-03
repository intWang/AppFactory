# Services Foundation V1 Design

**Date:** 2026-05-03  
**Status:** Draft for implementation  
**Scope:** Shared technical foundation for account, upgrade, and local service management  
**Primary stack:** Go services, PostgreSQL, Redis  
**Local orchestration:** native process manager + Docker Compose  
**Future extension:** C++ performance modules, product-specific servers, PM/UD-led web account portal

## 1. Overview

This project adds a shared backend foundation to AppFactory so multiple products can rely on the same account, upgrade, and local service-management capabilities. The first slice is technical-infrastructure-first rather than product-first. It prioritizes stable service boundaries, local startup, health visibility, version governance, and a clean path toward a multi-service architecture.

The foundation includes three services:

- `account-service`
- `upgrade-service`
- `service-manager`

These services are shared infrastructure and therefore live at repository root under `services/`, not inside any single product. Product-specific servers can still be created later under `products/<product-slug>/server/`, but only when PM and AM decide that product-specific business logic must diverge from the shared baseline.

## 2. Goals and Non-Goals

### Goals

- Provide a shared account-service with local account flows and social-login provider interfaces.
- Provide a shared upgrade-service that governs both client and service version decisions.
- Provide a service-manager that can start, stop, inspect, and switch local service targets.
- Support local startup with both native-process and Docker Compose workflows.
- Persist core state with PostgreSQL and operational short-lived state with Redis.
- Make upgrade logic explicit, including deployment records, active version pointers, and rollback history.
- Reserve integration points for future PM/UD-led web account and registration experiences.
- Maintain a strict shared-service boundary so products reuse common services instead of cloning them.

### Non-Goals

- Full production-grade social OAuth integrations in this slice.
- Full release orchestration, approvals, or store publishing.
- Full service discovery mesh, production-grade traffic shaping, or multi-region deployment.
- Full browser portal for end users. The future web account surface is tracked but not implemented here.

## 3. Repository Layout

The new backend foundation should live in this structure:

```text
services/
  account-service/
    cmd/account-service/
    internal/
    migrations/
    configs/
    tests/
  upgrade-service/
    cmd/upgrade-service/
    internal/
    migrations/
    configs/
    tests/
  service-manager/
    cmd/service-manager/
    internal/
    configs/
    tests/
  shared/
    go/
      config/
      health/
      httpx/
      logging/
      runtime/
    proto/
    configs/

ops/
  compose/
  scripts/

docs/services/
  specs/
  plans/
  runbooks/

todo/
  services-foundation-todo.md
```

## 4. Service Boundaries

### 4.1 account-service

`account-service` owns identity and account-domain truth:

- local account registration
- password login
- session issuance
- user profile basics
- provider registration for future social login adapters
- shared account APIs used by products and future web account surfaces

This service should not own product-specific billing, feature entitlements, or product-private business workflows.

### 4.2 upgrade-service

`upgrade-service` owns version governance across both apps and services:

- client version checks
- service version checks
- forced-upgrade rules
- deployment records
- active-version pointers
- switch history
- rollback history
- release metadata such as upgrade messages and target URLs

This service is not just a version-query endpoint. It is the shared control plane for upgrade state.

### 4.3 service-manager

`service-manager` is the local control entrypoint. It does not own product or account truth. It owns local operational coordination:

- local service registration
- start / stop / restart
- health-check aggregation
- current target visibility
- active-version visibility
- config profile switching
- hot-switch command routing
- local status inspection

The first version is a light platform manager, not a full service mesh.

## 5. Data Ownership and Storage

### 5.1 PostgreSQL

PostgreSQL stores durable business and operational truth.

`account-service` durable entities:

- users
- local_credentials
- auth_identities
- sessions
- provider_links

`upgrade-service` durable entities:

- release_channels
- release_versions
- deployment_records
- active_targets
- switch_events
- rollback_events

`service-manager` durable entities:

- registered_services
- service_profiles
- switch_policies

### 5.2 Redis

Redis stores short-lived coordination and runtime state:

- local service heartbeats
- health snapshots
- hot-switch transient state
- session cache or token acceleration if needed later
- in-flight operational locks

Redis is not the source of truth for releases or account records.

## 6. Runtime Architecture

### 6.1 Local Native Process Mode

The default developer path should support starting services as local processes through `service-manager`.

`service-manager` responsibilities in native mode:

- read registered local service config
- launch service processes
- track PID and command metadata
- query HTTP health endpoints
- expose aggregated service status

### 6.2 Docker Compose Mode

Compose should be provided for parity and onboarding:

- PostgreSQL
- Redis
- account-service
- upgrade-service
- service-manager

Compose is a supported mode, not the only mode.

## 7. API Contracts

### 7.1 account-service

First-slice APIs:

- `POST /v1/accounts/register`
- `POST /v1/accounts/login`
- `POST /v1/accounts/logout`
- `GET /v1/accounts/me`
- `GET /v1/providers`

`GET /v1/providers` returns provider availability and current enablement state for future social-login adapters.

### 7.2 upgrade-service

First-slice APIs:

- `POST /v1/upgrade/check-client`
- `POST /v1/upgrade/check-service`
- `POST /v1/releases`
- `POST /v1/deployments`
- `POST /v1/switches`
- `POST /v1/rollbacks`
- `GET /v1/targets/active`

The client upgrade check returns:

- current version
- latest version
- latest build number
- whether upgrade is available
- whether upgrade is forced
- upgrade message
- upgrade URL

### 7.3 service-manager

First-slice APIs:

- `GET /v1/services`
- `POST /v1/services/start`
- `POST /v1/services/stop`
- `POST /v1/services/restart`
- `GET /v1/services/health`
- `POST /v1/services/switch-profile`
- `GET /v1/services/targets`

## 8. Upgrade Policy

Upgrade checks are shared behavior and must not be reinvented per product.

Default policy:

- if latest build is not newer than current build: no prompt
- if latest build is newer but within 3 builds: optional prompt
- if latest build is more than 3 builds ahead: forced upgrade

This rule applies to client apps and should also inform service target management for local operational switching.

## 9. Versioning Model

The repository should standardize the display version format as:

`YY.Q.0M.BB`

Where:

- `YY` = two-digit year
- `Q` = quarter number
- `0M` = month encoded as a two-digit segment within the quarter context
- `BB` = release sequence within that month window

Example for May 2026:

- `26.2.20.01`
- `26.2.20.02`

However, forced-upgrade decisions must not rely on string comparison. They must rely on an explicit numeric build sequence maintained by `upgrade-service`.

## 10. Shared Web Account Future

The future browser-based account surface is a separate product workflow. It should:

- use shared account-service APIs
- support account registration and login
- support common social-login providers
- follow the full workflow:
  `PM draft -> UD -> PM final -> AM -> SD -> QA`

This slice only prepares the shared backend interfaces required by that future product.

## 11. Testing Strategy

### 11.1 account-service

- unit tests for registration and login handlers
- storage tests for account persistence boundaries
- provider-registry tests

### 11.2 upgrade-service

- unit tests for upgrade rules
- tests for deployment and active-target switching
- tests for rollback state transitions

### 11.3 service-manager

- tests for service registry loading
- tests for health aggregation
- tests for switch-profile behavior

### 11.4 integration

- local workflow tests for:
  - manager sees both services
  - services expose health
  - upgrade-service returns forced upgrade when build gap exceeds threshold

## 12. Delivery Sequence

The recommended implementation sequence is:

1. create service spec and todo
2. create implementation plan
3. scaffold shared Go workspace and directories
4. scaffold account-service
5. scaffold upgrade-service
6. scaffold service-manager
7. add local configs and compose
8. add first health and minimal business endpoints
9. add QA-facing runbook and verification notes

## 13. Risks

- Go toolchain is not currently installed on the local machine.
- Docker / Docker Compose is not currently installed on the local machine.
- Social provider integration can expand scope quickly if not contained.
- Local process hot-switch semantics can become over-engineered if not kept lightweight in V1.

## 14. Gate Result

Approved for implementation as a shared infrastructure slice inside AppFactory. Immediate next step is writing the implementation plan and scaffolding the repository structure, while recording local environment blockers in runbooks and todo tracking.
