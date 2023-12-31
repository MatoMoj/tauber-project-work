package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project-work-tauber/app/config"
	"project-work-tauber/app/internal/models"
	"strings"
)

type PSQLCoreDao struct {
	Conf  *config.AppConfig
	vault VaultPsqlCredentialProvider
}

func NewPSQLCoreDao(conf *config.AppConfig, vault VaultPsqlCredentialProvider) *PSQLCoreDao {
	return &PSQLCoreDao{
		Conf:  conf,
		vault: vault,
	}
}

func (p *PSQLCoreDao) getConnectionString(creds models.DBCredentials) (cs string) {
	cs = p.Conf.DBConfig.ConnectionString
	cs = strings.Replace(cs, "<username>", creds.Username, 1)
	cs = strings.Replace(cs, "<password>", creds.Password, 1)
	return
}

func (p *PSQLCoreDao) GetConnection(creds models.UserCredentials, permissionRaw string) (db *gorm.DB, err error) {

	permission, err := p.vault.getPermissionMapping(permissionRaw)
	if err != nil {
		return
	}

	dbCreds, err := p.vault.GetDBCredentials(creds, permission)
	if err != nil {
		return nil, err
	}

	cs := p.getConnectionString(dbCreds)
	db, err = gorm.Open(postgres.Open(cs), &gorm.Config{})
	return
}

func (p *PSQLCoreDao) DBAutoMigrate() (err error) {

	adminCreds := models.UserCredentials{
		Username: p.Conf.DBConfig.AdminUserName,
		Password: p.Conf.DBConfig.AdminUserPassword,
	}

	conn, err := p.GetConnection(adminCreds, p.Conf.DBConfig.PermissionAll)
	if err != nil {
		return err
	}

	err = conn.AutoMigrate(
		&models.Customer{},
		&models.Booking{},
	)

	return
}
