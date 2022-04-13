package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"test/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	config.AllowOrigins = []string{"http://localhost:63342"}

	r.Use(cors.New(config))
	//TODO: any router?
	r.GET("/miner", controller.GetMinerInfo)
	r.GET("/mining-data", controller.GetMiningData)
	return r
}

