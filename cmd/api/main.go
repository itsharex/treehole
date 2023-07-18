package main

import (
	"github.com/Jazee6/treehole/cmd/account/rpc"
	rpc2 "github.com/Jazee6/treehole/cmd/topic/rpc"
	_ "github.com/Jazee6/treehole/pkg/config"
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net"
)

const name = "gateway"

func main() {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	utils.InitJWT()
	utils.InitRecaptcha()
	initRouter(g)
	rpc.InitAccount()
	rpc2.InitTopic()

	err := g.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	sub := viper.Sub("server." + name)
	addr := net.JoinHostPort(sub.GetString("host"), sub.GetString("port"))
	err = g.Run(addr)
	if err != nil {
		panic(err)
	}
}
