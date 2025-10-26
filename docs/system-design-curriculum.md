# System Design Architect Learning Path

Capstone: Build a Small-Scale Shopping Portal ("Mini Amazon") while mastering distributed systems, microservices, reliability, availability, security, performance, observability, analytics, networking, deployments, queues, workflows, batch processing, containers, and Kubernetes.

This curriculum uses the attached transcript “Designing a Scalable and Global E-commerce Platform System” as thematic guidance. We’ll implement production-grade patterns—at small scale—so you truly learn by doing.

## Who this is for

- Engineers who know one backend language (Node/TS, Go, Java, or Python)
- Comfortable with Git, basic HTTP, REST/JSON, Docker
- Optional but helpful: SQL basics

## Learning outcomes

By completing this program you will be able to:

- Design microservice architectures with clear service boundaries, APIs, and data ownership
- Choose data stores by workload (relational for orders/payments, document for catalog, cache for speed, search for discovery)
- Apply CAP trade-offs, strong vs eventual consistency, and the Saga pattern
- Build event-driven systems with queues/streams and idempotent consumers
- Implement distributed locks (e.g., Redis SET NX + TTL) for hot-path correctness (inventory)
- Design for availability (rate limiting, backpressure, retries, circuit breakers), scalability (sharding, replicas, CDN, caches), and disaster recovery (RTO/RPO)
- Secure systems (JWT, TLS, OWASP Top 10, PCI-scope minimization via tokenization)
- Operate systems (metrics, logs, traces, dashboards, SLOs, on-call runbooks)
- Deploy to Kubernetes; manage releases with blue/green or canary
- Build an analytics pipeline isolated from OLTP

## How to use this curriculum

- Format: 16 weeks (flex to 12/20). ~6–8 hrs/week split between theory (30%) and hands-on (70%).
- Each module includes: Concepts, Build, Deliverables, Acceptance Criteria, Stretch.
- Default stack (feel free to swap equivalents):

  - Language: TypeScript/Node (Express/Fastify) or Go (net/http) or Java (Spring Boot). Pick ONE.
  - Data: Postgres (Users/Orders/Payments), MongoDB (Catalog), Redis (cache/locks), Elasticsearch/OpenSearch (Search), MinIO (S3-compatible object storage), Kafka or RabbitMQ (events), ClickHouse or DuckDB (analytics warehouse, local-friendly).
  - Infra: Docker Compose (weeks 1–7), then Kubernetes with Kind/Minikube (weeks 8+). CDN conceptually via Cloudflare; locally simulate with Nginx cache.
  - Observability: OpenTelemetry, Prometheus, Grafana, Loki, Jaeger/Tempo.
  - Auth: JWT via an auth service (Keycloak optional).
  - Testing/Load: k6, hey, Locust.

  ## Reference artifacts

  - Enterprise functional requirements: `docs/requirements/functional-enterprise.md`
  - Enterprise non-functional requirements: `docs/requirements/non-functional-enterprise.md`
  - Enterprise security requirements: `docs/requirements/security-enterprise.md`

## Capstone system overview (Mini Amazon)

Services (first-class):

- API Gateway/Edge: routing, rate limits
- User Service: accounts, JWT issuance
- Catalog Service: products (NoSQL), images in object storage
- Search Service: Elasticsearch indexer + query API
- Cart Service: per-user carts
- Inventory Service: stock per SKU/variant, Redis lock on reserve
- Order Service: create order, orchestrate Saga
- Payment Adapter: talks to “fake gateway” to tokenize and authorize (no raw card data in our system)
- Notification Service: email/Webhook (local: console or Mailhog)

Shared infra:

- Kafka/RabbitMQ topics: product-updated, order-placed, order-paid, inventory-reserved, payment-failed, etc.
- Databases: Postgres (users, orders, payments); Mongo (products); Redis (locks/cache);
- Object storage: MinIO for product images; CDN simulated by Nginx cache
- Observability: OTel SDKs + collectors → Prometheus, Grafana, Loki, Jaeger
- Analytics: ELT from OLTP/outbox → Kafka → ClickHouse/DuckDB

Key patterns from the transcript you’ll implement:

- API Gateway with rate limiting; JWT stateless auth
- Event-driven choreography; decoupled services via topics/queues
- Inventory correctness using Redis SET NX + TTL locks
- Saga for order → payment → inventory with compensation
- Separate search index; near-real-time updates via product-updated events
- DR thinking with RTO/RPO targets; multi-AZ/region simulation locally
- Observability pillars: metrics, logs, tracing; business KPIs (conversion, failure rate)

## Weekly plan (16 weeks)

### Week 1 — Foundations and single service (one tool)

- Concepts: Functional vs non-functional requirements; latency SLO (<200ms), availability targets (99.9 vs 99.99), capacity planning lite.
- Build: Pick language (Go) and create a single User service with `GET /health` only. No databases, gateway, broker, or object storage yet. Local run via `go run` (optional: a tiny Compose with just this service).
- Deliverables: Architecture doc v0 (C4 Context + Container) focused on a single service; service conventions; Makefile.
- Acceptance: `GET /health` returns 200 in <50ms locally; graceful shutdown works; unit test for health endpoint passes.
- Stretch: Add simple structured logging and readiness/liveness endpoints.

### Week 2 — Catalog API (MongoDB) (one tool)

- Concepts: Flexible product schema; validation; pagination.
- Build: Catalog service with CRUD backed by MongoDB; no images/CDN yet.
- Deliverables: Product schema; handler + repository tests; basic integration test.
- Acceptance: Create/list/update products; pagination works; tests pass.
- Stretch: Category facets; simple price filter.

### Week 3 — Frontend (Next.js) (one tool)

- Concepts: SSR/ISR page; backend API calls; .env handling.
- Build: Next.js app with a Products page listing items from Catalog API.
- Deliverables: Basic UI and an e2e/integration test for the products page.
- Acceptance: Visiting `/products` shows catalog items; test passes.
- Stretch: Simple product detail page.

### Week 4 — Search service and indexing pipeline

- Concepts: Dedicated search (Elasticsearch), event-driven indexing, eventual consistency acceptable for search.
- Build: product-updated event; indexer consumes and updates ES; search API with facets and autocomplete.
- Deliverables: Index mapping; freshness SLI (<5s lag).
- Acceptance: New/updated product searchable within SLA.
- Stretch: Typo tolerance, synonyms.

### Week 5 — Inventory service (naive) (one concept)

- Concepts: Stock per SKU/variant; sequential safety with mutexes; no external locks yet.
- Build: Inventory service with in-memory or simple store; reserve/decrement endpoints.
- Deliverables: API + concurrency tests demonstrating basic correctness.
- Acceptance: Under small concurrency, reserves/decrements are correct; tests pass.
- Stretch: Persist to a lightweight store (file/SQLite) without introducing new infra.

### Week 6 — Inventory service with distributed locks

- Concepts: Hotspot control; Redis SET NX + TTL; correctness over throughput for critical section.
- Build: Reserve stock endpoint acquires lock by SKU (or SKU+variant), re-reads stock, decrements, releases.
- Deliverables: Lock design notes (key, TTL, retry/backoff); failure modes.
- Acceptance: Last-item race test (10 concurrent buyers → ≤1 success).
- Stretch: Token bucket for “flash sale” gate before reaching inventory.

### Week 7 — Orders + Payments with Saga

- Concepts: Strong consistency for orders/payments; saga orchestration vs choreography; compensation.
- Build: Order service orchestrates steps: create order (pending) → call Payment Adapter (tokenized) → on success publish order-placed → Inventory reserves; on failure emit compensation (refund/cancel order).
- Deliverables: State machine; outbox pattern for reliability.
- Acceptance: Fault injection proves compensations work; no “stuck” partial orders.
- Stretch: Dead-letter queue handling and reprocessing tools.

### Week 8 — Observability deep dive

- Concepts: Metrics, logs, traces; exemplars; RED and USE methods; tracing across services.
- Build: OpenTelemetry SDKs; exporters to Collector → Prometheus, Loki, Jaeger; dashboards and alerts.
- Deliverables: Service dashboards; trace showing checkout path.
- Acceptance: Trace includes gateway → user → cart → order → payment → inventory; alert on 5xx >1%/5m.
- Stretch: Business KPIs dashboard (orders/min, auth failures, cart abandonment).

### Week 9 — Reliability: backpressure and resiliency patterns

- Concepts: Retries with jitter, timeouts, circuit breakers, bulkheads; queue depth alarms.
- Build: Circuit breaker in gateway for unstable downstream; consumer concurrency tuning; poison message handling.
- Deliverables: Runbook for “payment gateway slow/failing”.
- Acceptance: Under injected latency, system degrades gracefully; alerts fire with clear action.
- Stretch: Chaos experiments (kill a service; observe failover paths).

### Week 10 — Performance engineering

- Concepts: Caching layers (client/CDN/gateway/service/data), n+1 avoidance, connection pooling.
- Build: k6/hey scenarios: browse → search → add to cart → checkout; profile p95.
- Deliverables: Perf test scripts; optimization change log.
- Acceptance: p95 < 200ms for browse/search/cart at 50–100 RPS local; CPU<70%.
- Stretch: Read replicas for Postgres; cache aside pattern for catalog hot items.

### Week 11 — Data modeling, sharding, and replicas

- Concepts: Sharding keys (user_id/order_id), hot shard detection, read replicas; CAP trade-offs.
- Build: Simulate user_id-based partitioning (schema or app-level routing); add a read replica (logical) for reads.
- Deliverables: Partitioning plan; hot-shard mitigation notes.
- Acceptance: Reads routed to replica; write path unaffected; correctness tests pass.
- Stretch: Auto-shard rebalancing plan (design only).

### Week 12 — Security hardening (introduce API gateway + JWT)

- Concepts: TLS everywhere, OWASP Top 10, secrets management, JWT scopes/roles, rate limiting/WAF, PCI scope minimization via tokenization.
- Build: Introduce API Gateway (Nginx/Envoy) and enable JWT validation at the edge; mTLS between services (dev via mkcert); secret storage; input validation; audit logging.
- Deliverables: Threat model (STRIDE-lite); pentest checklist.
- Acceptance: Protected endpoints require valid JWT at the gateway; security tests pass; sensitive logs redacted.
- Stretch: OPA/ABAC for fine-grained auth.

### Week 13 — Deployments: Kubernetes

- Concepts: Pods, Services, Ingress, HPA, Requests/Limits, readiness/liveness, blue/green vs canary.
- Build: Port Compose stack to Kind/Minikube; Helm or Kustomize; HPA on CPU/RPS.
- Deliverables: Manifests/Charts; rollout procedure doc.
- Acceptance: Zero-downtime rollout of catalog service; auto-scale under load.
- Stretch: Service mesh (Linkerd/Istio) for mTLS and traffic shifting.

### Week 14 — Batch processing and workflows

- Concepts: Cron jobs, retries, idempotency; workflow engines (Temporal/Argo Workflows) vs DIY.
- Build: Nightly price sync batch; weekly cleanup job; optional: workflow engine to orchestrate ETL.
- Deliverables: Playbooks; backfill process.
- Acceptance: Batch runs idempotently; re-runnable without duplicates.
- Stretch: Scheduled inventory reconciliation report.

### Week 15 — Analytics and data platform

- Concepts: OLTP vs OLAP; ETL/ELT; CDC/outbox; warehouse isolation.
- Build: Stream order events to ClickHouse/DuckDB; create BI queries/dashboards (sales/day, conversion by category).
- Deliverables: Data models; freshness SLA; lineage doc.
- Acceptance: Analytics queries do not impact OLTP latency; dashboard renders in <3s.
- Stretch: AB testing framework skeleton.

### Week 16 — DR planning and final hardening

- Concepts: Multi-AZ/region strategies; RTO vs RPO trade-offs; backups/restore drills.
- Build: Backup/restore scripts; simulate region failover (two clusters or namespaces) via global ingress toggle.
- Deliverables: DR runbook; RTO/RPO targets and evidence.
- Acceptance: Recovery drill meets goals (e.g., RTO 10m, RPO 5m) in local simulation.
- Stretch: Active-active simulation with write fencing.

## TDD and weekly runnable outcomes

We practice test-driven development throughout. For every change:

- Write a failing unit/integration test first that captures the behavior.
- Implement the minimal code to pass the test.
- Refactor with tests green; commit tests and code together.

At the end of each week, the stack must start and be demoable with a repeatable script. Below is the explicit runnable state and tests per week.

Weeks 1–16 (core program):

- Week 1 Runnable: run the single service locally (`go run` or `docker compose up` with only the User service) and `GET /health` returns 200. Tests: health handler unit test.
- Week 2 Runnable: Catalog CRUD working against Mongo; unit + basic integration tests pass. Tests: model validation; repository/handler tests; pagination checks.
- Week 3 Runnable: Frontend lists products from Catalog. Tests: page renders list via API; simple e2e/integration test.
- Week 4 Runnable: Search service indexes `product-updated` and serves queries. Tests: indexing pipeline end-to-end; freshness SLI check (<5s) with polling test.
- Week 5 Runnable: Inventory service (naive) handles reserve/decrement correctly under small concurrency. Tests: race tests verifying correctness.
- Week 6 Runnable: Inventory reserve endpoint with Redis lock guarantees ≤1 success on last item. Tests: contention test (10 buyers/1 item); lock TTL and retry backoff.
- Week 7 Runnable: Order + Payment Saga executes; compensation works on failure. Tests: state machine transitions; outbox emission; refund/cancel compensation paths.
- Week 8 Runnable: Traces span gateway→services; metrics/logs visible in Grafana/Loki; alerts configured. Tests: tracing propagation test; metrics presence; alerting rule unit tests.
- Week 9 Runnable: Resiliency patterns enabled (timeouts, retries, circuit breaker); system degrades gracefully under injected faults. Tests: fault injection scenarios with assertions on fallbacks and error budgets.
- Week 10 Runnable: Performance scripts run; p95 meets targets at local RPS. Tests: automated perf check with thresholds; connection pool sizing sanity tests.
- Week 11 Runnable: Read routing to replica simulated; partitioning logic exercised. Tests: replica reads vs primary writes; shard key routing tests.
- Week 12 Runnable: mTLS (dev), input validation, secret management; logs redacted. Tests: security unit tests (validators), e2e TLS handshake, redaction checks.
- Week 13 Runnable: Kubernetes manifests deploy services; HPA scales under load; zero-downtime rollout. Tests: readiness/liveness probes; rollout integration test.
- Week 14 Runnable: Batch jobs run on schedule and ad-hoc; re-runnable idempotently. Tests: job idempotency; retry/backoff semantics.
- Week 15 Runnable: Analytics queries on ClickHouse/DuckDB populated from events; dashboards render. Tests: data freshness and correctness (sample aggregates).
- Week 16 Runnable: Backup/restore + failover drill completes within RTO/RPO. Tests: scripted restore and traffic switch; validation of data integrity post-restore.

Extension weeks (17–20) — optional but recommended:

- Week 17 Runnable: Message broker mediates events; consumers process reliably; DLQ configured. Tests: publish/consume with outage and replay; idempotent consumer tests.
- Week 18 Runnable: Payment Adapter service integrated with timeouts/retries. Tests: client timeout and retry behavior; emitted events on success/failure.
- Week 19 Runnable: Prometheus scrapes `/metrics`; starter dashboards available. Tests: metric emission assertions; scrape config sanity.
- Week 20 Runnable: Jaeger shows complete checkout traces across services. Tests: trace span/attribute assertions; error span events when faults occur.

## Milestones and acceptance criteria (summary)

1. Foundations (Wk1–3): Single service boot, catalog and frontend functional; search added in Wk4
2. Transactional core (Wk5–7): Inventory (naive + locks), order/payment Saga
3. Operability (Wk8–10): Observability, reliability, performance targets met
4. Scale/readiness (Wk11–13): Sharding plan, security hardening, Kubernetes deploys
5. Data/DR (Wk14–16): Batch/Workflows, Analytics pipeline, DR drill

Each milestone should include:

- Demo script and screencast
- Design doc delta (what changed and why)
- Test evidence (unit/integration/load)
- Operational docs (dashboards, alerts, runbooks)

## Events and topics (suggested Kafka/RabbitMQ topology)

- catalog.product-updated
- order.order-created
- order.order-paid
- order.order-cancelled
- inventory.stock-reserved
- inventory.stock-released
- payment.authorized
- payment.failed
- notification.email-requested

Design all consumers as idempotent; store a processed message ID and use outbox pattern for producers.

## Reliability patterns to implement

- Timeouts + retries with exponential backoff + jitter
- Circuit breakers at gateway and critical clients
- Bulkheads: isolated threadpools/connection pools per downstream
- Rate limiting (token bucket) at edge and service-level
- Dead letter queues and replay tooling

## Security checklist (from transcript themes)

- TLS everywhere (dev: mkcert); mTLS internal if mesh-less
- JWT with expiration, rotation (JWKS), minimal claims
- Input validation and output encoding; CSRF for non-API forms if any
- Least privilege for DB and message broker users
- Never store raw card data; use tokenization via fake gateway
- Secrets via env + sealed secrets (Kubernetes)
- Logging: redact PII and secrets; audit logs for auth and payments

## Observability checklist

- Metrics: RED (Rate, Errors, Duration) per endpoint; business KPIs (orders/min)
- Logs: structured JSON; correlation IDs propagated via headers (traceparent)
- Tracing: 100% traces in dev; sampling in prod-like
- Alerts: SLO-based with burn rates; on-call runbooks linked

## Performance targets (right-sized for dev)

- Browse/search/cart p95 < 200ms at 50–100 RPS locally
- Checkout p95 < 400ms without external payment latency; graceful degradation under injected slowness
- Index freshness < 5s; inventory lock hold time < 2s under contention

## DR targets (dev simulation)

- RTO: 10 minutes (restore and switch traffic)
- RPO: 5 minutes (replication/checkpoint frequency)

Use these to reason about costs/complexity if you later move to cloud multi-region.

## Assessment rubrics

- Design clarity: clear boundaries, contracts, and ownership
- Correctness under concurrency: tests for inventory and saga compensations
- Operability: dashboards, alerts, runbooks exist and are actionable
- Security posture: threat model complete; secrets handled; logs clean of sensitive data
- Resilience: chaos tests show graceful degradation
- Performance: targets met with evidence

## Stretch directions (post-capstone)

- Marketplace: third-party sellers, payouts, moderation
- Internationalization: multi-currency, tax rules, locales
- ML: recommendations, ranking, fraud scoring
- Cost management: autoscaling, right-sizing, storage lifecycle

## Suggested study references mapped to the transcript

- Availability targets and multi-region failover: “Let’s talk about those nines” and GLB/Anycast sections
- Microservices, API Gateway, rate limiting: request flow and edge sections
- Event-driven architecture and queues: order-placed event and decoupling rationale
- Data stores by workload: NoSQL for catalog; strong consistency RDBMS for orders/payments
- Inventory correctness: distributed locks with Redis SET NX + TTL
- Transactional integrity: Sagas vs 2PC; compensating actions
- Search: dedicated Elasticsearch and near-real-time indexing
- Sharding/replicas: keys and hot shard pitfalls; replicas for read scaling
- Security: JWT stateless auth; TLS; PCI scope minimization
- Observability: metrics, logs, tracing pillars; BI isolation
- DR/Backup: RTO/RPO trade-offs and drills

Use these anchors as “why” behind every build decision you take.

## Getting started next

- Decide your primary language (Node/Go/Java/Python)
- Clone this repo and create a `/services` folder with skeletons for: gateway, user, catalog, search, cart, inventory, order, payment, notification
- Add `/infra/docker-compose.yml` with Postgres, Mongo, Redis, Kafka/RabbitMQ, MinIO, Elasticsearch, Prometheus, Grafana, Loki, Jaeger
- Begin Week 1 tasks; keep notes in `/docs/` and wire dashboards early so you see progress

If you want, ask me to scaffold the initial folders and Compose stack and we’ll start implementing Week 1 right away.
