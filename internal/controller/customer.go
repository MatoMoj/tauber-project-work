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

type Customer struct {
	conf        config.AppConfig
	customerSrv cmd.CustomerService
	coreSrv     cmd.CoreService
}

func NewCustomer(conf config.AppConfig, customerSrv cmd.CustomerService, coreSrv cmd.CoreService) *Customer {
	return &Customer{
		conf:        conf,
		customerSrv: customerSrv,
		coreSrv:     coreSrv,
	}
}

func (c *Customer) Create(ctx *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = c.coreSrv.CheckPermission(credentials, c.conf.DBConfig.PermissionCreate)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var dto dtos.Customer
	if err := ctx.BindJSON(&dto); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := c.customerSrv.Create(dto, credentials)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (c *Customer) GetAll(ctx *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = c.coreSrv.CheckPermission(credentials, c.conf.DBConfig.PermissionGet)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	customerDtos, err := c.customerSrv.GetAll(credentials)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if len(customerDtos) == 0 {
		customerDtos = []dtos.Customer{}
	}

	ctx.JSON(http.StatusOK, customerDtos)
}

func (c *Customer) GetById(ctx *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = c.coreSrv.CheckPermission(credentials, c.conf.DBConfig.PermissionGet)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	customerIDStr := ctx.Param("id")
	customerID, err := uuid.Parse(customerIDStr)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	customerDTO, err := c.customerSrv.GetById(customerID, credentials)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, customerDTO)
}

func (c *Customer) Update(ctx *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = c.coreSrv.CheckPermission(credentials, c.conf.DBConfig.PermissionUpdate)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var customerDTO dtos.Customer
	if err := ctx.BindJSON(&customerDTO); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.customerSrv.Update(customerDTO, credentials)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Customer) Delete(ctx *gin.Context) {
	credentials, err := ParseBasicAuthCredentials(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = c.coreSrv.CheckPermission(credentials, c.conf.DBConfig.PermissionDelete)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	customerIDStr := ctx.Param("id")
	customerID, err := uuid.Parse(customerIDStr)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid customer ID"})
		return
	}

	err = c.customerSrv.Delete(customerID, credentials)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
