package domain

import "time"

type TransactionOtp struct {
	ID         string    `json:"id"`
	UserId     string    `json:"user_id"`
	Pin        string    `json:"otp"`
	Used       bool      `json:"used"`
	ExpiryDate time.Time `json:"expiry_date"`
}
