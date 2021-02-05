// Create policies to allow the vault agent running in K8s to read the secrets

path "gotempkv/data/database/arangodb/productsrv" {
  capabilities = ["read"]
}

path "gotempkv/data/broker/nats/productsrv" {
  capabilities = ["read"]
}