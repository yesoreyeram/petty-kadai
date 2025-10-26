# Week 12 — Cart service (in-memory) (one microservice)

tools-introduced: Cart service (Go/net/http), in-memory

concepts-covered:

- Cart semantics; idempotent add/remove; totals; guest vs authenticated

proposed-architecture:

- Add Cart service with in-memory store; per-user carts via JWT subject

changes-to-system-design:

- Define `/api/cart` routes; idempotency keys; basic validation

tasks-checklist:

- [ ] Implement endpoints: get cart, add item, remove item, clear
- [ ] Use idempotency key header for updates
- [ ] Compute totals using catalog price
- [ ] Update gateway routes to forward `/api/cart/*`

skills-required:

- Go handlers; state management; idempotency

prerequisites:

- Weeks 01–11 running

deliverables:

- Working cart operations with in-memory storage

acceptance-criteria:

- Concurrent adds/removes consistent; idempotent updates don’t duplicate

## Proposed architecture diagram

```mermaid
flowchart LR
  GW[Nginx API Gateway]
  CART[Cart Service (in-memory)]
  CATALOG[Catalog Service]

  GW --> CART --> CATALOG
```
