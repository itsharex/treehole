package utils

import (
	"errors"
	"github.com/Jazee6/treehole/cmd/account/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var (
	exp    time.Duration
	secret string
)

type Claims struct {
	Uid int32
	jwt.RegisteredClaims
}

func InitJWT() {
	h := viper.GetInt("token.expire")
	exp = time.Hour * time.Duration(h)
	secret = viper.GetString("token.secret")
}

func GenToken(u model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Uid: int32(u.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	})
	signedString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func ValidToken(token string) (*Claims, error) {
	tk, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !tk.Valid {
		return nil, errors.New("token invalid")
	}
	return tk.Claims.(*Claims), nil
}
