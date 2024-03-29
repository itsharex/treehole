package test

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"testing"
)

func TestMain(m *testing.M) {
	viper.SetConfigFile("../config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestT(t *testing.T) {
	url, err := redis.ParseURL(viper.GetString("db.redis.account.dsn"))
	if err != nil {
		panic(err)
	}
	r := redis.NewClient(url)

	result, err := r.HGet(context.Background(), "test", "test").Uint64()
	if err != nil {
		return
	}
	t.Log(result)
}

func TestT2(t *testing.T) {

}
