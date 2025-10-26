# Week 09 — CDN-style caching (one concept/tool config)

tools-introduced: Nginx caching (as CDN simulator)

concepts-covered:

- Edge caching for static media; cache keys; cache-control headers

proposed-architecture:

- Nginx serves images from MinIO with caching; respects Cache-Control

changes-to-system-design:

- Add caching location; set default TTL; purge strategy for updates (manual for now)

tasks-checklist:

- [ ] Configure Nginx cache path and keys
- [ ] Add headers in Catalog responses for images
- [ ] Verify cache HIT/MISS in logs
- [ ] Document cache busting approach

skills-required:

- Nginx caching directives; HTTP caching semantics

prerequisites:

- Weeks 01–08 running

deliverables:

- Image loads are cached at edge; subsequent requests faster

acceptance-criteria:

- First request MISS, subsequent HIT; measurable latency drop

## Proposed architecture diagram

```mermaid
flowchart LR
  UI[Next.js Web]
  CDN[Edge Cache (Nginx)]
  MINIO[(MinIO)]

  UI --> CDN --> MINIO
```
