---
name: app-factory-orchestrator
description: Use when a user wants to turn a lightweight tool app idea into a governed App Factory workflow from intake through QA
---

# App Factory Orchestrator

## Overview

Coordinate App Factory V1 from intake through QA. Default to automatic advancement and pause only on defined risk conditions.

## When to Use

- User wants a lightweight tool app taken from idea to compiled, testable output.
- User wants PM, AM, SD, and QA routed as a governed workflow.

## Workflow

1. Create `docs/projects/<project-slug>/00-intake.md`.
2. Route to PM for `01-prd.md`.
3. Route to AM for `02-architecture.md`.
4. Route to SD for `03-dev-plan.md` and implementation.
5. Route to QA for `04-test-cases.md` and `05-test-report.md`.
6. Stop only for blocked or rework conditions.

## Intake Artifact

Write `00-intake.md` with:

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
