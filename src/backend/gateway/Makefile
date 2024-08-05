run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/customer/pb/customer.proto --experimental_allow_proto3_optional
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/company/pb/company.proto --experimental_allow_proto3_optional
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/company/pb/payment.proto --experimental_allow_proto3_optional
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/company/pb/payment_type.proto --experimental_allow_proto3_optional
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/product/pb/product.proto --experimental_allow_proto3_optional
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/auth/pb/user.proto --experimental_allow_proto3_optional
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/auth/pb/auth.proto --experimental_allow_proto3_optional
