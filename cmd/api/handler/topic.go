package handler

import (
	pb "github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateTopicRequest struct {
	Content string `json:"content" binding:"required,max=1024,min=1"`
}

func CreateTopic(c *gin.Context) {
	var req CreateTopicRequest
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	claims, _ := c.Get("payload")
	resp, err := pb.TopicClient.CreateTopic(c, &pb.CreateTopicRequest{
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

type GetTopicRequest struct {
	Limit  uint32 `uri:"limit" binding:"required,min=1,max=20"`
	Offset uint32 `uri:"offset" binding:"gte=0"`
}

func GetTopic(c *gin.Context) {
	var req GetTopicRequest
	err := c.BindUri(&req)
	if err != nil {
		return
	}
	resp, err := pb.TopicClient.GetTopic(c, &pb.GetTopicRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if resp.Code != rpcs.Code_Success {
		Error(c, NewErr(int(resp.Code), rpcs.Code_name[int32(resp.Code)]))
		return
	}
	Success(c, resp.Topics)
}
