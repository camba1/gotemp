// Create policies to allow the vault agent running in K8s to read the secrets

path "gotempkv/data/database/postgresql/usersrv" {
  capabilities = ["read"]
}

path "gotempkv/data/broker/nats/usersrv" {
  capabilities = ["read"]
}