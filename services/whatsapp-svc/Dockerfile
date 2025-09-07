# Stage 1: Dependency Cache
FROM golang:1.22.4-alpine AS deps

# Set the workspace directory
WORKDIR /workspace

# Copy only the  go.mod files
COPY whatsapp-svc/go.mod whatsapp-svc/go.mod

# Download dependencies for whatsapp-svc
WORKDIR /workspace/whatsapp-svc
RUN go mod tidy -e

# Stage 2: Build
FROM golang:1.22.4-alpine AS build

# Set the workspace directory
WORKDIR /workspace

# Copy the dependency cache and service files
COPY --from=deps /workspace /workspace
COPY whatsapp-svc/ whatsapp-svc/

# Build the binary for whatsapp-svc
WORKDIR /workspace/whatsapp-svc
RUN go build -o ./out/dist ./cmd

# Stage 3: Runtime
FROM alpine:3.21.0

# Copy the compiled binary from the build stage
COPY --from=build /workspace/whatsapp-svc/out/dist /app/dist

# Copy the .env file to the runtime container
COPY whatsapp-svc/.env /app/.env

# Set working directory
WORKDIR /app

# Set the entrypoint
CMD ["/app/dist"]