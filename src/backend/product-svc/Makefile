run:
	go run cmd/main.go

mongo:
	docker run --name product-db -p 8082:8081 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -d mongo:tag

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/pb/*.proto --experimental_allow_proto3_optional
