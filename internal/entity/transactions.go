package entity

import "time"

type Transaction struct {
	ID              int64    `gorm:"primaryKey"`
	UserID          int64    `gorm:"column:user_id"`
	PaymentID       int64    `gorm:"column:payment_id"`
	TransactionDate string   `gorm:"column:transaction_date"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	User            *User
	Payment         *TransactionPayment
	Details         []TransactionDetail
}

func (Transaction) TableName() string {
	return "transactions"
}