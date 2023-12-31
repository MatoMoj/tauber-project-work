package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"project-work-tauber/app/internal/dtos"
)

func ParseBasicAuthCredentials(c *gin.Context) (credentials dtos.UserCredentials, err error) {
	user, password, ok := c.Request.BasicAuth()
	if !ok {
		return credentials, errors.New("error parsing credentials")
	}
	credentials.Username = user
	credentials.Password = password
	return
}
