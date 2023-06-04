package handler

import (
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/cmd/api/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegRequest struct {
	NickName string `json:"nick_name"  binding:"required,max=16"`
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password" binding:"required,sha256"`
	Captcha  string `json:"captcha" binding:"required,len=6"`
}

func Register(c *gin.Context) {
	var req RegRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	resp, err := rpc.AccountClient.AccountRegister(c, &pb.RegisterRequest{
		Nickname: req.NickName,
		Email:    req.Email,
		Password: req.Password,
		Captcha:  req.Captcha,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if resp.Code != rpcs.Code_Success {
		Error(c, NewErr(int(resp.Code), rpcs.Code_name[int32(resp.Code)]))
		return
	}
	Success(c, resp)
}

type LoginRequest struct {
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password" binding:"required,sha256"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	resp, err := rpc.AccountClient.AccountLogin(c, &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if resp.Code != rpcs.Code_Success {
		Error(c, NewErr(int(resp.Code), rpcs.Code_name[int32(resp.Code)]))
		return
	}
	Success(c, resp)
}

type CaptchaRequest struct {
	Token string `json:"token"  binding:"required"`
	Email string `json:"email"  binding:"required,email"`
}

func Captcha(c *gin.Context) {
	var req CaptchaRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	resp, err := rpc.AccountClient.SendCaptcha(c, &pb.SendCaptchaRequest{
		Email: req.Email,
		Token: req.Token,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if resp.Code != rpcs.Code_Success {
		Error(c, NewErr(int(resp.Code), rpcs.Code_name[int32(resp.Code)]))
		return
	}
	Success(c, resp)
}
