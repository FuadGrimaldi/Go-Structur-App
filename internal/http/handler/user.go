package handler

import (
	"go-app/internal/common"
	"go-app/internal/dto"
	"go-app/internal/service"
	"go-app/internal/util"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GeneratePassword(c echo.Context) error {

	var request struct {
		Password string `json:"password"`
	}

	if err := c.Bind(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "successfully generate password", encodedPassword)
}

func (h *UserHandler) FindAllUser(c echo.Context) error {
	users, err := h.userService.FindAll(c.Request().Context())
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*common.JwtCustomClaims)
	// Check if the user's id same with jwt id
	if claims.Role != "admin" {
		return util.JSONResponse(c, http.StatusForbidden, "You don't have access to this user", nil)
	}
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
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*common.JwtCustomClaims)
	// Check if the user's id same with jwt id
	if claims.Role != "admin" {
		if id != claims.ID {
			return util.JSONResponse(c, http.StatusForbidden, "You don't have access to this user", nil)
		}
	}
	user, err := h.userService.FindOne(c.Request().Context(), id)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Successfully read one user", user)
}


func (h *UserHandler) CreateUser(c echo.Context) error {
	var request dto.NewUser

	if err := c.Bind(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	if err := h.userService.Create(c.Request().Context(), request); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	
	return util.JSONResponse(c, http.StatusCreated, "Succsesfully create user", request)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var request dto.UpdateUser
	
	if err := c.Bind(&request); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || userID == 0 {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
	}

	request.ID = userID
	
	if err := h.userService.Update(c.Request().Context(), request); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succsesfully update user", request)
	
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
	}
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*common.JwtCustomClaims)
	// Check if the user's id same with jwt id
	if claims.Role != "admin" {
		return util.JSONResponse(c, http.StatusForbidden, "You don't have access to this user", nil)
	}
	if err := h.userService.Delete(c.Request().Context(), id); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succsesfully delete user", nil)
}
