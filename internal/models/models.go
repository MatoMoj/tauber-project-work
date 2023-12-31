package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project-work-tauber/app/internal/dtos"
	"time"
)

type DBCredentials struct {
	Username string
	Password string
}

type UserCredentials struct {
	Username string
	Password string
}

func (uc *UserCredentials) ToDTO() dtos.UserCredentials {
	return dtos.UserCredentials{
		Username: uc.Username,
		Password: uc.Password,
	}
}

func (uc *UserCredentials) FromDTO(dto dtos.UserCredentials) {
	uc.Username = dto.Username
	uc.Password = dto.Password
}

type PSQLConfig struct {
	Username         string
	Password         string
	ConnectionString string
}

type Customer struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name       string    `gorm:"not null"`
	Address    Address   `gorm:"embedded;embeddedPrefix:address_"`
	CampingVan Van       `gorm:"embedded;embeddedPrefix:campingvan_"`
	Bookings   []Booking // Relationship for customer bookings
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}

func (c *Customer) ToDTO(bookingsMdls []Booking) dtos.Customer {

	var bookingsDtos []dtos.Booking

	for _, b := range bookingsMdls {
		bookingsDtos = append(bookingsDtos, b.ToDTO())
	}

	return dtos.Customer{
		ID:         c.ID,
		Name:       c.Name,
		Address:    c.Address.ToDTO(),
		CampingVan: c.CampingVan.ToDTO(),
		Bookings:   bookingsDtos,
	}
}

func (c *Customer) FromDTO(cusDto dtos.Customer, bookingDtos []dtos.Booking) {

	var bookingMdls []Booking

	for _, b := range bookingDtos {

		bkMdl := Booking{}
		bkMdl.FromDTO(b)
		bookingMdls = append(bookingMdls, bkMdl)
	}

	c.ID = cusDto.ID
	c.Name = cusDto.Name
	c.Address.FromDTO(cusDto.Address)
	c.CampingVan.FromDTO(cusDto.CampingVan)
	c.Bookings = bookingMdls
}

type Address struct {
	Street     string `gorm:"not null"`
	City       string `gorm:"not null"`
	Country    string `gorm:"not null"`
	PostalCode string `gorm:"not null"`
}

func (a *Address) ToDTO() dtos.Address {
	return dtos.Address{
		Street:     a.Street,
		City:       a.City,
		Country:    a.Country,
		PostalCode: a.PostalCode,
	}
}

func (a *Address) FromDTO(dto dtos.Address) {
	a.Street = dto.Street
	a.City = dto.City
	a.Country = dto.Country
	a.PostalCode = dto.PostalCode
}

type Van struct {
	LengthMeters float64 `gorm:"not null"`
	HeightMeters float64 `gorm:"not null"`
	WidthMeters  float64 `gorm:"not null"`
}

func (v *Van) ToDTO() dtos.Van {
	return dtos.Van{
		LengthMeters: v.LengthMeters,
		HeightMeters: v.HeightMeters,
		WidthMeters:  v.WidthMeters,
	}
}

func (v *Van) FromDTO(dto dtos.Van) {
	v.LengthMeters = dto.LengthMeters
	v.HeightMeters = dto.HeightMeters
	v.WidthMeters = dto.WidthMeters
}

type Booking struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerID    uuid.UUID `gorm:"type:uuid;index"`
	Customer      Customer  // Relationship to customer
	BeginDate     time.Time `gorm:"not null"`
	EndDate       time.Time `gorm:"not null"`
	PriceEuros    int       `gorm:"not null"`
	AmountPersons int       `gorm:"not null"`
	Remarks       string    `gorm:"size:255"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return nil
}

func (b *Booking) ToDTO() dtos.Booking {
	return dtos.Booking{
		ID:            b.ID,
		CustomerID:    b.CustomerID,
		BeginDate:     b.BeginDate,
		EndDate:       b.EndDate,
		PriceEuros:    b.PriceEuros,
		AmountPersons: b.AmountPersons,
		Remarks:       b.Remarks,
	}
}

func (b *Booking) FromDTO(dto dtos.Booking) {
	b.ID = dto.ID
	b.CustomerID = dto.CustomerID
	b.BeginDate = dto.BeginDate
	b.EndDate = dto.EndDate
	b.PriceEuros = dto.PriceEuros
	b.AmountPersons = dto.AmountPersons
	b.Remarks = dto.Remarks
}
