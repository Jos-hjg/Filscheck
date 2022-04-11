package Filscout

import (
	"test/config"
)

type MiningInfoResp struct {
	Code int `json:"code"`
	Message string `json:"message"`
	MiningData MiningData `json:"data"`
}
type MiningData struct {
	Miner string `json:"miner"`
	QualityPowerGrowth int `json:"qualityPowerGrowth"`
	QualityPowerGrowthStr string `json:"qualityPowerGrowthStr"`
	ProvingPower int `json:"provingPower"`
	ProvingPowerStr string `json:"provingPowerStr"`
	MiningEfficiencyFloat float64 `json:"miningEfficiencyFloat"`
	MiningEfficiency string `json:"miningEfficiency"`
	MachinesNum int `json:"machinesNum"`
	Blocks int `json:"blocks"`
	BlockReward float64 `json:"blockReward"`
	BlockRewardStr string `json:"blockRewardStr"`
	BlockRewardPercent string `json:"blockRewardPercent"`
	LuckyValue float64 `json:"luckyValue"`
	StatsType string `json:"statsType"`
}

func (ff FILSCOUT) GetMiningData(Miner string) ([]byte, error) {
	payload := map[string]string{
		"statsType":"24h",
	}
	resp, err := ff.Requests.Post( config.C.Filscout.Scheme + "://" + config.C.Filscout.Host + config.C.Filscout.URI + Miner + "/mining-data", nil, payload)
	if err != nil{
		return make([]byte, 0), err
	}

	return resp, nil
}
