package handler

import (
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Err struct {
	Code    int
	Message string
}

func NewErr(code int, message string) *Err {
	return &Err{Code: code, Message: message}
}

var (
	ErrRecaptcha = NewErr(40100, "recaptcha failed")
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, err *Err) {
	c.JSON(http.StatusOK, Response{
		Code:    err.Code,
		Message: err.Message,
		Data:    nil,
	})
}

func GetUid(c *gin.Context) uint32 {
	claims, _ := c.Get("payload")
	return uint32(claims.(*utils.Claims).Uid)
}

func GetPUid(c *gin.Context) uint32 {
	token := c.GetHeader("Authorization")
	if token == "" {
		return 0
	}
	payload, err := utils.ValidToken(token[7:])
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0
	}
	return uint32(payload.Uid)
}
