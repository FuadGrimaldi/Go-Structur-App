package main

import (
	"fmt"
	"go-app/config"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	fmt.Println(cfg)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}