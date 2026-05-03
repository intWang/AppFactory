---
name: app-factory-orchestrator
description: Use when a user wants to turn a lightweight tool app idea into a governed App Factory workflow from intake through QA with product, design, architecture, development, and testing handoffs
---

# App Factory Orchestrator

## Overview

Coordinate App Factory V1 from intake through QA. Default to automatic advancement and pause only on defined risk conditions.

## When to Use

- User wants a lightweight tool app taken from idea to compiled, testable output.
- User wants PM, UD, AM, SD, and QA routed as a governed workflow.

## Workflow

1. Create `products/<project-slug>/docs/00-intake.md`.
2. Route to PM for `products/<project-slug>/docs/01-prd-draft.md`.
3. Route to UD for `products/<project-slug>/docs/02-ux-spec.md` and `products/<project-slug>/design/figma-link.md`.
4. Route back to PM for `products/<project-slug>/docs/03-prd-final.md`.
5. Route to AM for `products/<project-slug>/docs/04-architecture.md`.
6. Route to SD for `products/<project-slug>/docs/05-dev-plan.md` and implementation.
7. Route to QA for `products/<project-slug>/docs/06-test-cases.md` and `products/<project-slug>/docs/07-test-report.md`.
8. Stop only for blocked or rework conditions.

## Product Initialization

Before routing work, initialize the product tree with at least:

- `products/<project-slug>/pubspec.yaml`
- `products/<project-slug>/lib/`
- `products/<project-slug>/test/unit/`
- `products/<project-slug>/test/widget/`
- `products/<project-slug>/config/env/`
- `products/<project-slug>/docs/`
- `products/<project-slug>/design/exports/`
- `products/<project-slug>/build/outputs/`

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

## Initialization Checks

Do not advance beyond intake until:

- the product exists under `products/<project-slug>/`
- all required `docs/` files exist
- `design/figma-link.md` exists
- `design/exports/` exists
- `build/outputs/` exists
- the product does not depend on another product directory

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
