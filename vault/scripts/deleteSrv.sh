#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No service name supplied" >&2
    exit 1
fi

# Delete role
vault delete auth/kubernetes/role/gotemp-"$1"
# Delete policy
vault policy delete gotemp-"$1"

#Delete secrets (all versions and metadata)
vault kv metadata delete gotempkv/database/postgresql/"$1"
vault kv metadata delete gotempkv/database/redis/"$1"
vault kv metadata delete gotempkv/database/arangodb/"$1"
vault kv metadata delete gotempkv/database/timescaledb/"$1"
vault kv metadata delete gotempkv/broker/nats/"$1"