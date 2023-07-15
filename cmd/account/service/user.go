package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/model"
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AccountService struct{}

func (c *AccountService) AccountRegister(ctx context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	result, err := r.Get(ctx, request.Email).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if err == redis.Nil {
		return &rpc.RegisterResponse{
			Code: rpcs.Code_ErrCaptchaNil,
		}, nil
	}
	if result != request.Captcha {
		return &rpc.RegisterResponse{
			Code: rpcs.Code_ErrCaptchaErr,
		}, nil
	}

	q := dao.Q.User
	user, err := q.Where(q.Email.Eq(request.Email)).Take()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if user != nil {
		return &rpc.RegisterResponse{
			Code: rpcs.Code_ErrUserExist,
		}, nil
	}
	s := sha256.New()
	s.Write([]byte(request.Password + salt))
	usr := &model.User{
		Email:    request.Email,
		Password: hex.EncodeToString(s.Sum(nil)),
		CampusID: request.CampusId,
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
		Code:  rpcs.Code_Success,
		Token: tk,
	}, nil
}

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

func (c *AccountService) GetTopicInfo(_ context.Context, req *rpc.TopicInfoReq) (*rpc.TopicInfoResp, error) {
	q := dao.Q.User
	ca := dao.Q.Campu
	var topicInfo []*rpc.TopicInfo
	err := q.Where(q.ID.In(req.Uid...)).LeftJoin(ca, q.CampusID.EqCol(ca.ID)).Select(q.ID.As("uid"), ca.Name.As("CampusName"), q.Verified).Scan(&topicInfo)
	if err != nil {
		return nil, err
	}
	return &rpc.TopicInfoResp{
		TopicInfo: topicInfo,
	}, nil
}
