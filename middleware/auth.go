package middleware

import (
	"dts/learn_middleware/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "You need to be authorized to access this route",
			})
			c.Abort()
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
