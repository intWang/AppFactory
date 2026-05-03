---
name: app-factory-am
description: Use when an approved App Factory PRD needs technical review, capability selection, and a modular architecture design
---

# App Factory AM

## Overview

Review the PRD from a technical perspective and convert it into a modular Flutter architecture that preserves reuse and safety boundaries.

## Monorepo Boundary Rules

- Decide explicitly whether new code belongs in `packages/`, `products/<product-slug>/`, or `products/<product-slug>/lib/integrations/`.
- Only move code into `packages/` when it is suitable for at least two products.
- Reject abstractions that carry product-specific names, flows, or copy into public packages.
- Protect product isolation. `products/a` must not depend directly on `products/b`.

## Output Requirements

- Write `02-architecture.md`.
- Review the PRD for technical risks.
- Select capability modules before feature modules.
- Define App Shell, capability boundaries, and feature boundaries.
- Push requirement revisions back to PM when reuse or safety is threatened.
- State where each new responsibility belongs in the monorepo.

## Required Sections

- `## Architecture Goals`
- `## Platform Scope`
- `## App Shell`
- `## Capability Modules`
- `## Feature Modules`
- `## Monorepo Placement`
- `## Route Strategy`
- `## State Strategy`
- `## Data Flow`
- `## Safety Boundaries`
- `## Risks`
- `## Gate Result`
