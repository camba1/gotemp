// Create policies to allow the vault agent running in K8s to read the secrets

path "gotempkv/data/database/arangodb/customersrv" {
  capabilities = ["read"]
}

path "gotempkv/data/broker/nats/customersrv" {
  capabilities = ["read"]
}
