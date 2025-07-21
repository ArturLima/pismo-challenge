#!/bin/sh
set -e

echo "🚀 Executando migrations..."

DATABASE_URL="postgres://${PISMO_DATABASE_USER}:${PISMO_DATABASE_PASSWORD}@${PISMO_DATABASE_HOST}:${PISMO_DATABASE_PORT}/${PISMO_DATABASE_NAME}?sslmode=disable"

tern migrate \
  --conn-string "$DATABASE_URL" \
  --migrations /app/migrations

echo "✅ Migrations executadas!"

echo "🚀 Iniciando aplicação..."
exec /app/app
