include .env

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

docker-compose-up-prod:
	docker-compose -f infra/production/docker/docker-compose.yml --env-file .env -p "business-manager" up -d

docker-compose-down-prod:
	docker-compose -f infra/production/docker/docker-compose.yml --env-file .env -p "business-manager" down

docker-compose-up-dev:
	docker-compose -f infra/development/docker/docker-compose.yml --env-file .env -p "business-manager" up -d

docker-compose-down-dev:
	docker-compose -f infra/development/docker/docker-compose.yml --env-file .env -p "business-manager" down

SERVICES = docker-auth-svc docker-company-svc docker-customer-svc docker-file-svc docker-gateway docker-order-svc docker-product-svc docker-whatsapp-svc
REGISTRY = edorguez

.PHONY: build-all $(SERVICES) clean list

# Build all services sequentially
docker-image-build-all: $(SERVICES)
	@echo "All services built successfully!"

docker-auth-svc:
	@echo "Building auth-svc..."
	docker build -t $(REGISTRY)/auth-svc -f infra/production/docker/auth-svc.Dockerfile .

docker-company-svc:
	@echo "Building company-svc..."
	docker build -t $(REGISTRY)/company-svc -f infra/production/docker/company-svc.Dockerfile .

docker-customer-svc:
	@echo "Building customer-svc..."
	docker build -t $(REGISTRY)/customer-svc -f infra/production/docker/customer-svc.Dockerfile .

docker-file-svc:
	@echo "Building file-svc..."
	docker build -t $(REGISTRY)/file-svc -f infra/production/docker/file-svc.Dockerfile .

docker-gateway:
	@echo "Building gateway..."
	docker build -t $(REGISTRY)/gateway -f infra/production/docker/gateway.Dockerfile .

docker-order-svc:
	@echo "Building order-svc..."
	docker build -t $(REGISTRY)/order-svc -f infra/production/docker/order-svc.Dockerfile .

docker-product-svc:
	@echo "Building product-svc..."
	docker build -t $(REGISTRY)/product-svc -f infra/production/docker/product-svc.Dockerfile .

docker-whatsapp-svc:
	@echo "Building whatsapp-svc..."
	docker build -t $(REGISTRY)/whatsapp-svc -f infra/production/docker/whatsapp-svc.Dockerfile .

# Sync Databases
migrateup-all: migrateup-auth-svc migrateup-company-svc migrateup-customer-svc migrateup-whatsapp-svc

migratedown-all: migratedown-auth-svc migratedown-company-svc migratedown-customer-svc migratedown-whatsapp-svc

migrateup-auth-svc:
	migrate -path services/auth-svc/pkg/db/migration -database ${AUTH_DB_SOURCE_DEVELOPMENT} --verbose up

migratedown-auth-svc:
	migrate -path services/auth-svc/pkg/db/migration -database ${AUTH_DB_SOURCE_DEVELOPMENT} --verbose down

migrateup-company-svc:
	migrate -path services/company-svc/pkg/db/migration -database ${COMPANY_DB_SOURCE_DEVELOPMENT} --verbose up

migratedown-company-svc:
	migrate -path services/company-svc/pkg/db/migration -database ${COMPANY_DB_SOURCE_DEVELOPMENT} --verbose down

migrateup-customer-svc:
	migrate -path services/customer-svc/pkg/db/migration -database ${CUSTOMER_DB_SOURCE_DEVELOPMENT} --verbose up

migratedown-customer-svc:
	migrate -path services/customer-svc/pkg/db/migration -database ${CUSTOMER_DB_SOURCE_DEVELOPMENT} --verbose down

migrateup-whatsapp-svc:
	migrate -path services/whatsapp-svc/pkg/db/migration -database ${WHATSAPP_DB_SOURCE_DEVELOPMENT} --verbose up

migratedown-whatsapp-svc:
	migrate -path services/whatsapp-svc/pkg/db/migration -database ${WHATSAPP_DB_SOURCE_DEVELOPMENT} --verbose down

