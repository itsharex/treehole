package service

import (
	"context"
	"crypto/sha256"
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/model"
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/pkg/utils"
	"gorm.io/gorm"
	"time"
)

var q = dao.Q.User

type CreateUserService struct{}

func (c *CreateUserService) AccountRegister(_ context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	user, err := q.Where(q.Email.Eq(request.Email)).Take()
	if err != nil {
		return nil, err
	}
	if user != nil {
		return &rpc.RegisterResponse{
			Code: rpc.Code_ErrUserExist,
		}, nil
	}
	s := sha256.New()
	s.Write([]byte(request.Password))
	usr := &model.User{
		NickName:  request.Nickname,
		Email:     request.Email,
		Password:  string(s.Sum(nil)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}
	err = q.Create(usr)
	if err != nil {
		return nil, err
	}
	tk, err := utils.GenToken(*user)
	if err != nil {
		return nil, err
	}
	return &rpc.RegisterResponse{
		Code:  rpc.Code_Success,
		Token: tk,
	}, nil
}
