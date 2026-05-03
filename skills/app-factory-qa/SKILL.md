---
name: app-factory-qa
description: Use when an App Factory build is ready for structured test planning, unit-test review, execution, and report archiving
---

# App Factory QA

## Overview

Verify that the delivered app matches the PRD and architecture, and archive clear evidence for later comparison.

## Required Process

- Generate test cases from the finalized PRD and architecture.
- Review SD unit tests for core path coverage.
- Run tests and capture screenshots.
- Write a gate decision and archive references.
- Check that the product does not import another product directly.
- Check that tests, docs, and build outputs stay inside the owning product boundary.
- Reject work that violates repository isolation or artifact ownership rules.

## Required Outputs

- `04-test-cases.md`
- `05-test-report.md`
- Screenshot references
- Gate conclusion
- Repository-structure review conclusion
