package main

import (
	"go-app/config"
	"go-app/internal/builder"
	"go-app/pkg/database"
	"go-app/pkg/server"
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

	_, err = database.ConnectToPostgres(cfg)
	checkError(err)

	_ = builder.BuildPublicRoutes()
	_ = builder.BuildPrivateRoutes()

	srv := server.NewServer(cfg)
	srv.Run()
	srv.GracefulShutdown()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}