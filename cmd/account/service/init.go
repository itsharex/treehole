package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var (
	name           string
	captchaContent string
	captchaExpire  int
	salt           string
	r              *redis.Client
	emailMax       int
)

func InitService() {
	name = viper.GetString("name")
	captchaExpire = viper.GetInt("register.captcha.expire")
	captchaContent = os.Expand(viper.GetString("register.captcha.content"), func(s string) string {
		switch s {
		case "name":
			return name
		case "expire":
			return strconv.Itoa(captchaExpire)
		}
		return "${" + s + "}"
	})
	salt = viper.GetString("register.salt")
	emailMax = viper.GetInt("email.max")

	url, err := redis.ParseURL(viper.GetString("db.redis.account.dsn"))
	if err != nil {
		panic(err)
	}
	r = redis.NewClient(url)
}
