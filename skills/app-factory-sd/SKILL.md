---
name: app-factory-sd
description: Use when an App Factory architecture has been approved and a Flutter app must be implemented with TDD and build verification
---

# App Factory SD

## Overview

Check the development environment, implement the app against the approved architecture, and keep the work test-first.

## Required Process

- Check Flutter, Dart, iOS, Android, and Web environment support first.
- Do not write implementation before defining failing tests.
- Produce a compilable app and automated tests.
- Record build commands and artifact locations.
- Default new business code to `products/<product-slug>/`.
- Modify `packages/` only when the approved architecture explicitly calls for shared extraction.
- Do not introduce direct imports from one product into another product.
- Keep build outputs under the owning product's `build/outputs/`.
- Keep product docs and tests inside the owning product directory.
- Implement against `products/<product-slug>/docs/04-architecture.md`, `products/<product-slug>/docs/02-ux-spec.md`, `products/<product-slug>/design/figma-link.md`, and exported screens.
- Do not invent UI structure from the PRD alone when UD artifacts exist.

## Required Outputs

- `products/<product-slug>/docs/05-dev-plan.md`
- Buildable Flutter project changes
- Automated tests
- Verification commands
- Build artifact locations
- Evidence that repository boundaries were preserved
