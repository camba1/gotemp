#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

echo "Creating promotionsrv artifacts...."

# Export token used to login to vault
export VAULT_TOKEN=$1

# Create key value pair secrets
vault kv put gotempkv/database/postgresql/promotionsrv username="postgres" password="TestDB@home2" application_name="promotionSrv" server="pgdb" dbname="postgres"

vault kv put gotempkv/broker/nats/promotionsrv username="natsUser" password="TestDB@home2" server="nats"

vault kv put gotempkv/database/redis/promotionsrv password="TestDB@home2" server="redis"

# Create Vault Policy
vault policy write gotemp-promotionsrv /vault/file/policies/promotionsrv.hcl


# Create Vault K8s role to associate service account to the appropriate policy
vault write auth/kubernetes/role/gotemp-promotionsrv \
    bound_service_account_names=gotemp-promotionsrv \
    bound_service_account_namespaces=default \
    policies=gotemp-promotionsrv \
    ttl=24h