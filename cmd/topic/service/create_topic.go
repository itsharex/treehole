package service

import (
	"context"
	"github.com/Jazee6/treehole/cmd/topic/dao"
	"github.com/Jazee6/treehole/cmd/topic/model"
	"github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"time"
)

type TopicService struct{}

func (t TopicService) CreateTopic(_ context.Context, request *rpc.CreateTopicRequest) (*rpc.CreateTopicResponse, error) {
	q := dao.Q.Topic
	now := time.Now()
	err := q.Create(&model.Topic{
		UID:       int32(request.Uid),
		Content:   request.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return nil, err
	}
	return &rpc.CreateTopicResponse{
		Code: rpcs.Code_Success,
	}, nil
}
