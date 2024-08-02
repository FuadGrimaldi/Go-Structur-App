package main

import (
	"fmt"
	"go-app/config"
	"go-app/pkg/database"
)

type User struct {
	ID int64
	Name string
	Address string
	Gender string
	Email string
}

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	db, err := database.ConnectToPostgres(cfg)
	checkError(err)

	users := make([]User, 0)
	err = db.Table("user_tb").Find(&users).Error
	checkError(err)

	fmt.Println(users)
	fmt.Println(cfg)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}