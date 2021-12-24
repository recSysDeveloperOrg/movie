package main

import (
	"movie/config"
	"movie/model"
	"movie/rpc"
)

func main() {
	if err := config.InitConfig(); err != nil {
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
}
