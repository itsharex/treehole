package main

import (
	"github.com/Jazee6/treehole/cmd/api/handler"
	_ "github.com/Jazee6/treehole/pkg/configs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	initRouter(g)
	handler.InitHandler()

	err := g.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	err = g.Run(viper.GetString("server.gateway"))
	if err != nil {
		panic(err)
	}
}
