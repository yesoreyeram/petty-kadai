# Week 06 — Catalog service (in-memory) (one microservice)

tools-introduced: Catalog service (Go/net/http), no DB yet

concepts-covered:

- Service boundaries; product model; versioned API

proposed-architecture:

- Add Catalog service with in-memory store for list/get product

changes-to-system-design:

- Define `/api/catalog` routes; no persistence; seed sample products on boot

tasks-checklist:

- [ ] Implement Catalog HTTP handlers: list, get by ID
- [ ] In-memory store with seed data
- [ ] Validate inputs and return problem+json errors
- [ ] Update gateway routes to forward `/api/catalog/*`

skills-required:

- Go handlers, routing, validation

prerequisites:

- Weeks 01–05 running

deliverables:

- Catalog endpoints responding with sample data via gateway

acceptance-criteria:

- `GET /api/catalog/products` returns seeded products; p95 < 100ms locally

## Proposed architecture diagram

```mermaid
flowchart LR
  UI[Next.js Web]
  GW[Nginx API Gateway]
  USER[User Service]
  CATALOG[Catalog Service (in-memory)]
  UI --> GW --> USER
  GW --> CATALOG
```
