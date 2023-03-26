package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/model"
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type CreateUserService struct{}

func (c *CreateUserService) AccountRegister(_ context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	q := dao.Q.User
	user, err := q.Where(q.Email.Eq(request.Email)).Take()
	if err != nil && err != gorm.ErrRecordNotFound {
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
		Password:  hex.EncodeToString(s.Sum(nil)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}
	err = q.Create(usr)
	if err != nil {
		return nil, err
	}
	tk, err := utils.GenToken(*usr)
	if err != nil {
		return nil, err
	}
	return &rpc.RegisterResponse{
		Code:  rpc.Code_Success,
		Token: tk,
	}, nil
}
