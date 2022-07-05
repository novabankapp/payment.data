package domain

import "time"

type UserToPaymentLink struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p UserToPaymentLink) IsNoSQLEntity() bool {
	return true
}
