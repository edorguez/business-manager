PROTO_DIR=./proto
OUT_DIR_AUTH=./auth-svc/pkg/pb
OUT_DIR_COMPANY=./company-svc/pkg/pb
OUT_DIR_CUSTOMER=./customer-svc/pkg/pb
OUT_DIR_ORDER=./order-svc/pkg/pb
OUT_DIR_FILE=./file-svc/pkg/pb
OUT_DIR_WHATSAPP=./whatsapp-svc/pkg/pb
OUT_DIR_PRODUCT=./product-svc/pkg/pb
OUT_DIR_GATEWAY=./gateway/pkg/pb

# Install required tools
PROTOC_GEN_GO=google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.0
PROTOC_GEN_GO_GRPC=google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

tools:
	@echo "Installing tools..."
	go install $(PROTOC_GEN_GO)
	go install $(PROTOC_GEN_GO_GRPC)

# Generate gRPC code for company-svc
company-proto:
	@echo "Generating gRPC code for company-proto..."
	mkdir -p $(OUT_DIR_COMPANY)
	protoc --go_out=$(OUT_DIR_COMPANY) \
	       --go-grpc_out=$(OUT_DIR_COMPANY) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/company.proto

payment-proto:
	@echo "Generating gRPC code for payment-proto..."
	mkdir -p $(OUT_DIR_COMPANY)
	protoc --go_out=$(OUT_DIR_COMPANY) \
	       --go-grpc_out=$(OUT_DIR_COMPANY) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/payment.proto

payment-type-proto:
	@echo "Generating gRPC code for payment-type-proto..."
	mkdir -p $(OUT_DIR_COMPANY)
	protoc --go_out=$(OUT_DIR_COMPANY) \
	       --go-grpc_out=$(OUT_DIR_COMPANY) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/payment_type.proto

# Generate gRPC code for auth-svc
auth-proto:
	@echo "Generating gRPC code for auth-proto..."
	mkdir -p $(OUT_DIR_AUTH)
	protoc --go_out=$(OUT_DIR_AUTH) \
	       --go-grpc_out=$(OUT_DIR_AUTH) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/auth.proto

role-proto:
	@echo "Generating gRPC code for role-proto..."
	mkdir -p $(OUT_DIR_AUTH)
	protoc --go_out=$(OUT_DIR_AUTH) \
	       --go-grpc_out=$(OUT_DIR_AUTH) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/role.proto

user-proto:
	@echo "Generating gRPC code for user-proto..."
	mkdir -p $(OUT_DIR_AUTH)
	protoc --go_out=$(OUT_DIR_AUTH) \
	       --go-grpc_out=$(OUT_DIR_AUTH) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/user.proto

# Generate gRPC code for customer-svc
customer-proto:
	@echo "Generating gRPC code for customer-proto..."
	mkdir -p $(OUT_DIR_CUSTOMER)
	protoc --go_out=$(OUT_DIR_CUSTOMER) \
	       --go-grpc_out=$(OUT_DIR_CUSTOMER) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/customer.proto

# Generate gRPC code for product-svc
product-proto:
	@echo "Generating gRPC code for product-proto..."
	mkdir -p $(OUT_DIR_PRODUCT)
	protoc --go_out=$(OUT_DIR_PRODUCT) \
	       --go-grpc_out=$(OUT_DIR_PRODUCT) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/product.proto

# Generate gRPC code for order-svc
order-proto:
	@echo "Generating gRPC code for order-proto..."
	mkdir -p $(OUT_DIR_ORDER)
	protoc --go_out=$(OUT_DIR_ORDER) \
	       --go-grpc_out=$(OUT_DIR_ORDER) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/order.proto

# Generate gRPC code for file-svc
file-proto:
	@echo "Generating gRPC code for file-proto..."
	mkdir -p $(OUT_DIR_FILE)
	protoc --go_out=$(OUT_DIR_FILE) \
	       --go-grpc_out=$(OUT_DIR_FILE) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/file.proto

# Generate gRPC code for whatsapp-svc
whatsapp-proto:
	@echo "Generating gRPC code for whatsapp-proto..."
	mkdir -p $(OUT_DIR_WHATSAPP)
	protoc --go_out=$(OUT_DIR_WHATSAPP) \
	       --go-grpc_out=$(OUT_DIR_WHATSAPP) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/whatsapp.proto

# Generate gRPC code for Gateway
gateway-proto:
	@echo "Generating gRPC code for gateway..."
	mkdir -p $(OUT_DIR_GATEWAY)
	protoc --go_out=$(OUT_DIR_GATEWAY) \
	       --go-grpc_out=$(OUT_DIR_GATEWAY) \
		   --experimental_allow_proto3_optional \
	       $(PROTO_DIR)/company.proto \
	       $(PROTO_DIR)/auth.proto \
	       $(PROTO_DIR)/customer.proto \
	       $(PROTO_DIR)/order.proto \
	       $(PROTO_DIR)/file.proto \
	       $(PROTO_DIR)/payment_type.proto \
	       $(PROTO_DIR)/payment.proto \
	       $(PROTO_DIR)/product.proto \
	       $(PROTO_DIR)/role.proto \
	       $(PROTO_DIR)/user.proto

# Generate all proto files
proto-all: tools gateway-proto company-proto payment-proto payment-type-proto auth-proto role-proto user-proto customer-proto product-proto order-proto file-proto whatsapp-proto
	@echo "All services gRPC code generated successfully!"
