# Week 14 — Inventory service (one microservice)

tools-introduced: Inventory service (Go/net/http)

concepts-covered:

- Stock model per SKU/variant; single source of truth

proposed-architecture:

- Add Inventory service with in-memory stock map (will add locks later)

changes-to-system-design:

- Define `/api/inventory` routes: get stock, reserve mock (no lock yet)

tasks-checklist:

- [ ] Implement get stock by SKU and reserve endpoint (returns tentative response)
- [ ] Input validation; consistent errors
- [ ] Update gateway routes for inventory
- [ ] Add basic tests for race scenarios (expected failures without locks)

skills-required:

- Go handlers; concurrency basics

prerequisites:

- Weeks 01–13 running

deliverables:

- Inventory endpoints available; known race limitations documented

acceptance-criteria:

- Endpoints respond correctly under single-threaded use; race tests demonstrate need for locks

## Proposed architecture diagram

```mermaid
flowchart LR
  GW[Nginx API Gateway]
  INVENTORY[Inventory Service (in-memory)]

  GW --> INVENTORY
```
