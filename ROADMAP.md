# Project Roadmap

This document tracks the high-level technical improvement tasks for the `api-go` project. Each item represents a significant step towards a more robust, maintainable, and performant application aligned with Go best practices and Hexagonal Architecture principles.

---

### Task List

- [ ] **1. Refactor: Centralize Error Handling with Middleware**
  - **Description:** Error handling is currently duplicated across HTTP handlers, coupling them to implementation details and HTTP status codes. This violates the separation of concerns.
  - **Goal:** Implement a central middleware to inspect domain errors returned from the core layer and map them to appropriate HTTP responses. This will clean up handlers and strengthen the architectural boundary.

- [ ] **2. Refactor: Adopt Environment Variables for Configuration**
  - **Description:** Configuration is managed via static JSON files, which is inflexible and not ideal for containerized environments.
  - **Goal:** Transition to a 12-Factor App approach by loading configuration from environment variables. This decouples the application from the filesystem and improves security.

- [ ] **3. Refactor: Simplify Database Adapter Abstraction**
  - **Description:** The current database adapter structure seems overly complex, suggesting a "leaky" abstraction that might expose underlying database details.
  - **Goal:** Consolidate and simplify the database adapter to ensure it provides a clean, impermeable boundary, completely hiding the database technology from the application's core.

- [ ] **4. Feat: Implement Context Propagation Through All Layers**
  - **Description:** `context.Context` is not consistently passed through all application layers.
  - **Goal:** Enforce `context.Context` propagation from the initial request down to the database. This is critical for managing timeouts, cancellation, and enabling distributed tracing, which are fundamental for a resilient, production-grade service.
