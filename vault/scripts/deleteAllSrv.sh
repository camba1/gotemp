#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

# Export token used to login to vault
export VAULT_TOKEN=$1

# Delete all services artifacts
/vault//file/scripts/deleteSrv.sh auditsrv
/vault//file/scripts/deleteSrv.sh customersrv
/vault//file/scripts/deleteSrv.sh productsrv
/vault//file/scripts/deleteSrv.sh promotionsrv
/vault//file/scripts/deleteSrv.sh usersrv