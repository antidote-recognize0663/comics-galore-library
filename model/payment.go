package model

import "github.com/appwrite/sdk-for-go/models"

// PaymentStatus represents the payment status response
type PaymentStatus struct {
	PaymentID       string            `json:"payment_id"`
	PaymentStatus   PaymentStatusEnum `json:"payment_status"`
	PayAddress      string            `json:"pay_address"`
	PriceAmount     float64           `json:"price_amount"`
	PriceCurrency   string            `json:"price_currency"`
	PayAmount       float64           `json:"pay_amount"`
	ActuallyPaid    float64           `json:"actually_paid"`
	PayCurrency     string            `json:"pay_currency"`
	CreatedAt       string            `json:"created_at"`
	UpdatedAt       string            `json:"updated_at"`
	PurchaseID      string            `json:"purchase_id"`
	OutcomeCurrency string            `json:"outcome_currency"`
	OutcomeAmount   float64           `json:"outcome_amount"`
}

type PaymentList struct {
	*models.DocumentList
	*PaymentData
}

type PaymentData struct {
	UserId           string `json:"user_id"`
	OrderId          string `json:"order_id"`
	QrCodeUrl        string `json:"qr_code_url"`
	PaymentId        string `json:"payment_id"`
	PayAmount        string `json:"pay_amount"`
	PayAddress       string `json:"pay_address"`
	PayCurrency      string `json:"pay_currency"`
	PriceAmount      string `json:"price_amount"`
	PaymentStatus    string `json:"payment_status"`
	PriceCurrency    string `json:"price_currency"`
	OrderDescription string `json:"order_description"`
	ExpirationDate   string `json:"expiration_date"`
	Expired          bool   `json:"expired"`
}

type Payment struct {
	*models.Document
	*PaymentData
}
