package entity

import "time"

type Product struct {
	ID              int64    `gorm:"primaryKey"`
	Title           string   `gorm:"column:title"`
	Author          string   `gorm:"column:author"`
	Publicatio_year int64    `gorm:"column:publication_year"`
	Description     string   `gorm:"column:description"`
	Category        string   `gorm:"column:category"`
	ISBN            string   `gorm:"column:isbn"`
	Stoct           int64    `gorm:"column:stoct"`
	Price           float64    `gorm:"column:price"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (Product) TableName() string {
	return "product_tb"
}