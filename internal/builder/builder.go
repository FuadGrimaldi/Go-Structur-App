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
func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	
	userService := service.NewUserService(cfg, userRepository)
	productService := service.NewProductService(productRepository)
	transactionService := service.NewTransactionService(transactionRepository)

	userHandler := handler.NewHandler(userService)
	productHandler := handler.NewProductHanlder(productService)
	transactionHandler := handler.NewTransactionHandler(transactionService)


	return router.PrivateRoutes(userHandler, productHandler, transactionHandler)
}

// no need login
func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route{
	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)

	productService := service.NewProductService(productRepository)
	authService := service.NewAuthService(cfg, userRepository)
	userService := service.NewUserService(cfg, userRepository)

	productHandler := handler.NewProductHanlder(productService)
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewHandler(userService)

	return router.PublicRoutes(userHandler, authHandler, productHandler)
}