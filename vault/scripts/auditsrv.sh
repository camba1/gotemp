#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

echo "Creating auditsrv artifacts...."

# Export token used to login to vault
export VAULT_TOKEN=$1

# Create key value pair secrets
vault kv put gotempkv/database/postgresql/auditsrv username="postgres" password="TestDB@home2" application_name="auditSrv" server="timescaledb" dbname="postgres"

vault kv put gotempkv/broker/nats/auditsrv username="natsUser" password="TestDB@home2" server="nats"

# Create Vault Policy
vault policy write gotemp-auditsrv - <<EOF
path "gotempkv/data/database/postgresql/auditsrv"  {
  capabilities = ["read"]
}

path "gotempkv/data/broker/nats/auditsrv" {
  capabilities = ["read"]
}
EOF

# Create Vault K8s role to associate service account to the appropriate policy
vault write auth/kubernetes/role/gotemp-auditsrv \
    bound_service_account_names=gotemp-auditsrv \
    bound_service_account_namespaces=default \
    policies=gotemp-auditsrv \
    ttl=24h