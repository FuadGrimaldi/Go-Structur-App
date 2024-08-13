package router

import (
	"go-app/internal/http/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	Admin = "admin"
	User = "user"
)

var (
	allRoles  = []string{Admin, User}
	onlyAdmin = []string{Admin}
	onlyUser  = []string{User}
)
type Route struct {
	Method 	string
	Path 	string
	Handler	echo.HandlerFunc
	Roles   []string
}



func PublicRoutes(UserHandler *handler.UserHandler, AuthHandler *handler.AuthHandler, ProductHandler *handler.ProductHandler) []*Route {
	return []*Route{
		{
			Method: http.MethodPost,
			Path: "/login",
			Handler: AuthHandler.Login,
			Roles: allRoles,
		},
		{
			Method: http.MethodPost,
			Path: "/users",
			Handler: UserHandler.CreateUser,
			Roles: onlyUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/generate-password",
			Handler: UserHandler.GeneratePassword,
		},
	}
}

func PrivateRoutes(UserHandler *handler.UserHandler, ProductHandler *handler.ProductHandler) []*Route {
	return []*Route{
		{
			Method: http.MethodGet,
			Path: "/users/:id",
			Handler: UserHandler.FindOneUser,
			Roles: allRoles,
		},
		{
			Method: http.MethodGet,
			Path: "/users",
			Handler: UserHandler.FindAllUser,
			Roles:   onlyAdmin,
		},
		{
			Method: http.MethodPut,
			Path: "/users/:id",
			Handler: UserHandler.UpdateUser,
			Roles: allRoles,
		},
		{
			Method: http.MethodDelete,
			Path: "/users/:id",
			Handler: UserHandler.DeleteUser,
			Roles: onlyAdmin,
		},
		{
			Method: http.MethodPost,
			Path: "/products",
			Handler: ProductHandler.Create,
			Roles: onlyAdmin,
		},
		{
			Method: http.MethodGet,
			Path: "/products",
			Handler: ProductHandler.FindAllProduct,
			Roles:   onlyAdmin,
		},
		{
			Method: http.MethodGet,
			Path: "/products/:id",
			Handler: ProductHandler.FindOneProductById,
			Roles: allRoles,
		},
	}
}