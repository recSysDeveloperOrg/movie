package main

import (
	"movie/config"
	"movie/model"
	"movie/rpc"
	"movie/search"
)

func main() {
	if err := config.InitConfig(config.CfgFileMain); err != nil {
		panic(err)
	}
	if err := model.InitModel(); err != nil {
		panic(err)
	}
	if err := model.Disconnect(); err != nil {
		panic(err)
	}
	if err := rpc.InitRpcClients(); err != nil {
		panic(err)
	}
	if err := search.InitES(); err != nil {
		panic(err)
	}
}
