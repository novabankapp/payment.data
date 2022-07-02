package domain

import "time"

type PaymentProvider struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	IsInternal bool      `json:"is_internal"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `json:"created_at"`
}

func (p PaymentProvider) IsNoSQLEntity() bool {
	return true
}
