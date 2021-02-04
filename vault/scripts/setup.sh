#!/bin/sh

# Check if we got a vault token to be able to login
if [ "$#" -ne 1 ]
  then
    echo "No vault token supplied" >&2
    exit 1
fi

# Export token used to login to vault
export VAULT_TOKEN=$1

# Enable the secret engine at the correct path
vault secrets enable -path=gotempkv kv-v2

# Enable vault K8s auth method
vault auth enable kubernetes

# Configure Authentication using a service account token and its certificate
vault write auth/kubernetes/config \
    token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
    kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443" \
    kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt