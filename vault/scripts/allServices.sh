#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

# Create all services artifacts
/vault/file/scripts/auditsrv.sh "$1"
/vault/file/scripts/customersrv.sh "$1"
/vault/file/scripts/productsrv.sh "$1"
/vault/file/scripts/promotionsrv.sh "$1"
/vault/file/scripts/usersrv.sh "$1"