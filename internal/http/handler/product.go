package handler

import (
	"go-app/internal/service"
	"go-app/internal/util"
	"net/http"

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