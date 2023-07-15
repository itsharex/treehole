package service

import (
	"context"
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/cmd/topic/dao"
	"github.com/Jazee6/treehole/cmd/topic/model"
	"github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"log"
)

type TopicService struct{}

func (t TopicService) CreateTopic(_ context.Context, request *rpc.CreateTopicRequest) (*rpc.CreateTopicResponse, error) {
	q := dao.Q.Topic
	err := q.Create(&model.Topic{
		UID:     request.Uid,
		Content: request.Content,
	})
	if err != nil {
		return nil, err
	}
	return &rpc.CreateTopicResponse{
		Code: rpcs.Code_Success,
	}, nil
}

func (t TopicService) GetTopic(ctx context.Context, request *rpc.GetTopicRequest) (*rpc.GetTopicResponse, error) {
	q := dao.Q.Topic
	finds, err := q.Limit(int(request.Limit)).Offset(int(request.Offset)).Order(q.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}
	var req pb.TopicInfoReq
	req.Uid = make([]uint32, len(finds))
	for i, find := range finds {
		req.Uid[i] = find.UID
	}
	info, err := rpc.AccountClient.GetTopicInfo(ctx, &req)
	if err != nil {
		return nil, err
	}
	var topics = make([]*rpc.Topic, len(finds))
	log.Println(info.TopicInfo)
	for i, find := range finds {
		topicInfo := appendTopicInfo(find.UID, info)
		topics[i] = &rpc.Topic{
			Id:        find.ID,
			Content:   find.Content,
			CreatedAt: find.CreatedAt.String(),
			Campus:    topicInfo.campusName,
			Verified:  topicInfo.verified,
		}
	}
	log.Println(topics)
	return &rpc.GetTopicResponse{
		Code:   rpcs.Code_Success,
		Topics: topics,
	}, nil
}

type TopicInfo struct {
	campusName string
	verified   bool
}

func appendTopicInfo(uid uint32, info *pb.TopicInfoResp) TopicInfo {
	var topicInfo TopicInfo
	for _, user := range info.TopicInfo {
		if user.Uid == uid {
			topicInfo.campusName = user.CampusName
			topicInfo.verified = user.Verified
			break
		}
	}
	return topicInfo
}
