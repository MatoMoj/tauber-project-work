package service

import (
	"github.com/google/uuid"
	"project-work-tauber/app/cmd"
	"project-work-tauber/app/internal/dtos"
)

type Customer struct {
	cusDao cmd.CustomerDao
	bkgDao cmd.BookingDao
}

func NewCustomer(cusDao cmd.CustomerDao, bkgDao cmd.BookingDao) *Customer {
	return &Customer{
		cusDao: cusDao,
		bkgDao: bkgDao,
	}
}

func (c *Customer) GetById(id uuid.UUID, credentials dtos.UserCredentials) (dto dtos.Customer, err error) {
	credMdl := ParseCredentials(credentials)
	customerMdl, err := c.cusDao.GetById(id, credMdl)
	if err != nil {
		return
	}
	bookingMdls, err := c.bkgDao.GetAllByCustomerID(customerMdl.ID, credMdl)
	if err != nil {
		return
	}
	return customerMdl.ToDTO(bookingMdls), nil
}

func (c *Customer) GetAll(credentials dtos.UserCredentials) (dto []dtos.Customer, err error) {
	credMdl := ParseCredentials(credentials)
	customerMdls, err := c.cusDao.GetAll(credMdl)
	if err != nil {
		return nil, err
	}
	var customers []dtos.Customer
	for _, customer := range customerMdls {
		bookingMdls, err := c.bkgDao.GetAllByCustomerID(customer.ID, credMdl)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer.ToDTO(bookingMdls))
	}
	return customers, nil
}

func (c *Customer) Create(dto dtos.Customer, credentials dtos.UserCredentials) (id uuid.UUID, err error) {
	credMdl := ParseCredentials(credentials)
	customerMdl := ParseCustomer(dto)
	return c.cusDao.Create(customerMdl, credMdl)
}

func (c *Customer) Update(dto dtos.Customer, credentials dtos.UserCredentials) (err error) {
	credMdl := ParseCredentials(credentials)
	customerMdl := ParseCustomer(dto)
	return c.cusDao.Update(customerMdl, credMdl)
}

func (c *Customer) Delete(id uuid.UUID, credentials dtos.UserCredentials) (err error) {
	credMdl := ParseCredentials(credentials)
	return c.cusDao.Delete(id, credMdl)
}
