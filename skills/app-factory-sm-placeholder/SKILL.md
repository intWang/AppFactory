---
name: app-factory-sm-placeholder
description: Use when an App Factory product needs its future agile-management and retrospective boundary documented without running a full process-improvement cycle yet
---

# App Factory SM Placeholder

## Overview

This is a placeholder skill for recording future scrum or retrospective responsibilities. It captures the boundary now so later process work has a clear landing zone.

## When to Use

- The team wants to note retrospective or process-improvement responsibilities for later.
- A project needs a place to capture factory-memory candidates without running a full agile ceremony.
- Another role needs to document improvement inputs without owning long-term process management.

## Workflow

1. Record retrospective or process-improvement inputs in the owning product docs, usually `products/<product-slug>/docs/09-retro-input.md`.
2. Separate immediate delivery blockers from longer-term process improvements.
3. Identify items that may deserve promotion into shared factory memory later.
4. End by clarifying that the boundary is documented, but no full retrospective facilitation has been run yet.

## Required Outputs

- update or create `products/<product-slug>/docs/09-retro-input.md`
- list of delivery blockers versus longer-term improvement ideas
- candidate items for shared factory memory
- gate note stating that no full retrospective was facilitated

## Required Sections

- `## Immediate Delivery Issues`
- `## Process Improvement Ideas`
- `## Memory Candidates`
- `## Recommended Next Owner`
- `## Gate Result`

## Reserved Scope

- retrospective facilitation
- process improvement capture
- project memory escalation to factory memory

## Out Of Scope Now

- running a formal retrospective workshop
- rewriting shared factory process immediately
- claiming that memory updates were already completed

## Pause Conditions

- A user expects a completed retrospective rather than a placeholder boundary note.
- The improvement item requires immediate cross-product process changes instead of documentation only.

## Handoff

- Hand off urgent blockers to the active delivery owner.
- Hand off shared process opportunities to the future SM or factory-memory owner.
