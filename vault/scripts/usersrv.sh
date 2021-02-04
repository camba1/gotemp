#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

# Export token used to login to vault
export VAULT_TOKEN=$1

# Create key value pair secrets
 vault kv put gotempkv/database/postgresql/usersrv username="postgres" password="TestDB@home2" application_name="userSrv" server="pgdb" dbname="appuser"

 vault kv put gotempkv/broker/nats/usersrv username="natsUser" password="TestDB@home2" server="nats"


 # Create Vault Policy
 vault policy write gotemp-usersrv - <<EOF
path "gotempkv/data/database/postgresql/usersrv" {
  capabilities = ["read"]
}

path "gotempkv/data/broker/nats/usersrv" {
  capabilities = ["read"]
}

EOF

# Create Vault K8s role to associate service account to the appropriate policy
vault write auth/kubernetes/role/gotemp-usersrv \
    bound_service_account_names=gotemp-usersrv \
    bound_service_account_namespaces=default \
    policies=gotemp-usersrv \
    ttl=24h