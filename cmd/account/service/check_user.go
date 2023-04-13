package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"github.com/Jazee6/treehole/pkg/utils"
	"gorm.io/gorm"
)

func (c *AccountService) AccountLogin(_ context.Context, request *rpc.LoginRequest) (*rpc.LoginResponse, error) {
	q := dao.Q.User
	u, err := q.Where(q.Email.Eq(request.Email)).Take()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &rpc.LoginResponse{
				Code: rpcs.Code_ErrUserNotExist,
			}, nil
		}
		return nil, err
	}
	s := sha256.New()
	s.Write([]byte(request.Password + salt))
	if u.Password != hex.EncodeToString(s.Sum(nil)) {
		return &rpc.LoginResponse{
			Code: rpcs.Code_ErrPasswordErr,
		}, nil
	}
	tk, err := utils.GenToken(*u)
	if err != nil {
		return nil, err
	}
	return &rpc.LoginResponse{
		Code:  rpcs.Code_Success,
		Token: tk,
	}, nil
}
