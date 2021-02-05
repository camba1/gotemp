// Create policies to allow the vault agent running in K8s to read the secrets

path "gotempkv/data/database/postgresql/promotionsrv" {
  capabilities = ["read"]
}

path "gotempkv/data/broker/nats/promotionsrv" {
  capabilities = ["read"]
}

path "gotempkv/data/database/redis/promotionsrv" {
  capabilities = ["read"]
}