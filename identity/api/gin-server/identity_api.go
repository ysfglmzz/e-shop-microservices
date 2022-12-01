package ginserver

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/dto"
	usecase "github.com/ysfglmzz/e-shop-microservices/identity/internal/app/use-case"
)

type IdentityApi struct {
	logger          *logrus.Logger
	identityService usecase.IIdentityService
}

func NewIdentityApi(identityService usecase.IIdentityService, logger *logrus.Logger) *IdentityApi {
	return &IdentityApi{identityService: identityService, logger: logger}
}

// @Tags AuthApi
// @Security ApiKeyAuth
// @Summary User Register
// @Param data body dto.CreateUserRequest true "CreateUserRequest"
// @Success 201 string Success "{"success":true,"msg":"Success"}"
// @Router /auth/register [post]
func (i *IdentityApi) CreateUser(c *gin.Context) {
	var createUserRequest dto.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		i.logger.WithFields(logrus.Fields{"body": c.Request.Body}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := i.identityService.CreateUser(createUserRequest); err != nil {
		i.logger.WithFields(logrus.Fields{"body": createUserRequest}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	i.logger.WithFields(logrus.Fields{"body": createUserRequest}).Info()
	c.JSON(http.StatusCreated, "Created successfully")
}

// @Tags AuthApi
// @Security ApiKeyAuth
// @Summary Login
// @Param data body dto.LoginUserRequest true "LoginUserRequest"
// @Success 201 string Success "{"success":true,"msg":"Success"}"
// @Router /auth/login [post]
func (i *IdentityApi) Login(c *gin.Context) {
	var loginUserRequest dto.LoginUserRequest
	if err := c.ShouldBindJSON(&loginUserRequest); err != nil {
		i.logger.WithFields(logrus.Fields{"body": c.Request.Body}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tokenResponse, err := i.identityService.LoginUser(loginUserRequest)
	if err != nil {
		i.logger.WithFields(logrus.Fields{"body": loginUserRequest}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	i.logger.WithFields(logrus.Fields{"body": loginUserRequest}).Info()
	c.JSON(http.StatusOK, map[string]any{"data": tokenResponse, "message": "Login successfully"})
}
