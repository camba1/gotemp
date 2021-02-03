path "gotempkv/data/database/postgresql/promotionsrv" {
  capabilities = ["read"]
}

path "gotempkv/data/broker/nats/promotionsrv" {
  capabilities = ["read"]
}

path "gotempkv/data/database/redis/promotionsrv" {
  capabilities = ["read"]
}