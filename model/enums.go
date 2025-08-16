package model

import (
	"errors"
)

type PaymentStatusEnum string

const (
	StatusWaiting       PaymentStatusEnum = "waiting"
	StatusConfirming    PaymentStatusEnum = "confirming"
	StatusConfirmed     PaymentStatusEnum = "confirmed"
	StatusSending       PaymentStatusEnum = "sending"
	StatusPartiallyPaid PaymentStatusEnum = "partially_paid"
	StatusFinished      PaymentStatusEnum = "finished"
	StatusFailed        PaymentStatusEnum = "failed"
	StatusRefunded      PaymentStatusEnum = "refunded"
	StatusExpired       PaymentStatusEnum = "expired"
)

func (p PaymentStatusEnum) IsValid() bool {
	switch p {
	case StatusWaiting, StatusConfirming, StatusConfirmed, StatusSending,
		StatusPartiallyPaid, StatusFinished, StatusFailed, StatusRefunded, StatusExpired:
		return true
	}
	return false
}

func (p PaymentStatusEnum) Validate() error {
	if !p.IsValid() {
		return errors.New("invalid payment status")
	}
	return nil
}
