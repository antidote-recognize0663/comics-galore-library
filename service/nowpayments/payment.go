package nowpayments

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"resty.dev/v3"
	"time"
)

type NowPayments interface {
	GetAvailableCurrencies() (*model.CurrenciesResponse, error)
	GetMerchantCoins() (*model.MerchantCoins, error)
	GetApiStatus() (*model.StatusResponse, error)
	CreateNowPayment(priceIndex int, priceCurrency string, payAmount float64, baseUrl string) (*model.PaymentResponse, error)
	GetEstimatedPrice(amount float64, currencyFrom, currencyTo string) (*model.EstimatedPrice, error)
	GetPaymentStatus(paymentID string) (*model.NowPaymentsIPN, error)
}

type nowPayments struct {
	apiKey            string
	endpoint          string
	subscriptionPlans []config.SubscriptionPlan
}

func NewNowPaymentsWithConfig(cfg *config.Config) NowPayments {
	return &nowPayments{
		apiKey:            cfg.Appwrite.ApiKey,
		endpoint:          cfg.Appwrite.Endpoint,
		subscriptionPlans: *cfg.Application.GetSubscriptionPlans(),
	}
}

func NewNowPayments(opts ...Option) NowPayments {
	cfg := &Config{
		endpoint:          "https://api.nowpayments.io/v1",
		subscriptionPlans: *config.NewSubscriptionPlans(),
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &nowPayments{
		apiKey:            cfg.apiKey,
		endpoint:          cfg.endpoint,
		subscriptionPlans: cfg.subscriptionPlans,
	}
}

func (n *nowPayments) GetAvailableCurrencies() (*model.CurrenciesResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("x-nowPayments-key", n.apiKey).
		SetResult(&model.CurrenciesResponse{}).
		Get(n.endpoint + "/currencies?fixed_rate=true")
	if err != nil {
		return nil, fmt.Errorf("GetAvailableCurrencies request failed: %w", err)
	}
	currenciesResponse := resp.Result().(*model.CurrenciesResponse)
	if currenciesResponse.Currencies == nil || len(currenciesResponse.Currencies) == 0 {
		return nil, fmt.Errorf("received an empty or missing currencies list")
	}
	return currenciesResponse, nil
}

func (n *nowPayments) GetMerchantCoins() (*model.MerchantCoins, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("x-nowPayments-key", n.apiKey).
		SetResult(&model.MerchantCoins{}).
		Get(n.endpoint + "/merchant/coins")
	if err != nil {
		return nil, fmt.Errorf("GetMerchantCoins request failed: %w", err)
	}
	currenciesResponse := resp.Result().(*model.MerchantCoins)
	if currenciesResponse.SelectedCurrencies == nil || len(currenciesResponse.SelectedCurrencies) == 0 {
		return nil, fmt.Errorf("received an empty or missing currencies list")
	}
	return currenciesResponse, nil
}

func (n *nowPayments) GetApiStatus() (*model.StatusResponse, error) {
	client := resty.New()
	resp, err := client.R().SetResult(&model.StatusResponse{}).Get(n.endpoint + "/status")
	if err != nil {
		return nil, fmt.Errorf("GetApiStatus request failed: %w", err)
	}
	return resp.Result().(*model.StatusResponse), nil
}

func (n *nowPayments) CreateNowPayment(priceIndex int, priceCurrency string, payAmount float64, baseUrl string) (*model.PaymentResponse, error) {

	description := n.subscriptionPlans[priceIndex].GetName()
	priceAmount := n.subscriptionPlans[priceIndex].GetPrice()

	ipnCallbackURL := fmt.Sprintf("%s/nowpayments/ipn", baseUrl)

	request := model.NewPaymentRequest(priceAmount, "usd", payAmount, priceCurrency, ipnCallbackURL, description)

	client := resty.New()
	resp, err := client.R().
		SetHeader("x-nowPayments-key", n.apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&model.PaymentResponse{}).
		Post(n.endpoint + "/payment")
	if err != nil {
		return nil, fmt.Errorf("CreateNowPayment request failed: %w", err)
	}
	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("CreateNowPayment returned a non-201 status code: %d", resp.StatusCode())
	}
	return resp.Result().(*model.PaymentResponse), nil
}

func (n *nowPayments) GetEstimatedPrice(amount float64, currencyFrom, currencyTo string) (*model.EstimatedPrice, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("x-nowPayments-key", n.apiKey).
		SetQueryParams(map[string]string{
			"amount":        fmt.Sprintf("%.4f", amount),
			"currency_from": currencyFrom,
			"currency_to":   currencyTo,
		}).
		SetResult(&model.EstimatedPrice{}).
		Get(n.endpoint + "/estimate")
	if err != nil {
		return nil, fmt.Errorf("GetEstimatedPrice request failed: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("GetEstimatedPrice returned a non-200 status code: %d", resp.StatusCode())
	}
	return resp.Result().(*model.EstimatedPrice), nil
}

// GetPaymentStatus fetches the payment status from the NowPayments API
func (n *nowPayments) GetPaymentStatus(paymentID string) (*model.NowPaymentsIPN, error) {
	if paymentID == "" {
		return nil, fmt.Errorf("paymentID cannot be empty")
	}
	client := resty.New().SetTimeout(10 * time.Second)
	resp, err := client.R().
		SetHeader("x-nowPayments-key", n.apiKey).
		SetResult(&model.NowPaymentsIPN{}).
		Get(fmt.Sprintf("%s/payment/%s", n.endpoint, paymentID))
	if err != nil {
		return nil, fmt.Errorf("GetPaymentStatus request to NowPayments failed: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("GetPaymentStatus returned a non-200 status code: %d", resp.StatusCode())
	}
	result, ok := resp.Result().(*model.NowPaymentsIPN)
	if !ok || result == nil {
		return nil, fmt.Errorf("failed to decode GetPaymentStatus response for paymentID: %s", paymentID)
	}
	return result, nil
}

func WithApiKey(apiKey string) Option {
	return func(c *Config) {
		c.apiKey = apiKey
	}
}

func WithEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.endpoint = endpoint
	}
}

type Config struct {
	apiKey            string
	endpoint          string
	subscriptionPlans []config.SubscriptionPlan
}

type Option func(*Config)
