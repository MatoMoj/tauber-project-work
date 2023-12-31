package infrastructure

import (
	"github.com/google/uuid"
	"project-work-tauber/app/internal/models"
)

type PSQLBookingDao struct {
	cDao *PSQLCoreDao
}

func NewPSQLBookingDao(cDao *PSQLCoreDao) *PSQLBookingDao {
	return &PSQLBookingDao{cDao: cDao}
}

func (p *PSQLBookingDao) GetById(id uuid.UUID, credentials models.UserCredentials) (dto models.Booking, err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionGet)
	if err != nil {
		return dto, err
	}

	if err = db.First(&dto, id).Error; err != nil {
		return dto, err
	}

	return dto, nil
}

func (p *PSQLBookingDao) GetAll(credentials models.UserCredentials) (dto []models.Booking, err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionGet)
	if err != nil {
		return dto, err
	}

	if err = db.Find(&dto).Error; err != nil {
		return dto, err
	}

	return dto, nil
}

func (p *PSQLBookingDao) GetAllByCustomerID(customerID uuid.UUID, credentials models.UserCredentials) (bookings []models.Booking, err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionGet)
	if err != nil {
		return bookings, err
	}

	if err := db.Where("customer_id = ?", customerID).Find(&bookings).Error; err != nil {
		return bookings, err
	}

	return bookings, nil
}

func (p *PSQLBookingDao) Create(dto models.Booking, credentials models.UserCredentials) (id uuid.UUID, err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionCreate)
	if err != nil {
		return uuid.Nil, err
	}

	if err = db.Create(&dto).Error; err != nil {
		return uuid.Nil, err
	}

	return dto.ID, nil
}

func (p *PSQLBookingDao) Update(dto models.Booking, credentials models.UserCredentials) (err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionUpdate)
	if err != nil {
		return err
	}

	if err = db.Save(&dto).Error; err != nil {
		return err
	}

	return nil
}

func (p *PSQLBookingDao) Delete(id uuid.UUID, credentials models.UserCredentials) (err error) {
	db, err := p.cDao.GetConnection(credentials, p.cDao.Conf.DBConfig.PermissionDelete)
	if err != nil {
		return err
	}

	if err = db.Delete(&models.Booking{}, id).Error; err != nil {
		return err
	}

	return nil
}
