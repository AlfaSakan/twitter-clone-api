package middlewares

import (
	"github.com/gin-gonic/gin"
)

func RequireUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("User")
		if !exists {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
