#!/bin/sh
set -e

echo "⏳ Aguardando o banco ficar pronto..."
until pg_isready -h "$PISMO_DATABASE_HOST" -p "$PISMO_DATABASE_PORT" -U "$PISMO_DATABASE_USER"; do
  sleep 2
done

echo "🚀 Executando migrations..."

DATABASE_URL="postgres://${PISMO_DATABASE_USER}:${PISMO_DATABASE_PASSWORD}@${PISMO_DATABASE_HOST}:${PISMO_DATABASE_PORT}/${PISMO_DATABASE_NAME}?sslmode=disable"

tern migrate \
  --conn-string "$DATABASE_URL" \
  --migrations /app/migrations

echo "✅ Migrations executadas!"
echo "🚀 Iniciando aplicação..."

exec /app/app
