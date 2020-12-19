#!/bin/bash
set -e


POSTGRES="psql --username postgres"

echo "Creating database: appuser"

$POSTGRES <<EOSQL
CREATE DATABASE appuser;
EOSQL

echo "Creating schema..."
psql -d appuser -a -U postgres -f ./docker-entrypoint-initdb.d/userSchema_sql.txt
