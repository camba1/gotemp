#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

echo "Creating customersrv artifacts...."

# Export token used to login to vault
export VAULT_TOKEN=$1

# Create key value pair secrets
vault kv put gotempkv/database/arangodb/customersrv username="customerUser" password="TestDB@home2" server="arangodb:8529"

vault kv put gotempkv/broker/nats/customersrv username="natsUser" password="TestDB@home2" server="nats"

# Create Vault Policy
vault policy write gotemp-customersrv /vault/file/policies/customersrv.hcl


# Create Vault K8s role to associate service account to the appropriate policy
vault write auth/kubernetes/role/gotemp-customersrv \
    bound_service_account_names=gotemp-customersrv \
    bound_service_account_namespaces=default \
    policies=gotemp-customersrv \
    ttl=24h

