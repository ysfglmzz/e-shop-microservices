package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ysfglmzz/e-shop-microservices/basket/pkg/constants"
)

func Authorization(expectedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
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
		roleList := claims["roles"].([]interface{})
		for _, role := range expectedRoles {
			for _, roleListRole := range roleList {
				if role == roleListRole {
					ctx.Next()
					return
				}
			}
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}
}
