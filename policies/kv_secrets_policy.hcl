# List, create, update, and delete secrets in the KV secrets engine
path "kv/data/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}

# Additional necessary permissions for KV Version 2
path "kv/metadata/*" {
  capabilities = ["list", "read", "delete"]
}

path "kv/destroy/*" {
  capabilities = ["update"]
}

path "kv/undelete/*" {
  capabilities = ["update"]
}

path "kv/versions/*" {
  capabilities = ["read"]
}