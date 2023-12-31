package config

type AppConfig struct {
	DBConfig    DBConfig    `yaml:"DBConfig"`
	VaultConfig VaultConfig `yaml:"VaultConfig"`
}

type DBConfig struct {
	ConnectionString  string `yaml:"ConnectionString"`
	AdminUserName     string `yaml:"AdminUserName"`
	AdminUserPassword string `yaml:"AdminUserPassword"`
	PermissionAll     string `yaml:"PermissionAll"`
	PermissionGet     string `yaml:"PermissionGet"`
	PermissionCreate  string `yaml:"PermissionCreate"`
	PermissionUpdate  string `yaml:"PermissionUpdate"`
	PermissionDelete  string `yaml:"PermissionDelete"`
}

type VaultConfig struct {
	ServerAddress       string                  `yaml:"ServerAddress"`
	CredentialsBasePath string                  `yaml:"CredentialsBasePath"`
	Permissions         []VaultPermissionConfig `yaml:"Permissions"`
}

type VaultPermissionConfig struct {
	Permission     string `yaml:"Permission"`
	CredentialName string `yaml:"CredentialName"`
}
