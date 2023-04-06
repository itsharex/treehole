package etcd

import (
	"fmt"
	clientV3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func WatchGrpc(etcdTarget, serviceName string) (*grpc.ClientConn, error) {
	c, err := clientV3.NewFromURL(etcdTarget)
	if err != nil {
		return nil, err
	}
	r, err := resolver.NewBuilder(c)
	if err != nil {
		return nil, err
	}
	dial, err := grpc.Dial(fmt.Sprintf("etcd:///%s", serviceName),
		grpc.WithResolvers(r),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		return nil, err
	}
	return dial, nil
}
