package middlewares

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrCodeValidToken)
			c.Abort()
			return
		}
		c.Next()
	}
}
