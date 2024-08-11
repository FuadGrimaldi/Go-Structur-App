package entity

import "time"

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Name      string `gorm:"column:name"`
	Address   string `gorm:"column:address"`
	Gender    string `gorm:"column:gender"`
	Email     string `gorm:"column:email"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Role 	  string `gorm:"column:role"`
}

func (User) TableName() string {
	return "user_tb"
}
