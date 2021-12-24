package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"movie/idl/gen/recommend"
)

var (
	RecommendClient recommend.RecommenderClient
)

var (
	ClientIP            = "127.0.0.1"
	RecommendClientPort = 50000
)

func InitRpcClients() error {
	recConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ClientIP, RecommendClientPort),
		grpc.WithInsecure())
	if err != nil {
		return err
	}
	RecommendClient = recommend.NewRecommenderClient(recConn)

	return nil
}
