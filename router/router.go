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
	config.AllowOrigins = []string{"http://localhost"}
	config.AllowHeaders = []string{
		"authorization",
	}
	r.Use(cors.New(config))
	//TODO: any router?
	r.GET("/miner", controller.GetMinerInfo)
	r.GET("/mining-data", controller.GetMiningData)
	return r
}

