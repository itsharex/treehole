package rpc

import (
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/pkg/etcd"
	"github.com/spf13/viper"
	"net"
)

var AccountClient pb.AccountServiceClient

func InitAccount() {
	sub := viper.Sub("server.etcd")
	addr := net.JoinHostPort(sub.GetString("host"), sub.GetString("port"))
	dail, err := etcd.WatchGrpc(addr, "account")
	if err != nil {
		panic(err)
	}
	AccountClient = pb.NewAccountServiceClient(dail)
}
