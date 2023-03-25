package handler

import (
	"github.com/gin-gonic/gin"
)

type RegRequest struct {
	NickName string `form:"nick_name"  binding:"required,max=16"`
	Email    string `form:"email"  binding:"required,email"`
	Password string `form:"password" binding:"required,sha256"`
}

func Register(c *gin.Context) {
	var req RegRequest
	err := c.Bind(&req)
	if err != nil {
		Error(c, ErrValidate)
		return
	}
	Success(c, nil)
}

func Login(c *gin.Context) {

}
