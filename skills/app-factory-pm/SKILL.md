---
name: app-factory-pm
description: Use when a lightweight tool app idea must be shaped into an App Factory draft PRD and then finalized after UX feedback
---

# App Factory PM

## Overview

Turn rough demand into a small, practical tool-app PRD that first guides design and then becomes the final product definition after UX feedback.

## When to Use

- A new tool-style app idea needs scope definition before design starts.
- UD has finished design work and PM must reconcile UX feedback into a final product definition.
- The team needs explicit non-goals, page inventory, and scope boundaries before AM starts architecture.

## Required Inputs

- user request or intake notes from `products/<product-slug>/docs/00-intake.md`
- UX feedback from `products/<product-slug>/docs/02-ux-spec.md` when finalizing
- PM write-back items from UD

## Workflow

1. Read the intake and restate the user problem, target user, and smallest useful scope.
2. Write `products/<product-slug>/docs/01-prd-draft.md` for UX design, focusing on core value, core flow, page list, and non-goals.
3. Keep the draft intentionally narrow. Add monetization reserve only when it does not expand the first release scope.
4. After UD finishes, compare the draft PRD with `products/<product-slug>/docs/02-ux-spec.md` and the PM write-back notes.
5. Update `products/<product-slug>/docs/03-prd-final.md` so it matches the designed flows, states, and constraints.
6. Call out unresolved conflicts instead of smoothing them over.
7. End with a gate result that either clears AM to architect or returns explicit rework items.

## Output Requirements

- Write `products/<product-slug>/docs/01-prd-draft.md` before UX design.
- Update `products/<product-slug>/docs/03-prd-final.md` after UX design.
- Focus on tool-style apps.
- Keep the scope small.
- Include monetization reserve points without expanding scope.
- Define explicit non-goals.
- Absorb UD write-back items before handing work to AM.

## Draft Sections

- `## Product Overview`
- `## Target Users`
- `## User Pain Points`
- `## Core Value`
- `## Core Flow`
- `## Feature List`
- `## Page List`
- `## Interaction Principles`
- `## Monetization Reserve`
- `## Non-Functional Requirements`
- `## Non-Goals`
- `## Gate Result`

## Final PRD Additions

- `## UX Changes From UD`
- `## Final Page Inventory`
- `## Final Interaction Constraints`
- `## AM Handoff Notes`

## Pause Conditions

- The request is not a lightweight tool app and needs a broader product strategy.
- The draft scope has multiple unrelated core flows.
- UD feedback changes the core value proposition instead of refining delivery details.
- A monetization idea adds net-new first-release features.
- The final PRD and UX spec disagree on page inventory or key state behavior.

## Handoff

- Hand off `docs/01-prd-draft.md` to UD only after the scope is small, coherent, and explicit.
- Hand off `docs/03-prd-final.md` to AM only after UX write-back items are absorbed and unresolved conflicts are listed clearly.
