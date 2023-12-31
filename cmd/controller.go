package cmd

import (
	"github.com/google/uuid"
	"project-work-tauber/app/internal/dtos"
)

type BookingService interface {
	GetById(id uuid.UUID, credentials dtos.UserCredentials) (dto dtos.Booking, err error)
	GetAll(credentials dtos.UserCredentials) (dtos []dtos.Booking, err error)
	Create(dto dtos.Booking, credentials dtos.UserCredentials) (id uuid.UUID, err error)
	Update(dto dtos.Booking, credentials dtos.UserCredentials) (err error)
	Delete(id uuid.UUID, credentials dtos.UserCredentials) (err error)
}

type CustomerService interface {
	GetById(id uuid.UUID, credentials dtos.UserCredentials) (dto dtos.Customer, err error)
	GetAll(credentials dtos.UserCredentials) (dtos []dtos.Customer, err error)
	Create(dto dtos.Customer, credentials dtos.UserCredentials) (id uuid.UUID, err error)
	Update(dto dtos.Customer, credentials dtos.UserCredentials) (err error)
	Delete(id uuid.UUID, credentials dtos.UserCredentials) (err error)
}

type CoreService interface {
	DBAutoMigrate() (err error)
	CheckPermission(credentials dtos.UserCredentials, permission string) (err error)
}
