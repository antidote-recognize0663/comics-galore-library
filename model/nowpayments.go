package model

import (
	"github.com/appwrite/sdk-for-go/id"
)

type CurrenciesResponse struct {
	Currencies []Currency `json:"currencies"`
}

type MerchantCoins struct {
	SelectedCurrencies []string
}

type Currency struct {
	MinAmount float64 `json:"min_amount"`
	MaxAmount float64 `json:"max_amount"`
	Currency  string  `json:"currency"`
}

type StatusResponse struct {
	Message string `json:"message"`
}

type PaymentRequest struct {
	PriceAmount      float64  `json:"price_amount"`                  // (required) The fiat equivalent of the price to be paid in crypto
	PriceCurrency    string   `json:"price_currency"`                // (required) The fiat currency in which price_amount is specified (e.g., usd, eur)
	PayAmount        *float64 `json:"pay_amount,omitempty"`          // (optional) The crypto amount users need to pay; if empty, calculated automatically
	PayCurrency      string   `json:"pay_currency"`                  // (required) The cryptocurrency (e.g., btc, eth) or enabled fiat currency (e.g., usd, eur)
	IPNCallbackURL   *string  `json:"ipn_callback_url,omitempty"`    // (optional) URL to receive callbacks, e.g., "https://example.com"
	OrderID          *string  `json:"order_id,omitempty"`            // (optional) Inner store order fileId, e.g., "RGDBP-21314"
	OrderDescription *string  `json:"order_description,omitempty"`   // (optional) Inner store order description, e.g., "Apple Macbook Pro 2019 x 1"
	PayoutAddress    *string  `json:"payout_address,omitempty"`      // (optional) Address where funds are sent if different from the default
	PayoutCurrency   *string  `json:"payout_currency,omitempty"`     // (optional) Currency of the external payout_address, required when using payout_address
	PayoutExtraID    *string  `json:"payout_extra_id,omitempty"`     // (optional) Extra fileId, memo, or tag for the external payout_address
	IsFixedRate      *bool    `json:"is_fixed_rate,omitempty"`       // (optional) Boolean, true or false; required for fixed-rate exchanges
	IsFeePaidByUser  *bool    `json:"is_fee_paid_by_user,omitempty"` // (optional) Boolean, true or false; applicable for fixed-rate exchanges with user-paid fees
}

func NewPaymentRequest(priceAmount float64, priceCurrency string, payAmount float64, payCurrency string, ipnCallbackURL string, orderDescription string) *PaymentRequest {
	orderID := id.Unique()
	return &PaymentRequest{
		OrderID:          &orderID,
		PayAmount:        &payAmount,
		PriceAmount:      priceAmount,
		PayCurrency:      payCurrency,
		PriceCurrency:    priceCurrency,
		IPNCallbackURL:   &ipnCallbackURL,
		OrderDescription: &orderDescription,
	}
}

type PaymentResponse struct {
	PaymentID        string            `json:"payment_id"`
	PaymentStatus    PaymentStatusEnum `json:"payment_status"`
	PayAddress       string            `json:"pay_address"`
	PriceAmount      float64           `json:"price_amount"`
	PriceCurrency    string            `json:"price_currency"`
	PayAmount        float64           `json:"pay_amount"`
	PayCurrency      string            `json:"pay_currency"`
	OrderID          string            `json:"order_id"`
	OrderDescription string            `json:"order_description"`
	IPNCallbackURL   string            `json:"ipn_callback_url"`
}

type EstimatedPrice struct {
	AmountFrom      float64 `json:"amount_from"`
	CurrencyFrom    string  `json:"currency_from"`
	CurrencyTo      string  `json:"currency_to"`
	EstimatedAmount string  `json:"estimated_amount"`
}

type NowPaymentsIPN struct {
	PaymentID        int64             `json:"payment_id"`
	InvoiceID        string            `json:"invoice_id,omitempty"`
	PaymentStatus    PaymentStatusEnum `json:"payment_status"`
	PayAddress       string            `json:"pay_address"`
	PayinExtraID     string            `json:"payin_extra_id,omitempty"`
	PriceAmount      float64           `json:"price_amount"`
	PriceCurrency    string            `json:"price_currency"`
	PayAmount        float64           `json:"pay_amount"`
	ActuallyPaid     float64           `json:"actually_paid"`
	PayCurrency      string            `json:"pay_currency"`
	OrderID          string            `json:"order_id,omitempty"`
	OrderDescription string            `json:"order_description,omitempty"`
	PurchaseID       int64             `json:"purchase_id"`
	OutcomeAmount    float64           `json:"outcome_amount"`
	OutcomeCurrency  string            `json:"outcome_currency"`
	PayoutHash       string            `json:"payout_hash,omitempty"`
	PayinHash        string            `json:"payin_hash,omitempty"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
	BurningPercent   string            `json:"burning_percent,omitempty"`
	Type             string            `json:"type"`
	PaymentExtraIDs  []int64           `json:"payment_extra_ids"`
}
