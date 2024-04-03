run:
	go run cmd/main.go

postgres:
	docker run --name company-db -p 5435:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/pb/*.proto --experimental_allow_proto3_optional

new_migration:
	migrate create -ext sql -dir pkg/db/migration -seq $(name)

migrateup:
	migrate -path pkg/db/migration -database "postgresql://postgres:postgres@localhost:5435/company-db?sslmode=disable" --verbose up

migratedown:
	migrate -path pkg/db/migration -database "postgresql://postgres:postgres@localhost:5435/company-db?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate
