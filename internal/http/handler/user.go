package handler

import (
	"go-app/internal/service"
	"go-app/internal/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) FindAllUser(c echo.Context) error {
	users, err := h.userService.FindAll(c.Request().Context())
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succsesfully read all users", users)
}

func (h *UserHandler) FindOneUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
	}
	user, err := h.userService.FindOne(c.Request().Context(), id)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succsesfully read one user", user)
}