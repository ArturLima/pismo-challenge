FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux       go install github.com/jackc/tern@latest

COPY . .
RUN CGO_ENABLED=0 GOOS=linux       go build -o app ./cmd/api

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=0 /app/app .
COPY --from=0 /go/bin/tern /usr/bin/tern

COPY --from=0 /app/internal/store/pgstore/migrations /app/migrations

COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

CMD ["/app/start.sh"]
