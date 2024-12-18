# Stage 1: Dependency Cache
FROM golang:1.22.4-alpine AS deps

# Set the workspace directory
WORKDIR /workspace

# Copy only the specific go.work file and relevant go.mod files
COPY order-svc.go.work go.work
COPY order-svc/go.mod order-svc/go.mod
COPY customer-svc/go.mod customer-svc/go.mod

# Download dependencies for order-svc
WORKDIR /workspace/order-svc
RUN go mod tidy -e

# Download dependencies for customer-svc (if referenced by order-svc)
WORKDIR /workspace/customer-svc
RUN go mod tidy -e

# Stage 2: Build
FROM golang:1.22.4-alpine AS build

# Set the workspace directory
WORKDIR /workspace

# Copy the dependency cache and service files
COPY --from=deps /workspace /workspace
COPY order-svc/ order-svc/
COPY customer-svc/ customer-svc/

# Build the binary for order-svc
WORKDIR /workspace/order-svc
RUN go build -o ./out/dist ./cmd

# Stage 3: Runtime
FROM alpine:3.21.0

# Copy the compiled binary from the build stage
COPY --from=build /workspace/order-svc/out/dist /app/dist

# Copy the .env file to the runtime container
COPY order-svc/.env /app/.env

# Set working directory
WORKDIR /app

# Set the entrypoint
CMD ["/app/dist"]