package models

type Transaction struct {
	TransactionID string `gorm:"primaryKey" json:"transaction_id"`
	UserID        int    `json:"user_id"`
	Amount        int    `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
	ExpiresAt     string `json:"expires_at"`
	Balance       int    `json:"balance"`
}
