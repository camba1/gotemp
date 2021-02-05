// Create policies to allow the vault agent running in K8s to read the secrets

path "gotempkv/data/database/timescaledb/auditsrv"  {
  capabilities = ["read"]
}

path "gotempkv/data/broker/nats/auditsrv" {
  capabilities = ["read"]
}