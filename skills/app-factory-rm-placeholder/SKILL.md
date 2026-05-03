---
name: app-factory-rm-placeholder
description: Use when an App Factory product needs its future release-management boundary documented without performing real store release work yet
---

# App Factory RM Placeholder

## Overview

This is a placeholder skill for documenting release-management scope now so later store-delivery work has a clear boundary. It does not perform actual release, signing, or money-sensitive operations.

## When to Use

- The product needs a documented release-management boundary in current planning artifacts.
- Store delivery, signing, or cost review is known to matter later but is intentionally out of scope for the current build.
- Another role needs to record release risks without taking release actions.

## Workflow

1. Record the reserved release-management scope in the owning product docs, usually `products/<product-slug>/docs/08-release-gate.md`.
2. Note which items are future release work versus current engineering or QA work.
3. Flag any request that would require real signing, account access, payment, or store submission as a pause condition.
4. End by stating that the release-management boundary is documented but not yet executed.

## Required Outputs

- update or create `products/<product-slug>/docs/08-release-gate.md`
- explicit list of future release responsibilities
- explicit list of out-of-scope actions for the current phase
- gate note stating that no real release action was performed

## Required Sections

- `## Release Scope Boundary`
- `## Future Release Responsibilities`
- `## Current Non-Release Scope`
- `## Risks`
- `## Gate Result`

## Reserved Scope

- App Store and Play Store release details
- release cost review
- signing, certificate, account, and money-related safety checks
- store compliance review

## Out Of Scope Now

- generating signing assets
- uploading binaries
- touching billing accounts
- approving money-sensitive release actions

## Pause Conditions

- A task requires real credentials, certificates, payment methods, or store console access.
- A user expects this placeholder to execute release work instead of documenting the boundary.

## Handoff

- Hand off documented release risks to QA or the future RM owner.
- Do not hand off a fake "release complete" signal to any downstream role.
