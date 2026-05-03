# UX Spec

## Project Name
Example Tool App

## Project Slug
example-tool-app

## Input Source
Derived from `01-prd-draft.md`.

## User Flows
- Launch app
- Land on home screen
- See primary tool entry

## Page Inventory
- Home
- Loading
- Empty
- Error

## Key States
- Loading while app services initialize
- Empty state when no user-created content exists
- Error state for failed startup or missing capability registration
- Success state when the shell loads normally

## Interaction Notes
- Startup should feel immediate and uncluttered.
- Primary action should remain visible without nested navigation.

## Motion Notes
- Keep motion minimal in V1.
- Prefer short fade or scale transitions over complex choreography.

## Component Guidance
- Reuse shared theme primitives from `app_factory_ui`.
- Favor product-level composition before extracting new public components.

## PM Write-back Items
- Final PRD should explicitly include loading, empty, and error states.
- Final PRD should lock the home screen as the only V1 page.

## Gate Result
Pass
