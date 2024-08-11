package router

import (
	"go-app/internal/http/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Route struct {
	Method 	string
	Path 	string
	Handler	echo.HandlerFunc
}

func PublicRoutes(UserHandler *handler.UserHandler, AuthHandler *handler.AuthHandler) []*Route {
	return []*Route{
		{
			Method: http.MethodPost,
			Path: "/login",
			Handler: AuthHandler.Login,
		},
		{
			Method: http.MethodGet,
			Path: "/users",
			Handler: UserHandler.FindAllUser,
		},
	}
}

func PrivateRoutes(UserHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method: http.MethodGet,
			Path: "/users/:id",
			Handler: UserHandler.FindOneUser,
		},
		{
			Method: http.MethodPost,
			Path: "/users",
			Handler: UserHandler.CreateUser,
		},
		{
			Method: http.MethodPut,
			Path: "/users/:id",
			Handler: UserHandler.UpdateUser,
		},
		{
			Method: http.MethodDelete,
			Path: "/users/:id",
			Handler: UserHandler.DeleteUser,
		},
	}
}