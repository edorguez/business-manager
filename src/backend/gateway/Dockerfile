# Stage 1: Dependency Cache
FROM golang:1.22.4-alpine AS deps

# Set the workspace directory
WORKDIR /workspace

# Copy only the specific go.work file and relevant go.mod files
COPY gateway.go.work go.work
COPY gateway/go.mod gateway/go.mod
COPY auth-svc/go.mod auth-svc/go.mod
COPY company-svc/go.mod company-svc/go.mod
COPY customer-svc/go.mod customer-svc/go.mod
COPY order-svc/go.mod order-svc/go.mod
COPY product-svc/go.mod product-svc/go.mod

# Copy .env file
COPY gateway/.env gateway/.env

# Download dependencies for gateway
WORKDIR /workspace/gateway
RUN go mod tidy -e

# Download dependencies for auth-svc (if referenced by gateway)
WORKDIR /workspace/auth-svc
RUN go mod tidy -e

# Download dependencies for company-svc (if referenced by gateway)
WORKDIR /workspace/company-svc
RUN go mod tidy -e

# Download dependencies for customer-svc (if referenced by gateway)
WORKDIR /workspace/customer-svc
RUN go mod tidy -e

# Download dependencies for order-svc (if referenced by gateway)
WORKDIR /workspace/order-svc
RUN go mod tidy -e

# Download dependencies for product-svc (if referenced by gateway)
WORKDIR /workspace/product-svc
RUN go mod tidy -e

# Stage 2: Build
FROM golang:1.22.4-alpine AS build

# Set the workspace directory
WORKDIR /workspace

# Copy the dependency cache and service files
COPY --from=deps /workspace /workspace
COPY gateway/ gateway/
COPY auth-svc/ auth-svc/
COPY company-svc/ company-svc/
COPY customer-svc/ customer-svc/
COPY order-svc/ order-svc/
COPY product-svc/ product-svc/

# Build the binary for gateway
WORKDIR /workspace/gateway
RUN go build -o ./out/dist ./cmd

# Stage 3: Runtime
FROM alpine:3.21.0

# Copy the compiled binary from the build stage
COPY --from=build /workspace/gateway/out/dist /app/dist

# Copy the .env file to the runtime container
COPY gateway/.env /app/.env

# Set working directory
WORKDIR /app

# Set the entrypoint
CMD ["/app/dist"]