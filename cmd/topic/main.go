package main

import (
	"github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/cmd/topic/service"
	_ "github.com/Jazee6/treehole/pkg/configs"
	"github.com/Jazee6/treehole/pkg/etcd"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const name = "topic"

func main() {
	sub := viper.Sub("server." + name)
	host := sub.GetString("host")
	port := sub.GetString("port")
	addr := net.JoinHostPort(host, port)
	log.Println(name+" server listen on", addr)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	sub = viper.Sub("server.etcd")
	addr = net.JoinHostPort(sub.GetString("host"), sub.GetString("port"))
	err = etcd.Register(host, port, name, addr, 60) //time.second
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	rpc.RegisterTopicServiceServer(s, &service.TopicService{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
