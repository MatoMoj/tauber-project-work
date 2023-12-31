package main

import (
	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"
	"log"
	"project-work-tauber/app/config"
	"project-work-tauber/app/internal/controller"
	"project-work-tauber/app/internal/infrastructure"
	"project-work-tauber/app/internal/service"
)

func main() {

	// Parse App Config
	var appConfig config.AppConfig
	config.ReadConfig(&appConfig, "config/app_config.yml")

	// Initialize REST Controller
	r := gin.Default()
	v1 := r.Group("api/v1")

	// Initialize Middleware
	xssMdlwr := &xss.XssMw{
		FieldsToSkip: []string{"password"},
		BmPolicy:     "UGCPolicy",
	}
	r.Use(xssMdlwr.RemoveXss())
	v1.Use(xssMdlwr.RemoveXss())

	// Initialize controller, service, dao
	vaultProvider := infrastructure.NewVaultPsqlCredentialProvider(&appConfig.VaultConfig)

	coreDao := infrastructure.NewPSQLCoreDao(&appConfig, *vaultProvider)
	customerDao := infrastructure.NewPSQLCustomerDao(coreDao)
	bookingDao := infrastructure.NewPSQLBookingDao(coreDao)

	coreSrv := service.NewCore(coreDao)

	// DB Auto migration
	err := coreSrv.DBAutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	customerSrv := service.NewCustomer(customerDao, bookingDao)
	bookingSrv := service.NewBooking(bookingDao)

	customerCtrl := controller.NewCustomer(appConfig, customerSrv, coreSrv)
	bookingCtrl := controller.NewBooking(appConfig, bookingSrv, coreSrv)

	v1.POST("/customers", customerCtrl.Create)
	v1.GET("/customers", customerCtrl.GetAll)
	v1.GET("/customers/:id", customerCtrl.GetById)
	v1.PUT("/customers/:id", customerCtrl.Update)
	v1.DELETE("/customers/:id", customerCtrl.Delete)

	v1.POST("/bookings", bookingCtrl.Create)
	v1.GET("/bookings", bookingCtrl.GetAll)
	v1.GET("/bookings/:id", bookingCtrl.GetById)
	v1.PUT("/bookings/:id", bookingCtrl.Update)
	v1.DELETE("/bookings/:id", bookingCtrl.Delete)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
