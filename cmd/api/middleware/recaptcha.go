package middleware

import (
	"github.com/Jazee6/treehole/cmd/api/handler"
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Recaptcha")
		if token == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		err := utils.Recaptcha(token)
		if err != nil {
			if err.Error() == "recaptcha failed" {
				handler.Error(c, handler.ErrRecaptcha)
				return
			}
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Next()
	}
}
