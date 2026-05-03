# Contributing

AppFactory is a multi-product monorepo. The repository hosts multiple independent products while reusing one shared capability layer. Every contributor must preserve the boundary between public packages and product-specific code.

## Client and Server Decision Rule

- A product may be client-only when it is a local-first tool with no shared accounts, no centralized business rules, no paid entitlements, and no shared data.
- AppFactory products should default to reserving server capability because most products are expected to add account, entitlement, subscription, or monetization entry points later.
- Shared account management and payment or entitlement logic should prefer shared services instead of product-specific reimplementation.
- Product-specific server logic is optional and must be decided by PM and AM based on the product's business requirements, data model, and security needs.
- If a product does not need a dedicated server yet, it should still keep room in its structure for a future `server/` directory.
- App shell launch must run a shared upgrade check before normal product use.
- When a newer release exists, the app should prompt the user to upgrade.
- When the latest released build is more than 3 builds ahead of the installed build, the app must force upgrade and exit instead of continuing.

## Repository Layers

- `packages/` stores shared capabilities, reusable UI, tooling, and public modules.
- `products/` stores concrete products. Every product must live under `products/<product-slug>/`.
- `templates/` stores reusable product shells, scaffolds, and document templates.
- Root-level product code is not allowed.

## Product Isolation Rules

- Each product must own its own source code, config, tests, docs, and build outputs.
- `products/a` must not import or depend directly on `products/b`.
- Cross-product reuse must flow through `packages/`.
- Logic that serves only one product must remain inside that product directory.

## Public Package Admission Rules

- Code may enter `packages/` only when it is reusable across at least two products and has been reviewed by AM.
- Public capabilities must define clear interfaces, dependency direction, config shape, and downgrade behavior.
- Product names, product copy, and product-specific flows must not leak into public packages.

## Required Product Structure

Every product must include at least:

- `client/`
- `docs/`
- `design/`
- `build/outputs/`

Recommended structure:

```text
products/<product-slug>/
  client/
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
  server/
    src/
    include/
    tests/
    config/
    deploy/
  docs/
  design/
    figma-link.md
    exports/
  build/
    outputs/
```

## Build and Artifact Ownership

- Build outputs must be written to the owning product's `build/outputs/`.
- Client and server deploy artifacts must be written to product-local output paths and must not be mixed with other products.
- Test reports, screenshots, and validation records must be written to the owning product's `docs/` or matching archive location.
- Figma links, exported design screens, and UX specs must stay with the owning product.
- Multiple products must not share a single output directory.

## Documentation and Test Ownership

- Product docs stay with the product.
- Product design assets stay with the product.
- Product tests stay with the product.
- Public package tests stay with the public package.
- Test files and docs must not be mixed across products.

## Role Alignment

### UD

- Convert `01-prd-draft` into interaction-design deliverables before architecture starts.
- Maintain product-local Figma references, exported screens, and UX specifications.
- Record design decisions that PM must write back into the final PRD.
- Surface reusable UI patterns to AM before shared-component decisions are made.

### AM

- Use the final PRD plus UD deliverables as architecture inputs.
- Decide whether new code belongs in `packages/`, `products/<product-slug>/`, or the product integration layer.
- Decide whether the product needs a dedicated `server/` now or only a reserved server boundary for later.
- Keep shared account, payment, and entitlement capabilities out of ad-hoc product server implementations when a shared service can own them.
- Keep upgrade-check policy in shared growth capabilities instead of redefining version rules per product.
- Reject abstractions that carry product-specific semantics into public packages.
- Protect cross-product isolation and dependency direction.

### SD

- Default all client implementation to `products/<product-slug>/client/`.
- Add `products/<product-slug>/server/` only when the approved architecture requires product-specific backend logic.
- Change `packages/` only when the architecture explicitly calls for shared extraction.
- Prevent cross-product imports and misplaced outputs.
- Keep build, test, and doc artifacts inside the owning product boundary.
- Implement UI against UD's Figma and exported design references, not from PRD alone.
- Wire the shared upgrade capability into app-shell launch when architecture requires upgrade governance.

### QA

- Verify repository-boundary compliance in addition to feature correctness.
- Check import boundaries, test placement, doc placement, and build-output placement.
- Check design consistency against UD exports and UX specifications.
- Verify optional upgrade prompts and forced-upgrade blocking behavior when version policy applies.
- Reject work that violates product isolation or artifact ownership rules.

## Non-Compliant Examples

The following are not allowed:

- One product importing another product directly.
- Shared logic copied into multiple products without AM review.
- Product-specific flows moved into `packages/`.
- Build outputs written outside the product's own output path.
- Product docs, design assets, or tests stored outside the product directory.
- Product-specific account or payment services created without first checking whether they belong in a shared service.
- Product-specific version-gap rules or skipped launch upgrade checks without AM approval.
