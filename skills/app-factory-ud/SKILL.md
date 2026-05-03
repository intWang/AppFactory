---
name: app-factory-ud
description: Use when an App Factory product has a drafted PRD and needs UX structure, Figma references, and design handoff artifacts before architecture and implementation
---

# App Factory UD

## Overview

Turn a drafted PRD into a stable interaction-design handoff that PM, AM, SD, and QA can all use consistently.

## When to Use

- PM has produced `products/<product-slug>/docs/01-prd-draft.md`.
- The team needs user flows, page inventory, and key states before architecture and implementation.
- PM is expecting UX write-back items to refine the final PRD.

## Required Inputs

- `products/<product-slug>/docs/01-prd-draft.md`
- any intake assumptions that still affect flow or scope

## Workflow

1. Read the drafted PRD and confirm the product's core flow, target user, and explicit non-goals.
2. Define user flows, page structure, key states, and high-fidelity interaction direction.
3. Produce a Figma primary link and repository-local design references.
4. Export the screens or states that SD and QA will need for implementation and verification.
5. Record design decisions that PM must write back into the final PRD.
6. Highlight reusable UI patterns that AM may need to treat as shared components.
7. Finish with a gate result that either clears PM finalization and AM review or lists blocking design gaps.

## Required Process

- Read `products/<product-slug>/docs/01-prd-draft.md` before starting design work.
- Define user flows, page structure, key states, and high-fidelity direction.
- Produce a Figma primary link and repository-local design references.
- Record design decisions that PM must write back into the final PRD.
- Highlight reusable UI patterns that AM may need to treat as shared components.
- Give SD and QA a stable design baseline through links, exported screens, and UX notes.

## Required Outputs

- `products/<product-slug>/design/figma-link.md`
- `products/<product-slug>/docs/02-ux-spec.md`
- `products/<product-slug>/design/exports/`
- PM write-back notes for `products/<product-slug>/docs/03-prd-final.md`

## Required Sections

- `## User Flows`
- `## Page Inventory`
- `## Key States`
- `## Interaction Notes`
- `## Motion Notes`
- `## Component Guidance`
- `## PM Write-back Items`
- `## Gate Result`

## Pause Conditions

- The draft PRD has multiple conflicting core flows.
- A key state or edge case cannot be inferred from the draft PRD.
- A required design export or primary Figma reference cannot be produced.
- The design direction would materially change the product scope instead of clarifying it.
