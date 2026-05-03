---
name: app-factory-qa
description: Use when an App Factory build is ready for structured test planning, execution, evidence capture, and release-gate reporting
---

# App Factory QA

## Overview

Verify that the delivered app matches the PRD, UX spec, and architecture, then archive clear evidence and a gate decision for later comparison.

## When to Use

- SD has produced an implementation and build artifacts.
- The team needs test cases, execution evidence, and a pass or rework decision.
- Repository-boundary, upgrade-path, and design-consistency checks must be included in the gate.

## Required Inputs

- `products/<product-slug>/docs/03-prd-final.md`
- `products/<product-slug>/docs/02-ux-spec.md`
- `products/<product-slug>/docs/04-architecture.md`
- implementation outputs under `products/<product-slug>/client/`
- exported design screens and `products/<product-slug>/design/figma-link.md`

## Workflow

1. Generate test cases from the finalized PRD, UX spec, and architecture.
2. Review SD unit and widget tests for core-path coverage and obvious gaps.
3. Run automated checks and capture the exact verification commands and outcomes.
4. Execute focused manual checks for primary flows, critical states, and visible UX fidelity.
5. Capture screenshots and record where the evidence lives.
6. Check that the product does not import another product directly and that docs, tests, and build outputs stay inside the owning product boundary.
7. Verify launch upgrade checks, including optional upgrade prompts and forced-upgrade blocking when the latest build is more than 3 builds ahead.
8. Write a gate decision in `products/<product-slug>/docs/07-test-report.md` that clearly says pass, blocked, or needs rework.

## Required Outputs

- `products/<product-slug>/docs/06-test-cases.md`
- `products/<product-slug>/docs/07-test-report.md`
- screenshot references
- gate conclusion
- repository-structure review conclusion

## Required Report Coverage

- automated test summary
- manual test summary
- design-consistency summary
- version-upgrade behavior summary
- repository-boundary summary
- known issues
- gate result

## Pause Conditions

- A required build, test, or screenshot artifact cannot be accessed.
- The build is not runnable enough to complete the planned QA pass.
- The approved PRD, UX spec, and architecture disagree on the expected behavior.
- A critical blocker is found that must be fixed before the remaining checks are meaningful.

## Fail Conditions

- Core flow behavior conflicts with the finalized PRD or UX spec.
- Required automated checks were skipped without a written blocker.
- Shared version-gap rules are missing or behave incorrectly.
- Repository isolation is broken.
- Evidence is missing for a claimed pass result.
