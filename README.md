# TAUBER Project Work IMS22

## Hashicorp Vault Integration for dynamic DB credentials from PSQL

This sample application uses Hashicorp Vault to generate dynamic Database credentials for a PSQL database. Users authenticate using HTTP Basic Auth (Username, Password), the application then authorizes against the Vault instance using the user's credentials to retrieve temporary database credentials for performing the desired operation.

Credentials for the database are never exposed to the user directly and only temporary created. This workflow enhances the general security and makes easy credential revocation for users possible without exposing the actual database credentials at any time.

## How to use

The database and vault server run as docker images. To configure them, use the .env file in folder "deploy". Afterwards run 
```bash
docker-compose up -d
```
to start the services.

The vault instance is automatically configured after startup using the shell script "vault_init.sh".

Start the go application and use the postman collection to send API requests.