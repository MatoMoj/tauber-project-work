package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"project-work-tauber/app/cmd"
	"project-work-tauber/app/config"
	"project-work-tauber/app/internal/dtos"
)

type Booking struct {
	conf       config.AppConfig
	bookingSrv cmd.BookingService
	coreSrv    cmd.CoreService
}

func NewBooking(conf config.AppConfig, bookingSrv cmd.BookingService, coreSrv cmd.CoreService) *Booking {
	return &Booking{
		conf:       conf,
		bookingSrv: bookingSrv,
		coreSrv:    coreSrv,
	}
}

func (b *Booking) Create(c *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = b.coreSrv.CheckPermission(credentials, b.conf.DBConfig.PermissionCreate)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var dto dtos.Booking
	if err := c.BindJSON(&dto); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := b.bookingSrv.Create(dto, credentials)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (b *Booking) GetAll(c *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = b.coreSrv.CheckPermission(credentials, b.conf.DBConfig.PermissionGet)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	bookingDtos, err := b.bookingSrv.GetAll(credentials)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if len(bookingDtos) == 0 {
		bookingDtos = []dtos.Booking{}
	}

	c.JSON(http.StatusOK, bookingDtos)
}

func (b *Booking) GetById(c *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = b.coreSrv.CheckPermission(credentials, b.conf.DBConfig.PermissionGet)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	bookingIDStr := c.Param("id")
	bookingID, err := uuid.Parse(bookingIDStr)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid booking ID"})
		return
	}

	bookingDTO, err := b.bookingSrv.GetById(bookingID, credentials)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, bookingDTO)
}

func (b *Booking) Update(c *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = b.coreSrv.CheckPermission(credentials, b.conf.DBConfig.PermissionUpdate)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var bookingDto dtos.Booking
	if err := c.BindJSON(&bookingDto); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = b.bookingSrv.Update(bookingDto, credentials)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (b *Booking) Delete(c *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = b.coreSrv.CheckPermission(credentials, b.conf.DBConfig.PermissionDelete)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	bookingIDStr := c.Param("id")
	bookingID, err := uuid.Parse(bookingIDStr)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid booking ID"})
		return
	}

	err = b.bookingSrv.Delete(bookingID, credentials)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
