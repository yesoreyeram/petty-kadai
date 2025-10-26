# Enterprise Non-Functional Requirements — Petty Kadai

> This document captures the quality attributes, guardrails, and operational targets that govern Petty Kadai across environments. It complements the functional requirements and must be met before production launch.

## 1. Reliability & Availability

- **Service Level Objectives (SLOs)**:
  - Auth, Catalog, Cart, Checkout, Order, Payment APIs: 99.95% monthly availability.
  - Search, Notification APIs: 99.9% monthly availability.
- **Service Level Indicators (SLIs)**: Success rate, latency (p50/p95), error budgets tracked per endpoint.
- **Redundancy**: Active-active deployment across at least two availability zones per region.
- **Failover**: Automated failover for stateless services via Kubernetes; database failover via managed services with < 60s switchover.
- **Maintenance**: Planned downtime windows communicated 14 days in advance; zero-downtime rolling deploys preferred.

## 2. Performance & Capacity

- **Latency Targets (p95)**:
  - `GET /catalog/products`: <= 250ms
  - `POST /checkout`: <= 400ms
  - `POST /auth/login`: <= 200ms
  - Web UI Largest Contentful Paint (LCP): <= 2.5s on broadband, <= 4s on 3G.
- **Throughput**: Support peak 500 orders/minute, 5k concurrent carts, 20k search requests/minute with <= 70% resource utilization.
- **Scalability**: Horizontal scaling for stateless services; auto-scaling policies based on CPU, latency, and custom queue depth metrics.
- **Capacity Planning**: Quarterly load and stress tests; maintain 30% headroom over projected peak.

## 3. Resilience & Fault Tolerance

- Implement circuit breakers, retries with jitter, and timeouts for all cross-service calls.
- Use Saga pattern with compensations to ensure eventual consistency on distributed transactions.
- Chaos testing (weekly in staging, quarterly in production) covering network partitions, dependency failures, and slow responses.
- Graceful degradation strategies (serve cached catalog, disable non-essential features) under partial outages.

## 4. Security & Privacy

- **Identity & Access Management**: RBAC enforced across services; least privilege; quarterly access reviews.
- **Data Protection**: Encrypt data in transit (TLS 1.2+), at rest (AES-256). Manage keys via KMS with rotation every 365 days.
- **Secrets Management**: Use vault solution; no secrets checked into source control.
- **Vulnerability Management**: Monthly SCA and SAST scans; patch critical vulnerabilities within 7 days, high within 30 days.
- **Privacy Controls**: GDPR/CCPA compliance; privacy impact assessments for major features.
- **Audit Logging**: Tamper-evident logs for sensitive actions; retained minimum 400 days.

## 5. Compliance & Regulatory

- **Frameworks**: SOC 2 Type II, PCI DSS SAQ A compliance, ISO 27001 alignment.
- **Data Residency**: Ability to host customer data within designated geographic regions (e.g., EU) with configurable policy enforcement.
- **Accessibility**: WCAG 2.1 AA compliance for all customer-facing interfaces.
- **Records Management**: Financial records retained 7 years; user activity logs 2 years; ability to enforce legal holds.

## 6. Observability & Monitoring

- **Metrics**: RED and USE metrics instrumented via OpenTelemetry; collected by Prometheus/Grafana stack.
- **Logs**: Structured JSON logs with correlation IDs, centralized in Loki/ELK; PII redaction rules enforced.
- **Tracing**: Distributed tracing across services with 100% sampling for critical workflows; Jaeger or Tempo as backend.
- **Alerting**: On-call alert thresholds aligned to error budgets; alerts routed via PagerDuty/ops channel with runbook links.
- **SLO Reporting**: Monthly reviews; automated burn-rate alerts for fast/slow indicators.

## 7. Data Management & Quality

- **Consistency**: Strong consistency within services; eventual consistency across domains with well-defined reconciliation jobs.
- **Backups**: Automated backups daily (retention 35 days); point-in-time recovery (PITR) enabled for primary databases.
- **RPO/RTO**: RPO <= 5 minutes, RTO <= 30 minutes for Tier-1 systems; Tier-2 RPO <= 30 minutes, RTO <= 2 hours.
- **Data Quality Checks**: Validation pipelines with alerts on anomalies (e.g., negative inventory, orphaned orders).

## 8. Deployment & Release Management

- **Environment Strategy**: Dev → QA → Staging → Production with strict promotion criteria; staging mirrors production topology.
- **CI/CD**: Automated builds, tests, security scans; manual approval gates for production deploys.
- **Deployment Strategy**: Blue/green or canary rollout with automated rollback on SLO breach or error thresholds.
- **Release Cadence**: Weekly scheduled releases with emergency hotfix process (< 1 hour turnaround).

## 9. Testing & Quality Assurance

- **Testing Pyramid**: Unit tests (fast), integration/component tests, contract tests, end-to-end flows, non-functional tests (load, soak, security).
- **Coverage Requirements**: Critical services >= 80% unit coverage, >= 70% integration coverage, contract tests for all external APIs.
- **Shift-Left Security**: Threat modeling during design; security test cases in backlog; penetration testing twice per year.
- **Test Data Management**: Synthetic datasets; anonymized production copies only in guarded compliance-signed environments.

## 10. Operability & Support

- **Runbooks**: Required for each service (startup, shutdown, troubleshooting, dependencies).
- **Incident Response**: Define severity matrix; MTTA <= 5 minutes, MTTR <= 45 minutes for Sev-1 incidents.
- **Post-Incident Reviews**: Conduct blameless PIRs within 5 business days; track action items to closure.
- **Support Coverage**: 24x7 on-call rotation; follow-the-sun support for customer-facing issues.

## 11. Maintainability & Extensibility

- **Architecture Guardrails**: Domain-driven design; shared libraries for cross-cutting concerns; documented API contracts.
- **Code Quality**: Enforce linting, formatting, static analysis; PR reviews required (minimum two approvals for critical services).
- **Versioning**: Semantic versioning for services and APIs; backward compatibility maintained for at least two minor revisions.
- **Technical Debt**: Track in backlog with remediation SLAs; quarterly architecture review to assess debt items.

## 12. Scalability & Multi-Tenancy Readiness

- Ability to operate multiple tenants with logical isolation (tenant IDs propagated end-to-end).
- Configurable rate limits and quotas per tenant.
- Resource utilization dashboards per tenant for billing/chargeback.

## 13. Localization & Internationalization

- Support multiple locales and currencies; currency conversion refreshed hourly.
- Date/time handling via UTC internally with localized presentation.
- Content localization pipeline with fallback locales.

## 14. Environmental & Sustainability Goals

- Track infrastructure utilization and carbon footprint metrics quarterly.
- Prefer managed services with sustainability commitments; decommission unused resources automatically.

## 15. Change Control & Governance

- **Approvals**: Architectural changes reviewed by architecture board; security-sensitive changes require security team sign-off.
- **Configuration Management**: Immutable infrastructure via IaC; configuration changes tracked in Git with audit trail.
- **Documentation**: Update architecture diagrams, runbooks, and requirements documents alongside code changes.

## 16. Business Continuity & Disaster Recovery

- **DR Drills**: Bi-annual simulated regional failover; report recovery metrics and corrective actions.
- **Incident Communications**: Templates and playbooks for stakeholder updates (internal, customers, regulators).
- **Third-Party Dependencies**: Maintain vendor risk assessments; ensure SLAs and redundancy plans in vendor contracts.

## 17. Compliance Evidence & Auditing

- Automated evidence collection for SOC 2 (access reviews, change logs, backups, patching cadence).
- Quarterly internal audits; external audit annually.
- Centralized compliance calendar with control owners and due dates.

## 18. Document Lifecycle

- Review non-functional requirements bi-annually or upon major architectural change.
- Track revisions in Git (`docs/requirements/non-functional-enterprise.md`) with change log references.
