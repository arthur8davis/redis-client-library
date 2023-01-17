package bootstrap

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func NewServerGrpc() *grpc.ClientConn {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_GATEWAY_GRPC_HOST"), os.Getenv("REDIS_GATEWAY_GRPC_PORT"))

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}
