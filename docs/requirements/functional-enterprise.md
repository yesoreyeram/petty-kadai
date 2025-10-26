# Enterprise Functional Requirements — Petty Kadai

> This document enumerates the enterprise-grade functional capabilities required to operate Petty Kadai ("Mini Amazon") as a production e-commerce platform. It is the canonical reference for delivery squads and solution architects.

## 1. Scope

- **In Scope**: End-to-end commerce journey (browse, search, cart, checkout, fulfilment), administrative controls, observability hooks, compliance enablers, customer support tooling, partner integrations.
- **Out of Scope**: Direct payment gateway build (use third-party), physical fulfilment execution (handled by 3rd-party logistics), advanced ML merchandising (roadmap), in-house fraud scoring (integrate vendor later).

## 2. Personas & Primary Journeys

| Persona                     | Goals                                              | Critical Journeys                                                            |
| --------------------------- | -------------------------------------------------- | ---------------------------------------------------------------------------- |
| Shopper                     | Discover products, purchase, receive notifications | Browse catalog, search, maintain cart, checkout, track order, manage returns |
| Seller (first-party)        | Publish inventory, manage pricing/promotions       | Create SKUs, upload images, adjust stock, view sales analytics               |
| Operations Analyst          | Monitor health, reconcile orders/payments          | View operational dashboards, export reports, audit transaction ledger        |
| Support Agent               | Resolve customer issues                            | Search orders/users, update order state within policy, trigger refunds       |
| Security/Compliance Officer | Maintain controls, respond to audits               | Review access logs, download compliance reports, approve policy changes      |

## 3. Core Functional Requirements

### 3.1 Account & Identity

1. Support email-and-password registration with verification workflow.
2. Provide OAuth 2.0 social sign-in (Google, Apple) with account linking.
3. Deliver secure session management via JWT access tokens and refresh tokens.
4. Offer MFA enrollment (TOTP) with enforcement rules by risk tier.
5. Allow account management (profile updates, password reset, session termination, GDPR data export/delete).

### 3.2 Catalog Management

1. Create and update products with versioned schema (title, description, price, tax class, inventory policy, media).
2. Manage hierarchical categories, facets, tags, and collections.
3. Support rich media assets (images, video) with CDN links and responsive renditions.
4. Allow sellers/ops to schedule product availability and pricing promotions.
5. Track audit history for all catalog mutations (who, when, what changed).

### 3.3 Search & Discovery

1. Full-text product search with typo tolerance and stemming.
2. Faceted navigation (category, price range, brand, rating).
3. Personalized/curated landing pages ("Featured", "Trending").
4. Autosuggest for top queries and categories.
5. Search analytics events (queries, zero-result counts, click-throughs).

### 3.4 Cart & Checkout

1. Persistent cart tied to user ID; guest cart persisted via signed cookie for 30 days.
2. Idempotent operations (add, update quantity, remove, clear).
3. Cart validation before checkout (stock, pricing, promotions, minimum order value).
4. Support discount codes, gift cards, loyalty credits with stackable rule engine.
5. Checkout workflow collecting shipping, billing, payment details with pluggable payment adapters.
6. Display real-time tax and shipping estimates via partner APIs.
7. Generate order summary, confirmation, and receipt artifacts.

### 3.5 Inventory & Order Orchestration

1. Maintain per-SKU stock levels with optimistic allocation and eventual reconciliation.
2. Reserve inventory during checkout; auto-release on timeout or failure.
3. Support multiple fulfilment locations and safety stock buffers.
4. Order lifecycle states: Pending → Authorized → Fulfilled → Delivered → Returned/Refunded.
5. Implement Saga orchestration with compensating actions (payment void, restock inventory, notify customer).
6. Provide manual override workflows (force fulfilment, cancel after ship) subject to role-based controls.

### 3.6 Payments & Billing

1. Integrate with third-party payment gateways via tokenization; never store raw PAN data.
2. Support payment methods: credit/debit cards, digital wallets, cash-on-delivery (configurable per region).
3. Handle partial authorizations, captures, and refunds.
4. Store financial ledger entries for reconciliation, supporting double-entry accounting model.
5. Automatically retry failed payments based on configurable retry schedule.

### 3.7 Notifications & Communications

1. Event-driven notifications (email, SMS, push) for order lifecycle, promotions, password reset.
2. Template management with localization support (en, es, ta, hi).
3. Opt-in/opt-out management per channel complying with regional regulations (CAN-SPAM, GDPR).
4. Provide delivery status tracking and bounce handling.

### 3.8 Customer Support Tools

1. Unified support console enabling search across users, orders, payments, tickets.
2. Role-scoped actions (resend email, issue refund, adjust loyalty points) with approval workflow.
3. Interaction history timeline per customer; notes and tags.
4. SLA tracking for ticket queues and escalations.

### 3.9 Reporting & Analytics

1. Operational dashboards (order volume, conversion rates, GMV, refund rate, inventory turnover).
2. Exportable datasets (CSV, secure API) for finance reconciliation and BI ingestion.
3. Event instrumentation for user behavior (page views, add-to-cart, checkout funnel drop-off).
4. Data retention policies (raw events 30 days hot, 365 days warm storage).

### 3.10 Administration & Governance

1. Role-based access control (RBAC) with predefined roles (Admin, Ops, Support, Analyst, Vendor) and custom roles.
2. Policy management (password rules, session timeout, IP safelist) with versioned configuration.
3. Change management workflow with approval and implementation logs.
4. Audit trail query UI and export capability.

## 4. External Integrations

- **Payments**: Stripe/Adyen for primary market; plug-in architecture for regional providers.
- **Shipping**: Integrate with logistics APIs for label generation, tracking updates, delivery estimates.
- **Tax**: Avatax or similar for automated tax calculation.
- **Email/SMS**: SendGrid/Twilio with fallback providers.
- **Analytics**: Stream events to warehouse (Snowflake/ClickHouse) via Kafka or managed pipeline.

## 5. Data & Domain Model Overview

- **Core Entities**: User, Address, Product, Variant, Category, InventoryItem, Cart, CartItem, Order, OrderLine, PaymentIntent, Refund, Promotion, NotificationEvent, AuditLog.
- **Associations**: User has many Orders; Order has many OrderLines; Product has many Variants; Variant references InventoryItems by location.
- **Identifiers**: Use ULID for globally sortable IDs; expose short IDs (base58) for UX display.
- **Versioning**: Store product revisions and order state transitions with immutable history.

## 6. API & Contract Expectations

- RESTful JSON APIs with consistent error envelopes (`trace_id`, `code`, `detail`).
- Standard headers: `X-Request-ID`, `X-User-ID`, `X-Tenant-ID` for multi-tenant readiness.
- Pagination via cursor-based strategy; limit maximum 200 items per page.
- Webhook support for order events, inventory changes, and refund updates with HMAC signatures.
- Documentation via OpenAPI 3.1; contract tests executed in CI.

## 7. Observability & Supportability Hooks

- Emit domain events for significant actions (user.created, order.placed, inventory.reserved).
- Integrate structured logging with correlation IDs.
- Provide admin endpoints for health, readiness, and dependency status.
- Expose SLA dashboards and alert rules for critical flows (auth, checkout, payment, order fulfillment).

## 8. Compliance & Regulatory Enablement

- GDPR/CCPA support: consent tracking, right-to-access/export/delete flows, data minimization.
- PCI-DSS scope reduction via tokenization and network segmentation.
- SOC 2 control coverage: change management, access reviews, monitoring, incident response.
- Accessibility: WCAG 2.1 AA compliance for shopper-facing UI.

## 9. Acceptance Criteria & Validation

- **Readiness Reviews**: Architecture, security, privacy, data governance sign-offs per major release.
- **Test Coverage**: Unit and integration tests >80% coverage on critical flows; contract tests for external integrations.
- **Operational Playbooks**: Runbooks for incident handling, rollback, and maintenance windows.
- **Pilot Launch**: Controlled cohort (internal + 100 beta users) achieving order success rate >= 95% and refund accuracy >= 99%.

## 10. Roadmap & Dependencies

1. **Phase Alpha (Weeks 1-6)**: User service baseline, catalog CRUD, frontend listings, naive inventory, initial observability.
2. **Phase Beta (Weeks 7-12)**: Order saga, payment integration, cart resiliency, gateway + JWT, security hardening.
3. **Phase GA (Weeks 13-18)**: Kubernetes deployment, analytics pipeline, DR drills, compliance audits.

## 11. Change Management

- Changes to this document require review by architecture board and product leadership.
- Maintain version history via Git (docs/requirements/functional-enterprise.md) with semantic tags (v1.0, v1.1, etc.).
