package dtos

import (
	"github.com/google/uuid"
	"time"
)

type UserCredentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type Customer struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name" validate:"required"`
	Address    Address   `json:"address"`
	CampingVan Van       `json:"campingVan"`
	Bookings   []Booking `json:"bookings"`
}

type Address struct {
	Street     string `json:"street" validate:"required"`
	City       string `json:"city" validate:"required"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
}

type Van struct {
	LengthMeters float64 `json:"lengthMeters" validate:"required,gt=0"`
	HeightMeters float64 `json:"heightMeters" validate:"required,gt=0"`
	WidthMeters  float64 `json:"widthMeters" validate:"required,gt=0"`
}

type Booking struct {
	ID            uuid.UUID `json:"id"`
	CustomerID    uuid.UUID `json:"customerId" validate:"required"`
	BeginDate     time.Time `json:"beginDate" validate:"required"`
	EndDate       time.Time `json:"endDate" validate:"required,gtfield=BeginDate"`
	PriceEuros    int       `json:"priceEuros" validate:"required,gte=0"`
	AmountPersons int       `json:"amountPersons" validate:"required,gte=1"`
	Remarks       string    `json:"remarks"`
}
