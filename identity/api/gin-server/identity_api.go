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

// @Tags AuthApi
// @Security ApiKeyAuth
// @Summary Verify User
// @Param verifyCode body dto.VerifyCodeRequest true "Verify Code"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /auth/verify [put]
func (i *IdentityApi) VerifyUserByCode(c *gin.Context) {
	var verifyCode dto.VerifyCodeRequest
	if err := c.ShouldBindJSON(&verifyCode); err != nil {
		i.logger.WithFields(logrus.Fields{"body": c.Request.Body}).WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := i.identityService.VerifyUserByCode(verifyCode.VerifyCode)
	if err != nil {
		i.logger.WithError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	i.logger.Info()
	c.JSON(http.StatusOK, "Verification successfully")
}

// @Tags AuthApi
// @Security ApiKeyAuth
// @Summary Token Control
// @Param uuid path string true "uuid"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /auth/{uuid} [get]
func (i *IdentityApi) TokenControl(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if !i.identityService.CheckTokenExist(uuid) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Authorized"})
}
