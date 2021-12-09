package main

import (
	"movie/config"
	"movie/model"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
	if err := model.InitModel(); err != nil {
		panic(err)
	}
}
