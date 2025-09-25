FROM golang:1.25.1 AS builder
WORKDIR /app
COPY . .
WORKDIR /app/services/customer-svc
RUN CGO_ENABLED=0 GOOS=linux go build -o customer-svc ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/services/customer-svc/customer-svc .
CMD ["./customer-svc"]