run:
	go run cmd/main.go

new_migration:
	migrate create -ext sql -dir pkg/db/migration -seq $(name)

migrateup:
	migrate -path pkg/db/migration -database "postgresql://postgres:postgres@127.0.0.1:5436/whatsapp-db?sslmode=disable" --verbose up

migratedown:
	migrate -path pkg/db/migration -database "postgresql://postgres:postgres@127.0.0.1:5436/whatsapp-db?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate