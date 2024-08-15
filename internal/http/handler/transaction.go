package handler

import (
	"go-app/internal/service"
	"go-app/internal/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) FindTransactionByUserID(c echo.Context) error {
	idStr := c.Param("id")
	UserID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return util.JSONResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
	}

	transactions, err := h.transactionService.FindTransactionByUserID(c.Request().Context(), UserID)

	if err != nil {
		if err == service.ErrNoTransactionsFound {
			return util.JSONResponse(c, http.StatusNotFound, "No transactions found for the given user ID", nil)
		}
		return util.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return util.JSONResponse(c, http.StatusOK, "Successfully retrieved transaction", transactions)
}