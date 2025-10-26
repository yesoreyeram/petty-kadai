# Enterprise Security Requirements — Petty Kadai

> This document formalizes the security controls, standards, and assurance activities required for Petty Kadai across all environments. It supplements functional and non-functional requirements with explicit security expectations.

## 1. Governance & Policy

- Maintain an information security management system (ISMS) aligned with ISO 27001.
- Publish and enforce security policies (access control, acceptable use, data classification, incident response, vendor management) with annual reviews.
- Assign a security steering committee responsible for risk management, policy approval, and exception tracking.
- Require security awareness training for all personnel annually; specialized secure coding training for engineers every 12 months.

## 2. Risk Management

- Conduct formal risk assessments bi-annually and after major architectural changes; document mitigation plans and residual risks.
- Maintain a centralized risk register with owners, mitigation status, and review cadence.
- Integrate threat modeling into design reviews for all new services and high-risk changes (STRIDE-based or equivalent).

## 3. Identity & Access Management (IAM)

- Enforce least-privilege, role-based access control (RBAC) with periodic entitlement reviews (quarterly for privileged roles).
- Mandate multi-factor authentication (MFA) for all administrative access (VPN, cloud consoles, CI/CD, observability tooling).
- Centralize identity federation via SSO (SAML/OIDC) for staff applications.
- Log all privileged actions and access escalations; retain audit logs for minimum 400 days.

## 4. Application Security

- Adopt secure coding standards (OWASP ASVS Level 2 baseline) and enforce via PR reviews and automated linting.
- Integrate application security testing (SAST, DAST, dependency scanning) into CI/CD; block releases on critical findings.
- Execute threat modeling, security design review, and security test planning before implementation of critical features.
- Require manual code review by at least two qualified reviewers for security-sensitive changes (auth, payments, crypto).
- Implement runtime protections: input validation, output encoding, CSRF tokens, rate limiting, and API abuse detection.

## 5. Secrets & Key Management

- Store secrets in a centralized vault with access scoped per service; prohibit secrets in source control or container images.
- Rotate credentials at least every 90 days; rotate asymmetric keys annually or upon suspected compromise.
- Protect encryption keys using hardware security modules (HSM) or cloud key management services (KMS); track provenance and usage.

## 6. Data Protection & Privacy

- Classify data assets (Public, Internal, Confidential, Restricted) and apply handling rules accordingly.
- Enforce encryption in transit (TLS 1.2+, strong ciphers) and at rest (AES-256) for all restricted/confidential data.
- Implement data minimization, pseudonymization, and anonymization where feasible; maintain data flow diagrams.
- Provide technical capabilities for privacy rights (access, delete, rectify, restrict processing) with SLA-driven fulfillment (<30 days).

## 7. Infrastructure Security

- Apply infrastructure as code (IaC) with security guardrails (linting, policy-as-code) before provisioning resources.
- Enforce network segmentation: separate production, staging, development; isolate sensitive services (payments, auth) in dedicated subnets.
- Configure firewalls/security groups with deny-by-default posture; restrict outbound traffic to approved destinations.
- Ensure host hardening (minimal base images, CIS benchmarks, automatic security updates) across compute workloads.
- Deploy endpoint detection and response (EDR) agents on all servers and workstations with centralized monitoring.

## 8. Vulnerability & Patch Management

- Perform automated vulnerability scanning (infrastructure and application) weekly; manual penetration testing twice per year.
- Remediate critical vulnerabilities within 7 days, high within 30 days; require documented exceptions for delays.
- Track third-party advisories; maintain SBOM (software bill of materials) for services and dependencies.

## 9. Monitoring & Detection

- Implement security information and event management (SIEM) ingesting logs from infrastructure, applications, IAM, and third-party SaaS providers.
- Define detection rules for account takeover, privilege escalation, data exfiltration, suspicious API usage, and anomalous network traffic.
- Conduct log integrity verification (hash/sign) and ensure retention meets compliance (>=400 days for security events).
- Use UEBA (user and entity behavior analytics) or equivalent anomaly detection for high-risk accounts.

## 10. Incident Response

- Maintain an incident response plan covering detection, analysis, containment, eradication, recovery, and post-incident review.
- Establish severity levels with communication playbooks for stakeholders, regulators, and customers.
- Conduct incident response tabletop exercises at least quarterly; document lessons learned and action items.
- Enable forensic readiness: capture required logs, maintain disk snapshots/backups, and protect chain-of-custody procedures.

## 11. Business Continuity & Disaster Recovery Security

- Ensure DR sites and failover environments enforce the same security controls as primary regions.
- Encrypt backup data and restrict access to DR personnel with MFA.
- Test DR failover scenarios with security validation (access controls, logging, monitoring) at least annually.

## 12. Third-Party & Supply Chain Security

- Perform security due diligence and risk scoring for vendors before onboarding; include security obligations in contracts.
- Require SOC 2/ISO 27001 reports or equivalent assurances; track remediation commitments.
- Monitor third-party integrations for anomalous behavior; enforce least-privilege API scopes and credential rotation.
- Maintain SBOM and supply-chain attestation (e.g., SLSA Level 2+ goals) for build pipelines.

## 13. Secure SDLC & DevSecOps

- Embed security checkpoints in each SDLC phase: ideation (threat modeling), development (code scanning), testing (pen tests), deployment (runtime checks), operations (monitoring).
- Use policy-as-code to enforce environment-specific security controls (e.g., Terraform Sentinel, Open Policy Agent).
- Automate dependency updates with security context (Dependabot/Renovate) and security regression tests before merge.

## 14. Compliance & Audit

- Align with PCI DSS SAQ A, SOC 2 Type II, and regional regulations (GDPR, CCPA, LGPD) for applicable services.
- Maintain evidence repositories for audits (access logs, change tickets, test results) with continuous updates.
- Conduct internal audits semi-annually; track non-conformities and remediation deadlines.

## 15. Physical & Operational Security

- Restrict physical access to offices and data centers with badge readers; monitor entry logs and video surveillance.
- Implement clean desk and secure disposal policies (shredding, media destruction).
- Provide secure remote access via VPN with device posture checks.

## 16. Metrics & Reporting

- Track key security metrics: time to detect/respond, vulnerability remediation SLAs, phishing simulation success, training completion, privileged access counts.
- Report security posture quarterly to executive leadership; include risk trends and remediation progress.

## 17. Continuous Improvement

- Maintain bug bounty or coordinated vulnerability disclosure program with defined response SLAs.
- Encourage security champions within engineering squads; hold monthly security clinics.
- Review and update this document whenever new threats emerge, major architecture changes occur, or annually at minimum.
