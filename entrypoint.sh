#!/bin/sh

# Tunggu hingga database siap
until nc -z -v -w30 db 5432
do
  echo "Waiting for database connection..."
  sleep 1
done

# Jalankan migrasi database
echo "Running database migrations..."
migrate -path "./migrations" -database "$DATABASE_URL" -verbose up

# Jalankan testing
echo "Running tests..."
go test -v ./...

# Jalankan aplikasi
echo "Running App..."
exec ./main
