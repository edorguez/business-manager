FROM golang:1.25.1 AS builder
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
WORKDIR /app/services/file-svc
RUN CGO_ENABLED=0 GOOS=linux go build -o file-svc ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy only the binary from builder stage
COPY --from=builder /app/services/file-svc/file-svc .

CMD ["./file-svc"]
