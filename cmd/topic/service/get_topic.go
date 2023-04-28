package service

import (
	"context"
	"github.com/Jazee6/treehole/cmd/topic/dao"
	"github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
)

func (t TopicService) GetTopic(ctx context.Context, request *rpc.GetTopicRequest) (*rpc.GetTopicResponse, error) {
	q := dao.Q.Topic
	var topics []*rpc.Topic
	err := q.Limit(int(request.Limit)).Offset(int(request.Offset)).Order(q.CreatedAt.Desc()).Scan(&topics)
	if err != nil {
		return nil, err
	}
	return &rpc.GetTopicResponse{
		Code:   rpcs.Code_Success,
		Topics: topics,
	}, nil
}
