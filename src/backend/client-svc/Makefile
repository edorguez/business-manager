run:
	go run cmd/main.go

postgres:
	docker run --name client-db -p 5434:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/pb/*.proto

new_migration:
	migrate create -ext sql -dir pkg/db/migration -seq $(name)