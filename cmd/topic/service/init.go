package service

import (
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	r *redis.Client
)

func InitService() {
	url, err := redis.ParseURL(viper.GetString("db.redis.topic.dsn"))
	if err != nil {
		panic(err)
	}
	r = redis.NewClient(url)

	rpc.InitAccount()
}
