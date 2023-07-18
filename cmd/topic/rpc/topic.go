package rpc

import (
	"github.com/Jazee6/treehole/pkg/etcd"
	"github.com/spf13/viper"
	"net"
)

var TopicClient TopicServiceClient

func InitTopic() {
	sub := viper.Sub("server.etcd")
	addr := net.JoinHostPort(sub.GetString("host"), sub.GetString("port"))
	dail, err := etcd.WatchGrpc(addr, "topic")
	if err != nil {
		panic(err)
	}
	TopicClient = NewTopicServiceClient(dail)
}
