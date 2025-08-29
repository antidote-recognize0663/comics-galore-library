package model

import (
	"fmt"
	"github.com/appwrite/sdk-for-go/models"
	"time"
)

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

type PaymentData struct {
	UserID           string            `json:"user_id,omitempty"`
	OrderID          string            `json:"order_id,omitempty"`
	QRCodeURL        string            `json:"qr_code_url,omitempty"`
	PaymentID        string            `json:"payment_id"`
	PayAmount        float64           `json:"pay_amount"`
	PayAddress       string            `json:"pay_address,omitempty"`
	PayCurrency      string            `json:"pay_currency,omitempty"`
	PriceAmount      float64           `json:"price_amount,omitempty"`
	PriceCurrency    string            `json:"price_currency,omitempty"`
	OrderDescription string            `json:"order_description,omitempty"`
	Expired          bool              `json:"expired,omitempty"`
	PaymentStatus    PaymentStatusEnum `json:"payment_status"`
	ExpirationDate   string            `json:"expiration_date"`
	ExpiresAt        string            `json:"expires_at,omitempty"`
	PayinHash        string            `json:"payin_hash,omitempty"`
	PayoutHash       string            `json:"payout_hash,omitempty"`
	ActuallyPaid     float64           `json:"actually_paid,omitempty"`
}

type Payment struct {
	*models.Document
	*PaymentData
}

type PaymentList struct {
	*models.DocumentList
	Payments []Payment `json:"documents"`
}

func NewPayment(document *models.Document) (*Payment, error) {
	var paymentData PaymentData
	if err := document.Decode(&paymentData); err != nil {
		return nil, err
	}
	return &Payment{
		Document:    document,
		PaymentData: &paymentData,
	}, nil
}

func NewPaymentData(userId string, payment *PaymentResponse) (*PaymentData, error) {
	durationMap := map[float64]Duration{
		10:  {0, 1, 0}, // 1 month
		20:  {0, 2, 0}, // 2 months
		30:  {0, 3, 0}, // 3 months
		60:  {0, 6, 0}, // 6 months
		120: {1, 0, 0}, // 1 year
	}
	expiresAt, err := getExpirationDate(time.Now(), payment.PriceAmount, durationMap)
	if err != nil {
		return nil, err
	}
	//expiresAt := expiresAt.Format(time.RFC3339)
	qrCodeURL := buildQRCodeURL(payment.PayCurrency, payment.PayAddress, payment.PriceAmount, payment.PaymentID)
	return &PaymentData{
		UserID:           userId,
		OrderID:          payment.OrderID,
		PaymentID:        payment.PaymentID,
		PayAmount:        payment.PayAmount,
		PayAddress:       payment.PayAddress,
		PayCurrency:      payment.PayCurrency,
		PriceAmount:      payment.PriceAmount,
		PaymentStatus:    payment.PaymentStatus,
		PriceCurrency:    payment.PriceCurrency,
		OrderDescription: payment.OrderDescription,
		QRCodeURL:        qrCodeURL,
		ExpiresAt:        expiresAt,
	}, nil
}

type Duration struct {
	Years  int
	Months int
	Days   int
}

func getExpirationDate(startTime time.Time, amount float64, durationMap map[float64]Duration) (string, error) {
	duration, exists := durationMap[amount]
	if !exists {
		return "", fmt.Errorf("invalid amount: no duration configured for %.2f", amount)
	}
	expirationDate := startTime.AddDate(duration.Years, duration.Months, duration.Days)
	return expirationDate.Format(time.RFC3339), nil
}

func buildQRCodeURL(currency, address string, amount float64, id string) string {
	const defaultSize = "200"
	const defaultMargin = "0"
	return fmt.Sprintf("/qrcode/%s/%s?amount=%.2f&label=%s&size=%s&margin=%s",
		currency, address, amount, id, defaultSize, defaultMargin)
}
