PROTO_DIR := proto
PROTO_SRC := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT := shared/pb

.PHONY: generate-proto
proto-all:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		--experimental_allow_proto3_optional \
		$(PROTO_SRC)

docker-compose-build:
	docker-compose -f infra/production/docker/docker-compose.yml --env-file .env -p "business-manager" up -d

SERVICES = auth-svc company-svc customer-svc file-svc gateway order-svc product-svc whatsapp-svc
REGISTRY = edorguez

.PHONY: build-all $(SERVICES) clean list

# Build all services sequentially
docker-image-build-all: $(SERVICES)
	@echo "All services built successfully!"

auth-svc:
	@echo "Building auth-svc..."
	docker build -t $(REGISTRY)/auth-svc -f infra/production/docker/auth-svc.Dockerfile .

company-svc:
	@echo "Building company-svc..."
	docker build -t $(REGISTRY)/company-svc -f infra/production/docker/company-svc.Dockerfile .

customer-svc:
	@echo "Building customer-svc..."
	docker build -t $(REGISTRY)/customer-svc -f infra/production/docker/customer-svc.Dockerfile .

file-svc:
	@echo "Building file-svc..."
	docker build -t $(REGISTRY)/file-svc -f infra/production/docker/file-svc.Dockerfile .

gateway:
	@echo "Building gateway..."
	docker build -t $(REGISTRY)/gateway -f infra/production/docker/gateway.Dockerfile .

order-svc:
	@echo "Building order-svc..."
	docker build -t $(REGISTRY)/order-svc -f infra/production/docker/order-svc.Dockerfile .

product-svc:
	@echo "Building product-svc..."
	docker build -t $(REGISTRY)/product-svc -f infra/production/docker/product-svc.Dockerfile .

whatsapp-svc:
	@echo "Building whatsapp-svc..."
	docker build -t $(REGISTRY)/whatsapp-svc -f infra/production/docker/whatsapp-svc.Dockerfile .
