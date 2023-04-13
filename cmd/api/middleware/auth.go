package middleware

import (
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				return
			}
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		payload, err := utils.ValidToken(cookie)
		if err != nil {
			println(err.Error())
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}
		c.Set("payload", payload)
		c.Next()
	}
}
