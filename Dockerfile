# Stage 1: Builder
FROM golang:1.23 AS builder

WORKDIR /app

# Copia módulos e baixa deps
COPY go.mod go.sum ./
RUN go mod download

# Instala o tern COMO BINÁRIO ESTÁTICO
# (CGO_ENABLED=0 faz o Go gerar um EXE totalmente estático)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go install github.com/jackc/tern@latest

# Copia o resto do código e compila sua app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

# Stage 2: Runtime
FROM debian:bullseye-slim

WORKDIR /app

# Copia app e tern estático
COPY --from=builder /app/app .
COPY --from=builder /go/bin/tern /usr/bin/tern

# Copia migrations (path relativo a /app)
COPY --from=builder /app/internal/store/pgstore/migrations /app/migrations

# Script de startup
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

CMD ["/app/start.sh"]
