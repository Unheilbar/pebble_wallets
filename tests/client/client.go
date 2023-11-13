package client

import (
	"log"

	pb "github.com/Unheilbar/pebble_wallets/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(nodeurl string) pb.PebbleAPIClient {
	conn, err := grpc.Dial(nodeurl, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}...)
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewPebbleAPIClient(conn)
}
