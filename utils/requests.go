package utils

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Requests struct {
	Verbose bool
	Logger  *zap.Logger
	Headers map[string]string
}

func NewRequests(logger *zap.Logger, verbose bool, headers map[string]string) Requests {
	return Requests{
		Verbose: verbose,
		Logger:  logger,
		Headers: headers,
	}
}

func (r Requests) do(url, method string, param map[string]string, body interface{}) ([]byte, error) {
	var (
		err  error
		resp *resty.Response
		res  []byte
	)
	req := resty.New().SetDebug(r.Verbose).R().SetHeaders(r.Headers)
	switch method {
	case resty.MethodGet:
		resp, err = req.SetQueryParams(param).SetBody(body).Get(url)
	case resty.MethodPost:
		resp, err = req.SetQueryParams(param).SetBody(body).Post(url)
	default:
		return res, errors.New("method not support")
	}

	if err != nil {
		return res, err
	}

	return resp.Body(), nil
}

func (r Requests) Get(url string, param, body map[string]string) ([]byte, error) {
	resp, err := r.do(url, "GET", param, body)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (r Requests) Post(url string, param map[string]string, body interface{}) ([]byte, error) {
	resp, err := r.do(url, "POST", param, body)
	if err != nil {
		return nil, err
	}
	return resp, err
}
