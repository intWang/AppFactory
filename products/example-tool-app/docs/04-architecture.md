# Architecture Design

## Project Name
Example Tool App

## Project Slug
example-tool-app

## Input Source
Derived from the final example PRD, UX specification, and design references.

## Architecture Goals
Validate the shell, capability registry, and minimal feature rendering path.

## Platform Scope
Flutter template with iOS and Android first, Web-compatible where possible.

## App Shell
`MaterialApp` root with a single home route.

## Capability Modules
Foundation config, logging, error reporting, and growth entry points.

## Feature Modules
One home feature that renders the shell and proves bootstrap wiring.

## Route Strategy
Single root route for the example app.

## State Strategy
Bootstrap-owned app services with direct constructor injection.

## Data Flow
App services are created at startup and passed into the shell.

## Safety Boundaries
Feature code consumes bootstrap services and does not reach directly into platform-sensitive APIs.

## Risks
No active toolchain risk. Design exports are still placeholders until a real Figma file is attached.

## Gate Result
Pass
