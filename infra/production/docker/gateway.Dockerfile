FROM golang:1.25.1 AS builder
WORKDIR /app
COPY . .
WORKDIR /app/services/gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o gateway ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/services/gateway/gateway .
CMD ["./gateway"]