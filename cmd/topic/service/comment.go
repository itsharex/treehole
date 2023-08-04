package service

import (
	"context"
	"github.com/Jazee6/treehole/cmd/topic/dao"
	"github.com/Jazee6/treehole/cmd/topic/model"
	"github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"log"
	"strconv"
)

func (t TopicService) AddComment(ctx context.Context, req *rpc.AddCommentReq) (*rpc.AddCommentResp, error) {
	c := dao.Q.Comment
	result, err := r.HGet(ctx, strconv.Itoa(int(req.Tid)), strconv.Itoa(int(req.Uid))).Uint64()
	if err != nil {
		return nil, err
	}
	if result == 0 {
		return &rpc.AddCommentResp{
			Code: rpcs.Code_ErrTopicNotFound,
		}, nil
	}
	if err := c.Create(&model.Comment{
		Tid:      req.Tid,
		UID:      req.Uid,
		Content:  req.Content,
		Root:     req.Root,
		ToTempID: req.ToTempId,
		TempID:   uint32(result),
	}); err != nil {
		return nil, err
	}
	return &rpc.AddCommentResp{
		Code: rpcs.Code_Success,
	}, nil
}

func (t TopicService) GetCommentList(ctx context.Context, req *rpc.GetCommentListReq) (*rpc.GetCommentListResp, error) {
	c := dao.Q.Comment
	take, err := c.Where(c.Tid.Eq(req.Tid), c.Status.Eq(0), c.Root.IsNull()).Limit(int(req.Limit)).Offset(int(req.Offset)).Order(c.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}
	var ids = make([]uint32, len(take))
	for i, v := range take {
		ids[i] = v.ID
	}
	var subComments []*rpc.SubComment
	err = c.Where(c.Root.In(ids...)).Select(c.ALL, c.Root.Count()).Group(c.Root).Scan(&subComments)
	if err != nil {
		return nil, err
	}
	log.Println(subComments)
	return nil, nil
}
