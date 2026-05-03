# App Factory V1 Design

**Date:** 2026-05-02
**Status:** Draft for review
**Scope:** V1
**Default stack:** Flutter
**Target products:** Tool-style apps
**Primary platforms:** iOS and Android
**Optional platform:** Web

## 1. Overview

App Factory V1 is a skill family for turning a lightweight app idea into a compiled, testable mobile application through a governed multi-role workflow. The first version prioritizes platform capability and process reliability over broad business generation. It focuses on small, practical tool apps and uses Flutter as the default stack so shared capabilities, UI patterns, testing strategy, and project structure can be reused across projects.

The workflow is designed to run mostly automatically. The orchestrator skill advances the project through product, architecture, development, and QA stages, stopping only for high-risk decisions, missing inputs, environment blockers, or quality failures. Release management and agile retrospective roles are reserved as placeholders in V1 so the main pipeline can be stabilized before expanding into full store release governance and process optimization.

## 2. Goals and Non-Goals

### Goals

- Accept raw user requirements and establish a project record.
- Produce structured documents that can be handed cleanly between roles.
- Generate a minimal Flutter app that compiles and can be tested.
- Reuse platform capabilities through pluggable shared modules.
- Preserve clear safety boundaries around shared modules, business modules, credentials, and monetization touchpoints.
- Archive test evidence and project memory for reuse across future projects.

### Non-Goals

- Full App Store or Play Store publishing automation in V1.
- Support for large social, marketplace, or SEO-heavy content products.
- Exhaustive business generation across many app categories.
- Complex monetization implementation in the first release.
- Full retrospective facilitation by SM or full release execution by RM.

## 3. Supported Product Shape

V1 is optimized for tool-style apps. These apps should solve a focused real user problem, remain small in scope, and be practical to implement with a common Flutter template and reusable capability modules.

The design assumption is:

- Flutter is the default and recommended stack.
- iOS and Android are the primary delivery targets.
- Web support is optional and should reuse the same architecture where possible.
- Web-specific limitations must be handled through explicit capability downgrade rules.

## 4. Skill Family Structure

The skill family is organized into three layers.

### 4.1 Orchestration Layer

`app-factory-orchestrator`

This skill receives user requirements, initializes project state, routes work through the role skills, tracks outputs, enforces gates, and pauses only when necessary.

### 4.2 Role Layer

- `app-factory-pm`
- `app-factory-ud`
- `app-factory-am`
- `app-factory-sd`
- `app-factory-qa`
- `app-factory-rm-placeholder`
- `app-factory-sm-placeholder`

These skills own professional role behavior and outputs. The orchestrator coordinates them but does not replace their expertise.

### 4.3 Shared Asset Layer

This layer includes reusable project assets rather than standalone roles:

- Flutter base template
- Documentation templates
- Capability registration patterns
- Factory memory
- Project memory
- Test report archive structure
- Safety and release policy placeholders

## 5. Role Responsibilities and Boundaries

### 5.1 PM

PM translates rough ideas into a practical, user-centered PRD for a small tool app. The PM skill focuses on user need, key flows, interaction direction, product scope, and reserved monetization touchpoints.

PM does not make technical architecture decisions.

### 5.2 AM

AM reviews the PRD from a technical perspective and maps the product into shared capabilities and feature modules. AM chooses which public modules to connect first, splits business modules, defines the architecture, and pushes PRD revisions back to PM if the requirement is technically weak, unsafe, or harmful to reuse.

AM does not directly implement product code.

### 5.3 SD

SD checks the local environment, follows the approved architecture, and implements the product using TDD. SD is responsible for producing a compilable Flutter app, automated tests, and development delivery notes.

SD does not expand product scope or bypass test discipline.

### 5.4 QA

QA creates test cases based on the finalized PRD and architecture, reviews SD test coverage, executes testing, captures screenshots, writes the report, and archives evidence for future comparison.

QA does not redefine the product scope.

### 5.5 RM Placeholder

RM is reserved in V1 to represent future ownership of release detail, release cost control, store compliance, certificates, accounts, and monetary security checks. V1 records the release gate inputs without executing real release operations.

### 5.6 SM Placeholder

SM is reserved in V1 to represent future ownership of process observation, retrospective facilitation, and memory extraction. V1 records retro inputs without running a formal retrospective flow.

## 6. Platform Architecture

V1 uses a three-part architecture:

- App Shell
- Capability Modules
- Feature Modules

### 6.1 App Shell

The shell initializes the app, configures routes, wires dependencies, installs modules, selects environment settings, and applies theme and layout conventions. It knows which modules are enabled, but it should not contain business logic from the modules themselves.

### 6.2 Capability Modules

Capability modules are reusable shared platform blocks. They must be pluggable and independently enabled or disabled by the shell.

### 6.3 Feature Modules

Feature modules contain the specific business functionality of a given tool app. They may consume shared capabilities but must not reach across other feature modules in uncontrolled ways.

## 7. Shared Capability System

V1 should define public capabilities in four groups.

### 7.1 Foundation

Shared infrastructure with no product-specific language:

- logging
- environment config
- error capture
- network client
- local storage
- analytics interface
- device information
- permission wrapper

### 7.2 Account

Authentication and session capabilities:

- login status
- guest access
- user profile
- session handling
- auth interception

V1 can start lightweight but the interface should exist from the beginning.

### 7.3 Growth

Shared growth and monetization placeholders:

- upgrade prompts
- announcement modal
- rating prompt
- feedback entry
- subscription or ad placement slots
- campaign entry points

The goal is not a heavy monetization implementation in V1. The goal is to reserve standard attachment points so PM can define growth opportunities without each project inventing them from scratch.

### 7.4 Toolkit UI

Cross-project presentation layer:

- theme
- typography
- button patterns
- forms
- lists
- loading state
- empty state
- error state
- modal patterns
- navigation shell
- responsive layout handling

This layer is the main place to absorb platform visual differences while preserving a consistent cross-platform experience.

## 8. Safety Boundaries

Safety is a first-class design requirement.

- Capability modules must not depend on feature modules.
- Feature modules must consume sensitive platform services only through approved capability interfaces.
- Payment, subscription, credentials, release accounts, keys, and store-related assets must go through explicit protected interfaces.
- Web support must declare downgrade behavior for capabilities that do not translate directly from mobile platforms.
- Shared modules must remain generic enough to be reused across apps without hidden business coupling.

## 9. Standard Workflow

The main V1 workflow is:

`User Requirement -> PM Draft PRD -> UD UX Design -> PM Final PRD -> AM Architecture -> SD Implementation -> QA Verification`

Each stage must generate a standard artifact and pass a minimal quality gate before the orchestrator advances automatically.

## 10. Input/Output Contracts by Stage

### 10.1 Intake

Produced by the orchestrator.

Inputs:

- raw user requirement
- default factory assumptions

Outputs:

- project name
- project slug
- support-scope decision
- current state
- next role
- initial risks
- unknowns

### 10.2 PM Output

Produced by PM.

Inputs:

- raw requirement
- target platform assumptions
- tool-app scope rules

First-pass outputs:

- `products/<product-slug>/docs/01-prd-draft.md`
- product overview
- target users
- user pain points
- value proposition
- core user flow
- feature list
- page list
- interaction principles
- monetization reserve points
- non-functional requirements
- explicit non-goals
- points requiring architecture review

Final outputs after UD review:

- `products/<product-slug>/docs/03-prd-final.md`
- UX changes from UD
- final page inventory
- final interaction constraints

### 10.3 UD Output

Produced by UD.

Inputs:

- drafted PRD
- product scope assumptions

Outputs:

- `products/<product-slug>/docs/02-ux-spec.md`
- `products/<product-slug>/design/figma-link.md`
- exported design screens under `products/<product-slug>/design/exports/`
- PM write-back items for final PRD

### 10.4 AM Output

Produced by AM.

Inputs:

- approved final PRD
- UX specification
- Figma link
- exported design screens
- capability catalog
- template boundaries

Outputs:

- architecture goals
- platform scope
- app shell design
- capability module plan
- feature module split
- routing strategy
- state strategy
- data flow
- error handling and logging strategy
- testing architecture guidance
- safety boundary statement
- PRD revision requests
- technical risks and tradeoffs
- monorepo placement decisions

### 10.5 SD Output

Produced by SD.

Inputs:

- approved architecture
- UX specification
- Figma link and exported screens
- Flutter base template
- TDD rules
- testing baseline

Outputs:

- environment check result
- implementation plan
- integrated capability modules
- feature implementation
- automated tests
- build commands
- build artifact location
- development blockers if present

SD must report blockers explicitly if the environment is not ready.

### 10.6 QA Output

Produced by QA.

Inputs:

- final PRD
- UX specification
- design link and exported screens
- final architecture
- codebase and test assets
- executable build

Outputs:

- test case set
- unit test review conclusion
- execution results
- screenshots
- defect log
- pass/fail gate decision
- archive references
- recommended rollback target if needed

## 11. Rework Rules

Only three formal rework paths are allowed in V1:

- AM can send work back to PM when the requirement is unsafe, infeasible, too costly, or damaging to reuse.
- UD can send work back to PM when interaction design exposes requirement ambiguity or missing product states.
- QA can send work back to SD when implementation, build quality, or test coverage is insufficient.
- QA can send work back to AM when failures reveal architecture-level issues that cannot be solved as isolated defects.

This keeps the process governed without allowing noisy circular review loops.

## 12. Orchestrator State Machine

The orchestrator should maintain an explicit state machine.

### 12.1 States

- `intake`
- `producting`
- `designing`
- `architecting`
- `developing`
- `testing`
- `passed`
- `blocked`
- `needs-rework`

### 12.2 Nominal Flow

`intake -> producting -> designing -> producting -> architecting -> developing -> testing -> passed`

### 12.3 Exception Flow

- `architecting -> needs-rework -> producting`
- `designing -> needs-rework -> producting`
- `testing -> needs-rework -> developing`
- `testing -> needs-rework -> architecting`
- `any state -> blocked`

### 12.4 Automatic Advancement

The default behavior is to continue automatically once the current stage artifact is present and meets the minimum gate. The orchestrator should not ask to continue after every stage. It should only pause when a real blocker or high-risk decision appears.

### 12.5 Pause Conditions

Pause only when one of the following is true:

- the requirement is materially incomplete
- the work exceeds V1 scope
- the technical route conflicts with current factory assumptions
- money, credentials, signing, account, or release risk is involved
- the development environment is blocked
- the QA quality gate is not met

### 12.6 Recovery

The orchestrator must be able to resume from the last recorded project state and existing artifacts rather than restarting the entire pipeline.

## 13. Quality Gates

### 13.1 PM Gate

- target user is explicit
- core task is explicit
- scope is bounded
- non-goals are explicit
- monetization reserve exists

### 13.2 UD Gate

- core user flow is covered
- key states are defined
- figma link exists
- ux spec exists
- PM write-back items are explicit

### 13.3 AM Gate

- capability module decisions exist
- feature split exists
- route, state, and data flow are described
- technical risks are identified

### 13.4 SD Gate

- environment check is complete
- project compiles
- core flow runs
- automated tests run
- build output is handoff-ready

### 13.5 QA Gate

- test cases exist
- unit test review is concluded
- core flows pass
- report and screenshots are archived

## 14. Directory Structure

```text
app-factory/
  packages/
    app_factory_foundation/
    app_factory_account/
    app_factory_growth/
    app_factory_ui/
    app_factory_tooling/

  products/
    <product-slug>/
      pubspec.yaml
      lib/
        app/
        features/
        integrations/
      test/
        unit/
        widget/
        integration/
      config/
      assets/
      docs/
        01-prd-draft.md
        02-ux-spec.md
        03-prd-final.md
        04-architecture.md
        05-dev-plan.md
        06-test-cases.md
        07-test-report.md
        08-release-gate.md
        09-retro-input.md
      design/
        figma-link.md
        exports/
      build/
        outputs/

  skills/
    app-factory-orchestrator/
      SKILL.md
    app-factory-pm/
      SKILL.md
    app-factory-am/
      SKILL.md
    app-factory-sd/
      SKILL.md
    app-factory-qa/
      SKILL.md
    app-factory-rm-placeholder/
      SKILL.md
    app-factory-sm-placeholder/
      SKILL.md

  templates/
    flutter_tool_app/
    flutter_product_shell/
    docs/
      prd.template.md
      add.template.md
      test-cases.template.md
      test-report.template.md
      memory.template.md
  docs/
    specs/
    plans/

  memory/
    factory-memory.md
    project-memory/

  archives/
    test-reports/
    build-artifacts/
```

## 15. Naming Convention

Every project should use a stable `project-slug`.

Example:

- `products/pomodoro/`
- `products/pomodoro/docs/`
- `memory/project-memory/2026-05-02-pomodoro.md`
- `archives/test-reports/2026-05-02-pomodoro/`
- `archives/build-artifacts/2026-05-02-pomodoro/`

## 16. Standard Project Artifacts

Each product directory should contain:

- `products/<product-slug>/docs/00-intake.md`
- `products/<product-slug>/docs/01-prd-draft.md`
- `products/<product-slug>/docs/02-ux-spec.md`
- `products/<product-slug>/docs/03-prd-final.md`
- `products/<product-slug>/docs/04-architecture.md`
- `products/<product-slug>/docs/05-dev-plan.md`
- `products/<product-slug>/docs/06-test-cases.md`
- `products/<product-slug>/docs/07-test-report.md`
- `products/<product-slug>/docs/08-release-gate.md`
- `products/<product-slug>/docs/09-retro-input.md`
- `products/<product-slug>/design/figma-link.md`
- `products/<product-slug>/design/exports/`

These files provide deterministic handoff points between skills and future automation.

## 17. Documentation Format Rules

Each artifact must include:

- project name
- project slug
- input source
- conclusion or gate result

The following patterns are considered invalid:

- vague placeholders such as "TBD" or "later"
- soft language that hides missing decisions
- silently skipped sections without explaining impact

If something is missing, the document must say what is missing and why it matters.

## 18. Memory Model

V1 should distinguish between three levels of memory.

### 18.1 Project Memory

Records project-specific context, decisions, defects, and leftovers. This should not pollute cross-project guidance.

### 18.2 Factory Memory

Records stable patterns that help future projects, such as strong module splits, recurring QA findings, and recurring Flutter/Web compatibility notes.

### 18.3 Skill Memory

Records improvements to the skills themselves, such as which prompting structures produce more reliable role behavior.

Only high-value, repeat-relevant information should be written to memory.

## 19. Archive Requirements

After QA passes, the project should archive at minimum:

- the final test report
- key screenshots
- build artifact or build metadata

This provides a baseline for future comparison and regression work.

## 20. V1 Skill Inventory

V1 includes:

- `app-factory-orchestrator`
- `app-factory-pm`
- `app-factory-am`
- `app-factory-sd`
- `app-factory-qa`
- `app-factory-rm-placeholder`
- `app-factory-sm-placeholder`

## 21. Recommended Implementation Order

The first delivery slice should prioritize the factory backbone:

1. Write the V1 design spec.
2. Write the implementation plan.
3. Create the skill directories and shared templates.
4. Implement `orchestrator`, `pm`, `am`, and `sd`.
5. Create the Flutter base template.
6. Run one example tool-app end-to-end.
7. Add `qa`.
8. Add RM and SM placeholders.

## 22. Acceptance Criteria for V1

V1 is successful when:

- a user can provide a tool-app idea
- the orchestrator can initialize and track the project
- PM can produce an acceptable PRD
- AM can produce an acceptable architecture design
- SD can generate a minimal Flutter app that compiles
- QA can produce test cases and a test report against the output
- all project artifacts, memory entries, and archive paths are consistent and traceable

## 23. Open Future Extensions

V1 intentionally leaves room for later expansion:

- full RM release workflow
- full SM retro workflow
- stronger monetization components
- richer capability catalog
- broader app category support
- deeper Web-specific strategy

## 24. Review Notes

This spec is intentionally biased toward platform capability and process governance first. It trades early business variety for stronger reuse, better safety boundaries, and a higher chance of stabilizing the app factory core before expanding scope.

## 25. First Implementation Slice

The first implementation slice should include:

- the seven skill directories
- shared packages under `packages/`
- product-local docs under `products/<product-slug>/docs/`
- shared document templates
- a reusable Flutter tool-app template
- a reusable Flutter product-shell template
- one example project at `products/example-tool-app/`
