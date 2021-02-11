#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

# Export token used to login to vault
export VAULT_TOKEN=$1

# Disable the secret engine at the correct path
vault secrets disable gotempkv

# Disable vault K8s auth method
vault auth disable kubernetes