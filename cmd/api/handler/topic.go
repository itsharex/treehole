package handler

import (
	pb "github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"github.com/gin-gonic/gin"
	"log"
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
	resp, err := pb.TopicClient.CreateTopic(c, &pb.CreateTopicRequest{
		Uid:     GetUid(c),
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
	Limit  uint32 `uri:"limit" binding:"required,min=5,max=20"`
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
		Uid:    GetUid(c),
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

type PutStarRequest struct {
	Id uint32 `uri:"id" binding:"required,number"`
}

func PutStar(c *gin.Context) {
	var req PutStarRequest
	err := c.BindUri(&req)
	if err != nil {
		return
	}
	star, err := pb.TopicClient.PutStar(c, &pb.PutStarReq{
		Uid: GetUid(c),
		Tid: req.Id,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	Success(c, star.Code)
}

type GetStarListRequest struct {
	Limit  uint32 `uri:"limit" binding:"required,min=5,max=20"`
	Offset uint32 `uri:"offset" binding:"gte=0"`
}

func GetStarList(c *gin.Context) {
	var req GetStarListRequest
	err := c.BindUri(&req)
	if err != nil {
		return
	}
	resp, err := pb.TopicClient.GetStarList(c, &pb.GetStarListReq{
		Uid:    GetUid(c),
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	Success(c, resp.Topics)
}

type AddCommentReq struct {
	Tid      uint32  `json:"tid" binding:"required,number,min=1"`
	Content  string  `json:"content" binding:"required,max=1024,min=1"`
	Root     *uint32 `json:"root"`
	ToTempId *uint32 `json:"to_temp_id"`
}

func AddComment(c *gin.Context) {
	var req AddCommentReq
	err := c.BindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %v", err)
		return
	}
	resp, err := pb.TopicClient.AddComment(c, &pb.AddCommentReq{
		Uid:      GetUid(c),
		Tid:      req.Tid,
		Content:  req.Content,
		Root:     req.Root,
		ToTempId: req.ToTempId,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	Success(c, resp.Code)
}
