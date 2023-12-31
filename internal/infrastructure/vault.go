package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/api/auth/userpass"
	"project-work-tauber/app/config"
	"project-work-tauber/app/internal/models"
	"time"
)

// VaultPsqlCredentialProvider provides methods to retrieve credentials from Vault for PostgreSQL.
type VaultPsqlCredentialProvider struct {
	ctx  context.Context
	conf *config.VaultConfig
}

// NewVaultPsqlCredentialProvider creates a new instance of VaultPsqlCredentialProvider.
func NewVaultPsqlCredentialProvider(conf *config.VaultConfig) *VaultPsqlCredentialProvider {
	return &VaultPsqlCredentialProvider{
		ctx:  context.Background(),
		conf: conf,
	}
}

// getVaultSession retrieves a session from Vault using user credentials.
// It returns a session token or an error if the authentication fails.
func (v *VaultPsqlCredentialProvider) getVaultSession(credReq models.UserCredentials) (session *api.Secret, err error) {

	c := api.Config{
		Address: v.conf.ServerAddress,
	}

	authClient, err := api.NewClient(&c)
	if err != nil {
		return
	}

	userPassAuth, err := userpass.NewUserpassAuth(credReq.Username, &userpass.Password{
		FromString: credReq.Password,
	})
	if err != nil {
		return
	}

	session, err = userPassAuth.Login(v.ctx, authClient)
	if err != nil {
		return
	}

	return
}

// getDbSession retrieves database credentials from Vault using a valid session token.
// It returns the database credentials or an error if retrieval fails.
func (v *VaultPsqlCredentialProvider) getDbSession(session *api.Secret, permission string) (credentials models.DBCredentials, err error) {

	client, err := vault.New(
		vault.WithAddress(v.conf.ServerAddress),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		return
	}

	err = client.SetToken(session.Auth.ClientToken)
	if err != nil {
		return
	}

	path := fmt.Sprintf("%s/%s", v.conf.CredentialsBasePath, permission)

	secret, err := client.Read(v.ctx, path)
	if err != nil {
		return
	}

	username, usernameOk := secret.Data["username"].(string)
	password, passwordOk := secret.Data["password"].(string)
	if !usernameOk || !passwordOk {
		err = fmt.Errorf("credentials could not be retrieved from API response")
		return
	}

	credentials.Username = username
	credentials.Password = password

	return
}

// GetDBCredentials retrieves database credentials from Vault using user credentials and a permission string.
// It returns the database credentials or an error if retrieval fails.
func (v *VaultPsqlCredentialProvider) GetDBCredentials(credReq models.UserCredentials, permission string) (cred models.DBCredentials, err error) {

	vaultSession, err := v.getVaultSession(credReq)
	if err != nil {
		return
	}

	cred, err = v.getDbSession(vaultSession, permission)
	if err != nil {
		return
	}

	return
}

// getPermissionMapping retrieves the corresponding Vault permission based on the provided permission string.
// It returns the Vault permission name or an error if the permission is not found.
func (v *VaultPsqlCredentialProvider) getPermissionMapping(permission string) (vaultPermission string, err error) {

	for _, p := range v.conf.Permissions {
		if p.Permission == permission {
			return p.CredentialName, nil
		}
	}

	return "", errors.New(fmt.Sprintf("Permission %s, could not be found.", permission))
}
