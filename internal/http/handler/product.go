package handler

import (
	"go-app/internal/dto"
	"go-app/internal/service"
	"go-app/internal/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHanlder(productHandler service.ProductService) *ProductHandler {
	return &ProductHandler{productHandler}
}

func (ph *ProductHandler) FindAllProduct(c echo.Context) error {
	products, err := ph.productService.FindAll(c.Request().Context())
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Succsesfully read all users", products)
}

func (ph *ProductHandler) FindOneProductById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
	}
	product, err := ph.productService.FindOneById(c.Request().Context(), id)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Successfully read one product by id", product)
}

func (ph *ProductHandler) FindOneProductByTitle(c echo.Context) error {
	title := c.Param("title")
	
	product, err := ph.productService.FindOneByTitle(c.Request().Context(), title)
	if err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusOK, "Successfully read one product by title", product)
}

func (ph *ProductHandler) Create(c echo.Context) error {
	var req dto.NewProduct

	if err := c.Bind(&req); err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := ph.productService.Create(c.Request().Context(), req); err != nil {
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.JSONResponse(c, http.StatusCreated, "Succsesfully create user", req)
}