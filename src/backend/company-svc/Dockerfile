# Stage 1: Dependency Cache
FROM golang:1.22.4-alpine AS deps

# Set the workspace directory
WORKDIR /workspace

# Copy only the  go.mod files
COPY company-svc/go.mod company-svc/go.mod

# Download dependencies for company-svc
WORKDIR /workspace/company-svc
RUN go mod tidy -e

# Stage 2: Build
FROM golang:1.22.4-alpine AS build

# Set the workspace directory
WORKDIR /workspace

# Install protoc and plugins
RUN apk add --no-cache protobuf
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copy the dependency cache and service files
COPY --from=deps /workspace /workspace
COPY company-svc/ company-svc/

# Build the binary for company-svc
WORKDIR /workspace/company-svc
RUN go build -o ./out/dist ./cmd

# Stage 3: Runtime
FROM alpine:3.21.0

# Copy the compiled binary from the build stage
COPY --from=build /workspace/company-svc/out/dist /app/dist

# Set working directory
WORKDIR /app

# Set the entrypoint
CMD ["/app/dist"]