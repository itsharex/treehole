package service

import (
	"context"
	"errors"
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/cmd/topic/dao"
	"github.com/Jazee6/treehole/cmd/topic/model"
	"github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"gorm.io/gorm"
	"time"
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
	finds, err := q.Where(q.Status.Eq(0)).Limit(int(request.Limit)).Offset(int(request.Offset)).Order(q.CreatedAt.Desc()).Find()
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
	var stars []*model.Star
	if request.Uid != 0 {
		stars, err = s.Where(s.TopicID.In(tid...), s.UID.Eq(request.Uid)).Find()
		if err != nil {
			return nil, err
		}
	}
	var topics = make([]*rpc.Topic, len(finds))
	for i, find := range finds {
		topicInfo := appendTopicInfo(find.UID, info)
		if request.Uid != 0 {
			for _, star := range stars {
				if star.TopicID == find.ID {
					topicInfo.starred = true
					break
				}
			}
		}
		topics[i] = &rpc.Topic{
			Id:        find.ID,
			Content:   find.Content,
			CreatedAt: find.CreatedAt.Format(time.RFC3339),
			Campus:    topicInfo.campusName,
			Verified:  topicInfo.verified,
			Starred:   topicInfo.starred,
			StarCount: find.Stars,
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
	var star bool
	err := dao.Q.Transaction(func(tx *dao.Query) error {
		take, err := tx.Star.Where(tx.Star.UID.Eq(req.Uid), tx.Star.TopicID.Eq(req.Tid)).Take()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if take != nil {
			if _, err := tx.Star.Where(tx.Star.UID.Eq(req.Uid), tx.Star.TopicID.Eq(req.Tid)).Delete(); err != nil {
				return err
			}
			if _, err := tx.Topic.Where(tx.Topic.ID.Eq(req.Tid)).UpdateSimple(tx.Topic.Stars.Sub(1)); err != nil {
				return err
			}
			return nil
		}
		if err := tx.Star.Create(&model.Star{
			UID:     req.Uid,
			TopicID: req.Tid,
		}); err != nil {
			return err
		}
		if _, err := tx.Topic.Where(tx.Topic.ID.Eq(req.Tid)).UpdateSimple(tx.Topic.Stars.Add(1)); err != nil {
			return err
		}
		star = true
		return nil
	})
	if err != nil {
		return nil, err
	}
	if star {
		return &rpc.PutStarResp{
			Code: rpcs.Code_OKStar,
		}, nil
	}
	return &rpc.PutStarResp{
		Code: rpcs.Code_OKUnStar,
	}, nil
}

func (t TopicService) GetStarList(ctx context.Context, req *rpc.GetStarListReq) (*rpc.GetStarListResp, error) {
	q := dao.Q.Topic
	s := dao.Q.Star
	stars, err := s.Where(s.UID.Eq(req.Uid)).Select(s.TopicID, s.CreatedAt).Limit(int(req.Limit)).Offset(int(req.Offset)).Order(s.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}
	if len(stars) == 0 {
		return &rpc.GetStarListResp{
			Code: rpcs.Code_Success,
		}, nil
	}
	var tid = make([]uint32, len(stars))
	for i, star := range stars {
		tid[i] = star.TopicID
	}
	take, err := q.Where(q.Status.Eq(0), q.ID.In(tid...)).Find()
	if err != nil {
		return nil, err
	}
	if len(take) == 0 {
		return &rpc.GetStarListResp{
			Code: rpcs.Code_Success,
		}, nil
	}
	var takes = make([]*model.Topic, len(take))
	for i, find := range stars {
		for _, topic := range take {
			if find.TopicID == topic.ID {
				takes[i] = topic
				break
			}
		}
	}

	var request pb.TopicInfoReq
	request.Uid = make([]uint32, len(take))
	for i, find := range take {
		request.Uid[i] = find.UID
	}
	info, err := pb.AccountClient.GetTopicInfo(ctx, &request)
	if err != nil {
		return nil, err
	}

	var topics = make([]*rpc.Topic, len(take))
	for i, find := range takes {
		topicInfo := appendTopicInfo(find.UID, info)
		topics[i] = &rpc.Topic{
			Id:        find.ID,
			Content:   find.Content,
			CreatedAt: find.CreatedAt.Format(time.RFC3339),
			Campus:    topicInfo.campusName,
			Verified:  topicInfo.verified,
			Starred:   true,
			StarCount: find.Stars,
		}
	}
	return &rpc.GetStarListResp{
		Code:   rpcs.Code_Success,
		Topics: topics,
	}, nil
}
