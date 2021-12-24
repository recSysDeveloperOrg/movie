package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"movie/idl/gen/recommend"
	"movie/idl/gen/tag"
)

var (
	RecommendClient recommend.RecommenderClient
	TagClient       tag.TagServiceClient
)

var (
	ClientIP            = "127.0.0.1"
	RecommendClientPort = 50000
	TagClientPort       = 50001
)

func InitRpcClients() error {
	recConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ClientIP, RecommendClientPort))
	if err != nil {
		return err
	}
	RecommendClient = recommend.NewRecommenderClient(recConn)

	tagConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ClientIP, TagClientPort))
	if err != nil {
		return err
	}
	TagClient = tag.NewTagServiceClient(tagConn)

	return nil
}
