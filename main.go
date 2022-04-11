package main

import (
	"go.uber.org/zap"
	"test/Filscout"
	"test/config"
	"test/router"
)

func main() {
	config.Init("./config/config.yaml")
	Filscout.NewRequests(zap.L(), config.C.Filscout)

	r := router.InitRouter()
	r.Run(config.C.Router.Port)

}