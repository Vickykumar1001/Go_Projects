package middlewares

import (
	"net/http"
	"user_management/internal/services"

	"github.com/gin-gonic/gin"
)

var (
	jwtServiceObj = services.JWTService{}
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// We will check for the authorization header
		authorization := c.GetHeader("Authorization")
		if len(authorization) == 0 {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide Authorization token",
			})
			c.Abort()

			return

		}

		// We will check if the authorization header is valid
		_, err := jwtServiceObj.ValidateJWT(authorization)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide valid Authorization token",
			})
			c.Abort()
			return
		}

		// then we will proceed further
		c.Next()
	}
}
