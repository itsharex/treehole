package main

import (
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/cmd/account/service"
	_ "github.com/Jazee6/treehole/pkg/configs"
	"github.com/Jazee6/treehole/pkg/etcd"
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const Name = "account"

func main() {
	dao.InitDB()
	utils.InitJWT()
	utils.InitSMTP()
	service.InitService()

	sub := viper.Sub("server." + Name)
	host := sub.GetString("host")
	port := sub.GetString("port")
	addr := net.JoinHostPort(host, port)
	log.Println(Name+" server listen on", addr)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	sub = viper.Sub("server.etcd")
	addr = net.JoinHostPort(sub.GetString("host"), sub.GetString("port"))
	err = etcd.Register(host, port, Name, addr, 60) //time.second
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	rpc.RegisterAccountServiceServer(s, &service.AccountService{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
