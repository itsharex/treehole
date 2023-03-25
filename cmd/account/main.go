package main

import (
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/cmd/account/service"
	_ "github.com/Jazee6/treehole/pkg/configs"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	dsn := viper.GetString("mysql.dsn.account")
	open, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return
	}
	dao.SetDefault(open)

	listen, err := net.Listen("tcp", viper.GetString("server.account"))
	if err != nil {
		return
	}
	s := grpc.NewServer()
	rpc.RegisterAccountServiceServer(s, &service.CreateUserService{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
