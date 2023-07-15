package main

import (
	"github.com/Jazee6/treehole/cmd/api/handler"
	"github.com/Jazee6/treehole/cmd/api/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	api := g.Group("/v1")
	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)
		api.POST("/captcha", handler.Captcha)
		api.GET("/topic/:limit/:offset", handler.GetTopic)
		api.POST("/campus", handler.GetCampus)
	}

	auth := api.Group("/")
	auth.Use(middleware.Auth())
	{
		auth.POST("/topic", handler.CreateTopic)
	}
}
