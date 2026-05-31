## Database Access Patterns in Go: A Comprehensive Comparison

This guide compares four popular approaches for implementing CRUD operations in Go applications, each with distinct tradeoffs in performance, safety, and developer experience.

---

### 1. Database/SQL (Standard Library)

The `database/sql` package provides low-level, direct database access with minimal abstraction.

**Advantages:**
- **High performance** — Queries execute at near-native database speed with minimal overhead
- **Full control** — Direct manipulation of SQL execution and result mapping
- **Lightweight** — No external dependencies; part of the Go standard library
- **Transparent** — Query behavior is explicit and predictable

**Disadvantages:**
- **Manual mapping** — Results must be manually mapped to Go variables, creating boilerplate
- **Error-prone** — Field mapping errors and type mismatches are difficult to catch and debug
- **Runtime failures** — Errors are discovered only during execution, not compile-time
- **Not scalable** — Complex queries and large data models require significant manual work

**Best for:** Simple applications with straightforward queries where performance is critical.

---

### 2. GORM (ORM Framework)

GORM is a full-featured Object-Relational Mapping library that abstracts database operations.

**Advantages:**
- **Pre-built CRUD** — Create, Read, Update, Delete operations are fully implemented
- **Automatic mapping** — Struct fields are automatically mapped to database columns
- **Type safety** — Compile-time checks catch many errors before runtime
- **Rich features** — Includes migrations, hooks, query builders, and associations
- **Rapid development** — Significantly reduces boilerplate and accelerates feature delivery

**Disadvantages:**
- **Performance overhead** — Abstraction layer adds latency compared to raw SQL
- **Opaque queries** — Query optimization becomes difficult to reason about
- **Inefficient generation** — May generate suboptimal SQL for complex operations
- **Learning curve** — Requires understanding ORM patterns and conventions

**Best for:** Rapid application development where productivity matters more than raw performance.

---

### 3. SQLX (Enhanced Standard Library)

SQLX extends `database/sql` with convenient features while maintaining the SQL-first approach.

**Advantages:**
- **Struct mapping** — Automatic mapping of query results to struct fields reduces boilerplate
- **Near-native performance** — Minimal overhead compared to raw `database/sql`
- **Familiar API** — Extends standard library with intuitive extensions
- **Flexibility** — Write SQL directly while benefiting from convenient helper functions

**Disadvantages:**
- **Runtime failures** — Errors in SQL queries are caught only during execution, not before
- **Manual query management** — Still requires writing and maintaining SQL strings
- **Limited features** — No built-in migrations or relationship management
- **Type assertions** — Requires careful handling of Go type conversions

**Best for:** Projects needing better ergonomics than raw `database/sql` without full ORM overhead.

---

### 4. SQLC (Compile-Time Code Generation)

SQLC generates type-safe Go code from SQL queries at compile-time.

**Advantages:**
- **High performance** — Generated code is as fast as raw SQL with zero runtime overhead
- **Type safe** — SQL query errors are caught during code generation, not runtime
- **Automatic code generation** — Eliminates boilerplate by generating Go functions from SQL
- **SQL-first** — You write standard SQL; SQLC handles the Go integration
- **Compile-time validation** — Query errors surface immediately during development

**Disadvantages:**
- **Build step required** — Adds a code generation phase to the build pipeline
- **Learning curve** — Requires understanding SQL and the SQLC configuration
- **Less dynamic** — Not suitable for dynamically constructed queries
- **Limited features** — Focuses on queries only; no built-in migrations or advanced ORM features

**Best for:** Performance-critical applications where type safety and code generation are valued.

---

## Recommendation for This Project

**SQLC** is the preferred choice because it combines:
- **Safety**: Compile-time validation catches SQL errors before runtime
- **Performance**: Zero abstraction overhead; generated code runs at native speed
- **Developer experience**: Automatic code generation eliminates boilerplate while maintaining explicit SQL control
- **Maintainability**: Clear, type-safe generated functions make code reviews and debugging straightforward

If rapid prototyping is prioritized, **GORM** is an acceptable alternative. For maximum flexibility, **SQLX** offers a middle ground between simplicity and control.
