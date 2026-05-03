# Flutter Product Shell Template

This template is the baseline product shell for AppFactory's multi-product monorepo.

## Intended Product Layout

Use this template under `products/<product-slug>/` with these sibling folders:

```text
products/<product-slug>/
  client/
    lib/
    test/
    config/
    assets/
  server/
    src/
    tests/
    config/
    deploy/
  docs/
  design/
  build/outputs/
```

## Launch Upgrade Rule

The product shell should run a shared launch upgrade check. When a newer build exists, show an upgrade prompt. When the latest released build is more than 3 builds ahead, trigger a forced-upgrade path instead of normal app entry.

## Required Product Docs

- `docs/00-intake.md`
- `docs/01-prd-draft.md`
- `docs/02-ux-spec.md`
- `docs/03-prd-final.md`
- `docs/04-architecture.md`
- `docs/05-dev-plan.md`
- `docs/06-test-cases.md`
- `docs/07-test-report.md`
- `docs/08-release-gate.md`
- `docs/09-retro-input.md`

## Required Design Assets

- `design/figma-link.md`
- `design/exports/`

## Verification

Run:

```bash
flutter test
```
