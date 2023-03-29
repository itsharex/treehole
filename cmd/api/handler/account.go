package handler

import (
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/cmd/api/rpc"
	"github.com/gin-gonic/gin"
)

type RegRequest struct {
	NickName string `form:"nick_name"  binding:"required,max=16"`
	Email    string `form:"email"  binding:"required,email"`
	Password string `form:"password" binding:"required,sha256"`
	Captcha  string `form:"captcha" binding:"required,len=6"`
}

func Register(c *gin.Context) {
	var req RegRequest
	err := c.Bind(&req)
	if err != nil {
		Error(c, ErrValidate)
		return
	}
	resp, err := rpc.Client.AccountRegister(c, &pb.RegisterRequest{
		Nickname: req.NickName,
		Email:    req.Email,
		Password: req.Password,
		Captcha:  req.Captcha,
	})
	if err != nil {
		Error(c, ErrServer)
		return
	}
	if resp.Code != pb.Code_Success {
		Error(c, NewErr(int(resp.Code), pb.Code_name[int32(resp.Code)]))
		return
	}
	Success(c, resp)
}
