---
name: app-factory-am
description: Use when an App Factory product has a finalized PRD and UX handoff and needs architecture decisions, module boundaries, and monorepo placement
---

# App Factory AM

## Overview

Review the approved product and design inputs from a technical perspective, then produce a modular Flutter architecture that preserves reuse, isolation, and delivery safety.

## When to Use

- PM has produced `products/<product-slug>/docs/03-prd-final.md`.
- UD has produced `products/<product-slug>/docs/02-ux-spec.md` and design references.
- The next step is deciding module boundaries, server scope, and implementation constraints before SD starts coding.

## Required Inputs

- `products/<product-slug>/docs/03-prd-final.md`
- `products/<product-slug>/docs/02-ux-spec.md`
- `products/<product-slug>/design/figma-link.md`
- exported screens under `products/<product-slug>/design/exports/`

## Workflow

1. Read the final PRD and UX artifacts together and list the product's must-ship flows, technical risks, and open assumptions.
2. Decide the platform scope and app shell responsibilities, including navigation, startup, auth state, and upgrade handling.
3. Identify capability modules before feature modules.
4. Assign each responsibility to `packages/`, `products/<product-slug>/client/`, or `products/<product-slug>/client/lib/integrations/`.
5. Decide whether backend logic is unnecessary, reserved for later, or required now under `products/<product-slug>/server/`.
6. Document data flow, state strategy, route strategy, and safety boundaries in `products/<product-slug>/docs/04-architecture.md`.
7. End with a gate result that either clears SD to build or sends concrete revision requests back to PM or UD.

## Monorepo Boundary Rules

- Decide explicitly whether new code belongs in `packages/`, `products/<product-slug>/client/`, or `products/<product-slug>/client/lib/integrations/`.
- Only move code into `packages/` when it is suitable for at least two products.
- Reject abstractions that carry product-specific names, flows, or copy into public packages.
- Protect product isolation. `products/a` must not depend directly on `products/b`.
- Decide whether a product needs `client/` only, `client/ + reserved server boundary`, or `client/ + active server/`.
- Prefer shared account, payment, subscription, and entitlement services over product-specific reimplementation.
- Treat app-version upgrade checks as a shared growth capability, and define whether the product uses optional upgrade prompts, forced upgrade gates, or both.

## Output Requirements

- Write `products/<product-slug>/docs/04-architecture.md`.
- Review the final PRD and UX assets for technical risks.
- Select capability modules before feature modules.
- Define App Shell, capability boundaries, and feature boundaries.
- Push requirement revisions back to PM when reuse or safety is threatened.
- State where each new responsibility belongs in the monorepo.
- State whether server capability is required now, reserved for later, or unnecessary.
- Define launch upgrade behavior, including the forced-upgrade threshold and whether the app should block or exit when too far behind.

## Pause Conditions

- Final PRD and UX spec disagree on core flow or page inventory.
- Design exports are missing for a high-risk interaction or state.
- A shared-package extraction is desired but the second reuse case is not real yet.
- Backend scope is implied but no server responsibility has been stated clearly.
- Repository placement would violate product isolation.

## Required Sections

- `## Architecture Goals`
- `## Platform Scope`
- `## App Shell`
- `## Capability Modules`
- `## Feature Modules`
- `## Monorepo Placement`
- `## Client Server Decision`
- `## Upgrade Strategy`
- `## UX-Driven Technical Risks`
- `## Route Strategy`
- `## State Strategy`
- `## Data Flow`
- `## Safety Boundaries`
- `## Risks`
- `## Gate Result`
