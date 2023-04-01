package utils

import (
	"errors"
	"github.com/Jazee6/treehole/cmd/account/model"
	"github.com/Jazee6/treehole/pkg/consts"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var (
	exp    time.Duration
	secret string
)

func InitJWT() {
	h := viper.GetInt("token.expire")
	exp = time.Hour * time.Duration(h)
	secret = viper.GetString("token.secret")
}

func GenToken(u model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": u.ID,
		"exp": time.Now().Add(exp).Format(time.RFC3339),
	})
	signedString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func ValidToken(token string) (jwt.MapClaims, error) {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !tk.Valid {
		return nil, errors.New(consts.ErrTokenInvalid)
	}
	parse, err := time.Parse(time.RFC3339, tk.Claims.(jwt.MapClaims)["exp"].(string))
	if err != nil {
		return nil, err
	}
	if time.Now().After(parse) {
		return nil, errors.New(consts.ErrExpired)
	}
	return tk.Claims.(jwt.MapClaims), nil
}
