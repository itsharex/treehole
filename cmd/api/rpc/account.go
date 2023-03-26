package rpc

import (
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"google.golang.org/grpc"
)

var Client pb.AccountServiceClient

func init() {
	dial, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		return
	}

	//dial.Close()

	Client = pb.NewAccountServiceClient(dial)

	//_, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
}
