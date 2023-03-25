package rpc

import (
	context "context"
	pb "github.com/Jazee6/treehole/cmd/account/rpc"
	"google.golang.org/grpc"
	"time"
)

var client pb.AccountServiceClient

func CallCreateUser() {
	dial, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer func(dial *grpc.ClientConn) {
		err := dial.Close()
		if err != nil {

		}
	}(dial)

	client = pb.NewAccountServiceClient(dial)

	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

}
