package main

import (
	"github.com/Jazee6/treehole/cmd/api/handler"
	"github.com/Jazee6/treehole/cmd/api/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	api := g.Group("/v1")
	{
		api.POST("/register", handler.Register, middleware.Recaptcha())
		api.POST("/login", handler.Login)
		api.POST("/captcha", handler.Captcha)
		api.POST("/campus", handler.GetCampus)
	}

	auth := api.Group("/")
	auth.Use(middleware.Auth())
	{
		auth.GET("/topic/:limit/:offset", handler.GetTopic)
		auth.POST("/topic", handler.CreateTopic, middleware.Recaptcha())
		auth.PUT("/topic/star/:id", handler.PutStar)
		auth.GET("/account", handler.GetAccountInfo)
		auth.GET("/topic/star/:limit/:offset", handler.GetStarList)
		auth.POST("/topic/comment", handler.AddComment) //middleware.Recaptcha()
	}
}
