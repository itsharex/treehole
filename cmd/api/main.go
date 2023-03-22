package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	err := g.Run(":8080")
	if err != nil {
		return
	}
}
