package handler

import (
	"github.com/gin-gonic/gin"
)

type CreateTopicRequest struct {
	Content string `form:"content" binding:"required,max=1024,min=1"`
}

func CreateTopic(c *gin.Context) {

}
