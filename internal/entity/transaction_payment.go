package entity

type TransactionPayment struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}