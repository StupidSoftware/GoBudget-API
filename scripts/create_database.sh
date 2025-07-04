#!/bin/bash

DB_EXISTS=$(docker exec -t gobudget_postgres psql -U postgres -tAc "SELECT 1 FROM pg_database WHERE datname='gobudget'")

if [ "$DB_EXISTS" != "1" ]; then
  echo "Creating database 'gobudget'..."
  docker exec -it gobudget_postgres psql -U root -d postgres -c "CREATE DATABASE gobudget;"
else
  echo "Database 'gobudget' already exists."
fi
