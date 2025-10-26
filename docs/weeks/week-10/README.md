# Week 10 — Search service (skeleton) (one microservice)

tools-introduced: Search service (Go/chi), no ES yet

concepts-covered:

- Search API design; query params; pagination; result shape

proposed-architecture:

- Add Search service with stubbed in-memory results; will add ES later

changes-to-system-design:

- Define `/api/search` routes; simple match on product name (in-memory)

tasks-checklist:

- [ ] Implement Search HTTP handlers (query, facets placeholder)
- [ ] Integrate with Catalog API for in-memory matching
- [ ] Update gateway routes to forward `/api/search/*`
- [ ] Add basic response time SLO and measure

skills-required:

- Go handlers; pagination; basic query parsing

prerequisites:

- Weeks 01–09 running

deliverables:

- Search endpoints returning plausible results from in-memory catalog list

acceptance-criteria:

- `GET /api/search?q=foo` returns items where name contains `foo`; p95 < 150ms

## Proposed architecture diagram

```mermaid
flowchart LR
  GW[Nginx API Gateway]
  SEARCH[Search Service (in-memory)]
  CATALOG[Catalog Service]

  GW --> SEARCH --> CATALOG
```
