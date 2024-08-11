package main

import (
	"go-app/config"
	"go-app/internal/builder"
	"go-app/pkg/database"
	"go-app/pkg/server"
)


func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	db, err := database.ConnectToPostgres(cfg)
	checkError(err)

	publicRoutes := builder.BuildPublicRoutes(cfg, db)
	privateRoute := builder.BuildPrivateRoutes(cfg, db)

	srv := server.NewServer(cfg, publicRoutes, privateRoute)
	srv.Run()
	srv.GracefulShutdown()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}