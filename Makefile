postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=yuay -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root DBchain

dropdb:
	docker exec -it postgres dropdb DBchain

migrateup:
	migrate -path db/migration -database "postgresql://root:yuay@localhost:5432/DBchain?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:yuay@localhost:5432/DBchain?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
