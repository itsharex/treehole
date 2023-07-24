package middleware

import (
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		payload, err := utils.ValidToken(token[7:])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("payload", payload)
		c.Next()
	}
}
