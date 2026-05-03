# App Factory V1 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build the first working version of the App Factory skill family for Flutter-based tool apps, including the orchestrator, core role skills, shared templates, a reusable Flutter base app, and one end-to-end example project path.

**Architecture:** The implementation is split into two tracks that meet in a single example flow: a skill-and-docs backbone and a Flutter platform backbone. The skills create and govern project artifacts, while the Flutter template provides the reusable shell, capability registration pattern, and test baseline those skills target.

**Tech Stack:** Markdown skills, Markdown templates, Flutter, Dart, `flutter_test`, platform toolchains for iOS/Android/Web, shell utilities

---

## File Structure

### New files and directories

- Create: `skills/app-factory-orchestrator/SKILL.md`
- Create: `skills/app-factory-pm/SKILL.md`
- Create: `skills/app-factory-am/SKILL.md`
- Create: `skills/app-factory-sd/SKILL.md`
- Create: `skills/app-factory-qa/SKILL.md`
- Create: `skills/app-factory-rm-placeholder/SKILL.md`
- Create: `skills/app-factory-sm-placeholder/SKILL.md`
- Create: `templates/docs/prd.template.md`
- Create: `templates/docs/add.template.md`
- Create: `templates/docs/test-cases.template.md`
- Create: `templates/docs/test-report.template.md`
- Create: `templates/docs/memory.template.md`
- Create: `templates/docs/intake.template.md`
- Create: `templates/docs/dev-plan.template.md`
- Create: `templates/docs/release-gate.template.md`
- Create: `templates/docs/retro-input.template.md`
- Create: `templates/flutter_tool_app/README.md`
- Create: `templates/flutter_tool_app/pubspec.yaml`
- Create: `templates/flutter_tool_app/lib/main.dart`
- Create: `templates/flutter_tool_app/lib/app/app.dart`
- Create: `templates/flutter_tool_app/lib/app/bootstrap.dart`
- Create: `templates/flutter_tool_app/lib/app/routes.dart`
- Create: `templates/flutter_tool_app/lib/core/foundation/logging/logger.dart`
- Create: `templates/flutter_tool_app/lib/core/foundation/config/app_config.dart`
- Create: `templates/flutter_tool_app/lib/core/foundation/error/error_reporter.dart`
- Create: `templates/flutter_tool_app/lib/core/capabilities/capability.dart`
- Create: `templates/flutter_tool_app/lib/core/capabilities/capability_registry.dart`
- Create: `templates/flutter_tool_app/lib/core/growth/growth_entry_points.dart`
- Create: `templates/flutter_tool_app/lib/core/ui/app_theme.dart`
- Create: `templates/flutter_tool_app/lib/features/home/home_page.dart`
- Create: `templates/flutter_tool_app/test/app_bootstrap_test.dart`
- Create: `templates/flutter_tool_app/test/capability_registry_test.dart`
- Create: `templates/flutter_tool_app/test/widget_smoke_test.dart`
- Create: `docs/projects/example-tool-app/00-intake.md`
- Create: `docs/projects/example-tool-app/01-prd.md`
- Create: `docs/projects/example-tool-app/02-architecture.md`
- Create: `docs/projects/example-tool-app/03-dev-plan.md`
- Create: `docs/projects/example-tool-app/04-test-cases.md`
- Create: `docs/projects/example-tool-app/05-test-report.md`
- Create: `docs/projects/example-tool-app/06-release-gate.md`
- Create: `docs/projects/example-tool-app/07-retro-input.md`
- Create: `memory/factory-memory.md`
- Create: `memory/project-memory/example-tool-app.md`
- Create: `archives/test-reports/example-tool-app/.gitkeep`
- Create: `archives/build-artifacts/example-tool-app/.gitkeep`

### Existing reference files

- Modify: `docs/superpowers/specs/2026-05-02-app-factory-v1-design.md`

## Task 1: Create the Skill Skeleton and Shared Directory Layout

**Files:**
- Create: `skills/app-factory-orchestrator/SKILL.md`
- Create: `skills/app-factory-pm/SKILL.md`
- Create: `skills/app-factory-am/SKILL.md`
- Create: `skills/app-factory-sd/SKILL.md`
- Create: `skills/app-factory-qa/SKILL.md`
- Create: `skills/app-factory-rm-placeholder/SKILL.md`
- Create: `skills/app-factory-sm-placeholder/SKILL.md`
- Create: `memory/factory-memory.md`
- Create: `memory/project-memory/example-tool-app.md`
- Create: `archives/test-reports/example-tool-app/.gitkeep`
- Create: `archives/build-artifacts/example-tool-app/.gitkeep`

- [ ] **Step 1: Write the failing structure check**

Create a small shell verification file or checklist in your working notes that expects all of the directories and skill files above to exist. The check content should be:

```text
skills/app-factory-orchestrator/SKILL.md
skills/app-factory-pm/SKILL.md
skills/app-factory-am/SKILL.md
skills/app-factory-sd/SKILL.md
skills/app-factory-qa/SKILL.md
skills/app-factory-rm-placeholder/SKILL.md
skills/app-factory-sm-placeholder/SKILL.md
memory/factory-memory.md
memory/project-memory/example-tool-app.md
archives/test-reports/example-tool-app/.gitkeep
archives/build-artifacts/example-tool-app/.gitkeep
```

- [ ] **Step 2: Run the structure check and verify it fails**

Run:

```bash
find skills memory archives -maxdepth 4 -type f 2>/dev/null | sort
```

Expected: the command output is empty or missing most of the target files because nothing has been created yet.

- [ ] **Step 3: Create the minimal directory tree and placeholder files**

Use `mkdir -p` for directories, then create minimal markdown files with the exact headings below:

```markdown
---
name: app-factory-orchestrator
description: Use when a user wants to turn a lightweight app idea into a governed App Factory workflow from intake through QA
---

# App Factory Orchestrator
```

Use equivalent two-line skill skeletons for each role skill, changing the `name`, `description`, and heading to match the file name and role.

For memory files, use:

```markdown
# Factory Memory
```

and:

```markdown
# Example Tool App Memory
```

- [ ] **Step 4: Run the structure check and verify it passes**

Run:

```bash
find skills memory archives -maxdepth 4 -type f | sort
```

Expected: every file in Step 1 appears in the output.

- [ ] **Step 5: Commit**

```bash
git add skills memory archives
git commit -m "chore: scaffold app factory skill family"
```

## Task 2: Author the Orchestrator Skill

**Files:**
- Modify: `skills/app-factory-orchestrator/SKILL.md`
- Test: `docs/projects/example-tool-app/00-intake.md`

- [ ] **Step 1: Write the failing artifact contract**

Write the expected intake artifact content in your task notes so the skill has a concrete target:

```markdown
# Intake

## Project Name

## Project Slug

## User Requirement

## Default Assumptions

## Scope Decision

## Current State

## Next Role

## Risks

## Unknowns
```

- [ ] **Step 2: Verify the skill cannot produce the artifact yet**

Run:

```bash
sed -n '1,120p' skills/app-factory-orchestrator/SKILL.md
```

Expected: only the skeleton exists, so there is no state-machine, gate, or intake artifact guidance.

- [ ] **Step 3: Write the minimal orchestrator skill**

Expand the skill with these sections and concrete content:

```markdown
## Overview
Coordinate App Factory V1 from intake through QA. Default to automatic advancement and pause only on defined risk conditions.

## When to Use
- User wants a lightweight tool app taken from idea to compiled, testable output
- User wants PM, AM, SD, and QA routed as a governed workflow

## Workflow
1. Create `docs/projects/<project-slug>/00-intake.md`
2. Route to PM for `01-prd.md`
3. Route to AM for `02-architecture.md`
4. Route to SD for `03-dev-plan.md` and implementation
5. Route to QA for `04-test-cases.md` and `05-test-report.md`
6. Stop only for blocked or rework conditions

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
```

- [ ] **Step 4: Run the artifact contract check**

Run:

```bash
grep -n "## Current State" skills/app-factory-orchestrator/SKILL.md
```

Expected: the skill now contains enough guidance to create the intake artifact with a current-state field.

- [ ] **Step 5: Commit**

```bash
git add skills/app-factory-orchestrator/SKILL.md
git commit -m "feat: define app factory orchestrator skill"
```

## Task 3: Author the PM and AM Skills

**Files:**
- Modify: `skills/app-factory-pm/SKILL.md`
- Modify: `skills/app-factory-am/SKILL.md`
- Test: `docs/projects/example-tool-app/01-prd.md`
- Test: `docs/projects/example-tool-app/02-architecture.md`

- [ ] **Step 1: Write the failing PRD and architecture contracts**

Use these expected sections:

```markdown
# PRD

## Product Overview

## Target Users

## User Pain Points

## Core Flow

## Feature List

## Monetization Reserve

## Non-Goals
```

and:

```markdown
# Architecture Design

## Architecture Goals

## App Shell

## Capability Modules

## Feature Modules

## Route Strategy

## State Strategy

## Safety Boundaries

## Risks
```

- [ ] **Step 2: Verify the current skills are insufficient**

Run:

```bash
sed -n '1,120p' skills/app-factory-pm/SKILL.md
sed -n '1,120p' skills/app-factory-am/SKILL.md
```

Expected: both files still contain only minimal placeholders and do not define the output structure.

- [ ] **Step 3: Write the minimal PM and AM skill bodies**

PM must include:

```markdown
## Output Requirements
- Write `01-prd.md`
- Focus on tool-style apps
- Keep the scope small
- Include monetization reserve points without expanding scope
- Define explicit non-goals
```

AM must include:

```markdown
## Output Requirements
- Write `02-architecture.md`
- Review the PRD for technical risks
- Select capability modules before feature modules
- Define App Shell, capability boundaries, and feature boundaries
- Push requirement revisions back to PM when reuse or safety is threatened
```

- [ ] **Step 4: Run the contract checks**

Run:

```bash
grep -n "Monetization Reserve" skills/app-factory-pm/SKILL.md
grep -n "Capability Modules" skills/app-factory-am/SKILL.md
```

Expected: both terms are present in the right skill files.

- [ ] **Step 5: Commit**

```bash
git add skills/app-factory-pm/SKILL.md skills/app-factory-am/SKILL.md
git commit -m "feat: add pm and am role skills"
```

## Task 4: Author the SD and QA Skills

**Files:**
- Modify: `skills/app-factory-sd/SKILL.md`
- Modify: `skills/app-factory-qa/SKILL.md`
- Test: `docs/projects/example-tool-app/03-dev-plan.md`
- Test: `docs/projects/example-tool-app/04-test-cases.md`
- Test: `docs/projects/example-tool-app/05-test-report.md`

- [ ] **Step 1: Write the failing delivery contracts**

Use these target sections:

```markdown
# Development Plan

## Environment Check

## TDD Strategy

## Build Commands

## Artifact Output
```

and:

```markdown
# Test Report

## Test Environment

## Unit Test Review

## Results Summary

## Screenshot Index

## Gate Decision
```

- [ ] **Step 2: Verify the current skills are insufficient**

Run:

```bash
sed -n '1,120p' skills/app-factory-sd/SKILL.md
sed -n '1,120p' skills/app-factory-qa/SKILL.md
```

Expected: both files are still bare and lack TDD or report requirements.

- [ ] **Step 3: Write the minimal SD and QA skill bodies**

SD must include:

```markdown
## Required Process
- Check Flutter, Dart, iOS, Android, and Web environment support first
- Do not write implementation before defining failing tests
- Produce a compilable app and automated tests
- Record build commands and artifact locations
```

QA must include:

```markdown
## Required Process
- Generate test cases from the finalized PRD and architecture
- Review SD unit tests for core path coverage
- Run tests and capture screenshots
- Write a gate decision and archive references
```

- [ ] **Step 4: Run the contract checks**

Run:

```bash
grep -n "failing tests" skills/app-factory-sd/SKILL.md
grep -n "capture screenshots" skills/app-factory-qa/SKILL.md
```

Expected: both requirements are present.

- [ ] **Step 5: Commit**

```bash
git add skills/app-factory-sd/SKILL.md skills/app-factory-qa/SKILL.md
git commit -m "feat: add sd and qa role skills"
```

## Task 5: Add RM and SM Placeholder Skills

**Files:**
- Modify: `skills/app-factory-rm-placeholder/SKILL.md`
- Modify: `skills/app-factory-sm-placeholder/SKILL.md`

- [ ] **Step 1: Write the failing placeholder contracts**

Define the future responsibilities in notes:

```markdown
## RM Placeholder
- release assets
- cost control
- signing and credential safety
- store compliance
```

and:

```markdown
## SM Placeholder
- retro inputs
- process observations
- memory extraction
```

- [ ] **Step 2: Verify the current files do not describe the placeholders**

Run:

```bash
sed -n '1,80p' skills/app-factory-rm-placeholder/SKILL.md
sed -n '1,80p' skills/app-factory-sm-placeholder/SKILL.md
```

Expected: both files are still too minimal to clarify future boundaries.

- [ ] **Step 3: Write the minimal placeholder skills**

RM must include:

```markdown
## Reserved Scope
- app store and play store release details
- release cost review
- signing, certificate, account, and money-related safety checks
```

SM must include:

```markdown
## Reserved Scope
- retrospective facilitation
- process improvement capture
- project memory escalation to factory memory
```

- [ ] **Step 4: Run the contract checks**

Run:

```bash
grep -n "money-related safety checks" skills/app-factory-rm-placeholder/SKILL.md
grep -n "factory memory" skills/app-factory-sm-placeholder/SKILL.md
```

Expected: both placeholder boundaries are explicit.

- [ ] **Step 5: Commit**

```bash
git add skills/app-factory-rm-placeholder/SKILL.md skills/app-factory-sm-placeholder/SKILL.md
git commit -m "docs: define rm and sm placeholder skills"
```

## Task 6: Create the Documentation Templates

**Files:**
- Create: `templates/docs/intake.template.md`
- Create: `templates/docs/prd.template.md`
- Create: `templates/docs/add.template.md`
- Create: `templates/docs/dev-plan.template.md`
- Create: `templates/docs/test-cases.template.md`
- Create: `templates/docs/test-report.template.md`
- Create: `templates/docs/memory.template.md`
- Create: `templates/docs/release-gate.template.md`
- Create: `templates/docs/retro-input.template.md`

- [ ] **Step 1: Write the failing template contract**

Define the required filenames in notes:

```text
intake.template.md
prd.template.md
add.template.md
dev-plan.template.md
test-cases.template.md
test-report.template.md
memory.template.md
release-gate.template.md
retro-input.template.md
```

- [ ] **Step 2: Verify the templates do not exist yet**

Run:

```bash
find templates/docs -maxdepth 1 -type f 2>/dev/null | sort
```

Expected: no template files yet.

- [ ] **Step 3: Create the templates with exact headings**

Example intake template:

```markdown
# Intake

## Project Name

## Project Slug

## Input Source

## User Requirement

## Default Assumptions

## Scope Decision

## Current State

## Next Role

## Risks

## Unknowns

## Gate Result
```

Follow the same pattern for the other templates using the sections defined in the spec.

- [ ] **Step 4: Run the template check**

Run:

```bash
grep -n "## Gate Result" templates/docs/intake.template.md
grep -n "## Monetization Reserve" templates/docs/prd.template.md
grep -n "## Safety Boundaries" templates/docs/add.template.md
grep -n "## Screenshot Index" templates/docs/test-report.template.md
```

Expected: all required headings are present.

- [ ] **Step 5: Commit**

```bash
git add templates/docs
git commit -m "feat: add app factory document templates"
```

## Task 7: Build the Flutter Base Template Through TDD

**Files:**
- Create: `templates/flutter_tool_app/pubspec.yaml`
- Create: `templates/flutter_tool_app/lib/main.dart`
- Create: `templates/flutter_tool_app/lib/app/app.dart`
- Create: `templates/flutter_tool_app/lib/app/bootstrap.dart`
- Create: `templates/flutter_tool_app/lib/app/routes.dart`
- Create: `templates/flutter_tool_app/lib/core/foundation/logging/logger.dart`
- Create: `templates/flutter_tool_app/lib/core/foundation/config/app_config.dart`
- Create: `templates/flutter_tool_app/lib/core/foundation/error/error_reporter.dart`
- Create: `templates/flutter_tool_app/lib/core/capabilities/capability.dart`
- Create: `templates/flutter_tool_app/lib/core/capabilities/capability_registry.dart`
- Create: `templates/flutter_tool_app/lib/core/growth/growth_entry_points.dart`
- Create: `templates/flutter_tool_app/lib/core/ui/app_theme.dart`
- Create: `templates/flutter_tool_app/lib/features/home/home_page.dart`
- Create: `templates/flutter_tool_app/test/app_bootstrap_test.dart`
- Create: `templates/flutter_tool_app/test/capability_registry_test.dart`
- Create: `templates/flutter_tool_app/test/widget_smoke_test.dart`
- Create: `templates/flutter_tool_app/README.md`

- [ ] **Step 1: Write the failing tests first**

Create `test/capability_registry_test.dart` with this minimal expectation:

```dart
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_tool_app/core/capabilities/capability_registry.dart';

void main() {
  test('registers and resolves capabilities by type', () {
    final registry = CapabilityRegistry();

    registry.register<String>('logger');

    expect(registry.resolve<String>(), 'logger');
  });
}
```

Create `test/app_bootstrap_test.dart` with:

```dart
import 'package:flutter_tool_app/app/bootstrap.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  test('bootstrap returns app services', () async {
    final services = await bootstrapApp();

    expect(services.registry, isNotNull);
    expect(services.config.appName, isNotEmpty);
  });
}
```

Create `test/widget_smoke_test.dart` with:

```dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_tool_app/app/app.dart';
import 'package:flutter_tool_app/app/bootstrap.dart';

void main() {
  testWidgets('renders tool app shell', (tester) async {
    final services = await bootstrapApp();

    await tester.pumpWidget(ToolApp(services: services));

    expect(find.byType(MaterialApp), findsOneWidget);
    expect(find.text('Tool App Home'), findsOneWidget);
  });
}
```

- [ ] **Step 2: Run the tests to verify they fail**

Run:

```bash
cd templates/flutter_tool_app && flutter test
```

Expected: fail because the package, bootstrap, app widget, and registry do not exist yet.

- [ ] **Step 3: Write the minimal implementation**

Create the package with:

```yaml
name: flutter_tool_app
description: App Factory V1 Flutter tool app template
publish_to: none
environment:
  sdk: ">=3.0.0 <4.0.0"
dependencies:
  flutter:
    sdk: flutter
dev_dependencies:
  flutter_test:
    sdk: flutter
flutter:
  uses-material-design: true
```

Create a registry that supports simple type-based registration and resolution, an app config with a non-empty app name, a `bootstrapApp()` function that returns services containing the registry and config, and a `ToolApp` widget that renders a `MaterialApp` with a home page labeled `Tool App Home`.

- [ ] **Step 4: Run the tests to verify they pass**

Run:

```bash
cd templates/flutter_tool_app && flutter test
```

Expected: all three tests pass.

- [ ] **Step 5: Commit**

```bash
git add templates/flutter_tool_app
git commit -m "feat: add flutter app factory base template"
```

## Task 8: Create Example Project Artifacts from the Templates

**Files:**
- Create: `docs/projects/example-tool-app/00-intake.md`
- Create: `docs/projects/example-tool-app/01-prd.md`
- Create: `docs/projects/example-tool-app/02-architecture.md`
- Create: `docs/projects/example-tool-app/03-dev-plan.md`
- Create: `docs/projects/example-tool-app/04-test-cases.md`
- Create: `docs/projects/example-tool-app/05-test-report.md`
- Create: `docs/projects/example-tool-app/06-release-gate.md`
- Create: `docs/projects/example-tool-app/07-retro-input.md`

- [ ] **Step 1: Write the failing example-project contract**

Define the expected files in notes:

```text
00-intake.md
01-prd.md
02-architecture.md
03-dev-plan.md
04-test-cases.md
05-test-report.md
06-release-gate.md
07-retro-input.md
```

- [ ] **Step 2: Verify the example project does not exist yet**

Run:

```bash
find docs/projects/example-tool-app -maxdepth 1 -type f 2>/dev/null | sort
```

Expected: no files yet.

- [ ] **Step 3: Create the example project from templates**

Use `example-tool-app` as the slug and populate each document with at least:

```markdown
## Project Name
Example Tool App

## Project Slug
example-tool-app
```

For `05-test-report.md`, include:

```markdown
## Gate Decision
Pending execution
```

- [ ] **Step 4: Run the contract checks**

Run:

```bash
grep -n "example-tool-app" docs/projects/example-tool-app/00-intake.md
grep -n "Pending execution" docs/projects/example-tool-app/05-test-report.md
```

Expected: the example slug and pending QA gate text are present.

- [ ] **Step 5: Commit**

```bash
git add docs/projects/example-tool-app
git commit -m "docs: add example tool app project artifacts"
```

## Task 9: Align the Spec and Implementation Notes

**Files:**
- Modify: `docs/superpowers/specs/2026-05-02-app-factory-v1-design.md`

- [ ] **Step 1: Write the failing alignment checklist**

Check that the spec reflects:

```text
skills/
templates/docs/
templates/flutter_tool_app/
docs/projects/example-tool-app/
memory/
archives/
```

- [ ] **Step 2: Verify the current spec lacks implementation-level precision**

Run:

```bash
grep -n "example-tool-app" docs/superpowers/specs/2026-05-02-app-factory-v1-design.md
```

Expected: no mention of the example project path yet.

- [ ] **Step 3: Add the minimal alignment note**

Append a short section such as:

```markdown
## 25. First Implementation Slice

The first implementation slice should include:

- the seven skill directories
- shared document templates
- a reusable Flutter tool-app template
- one example project at `docs/projects/example-tool-app/`
```

- [ ] **Step 4: Run the alignment check**

Run:

```bash
grep -n "First Implementation Slice" docs/superpowers/specs/2026-05-02-app-factory-v1-design.md
grep -n "docs/projects/example-tool-app/" docs/superpowers/specs/2026-05-02-app-factory-v1-design.md
```

Expected: both lines are now present.

- [ ] **Step 5: Commit**

```bash
git add docs/superpowers/specs/2026-05-02-app-factory-v1-design.md
git commit -m "docs: align spec with first implementation slice"
```

## Task 10: Verify the End-to-End Baseline

**Files:**
- Modify: `docs/projects/example-tool-app/05-test-report.md`
- Modify: `docs/projects/example-tool-app/03-dev-plan.md`

- [ ] **Step 1: Write the failing verification checklist**

Define the end-to-end proof points:

```text
All skills exist
All templates exist
Flutter tests pass
Example project docs exist
Test report references the Flutter baseline
```

- [ ] **Step 2: Run the baseline verification before final edits**

Run:

```bash
find skills templates docs/projects/example-tool-app memory archives -maxdepth 4 -type f | sort
cd templates/flutter_tool_app && flutter test
```

Expected: if any file or test is still missing, this step fails and identifies the gap.

- [ ] **Step 3: Update the dev plan and test report with actual verification details**

Add to `03-dev-plan.md`:

```markdown
## Verification Command
cd templates/flutter_tool_app && flutter test
```

Add to `05-test-report.md`:

```markdown
## Verification Baseline
Flutter template tests must pass before QA marks the project ready.
```

- [ ] **Step 4: Re-run the baseline verification**

Run:

```bash
grep -n "flutter test" docs/projects/example-tool-app/03-dev-plan.md
grep -n "Verification Baseline" docs/projects/example-tool-app/05-test-report.md
cd templates/flutter_tool_app && flutter test
```

Expected: document checks pass and Flutter tests remain green.

- [ ] **Step 5: Commit**

```bash
git add docs/projects/example-tool-app/03-dev-plan.md docs/projects/example-tool-app/05-test-report.md
git commit -m "test: verify app factory v1 baseline"
```

## Self-Review

- [ ] **Spec coverage:** Confirm every major section of `docs/superpowers/specs/2026-05-02-app-factory-v1-design.md` maps to at least one task above.
- [ ] **Placeholder scan:** Search the plan for `TBD`, `TODO`, `implement later`, `similar to`, and vague wording. Replace any weak wording with explicit file content or commands.
- [ ] **Type consistency:** Keep naming consistent across the plan: `app-factory-orchestrator`, `example-tool-app`, `CapabilityRegistry`, `bootstrapApp`, and `ToolApp`.

## Execution Handoff

Plan complete and saved to `docs/superpowers/plans/2026-05-02-app-factory-v1.md`. Two execution options:

**1. Subagent-Driven (recommended)** - I dispatch a fresh subagent per task, review between tasks, fast iteration

**2. Inline Execution** - Execute tasks in this session using executing-plans, batch execution with checkpoints

**Which approach?**
