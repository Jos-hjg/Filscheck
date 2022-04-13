package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"test/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	//TODO: any router?
	r.GET("/miner", controller.GetMinerInfo)
	r.GET("/mining-data", controller.GetMiningData)
	return r
}
