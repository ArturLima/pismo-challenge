FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Força saída do tern no local desejado
RUN GOBIN=/app/bin CGO_ENABLED=0 GOOS=linux go install github.com/jackc/tern@latest

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

FROM debian:bullseye-slim

WORKDIR /app

# Instala o cliente postgres para usar o pg_isready no entrypoint
RUN apt-get update && \
    apt-get install -y postgresql-client && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/app .
COPY --from=builder /app/bin/tern /usr/bin/tern
COPY --from=builder /app/internal/store/pgstore/migrations /app/migrations

COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

CMD ["/app/start.sh"]
