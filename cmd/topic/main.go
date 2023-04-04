package main

import (
	"github.com/Jazee6/treehole/cmd/topic/rpc"
	"github.com/Jazee6/treehole/cmd/topic/service"
	_ "github.com/Jazee6/treehole/pkg/configs"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	//dao.InitDB()
	//utils.InitJWT()
	//utils.InitSMTP()
	//service.InitService()

	addr := viper.GetString("server.topic")
	log.Println("topic server listen on", addr)
	listen, err := net.Listen("tcp", addr)
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
