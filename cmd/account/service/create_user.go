package service

import (
	"crypto/sha256"
	"errors"
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/model"
	"github.com/Jazee6/treehole/pkg/i18n"
	"gorm.io/gorm"
	"time"
)

var q = dao.Q.User

type CreateUserService struct{}

func NewCreateUserService() *CreateUserService {
	return &CreateUserService{}
}

func (c *CreateUserService) Register(u *model.CreateUserRequest) error {
	user, err := q.Where(q.Email.Eq(u.Email)).Take()
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New(i18n.ErrUserExist)
	}
	s := sha256.New()
	s.Write([]byte(u.Password))
	usr := &model.User{
		NickName:  u.NickName,
		Email:     u.Email,
		Password:  string(s.Sum(nil)),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	err = q.Create(usr)
	if err != nil {
		return err
	}
	return nil
}
