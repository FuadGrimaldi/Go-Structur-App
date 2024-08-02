package entity

type User struct {
	ID      int64  `gorm:"primaryKey"`
	Name    string `gorm:"column:name"`
	Address string `gorm:"column:address"`
	Gender  string `gorm:"column:gender"`
	Email   string `gorm:"column:email"`
}

func (User) TableName() string {
	return "user_tb"
}
