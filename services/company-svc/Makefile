run:
	go run cmd/main.go

new_migration:
	migrate create -ext sql -dir pkg/db/migration -seq $(name)

migrateup:
	migrate -path pkg/db/migration -database "postgresql://postgres:postgres@127.0.0.1:5435/company-db?sslmode=disable" --verbose up

migrateup1:
	migrate -path pkg/db/migration -database "postgresql://postgres:postgres@127.0.0.1:5435/company-db?sslmode=disable" --verbose up 1

migratedown:
	migrate -path pkg/db/migration -database "postgresql://postgres:postgres@127.0.0.1:5435/company-db?sslmode=disable" --verbose down

migratedown1:
	migrate -path pkg/db/migration -database "postgresql://postgres:postgres@127.0.0.1:5435/company-db?sslmode=disable" --verbose down 1

sqlc:
	docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate
