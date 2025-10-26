<!-- Sync Impact Report:
Version change: N/A → 1.0.0
Modified principles: (new document — all principles introduced)
Added sections: Enterprise Constraints; Development Workflow & Quality Gates
Removed sections: None
Templates requiring updates: .specify/templates/plan-template.md ✅; .specify/templates/spec-template.md ✅; .specify/templates/tasks-template.md ✅
Follow-up TODOs: None
-->

# Petty Kadai Constitution

## Core Principles

### Principle I — Focused Weekly Increment (NON-NEGOTIABLE)

- Each curriculum week MUST introduce exactly one new capability (tool, service, or concept) end to end.
- Every increment MUST culminate in a runnable artifact (`go run`, `docker compose up`, or documented equivalent) and a demo checklist.
- Dependencies outside the current week MUST be simulated or deferred; no premature multi-tool coupling.
  Rationale: Narrow weekly scope preserves learner focus and keeps the system in a consistently demoable state.

### Principle II — Test-Driven Runnable Delivery

- Teams MUST author failing automated tests before implementing features (unit, integration, or contract as appropriate).
- CI MUST stay green; merges blocked unless tests cover the new capability and pass locally.
- Every artifact MUST document its run command and expected outcomes in-line with tests.
  Rationale: TDD plus runnable guidance guarantees repeatable learning outcomes and reliable services.

### Principle III — Documentation Mirrors Implementation

- Whenever code changes, associated docs (`docs/weeks/`, `docs/system-design-curriculum.md`, runbooks) MUST be updated in the same change set.
- Acceptance criteria and BDD notes MUST live alongside weekly READMEs; drift is unacceptable.
- Architecture decisions MUST be logged, referencing the impacted requirement IDs.
  Rationale: Learners rely on the documentation as the single source of truth; divergence erodes trust.

### Principle IV — Standard Interfaces & Observability by Default

- Go services MUST expose HTTP handlers via `net/http`, include `GET /health`, and emit structured logs.
- Observability hooks (request IDs, basic metrics placeholders, trace propagation) MUST be scaffolded as part of the first implementation pass.
- External interfaces (APIs, events) MUST declare contracts (OpenAPI, JSON schema, or protobuf) before implementation.
  Rationale: Shared primitives make multi-service evolution coherent and instrumentable.

### Principle V — Enterprise Guardrail Alignment

- Feature work MUST cross-check the enterprise functional, non-functional, and security requirements before acceptance.
- Any gaps against those requirements MUST be logged with remediation tasks and owners before merge.
- Infrastructure and data decisions MUST respect the documented compliance and resilience baselines.
  Rationale: The curriculum must model real enterprise expectations, not simplified prototypes.

## Enterprise Constraints

- Reference `docs/requirements/functional-enterprise.md`, `docs/requirements/non-functional-enterprise.md`, and `docs/requirements/security-enterprise.md` for canonical requirements.
- Performance targets: adhere to defined SLOs (e.g., `/health` <50 ms p95 in Week 1, `/checkout` ≤400 ms p95 when introduced).
- Compliance posture: PCI scope minimisation, GDPR/CCPA controls, SOC 2-aligned change management.
- Data governance: ULID identifiers, versioned schemas, audit logging, and retention rules as per requirement docs.
- Tooling: Prefer Docker Compose early, graduate to Kubernetes once observability and resilience scaffolding meet enterprise baselines.

## Development Workflow & Quality Gates

- Speckit artifacts (plan, spec, tasks) MUST include a Constitution Check section confirming alignment with all principles.
- Week kick-off requires: constitution-aligned plan approval, dependency mock strategy, and test strategy sign-off.
- Pull requests MUST show: tests added/updated, documentation updates, requirement IDs referenced, and observability touchpoints.
- Releases or demos MUST verify enterprise guardrail checklists and capture results in weekly README retro sections.

## Governance

- The constitution supersedes conflicting process guidance. Exceptions require architecture board approval and documented expiry.
- Amendments:
  - Proposals documented via PR referencing rationale and impact.
  - Semantic versioning: MAJOR for breaking/removing principles, MINOR for new principles or sections, PATCH for clarifications.
  - Approval requires at least one architecture lead, one curriculum owner, and one security representative.
- Compliance:
  - Quarterly audits ensure artifacts and implementations adhere to principles and enterprise constraints.
  - Violations trigger remediation tasks tracked to completion before subsequent weekly releases.
- Runtime guidance (README, docs) MUST be updated in the same change set whenever governance changes affect learners.

**Version**: 1.0.0 | **Ratified**: 2025-10-26 | **Last Amended**: 2025-10-26
