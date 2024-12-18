# Stage 1: Dependency Cache
FROM golang:1.22.4-alpine AS deps

# Set the workspace directory
WORKDIR /workspace

# Copy only the specific go.work file and relevant go.mod files
COPY auth-svc.go.work go.work
COPY auth-svc/go.mod auth-svc/go.mod
COPY company-svc/go.mod company-svc/go.mod

# Download dependencies for auth-svc
WORKDIR /workspace/auth-svc
RUN go mod tidy -e

# Download dependencies for company-svc (if referenced by auth-svc)
WORKDIR /workspace/company-svc
RUN go mod tidy -e

# Stage 2: Build
FROM golang:1.22.4-alpine AS build

# Set the workspace directory
WORKDIR /workspace

# Copy the dependency cache and service files
COPY --from=deps /workspace /workspace
COPY auth-svc/ auth-svc/
COPY company-svc/ company-svc/

# Build the binary for auth-svc
WORKDIR /workspace/auth-svc
RUN go build -o ./out/dist ./cmd

# Stage 3: Runtime
FROM alpine:3.21.0

# Copy the compiled binary from the build stage
COPY --from=build /workspace/auth-svc/out/dist /app/dist

# Set working directory
WORKDIR /app

# Set the entrypoint
CMD ["/app/dist"]
