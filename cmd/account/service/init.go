package service

import (
	"github.com/spf13/viper"
)

var (
	name           string
	captchaContent string
	captchaExpire  int
)

func InitService() {
	name = viper.GetString("name")
	captchaContent = viper.GetString("register.captcha.content")
	captchaExpire = viper.GetInt("register.captcha.expire")
}
