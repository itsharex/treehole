package utils

import (
	"github.com/Jazee6/treehole/cmd/account/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var exp time.Duration
var secret string

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
