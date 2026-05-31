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
.PHONY: postgres createdb dropdb migrateup migratedown