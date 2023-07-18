package handler

import (
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegRequest struct {
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password" binding:"required,sha256"`
	Captcha  string `json:"captcha" binding:"required,len=6"`
	CampusId uint32 `json:"campusId" binding:"required,number"`
}

func Register(c *gin.Context) {
	var req RegRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	resp, err := pb.AccountClient.AccountRegister(c, &pb.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Captcha:  req.Captcha,
		CampusId: req.CampusId,
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
	resp, err := pb.AccountClient.AccountLogin(c, &pb.LoginRequest{
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
	Email string `json:"email"  binding:"required,email"`
}

func Captcha(c *gin.Context) {
	var req CaptchaRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	resp, err := pb.AccountClient.SendCaptcha(c, &pb.SendCaptchaRequest{
		Email: req.Email,
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

type CampusRequest struct {
	Name string `json:"name"  binding:"required,max=64"`
}

func GetCampus(c *gin.Context) {
	var req CampusRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	resp, err := pb.AccountClient.GetCampusList(c, &pb.CampusListReq{
		Name: req.Name,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	Success(c, resp)
}
