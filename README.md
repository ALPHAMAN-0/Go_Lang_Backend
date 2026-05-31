# Go Learning Roadmap

A practical, trackable 12-week plan to learn Go (Golang) with hands-on projects and resources.

## How to use this file
- Use the task lists below to mark your progress. On GitHub you can check boxes directly; locally edit this file and change `- [ ]` to `- [x]` and commit.
- For local commits, run:

```bash
git add README.md
git commit -m "Update learning progress"
git push
```

## 12-Week Roadmap (high level)

- [ ] Week 1–2 — Foundations: setup, `go` command, variables, types, Hello World
- [ ] Week 3–4 — Core Language: arrays, slices, maps, structs, functions, pointers
- [ ] Week 5 — Methods & Interfaces: interfaces, type assertions, embedding
- [ ] Week 6 — Concurrency: goroutines, channels, `select`, context, race detector
- [ ] Week 7 — Stdlib & Toolchain: `net/http`, `encoding/json`, `go fmt`, `go vet`
- [ ] Week 8 — Testing & Benchmarking: `testing`, table-driven tests, `httptest`, coverage
- [ ] Week 9 — Databases: `database/sql`, `sqlx`/`sqlc`, migrations, integration tests
- [ ] Week 10 — Web APIs & Ecosystem: REST best practices, middleware, auth, gRPC basics
- [ ] Week 11 — Performance & Deployment: `pprof`, tracing, Docker, basic Kubernetes
- [ ] Week 12 — Advanced Topics & Capstone: generics, reflection, CGO, deploy capstone

## Practice Projects (progress)

- [ ] CLI utilities: flags, file I/O
- [ ] In-memory CRUD service (users/phones) with JSON
- [ ] HTTP REST API with tests and graceful shutdown
- [ ] DB-backed API (Postgres + sqlc or GORM) with migrations
- [ ] Concurrent worker pool and job processing
- [ ] Capstone: Todo or Expense Tracker API with auth, CI, Docker

## Measurable Milestones

- [ ] Working JSON CRUD + tests (Week 4)
- [ ] Concurrency demo + race-free (Week 6)
- [ ] DB-backed API with integration tests (Week 9)
- [ ] Capstone deployed + CI (Week 12)

## Resources

- Official Go docs: https://go.dev/doc
- A Tour of Go: https://tour.golang.org
- Go by Example: https://gobyexample.com
- Effective Go: https://go.dev/doc/effective_go
- sqlc: https://sqlc.dev
- gin: https://github.com/gin-gonic/gin

## Daily / Session Checklist

- Read one focused topic (30–60 min)
- 45–90 min coding exercise or tests
- Run linters and tests (`go test`, `go vet`, `golangci-lint`)
- Quick notes: what broke, how fixed (15 min)

## Next steps

- Check the boxes as you progress. Want me to create a weekly tracker file or GitHub Actions to enforce tests? Reply and I'll add it.
# Go_Lang_Backend

## Database Schema

Use [dbdiagram.io](https://dbdiagram.io/) to create and maintain the database schema for this project.

Add your schema diagram or DBML definition here so it stays close to the codebase and is easy to update.