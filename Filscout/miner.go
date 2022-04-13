package Filscout

import (
	"test/config"
)

type MinerInfoResp struct {
	Code int `json:"code"`
	Message string `json:"message"`
	MinerData MinerData `json:"data"`
}
type MinerData struct {
	Miner string `json:"miner"`
	Address string `json:"address"`
	IDAddress string `json:"idAddress"`
	RobustAddress string `json:"robustAddress"`
	ActorType string `json:"actorType"`
	Balance float64 `json:"balance"`
	BalanceStr string `json:"balanceStr"`
	QualityPower float64 `json:"qualityPower"`
	QualityPowerStr string `json:"qualityPowerStr"`
	QualityPowerPercent float64 `json:"qualityPowerPercent"`
	QualityPowerPercentStr string `json:"qualityPowerPercentStr"`
	RawPower float64 `json:"rawPower"`
	RawPowerStr string `json:"rawPowerStr"`
	PeerID string `json:"peerId"`
	Owner string `json:"owner"`
	Worker string `json:"worker"`
	Available float64 `json:"available"`
	AvailableStr string `json:"availableStr"`
	SectorsPledge float64 `json:"sectorsPledge"`
	SectorsPledgeStr string `json:"sectorsPledgeStr"`
	LockedFunds float64 `json:"lockedFunds"`
	LockedFundsStr string `json:"lockedFundsStr"`
	FeeDebtStr string `json:"feeDebtStr"`
	IP string `json:"ip"`
	Location string `json:"location"`
	SectorSize float64 `json:"sectorSize"`
	SectorSizeStr string `json:"sectorSizeStr"`
	SectorCount int `json:"sectorCount"`
	ActiveCount int `json:"activeCount"`
	FaultCount int `json:"faultCount"`
	RecoveryCount int `json:"recoveryCount"`
	Tag string `json:"tag"`
	IsVerified int `json:"isVerified"`
	MsgCount int `json:"msgCount"`
	Blocks int `json:"blocks"`
	WinCount int `json:"winCount"`
	BlockReward float64 `json:"blockReward"`
	BlockRewardStr string `json:"blockRewardStr"`
	PowerRank int `json:"powerRank"`
	CreateTime string `json:"createTime"`
}

func (ff FILSCOUT) GetMinerInfos(Miner string) ([]byte, error) {
	resp, err := ff.Requests.Post(config.C.Filscout.Scheme + "://" + config.C.Filscout.Host + config.C.Filscout.URI + Miner, nil, nil)
	if err != nil{
		return make([]byte, 0), err
	}
	return resp, nil
}