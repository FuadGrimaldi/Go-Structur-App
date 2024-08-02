package router

import (
	"go-app/internal/http/handler"

	"github.com/labstack/echo/v4"
)

type Route struct {
	Method 	string
	Path 	string
	Handler	echo.HandlerFunc
}

func PublicRoutes(UserHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method: echo.GET,
			Path: "/users",
			Handler: UserHandler.FindAllUser,
		},
	}
}

func PrivateRoutes() []*Route {
	return []*Route{}

}