package models

type User struct {
	UserID        int    `gorm:"primaryKey;column:user_id" json:"user_id"`
	Amount        int    `json:"amount"`
	PaymentMethod string `json:"payment_method"`
}

func (User) TableName() string {
	return "users"
}
