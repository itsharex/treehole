package service

import (
	"context"
	"github.com/Jazee6/treehole/cmd/topic/rpc"
)

type TopicService struct{}

func (t TopicService) CreateTopic(ctx context.Context, request *rpc.CreateTopicRequest) (*rpc.CreateTopicResponse, error) {
	return &rpc.CreateTopicResponse{
		Code: rpc.Code_Success,
	}, nil
}
