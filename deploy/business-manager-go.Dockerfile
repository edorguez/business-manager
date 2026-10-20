# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS build

RUN apk add --no-cache git ca-certificates

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY shared ./shared
COPY services ./services

ARG SERVICE
RUN test -n "$SERVICE" && \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" \
      -o /out/app ./services/${SERVICE}/cmd

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=build /out/app /app/app

ENTRYPOINT ["/app/app"]
