package rpcs

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func Log() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			log.Printf("Err %s %v\n", info.FullMethod, err)
		}
		//log.Printf("Req %s %v\n", info.FullMethod, req)
		return
	}
}
