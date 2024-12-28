postgres:
	docker run --name postgres17 -p 8000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=8520 -d postgres:17-alpine
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:8520@localhost:8000/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:8520@localhost:8000/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc