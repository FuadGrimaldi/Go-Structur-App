package builder

import (
	"go-app/config"
	"go-app/internal/http/handler"
	"go-app/internal/http/router"
	"go-app/internal/repository"
	"go-app/internal/service"

	"gorm.io/gorm"
)

// need login
func BuildPrivateRoutes() []*router.Route {
	return router.PrivateRoutes()
}

// no need login
func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route{
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(cfg, userRepository)
	userHandler := handler.NewHandler(userService)
	return router.PublicRoutes(userHandler)
}