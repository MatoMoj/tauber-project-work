package service

import (
	"project-work-tauber/app/cmd"
	"project-work-tauber/app/internal/dtos"
	"project-work-tauber/app/internal/models"
)

type Core struct {
	dao cmd.CoreDao
}

func NewCore(dao cmd.CoreDao) *Core {
	return &Core{dao: dao}
}

func (c *Core) DBAutoMigrate() (err error) {

	return c.dao.DBAutoMigrate()
}

func (c *Core) CheckPermission(credentials dtos.UserCredentials, permission string) (err error) {

	credMdl := models.UserCredentials{}
	credMdl.FromDTO(credentials)

	_, err = c.dao.GetConnection(credMdl, permission)
	return
}

func ParseBooking(dto dtos.Booking) models.Booking {

	bookingMdl := models.Booking{}
	bookingMdl.FromDTO(dto)
	return bookingMdl
}

func ParseCustomer(dto dtos.Customer) models.Customer {
	customerMdl := models.Customer{}
	customerMdl.FromDTO(dto, nil)
	return customerMdl
}

func ParseCredentials(credentials dtos.UserCredentials) models.UserCredentials {
	credMdl := models.UserCredentials{}
	credMdl.FromDTO(credentials)
	return credMdl
}
