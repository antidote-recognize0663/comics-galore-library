package config

import "github.com/antidote-recognize0663/comics-galore-library/config/utils"

type NowPaymentsConfig struct {
	ApiKey    string
	Endpoint  string
	IPNSecret string
}

func NewNowPaymentsConfig() *NowPaymentsConfig {
	return &NowPaymentsConfig{
		ApiKey:    utils.GetEnv("NOW_PAYMENTS_API_KEY", ""),
		Endpoint:  utils.GetEnv("NOW_PAYMENTS_ENDPOINT", ""),
		IPNSecret: utils.GetEnv("NOW_PAYMENTS_IPN_SECRET", ""),
	}
}
