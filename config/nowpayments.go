package config

type NowPaymentsConfig struct {
	ApiKey    string
	Endpoint  string
	IPNSecret string
}

func NewNowPaymentsConfig() *NowPaymentsConfig {
	return &NowPaymentsConfig{
		ApiKey:    GetEnv("NOW_PAYMENTS_API_KEY", ""),
		Endpoint:  GetEnv("NOW_PAYMENTS_ENDPOINT", ""),
		IPNSecret: GetEnv("NOW_PAYMENTS_IPN_SECRET", ""),
	}
}
