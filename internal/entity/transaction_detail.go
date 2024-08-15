package entity

type TransactionDetail struct {
	ID            int64   `gorm:"primaryKey"`
	TransactionID int64   `gorm:"column:transaction_id"`
	ProductID     int64   `gorm:"column:product_id"`
	Qty           int64   `gorm:"column:qty"`
	Price         float64 `gorm:"column:price"`
	Product       *Product
}