## SQLC Learning Notes

SQLC is a compile-time code generator for Go database access. Instead of writing manual `database/sql` boilerplate, you write SQL first and let SQLC generate type-safe Go methods and structs.

This example project uses SQLC for two tables:

- `users` with `id`, `name`, `email`, and `created_at`
- `phones` with `id`, `user_id`, `phone_number`, and `created_at`

The `phones.user_id` column references `users.id`, so a phone belongs to a user and is removed automatically when the user is deleted.

### How the project is organized

- `db/migration/` contains the schema migrations
- `db/query/` contains SQLC query definitions
- `db/sqlc/` contains the generated Go code
- `sqlc.yml` tells SQLC where to find queries, schema, and generated output
- `makefile` provides shortcuts for migrations and code generation

### SQLC workflow

1. Write or update the schema in `db/migration/`.
2. Add SQL statements in `db/query/*.sql`.
3. Run `sqlc generate`.
4. Use the generated methods from `db/sqlc` in your Go code.

SQLC reads query annotations like `-- name: CreateUser :one` and turns them into strongly typed methods. The suffix defines the result shape:

- `:one` returns a single row
- `:many` returns a slice of rows
- `:exec` runs a statement without returning rows

### Generated API

From the query files in this project, SQLC generates methods such as:

- `CreateUser`
- `GetUser`
- `ListUsers`
- `UpdateUserName`
- `DeleteUser`
- `CreatePhone`
- `GetPhone`
- `UpdatePhone`
- `DeletePhone`

It also generates Go structs for the tables, such as `User` and `Phone`, plus a `Querier` interface that makes the database layer easy to mock in tests.

### Configuration

The `sqlc.yml` file is the main configuration for generation:

- `engine: postgresql` tells SQLC which database dialect to use
- `queries: ./db/query/` points to the SQL files
- `schema: ./db/migration/` points to the migration files
- `out: ./db/sqlc` sets the output folder for generated code
- `emit_json_tags: true` adds JSON tags to generated structs
- `emit_interface: true` generates the `Querier` interface
- `emit_empty_slices: true` returns empty slices instead of `nil` for list queries
- `emit_pointers_for_null_types: true` uses pointer-friendly nullable types where needed

### Commands

Use the `makefile` targets during development:

- `make postgres` starts a Postgres container
- `make createdb` creates the application database
- `make migrateup` applies migrations
- `make migratedown` rolls back migrations
- `make sqlc` regenerates Go code from the SQL files

The database URL is defined as:

`postgresql://admin:secret@localhost:5432/myapp?sslmode=disable`

### Migration note

To create a new migration pair, use:

`migrate create -ext sql -dir db/migration -seq init_schema`

This creates the `up` and `down` migration files that SQLC reads as the schema source.

### Why SQLC fits this project

SQLC is a good fit here because it keeps SQL explicit while still giving you:

- compile-time validation of queries
- generated types instead of manual row scanning
- less boilerplate than raw `database/sql`
- better safety than hand-written query code

In short, this project uses SQLC to get the performance of SQL with much better developer ergonomics.

