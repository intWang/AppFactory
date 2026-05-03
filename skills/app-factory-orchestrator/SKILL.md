---
name: app-factory-orchestrator
description: Use when a user wants a lightweight tool app moved through the full App Factory workflow from intake to QA with governed role handoffs
---

# App Factory Orchestrator

## Overview

Coordinate App Factory V1 from intake through QA. Advance automatically when required artifacts and gates exist, and pause only on defined risk or rework conditions.

## When to Use

- A user wants a lightweight tool app taken from idea to compiled, testable output.
- The work should flow through PM, UD, AM, SD, and QA instead of ad hoc execution.
- The product may be client-only now or may need a reserved or active server boundary.

## Workflow

1. Create `products/<project-slug>/docs/00-intake.md`.
2. Initialize the product tree and baseline artifacts.
3. Route to PM for `products/<project-slug>/docs/01-prd-draft.md`.
4. Route to UD for `products/<project-slug>/docs/02-ux-spec.md`, `products/<project-slug>/design/figma-link.md`, and exported screens.
5. Route back to PM for `products/<project-slug>/docs/03-prd-final.md`.
6. Route to AM for `products/<project-slug>/docs/04-architecture.md`.
7. Route to SD for `products/<project-slug>/docs/05-dev-plan.md` and implementation.
8. Route to QA for `products/<project-slug>/docs/06-test-cases.md` and `products/<project-slug>/docs/07-test-report.md`.
9. Stop only for blocked or rework conditions and record the current state in the owning product docs.

## Product Initialization

Before routing work, initialize the product tree with at least:

- `products/<project-slug>/client/pubspec.yaml`
- `products/<project-slug>/client/lib/`
- `products/<project-slug>/client/test/unit/`
- `products/<project-slug>/client/test/widget/`
- `products/<project-slug>/client/config/env/`
- `products/<project-slug>/docs/`
- `products/<project-slug>/design/exports/`
- `products/<project-slug>/build/outputs/`

Create `products/<project-slug>/server/` when AM decides the product needs active backend logic. Otherwise reserve the boundary in architecture and docs for future enablement.

Create or copy these baseline product artifacts:

- `docs/00-intake.md`
- `docs/01-prd-draft.md`
- `docs/02-ux-spec.md`
- `docs/03-prd-final.md`
- `docs/04-architecture.md`
- `docs/05-dev-plan.md`
- `docs/06-test-cases.md`
- `docs/07-test-report.md`
- `docs/08-release-gate.md`
- `docs/09-retro-input.md`
- `design/figma-link.md`

Use `templates/docs/` for document content and `templates/flutter_product_shell/` for the product shell baseline.

Preferred initialization command:

```bash
python3 scripts/init_product.py <project-slug> "<App Name>"
```

## Initialization Checks

Do not advance beyond intake until:

- the product exists under `products/<project-slug>/`
- the product has a valid `client/` shell
- all required `docs/` files exist
- `design/figma-link.md` exists
- `design/exports/` exists
- `build/outputs/` exists
- the product does not depend on another product directory
- the server decision is explicitly recorded before SD starts implementation

## Intake Artifact

Write `products/<project-slug>/docs/00-intake.md` with:

- `## Project Name`
- `## Project Slug`
- `## Input Source`
- `## User Requirement`
- `## Default Assumptions`
- `## Scope Decision`
- `## Current State`
- `## Next Role`
- `## Risks`
- `## Unknowns`
- `## Gate Result`

## Handoff Rules

- Never advance to the next role without the current role's required output file.
- Route rework back to the role that owns the missing decision instead of patching around it downstream.
- Keep all docs, design references, tests, and build outputs inside the owning `products/<project-slug>/` boundary.
- Record blocked status and next unblocker in the current doc rather than silently stalling.

## State Machine

- `intake`
- `producting`
- `designing`
- `architecting`
- `developing`
- `testing`
- `passed`
- `blocked`
- `needs-rework`

## Pause Conditions

- Missing core requirement data
- Out-of-scope request
- Credentials, money, signing, account, or release risk
- Development environment blocker
- Failed QA gate
- Product docs or outputs are written outside the owning `products/<project-slug>/` boundary
- Design assets are missing before architecture or implementation handoff
