package handler

import (
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
