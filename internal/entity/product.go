package entity

type Product struct {
	ID              int64  `gorm:"primaryKey"`
	Title           string `gorm:"column:title"`
	Author          string `gorm:"column:author"`
	Publicatio_year int64  `gorm:"column:publication_year"`
	Description     string `gorm:"column:description"`
	Category        string `gorm:"column:category"`
	ISBN            string `gorm:"column:isbn"`
	Stoct           int64  `gorm:"column:stoct"`
	Price           int64  `gorm:"column:price"`
}

func (Product) TableName() string {
	return "product_tb"
}