PROTO_DIR=./proto
OUT_DIR_AUTH=./auth-svc/pkg/pb
OUT_DIR_COMPANY=./company-svc/pkg/pb
OUT_DIR_CUSTOMER=./customer-svc/pkg/pb
OUT_DIR_ORDER=./order-svc/pkg/pb
OUT_DIR_GATEWAY=./gateway/pkg/pb

# Install required tools
PROTOC_GEN_GO=google.golang.org/protobuf/cmd/protoc-gen-go@latest
PROTOC_GEN_GO_GRPC=google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

tools:
	@echo "Installing tools..."
	go install $(PROTOC_GEN_GO)
	go install $(PROTOC_GEN_GO_GRPC)

# Generate gRPC code for company-svc
company-svc:
	@echo "Generating gRPC code for company-svc..."
	mkdir -p $(OUT_DIR_COMPANY)
	protoc --go_out=$(OUT_DIR_COMPANY) \
	       --go-grpc_out=$(OUT_DIR_COMPANY) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/company.proto \

# Generate gRPC code for auth-svc
auth-svc:
	@echo "Generating gRPC code for auth-svc..."
	mkdir -p $(OUT_DIR_AUTH)
	protoc --go_out=$(OUT_DIR_AUTH) \
	       --go-grpc_out=$(OUT_DIR_AUTH) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/auth.proto

# Generate gRPC code for Gateway
gateway:
	@echo "Generating gRPC code for gateway..."
	mkdir -p $(OUT_DIR_GATEWAY)
	protoc --go_out=$(OUT_DIR_GATEWAY) \
	       --go-grpc_out=$(OUT_DIR_GATEWAY) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/company.proto \
	       $(PROTO_DIR)/auth.proto

# Generate for all
all: tools company-svc auth-svc gateway
	@echo "All services gRPC code generated successfully!"
