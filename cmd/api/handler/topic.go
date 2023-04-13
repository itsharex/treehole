package handler

import (
	"github.com/Jazee6/treehole/cmd/api/rpc"
	pb "github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateTopicRequest struct {
	Content string `form:"content" binding:"required,max=1024,min=1"`
}

func CreateTopic(c *gin.Context) {
	var req CreateTopicRequest
	err := c.Bind(&req)
	if err != nil {
		return
	}
	claims, _ := c.Get("payload")
	resp, err := rpc.TopicClient.CreateTopic(c, &pb.CreateTopicRequest{
		Uid:     uint32(claims.(*utils.Claims).Uid),
		Content: req.Content,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if resp.Code != rpcs.Code_Success {
		Error(c, NewErr(int(resp.Code), rpcs.Code_name[int32(resp.Code)]))
		return
	}
	Success(c, resp.Code)
}
