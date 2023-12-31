package infrastructure

import (
	"github.com/google/uuid"
	"project-work-tauber/app/internal/models"
)

type PSQLCustomerDao struct {
	cDao *PSQLCoreDao
}

func NewPSQLCustomerDao(cDao *PSQLCoreDao) *PSQLCustomerDao {
	return &PSQLCustomerDao{cDao: cDao}
}

func (p *PSQLCustomerDao) GetById(id uuid.UUID, credentials models.UserCredentials) (model models.Customer, err error) {

	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionGet)
	if err != nil {
		return
	}

	if err = db.First(&model, id).Error; err != nil {
		return
	}

	return
}

func (p *PSQLCustomerDao) GetAll(credentials models.UserCredentials) (model []models.Customer, err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionGet)
	if err != nil {
		return nil, err
	}

	err = db.Find(&model).Error
	return model, err
}

func (p *PSQLCustomerDao) Create(dto models.Customer, credentials models.UserCredentials) (id uuid.UUID, err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionCreate)
	if err != nil {
		return uuid.Nil, err
	}

	if err = db.Create(&dto).Error; err != nil {
		return uuid.Nil, err
	}

	return dto.ID, nil
}

func (p *PSQLCustomerDao) Update(dto models.Customer, credentials models.UserCredentials) (err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionUpdate)
	if err != nil {
		return err
	}

	if err = db.Save(&dto).Error; err != nil {
		return err
	}

	return nil
}

func (p *PSQLCustomerDao) Delete(id uuid.UUID, credentials models.UserCredentials) (err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionDelete)
	if err != nil {
		return err
	}

	if err = db.Delete(&models.Customer{}, id).Error; err != nil {
		return err
	}

	return nil
}
