package cmd

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project-work-tauber/app/internal/models"
)

type BookingDao interface {
	GetById(id uuid.UUID, credentials models.UserCredentials) (mdl models.Booking, err error)
	GetAll(credentials models.UserCredentials) (mdls []models.Booking, err error)
	GetAllByCustomerID(customerID uuid.UUID, credentials models.UserCredentials) (bookings []models.Booking, err error)
	Create(mdl models.Booking, credentials models.UserCredentials) (id uuid.UUID, err error)
	Update(mdl models.Booking, credentials models.UserCredentials) (err error)
	Delete(id uuid.UUID, credentials models.UserCredentials) (err error)
}

type CustomerDao interface {
	GetById(id uuid.UUID, credentials models.UserCredentials) (mdl models.Customer, err error)
	GetAll(credentials models.UserCredentials) (mdls []models.Customer, err error)
	Create(mdl models.Customer, credentials models.UserCredentials) (id uuid.UUID, err error)
	Update(mdl models.Customer, credentials models.UserCredentials) (err error)
	Delete(id uuid.UUID, credentials models.UserCredentials) (err error)
}

type CoreDao interface {
	DBAutoMigrate() (err error)
	GetConnection(creds models.UserCredentials, permissionRaw string) (db *gorm.DB, err error)
}
