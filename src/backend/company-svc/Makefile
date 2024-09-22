run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/pb/*.proto --experimental_allow_proto3_optional

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
