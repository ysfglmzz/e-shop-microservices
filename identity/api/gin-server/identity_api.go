package ginserver

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/dto"
	usecase "github.com/ysfglmzz/e-shop-microservices/identity/internal/app/use-case"
	"github.com/ysfglmzz/e-shop-microservices/identity/pkg/constants"
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
// @Summary Verify User
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /auth/checkToken [get]
func (i *IdentityApi) TokenControl(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	tokenString := strings.Split(auth, " ")[1]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(constants.TokenSecretKey), nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	userId := claims["user_id"].(float64)
	if !i.identityService.CheckTokenExist(int(userId)) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Authorized"})
}
