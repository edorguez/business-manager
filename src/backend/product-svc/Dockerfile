# Stage 1: Dependency Cache
FROM golang:1.22.4-alpine AS deps

# Set the workspace directory
WORKDIR /workspace

# Copy only the specific go.work file and relevant go.mod files
COPY product-svc.go.work go.work
COPY product-svc/go.mod product-svc/go.mod
COPY company-svc/go.mod company-svc/go.mod

# Download dependencies for product-svc
WORKDIR /workspace/product-svc
RUN go mod tidy -e

# Download dependencies for company-svc (if referenced by product-svc)
WORKDIR /workspace/company-svc
RUN go mod tidy -e

# Stage 2: Build
FROM golang:1.22.4-alpine AS build

# Set the workspace directory
WORKDIR /workspace

# Copy the dependency cache and service files
COPY --from=deps /workspace /workspace
COPY product-svc/ product-svc/
COPY company-svc/ company-svc/

# Build the binary for product-svc
WORKDIR /workspace/product-svc
RUN go build -o ./out/dist ./cmd

# Stage 3: Runtime
FROM alpine:3.21.0

# Copy the compiled binary from the build stage
COPY --from=build /workspace/product-svc/out/dist /app/dist

# Copy the .env file to the runtime container
COPY product-svc/.env /app/.env

# Set working directory
WORKDIR /app

# Set the entrypoint
CMD ["/app/dist"]