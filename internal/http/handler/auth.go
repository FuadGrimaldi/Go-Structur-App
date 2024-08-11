package handler

import (
	"fmt"
	"go-app/internal/dto"
	"go-app/internal/service"
	"go-app/internal/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}


func (h * AuthHandler) Login(c echo.Context) error {
	var request dto.LoginRequest

	if err := c.Bind(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	token, err := h.authService.Login(c.Request().Context(), request) 
	if err != nil {
		return util.JSONResponse(c, http.StatusUnauthorized, err.Error(), nil)
	}
	
	message := fmt.Sprintf("token: %s", token)

	return util.JSONResponse(c, http.StatusOK,message, nil)
}