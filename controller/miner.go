package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"test/Filscout"
	"test/config"
)

type minerRes struct {
	mu sync.Mutex
	wait sync.WaitGroup
	res []Filscout.MinerData
}


func GetMinerInfo(ctx *gin.Context) {
	res := minerRes{}
	res.wait.Add(len(config.C.Filscout.Miners))
	for _, miner := range config.C.Filscout.Miners {
		go func(miner string) {
			resp, err := Filscout.Filscout.GetMinerInfos(miner)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"StatusCode": http.StatusNotFound,
					"StatusDesc": err,
					"data": "",
				})
			}
			rr := Filscout.MinerInfoResp{}
			if err := json.Unmarshal(resp, &rr); err != nil {
				log.Println("unmarshal miner failed!", err)
			}
			res.mu.Lock()
			res.res = append(res.res, rr.MinerData)
			res.mu.Unlock()
			res.wait.Done()
		}(miner)

	}
	res.wait.Wait()
	ctx.JSON(http.StatusOK, gin.H{
		"StatusCode": http.StatusOK,
		"StatusDesc": "success",
		"data": res.res,
	})
}

