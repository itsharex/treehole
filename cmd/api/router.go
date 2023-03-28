package main

import (
	"github.com/Jazee6/treehole/cmd/api/handler"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	api := g.Group("/api/v1")

	{
		api.POST("/register", handler.Register)
	}

	//auth := api.Group("/auth")
	//
	//{
	//
	//}
}
