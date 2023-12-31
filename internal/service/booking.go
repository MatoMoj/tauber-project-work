package service

import (
	"github.com/google/uuid"
	"project-work-tauber/app/cmd"
	"project-work-tauber/app/internal/dtos"
)

type Booking struct {
	dao cmd.BookingDao
}

func NewBooking(dao cmd.BookingDao) *Booking {
	return &Booking{dao: dao}
}

func (b *Booking) GetById(id uuid.UUID, credentials dtos.UserCredentials) (dto dtos.Booking, err error) {

	credMdl := ParseCredentials(credentials)

	bookingMdl, err := b.dao.GetById(id, credMdl)
	if err != nil {
		return
	}
	return bookingMdl.ToDTO(), nil
}

func (b *Booking) GetAll(credentials dtos.UserCredentials) (dtos []dtos.Booking, err error) {

	credMdl := ParseCredentials(credentials)

	bookingMdls, err := b.dao.GetAll(credMdl)
	if err != nil {
		return
	}

	for _, b := range bookingMdls {
		dtos = append(dtos, b.ToDTO())
	}

	return
}

func (b *Booking) Create(dto dtos.Booking, credentials dtos.UserCredentials) (id uuid.UUID, err error) {

	credMdl := ParseCredentials(credentials)
	bookingMdl := ParseBooking(dto)
	return b.dao.Create(bookingMdl, credMdl)
}

func (b *Booking) Update(dto dtos.Booking, credentials dtos.UserCredentials) (err error) {

	credMdl := ParseCredentials(credentials)
	bookingMdl := ParseBooking(dto)
	return b.dao.Update(bookingMdl, credMdl)

}

func (b *Booking) Delete(id uuid.UUID, credentials dtos.UserCredentials) (err error) {

	credMdl := ParseCredentials(credentials)
	return b.dao.Delete(id, credMdl)

}
