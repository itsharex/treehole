package service

import (
	"context"
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/rpc"
)

func (c *AccountService) GetCampusList(_ context.Context, req *rpc.CampusListReq) (*rpc.CampusListResp, error) {
	q := dao.Q.Campu
	var campus []*rpc.Campus
	err := q.Where(q.Name.Like(req.Name + "%")).Limit(10).Scan(&campus)
	if err != nil {
		return nil, err
	}
	return &rpc.CampusListResp{
		Campus: campus,
	}, nil
}
