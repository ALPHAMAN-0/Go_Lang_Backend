DB_URL=postgresql://admin:secret@localhost:5432/myapp?sslmode=disable

postgres:
	docker run --name myapp_db -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -d postgres:15

createdb:
	docker exec -it myapp_db createdb --username=admin --owner=admin myapp

dropdb:
	docker exec -it myapp_db dropdb --username=admin myapp

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

# Run ALL tests with verbose output
test:
	go test -v -cover ./db/sqlc/...

# Run only User tests
test-user:
	go test -v -run TestUser ./db/sqlc/...

# Run only Phone tests
test-phone:
	go test -v -run TestPhone ./db/sqlc/...

# Run a specific test by name
# Usage: make test-one NAME=TestCreateUser
test-one:
	go test -v -run $(NAME) ./db/sqlc/...

# Run tests with race condition detection
test-race:
	go test -v -race ./db/sqlc/...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test test-user test-phone test-one test-race
