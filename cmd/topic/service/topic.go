package service

import (
	"context"
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/cmd/topic/dao"
	"github.com/Jazee6/treehole/cmd/topic/model"
	"github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"gorm.io/gorm"
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

	//Get Account Info
	var req pb.TopicInfoReq
	var tid = make([]uint32, len(finds))
	req.Uid = make([]uint32, len(finds))
	for i, find := range finds {
		req.Uid[i] = find.UID
		tid[i] = find.ID
	}
	info, err := pb.AccountClient.GetTopicInfo(ctx, &req)
	if err != nil {
		return nil, err
	}

	s := dao.Q.Star
	if request.Uid != 0 {

	}
	stars, err := s.Where(s.TopicID.In(tid...), s.UID.Eq(request.Uid)).Find()
	if err != nil {
		return nil, err
	}

	var topics = make([]*rpc.Topic, len(finds))
	for i, find := range finds {
		topicInfo := appendTopicInfo(find.UID, info)
		for _, star := range stars {
			if star.TopicID == find.ID {
				topicInfo.starred = true
				break
			}
		}
		topics[i] = &rpc.Topic{
			Id:        find.ID,
			Content:   find.Content,
			CreatedAt: find.CreatedAt.String(),
			Campus:    topicInfo.campusName,
			Verified:  topicInfo.verified,
			Starred:   topicInfo.starred,
		}
	}
	return &rpc.GetTopicResponse{
		Code:   rpcs.Code_Success,
		Topics: topics,
	}, nil
}

type TopicInfo struct {
	campusName string
	verified   bool
	starred    bool
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

func (t TopicService) PutStar(_ context.Context, req *rpc.PutStarReq) (*rpc.PutStarResp, error) {
	q := dao.Q.Star
	take, err := q.Where(q.UID.Eq(req.Uid), q.TopicID.Eq(req.Tid)).Take()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if take != nil {
		_, err := q.Where(q.UID.Eq(req.Uid), q.TopicID.Eq(req.Tid)).Delete()
		if err != nil {
			return nil, err
		}
		return &rpc.PutStarResp{
			Code: rpcs.Code_OKUnStar,
		}, nil
	}
	err = q.Create(&model.Star{
		UID:     req.Uid,
		TopicID: req.Tid,
	})
	if err != nil {
		return nil, err
	}
	return &rpc.PutStarResp{
		Code: rpcs.Code_OKStar,
	}, nil
}
