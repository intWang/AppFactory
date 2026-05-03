---
name: app-factory-sd
description: Use when an App Factory architecture is approved and the product must be implemented in Flutter with TDD, repository-boundary discipline, and build verification
---

# App Factory SD

## Overview

Check the development environment, implement the app against the approved architecture, and keep the work test-first.

## When to Use

- AM has approved `products/<product-slug>/docs/04-architecture.md`.
- The next step is building the product under the owning product directory.
- The work must stay TDD-first and produce verifiable build artifacts.

## Required Inputs

- `products/<product-slug>/docs/04-architecture.md`
- `products/<product-slug>/docs/02-ux-spec.md`
- `products/<product-slug>/design/figma-link.md`
- exported screens under `products/<product-slug>/design/exports/`

## Workflow

1. Check Flutter, Dart, iOS, Android, and Web environment support before implementation starts.
2. Translate the approved architecture into a short development plan in `products/<product-slug>/docs/05-dev-plan.md`.
3. Define a failing test before each behavior change. Do not write implementation first.
4. Implement the minimum code needed to satisfy the approved architecture and UX handoff.
5. Keep new client code under `products/<product-slug>/client/` unless the architecture explicitly calls for shared extraction.
6. Create or modify `products/<product-slug>/server/` only when the architecture explicitly requires product-specific backend logic.
7. Record verification commands, build outputs, and boundary-preservation evidence as work progresses.
8. Finish only when the app is buildable, tested, and aligned with the approved UX artifacts.

## Required Process

- Check Flutter, Dart, iOS, Android, and Web environment support first.
- Do not write implementation before defining failing tests.
- Produce a compilable app and automated tests.
- Record build commands and artifact locations.
- Default new client code to `products/<product-slug>/client/`.
- Only create or modify `products/<product-slug>/server/` when the approved architecture explicitly requires product-specific backend logic.
- Modify `packages/` only when the approved architecture explicitly calls for shared extraction.
- Do not introduce direct imports from one product into another product.
- Keep build outputs under the owning product's `build/outputs/`.
- Keep product docs and tests inside the owning product directory.
- Implement against `products/<product-slug>/docs/04-architecture.md`, `products/<product-slug>/docs/02-ux-spec.md`, `products/<product-slug>/design/figma-link.md`, and exported screens.
- Do not invent UI structure from the PRD alone when UD artifacts exist.
- Treat shared account, payment, and entitlement flows as shared-service integrations unless AM has explicitly approved a product-specific server path.
- Implement launch-time upgrade checks in the app shell when architecture requires upgrade governance.
- Enforce the shared version-gap rule so versions more than 3 builds behind trigger a forced-upgrade path instead of normal app entry.

## Required Outputs

- `products/<product-slug>/docs/05-dev-plan.md`
- buildable Flutter project changes
- automated tests
- verification commands
- build artifact locations
- evidence that repository boundaries were preserved

## Pause Conditions

- The environment is missing a required platform toolchain.
- The architecture leaves a major flow or state unspecified.
- The UX artifacts conflict with the architecture on navigation or state behavior.
- A proposed shared extraction has only one real consumer.
- A requested change would require another product to be imported directly.
