#!/usr/bin/env sh

set -x

HOST=${HOST:-"psql"}
PORT=${PORT:-"5432"}
DB_NAME=${DB_NAME:-"booking_db"}
DB_USERNAME=${DB_USERNAME:-"vault"}
DB_PASSWORD=${DB_PASSWORD:-"secure"}

# Configure audit log to file
vault audit enable file file_path=/var/log/vault_audit.log

# Enable database plugin
vault secrets enable database

# Configure database connection for psql db
vault write database/config/"${DB_NAME}" \
    plugin_name="postgresql-database-plugin" \
    allowed_roles="db_read,db_create,db_update,db_delete,db_all_permissions" \
    connection_url="postgresql://${DB_USERNAME}:${DB_PASSWORD}@${HOST}:${PORT}/${DB_NAME}" \
    username="${DB_USERNAME}" \
    password="${DB_PASSWORD}" \
    password_authentication="scram-sha-256"

## Rotate root-user's credentials after initial setup so that the root password is no longer accessible
### Had to be removed as Vault locks out itself after updating the credentials...
#vault write -force database/rotate-root/"${DB_NAME}"

# Configure roles for database access
## Read
vault write database/roles/db_read \
    db_name="${DB_NAME}" \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"

## Create
vault write database/roles/db_create \
    db_name="${DB_NAME}" \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT INSERT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"

## Update
vault write database/roles/db_update \
    db_name="${DB_NAME}" \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT UPDATE ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"

## Delete
vault write database/roles/db_delete \
    db_name="${DB_NAME}" \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT DELETE ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"

## All Permissions
vault write database/roles/db_all_permissions \
    db_name="${DB_NAME}" \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT ALL PRIVILEGES ON SCHEMA public TO \"{{name}}\"; \
        GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"


# Create policies for roles
## Read
vault policy write db_read -<<EOF
path "database/creds/db_read" {
  capabilities = ["read"]
}
EOF

## Create
vault policy write db_create -<<EOF
path "database/creds/db_create" {
  capabilities = ["read"]
}
EOF

## Update
vault policy write db_update -<<EOF
path "database/creds/db_update" {
  capabilities = ["read"]
}
EOF

## Delete
vault policy write db_delete -<<EOF
path "database/creds/db_delete" {
  capabilities = ["read"]
}
EOF

## All Permissions
vault policy write db_all_permissions -<<EOF
path "database/creds/db_all_permissions" {
  capabilities = ["read"]
}
EOF


# Configure Access
## Enable Username/Password Authentication
vault auth enable userpass

## Create users
vault write auth/userpass/users/alice password="alice123" policies="db_read"
vault write auth/userpass/users/bob password="bob123" policies="db_read,db_create,db_update,db_delete"
vault write auth/userpass/users/admin password="admin" policies="db_all_permissions"