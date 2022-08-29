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

migrateup1:
	migrate -path db/migration -database "postgresql://root:yuay@localhost:5432/DBchain?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:yuay@localhost:5432/DBchain?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrate1:
	migrate create -ext sql -dir db/migration -seq add_session

mockdb:
	mockgen -destination db/mock/store.go github.com/aydogduyusuf/DBchain/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc migrateup1 migratedown1 migrateinit migrate1 mockdb
