package Filscout

import (
	"go.uber.org/zap"
	"test/config"
	"test/utils"
)

type FILSCOUT struct {
	Requests utils.Requests
	logger   *zap.Logger
	cfg      config.FilscoutData
}

var Filscout *FILSCOUT = &FILSCOUT{}

func NewRequests(logger *zap.Logger, cfg config.FilscoutData) {
	headers := make(map[string]string)
	headers["Host"] = config.C.Filscout.Host
	headers["Accept"] = "*/*"
	headers["user-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36"
	headers["Content-Type"] = "application/json;charset=UTF-8"
	headers["Origin"] = "https://www.filscout.com"
	headers["Accept-Encode"] = "gzip, deflate, br"
	headers["Accept-Language"] = "zh-CN,zh;q=0.9"
	headers["Referer"] = "https://www.filscout.com/"
	headers["token"] = ""
	// todo 是否还需要加其他请求头？！！

	req := utils.NewRequests(logger, cfg.Verbose, headers)
	Filscout = &FILSCOUT{
		Requests: req,
		logger:   logger,
		cfg:      cfg,
	}
}
