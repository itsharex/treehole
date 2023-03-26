package main

import (
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/cmd/account/service"
	_ "github.com/Jazee6/treehole/pkg/configs"
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net"
)

func main() {
	dsn := viper.GetString("mysql.dsn.account")
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.Default(), logger.Config{
			IgnoreRecordNotFoundError: true,
		}),
	})
	if err != nil {
		return
	}
	dao.SetDefault(open)

	utils.InitJWT()

	addr := viper.GetString("server.account")
	log.Println("account server listen on", addr)
	listen, err := net.Listen("tcp", addr)
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
