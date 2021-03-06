package service

import (
	"movie/config"
	"movie/model"
	"movie/rpc"
	"movie/search"
	"testing"
)

const (
	ratedUser  = "100000000000000000000001"
	taggedUser = "800000000000000000000002"
	newUser    = "800000000000000000000001"
)

func TestMain(m *testing.M) {
	if err := config.InitConfig(config.CfgFileNested); err != nil {
		panic(err)
	}
	if err := model.InitModel(); err != nil {
		panic(err)
	}
	if err := rpc.InitRpcClients(); err != nil {
		panic(err)
	}
	if err := search.InitES(); err != nil {
		panic(err)
	}
	if code := m.Run(); code != 0 {
		panic(code)
	}
}
