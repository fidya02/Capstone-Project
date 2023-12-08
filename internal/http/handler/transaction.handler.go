package handler

import (
	"net/http"

	"github.com/fidya02/Capstone-Project/entity"

	"github.com/fidya02/Capstone-Project/common"
	"github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/fidya02/Capstone-Project/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService service.TransactionUseCase
	paymentService     service.PaymentUseCase
}

func NewTransactionHandler(transactionService service.TransactionUseCase, paymentService service.PaymentUseCase) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		paymentService:     paymentService,
	}
}

func (h *TransactionHandler) CreateOrder(ctx echo.Context) error {
	var input struct {
		OrderID string `json:"order_id" validate:"required"`
		Amount  int64  `json:"amount" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*common.JwtCustomClaims)

	transaction := entity.NewTransaction(input.OrderID, claims.ID, input.Amount, "unpaid")

	err := h.transactionService.Create(ctx.Request().Context(), transaction)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	paymentRequest := entity.NewPaymentRequest(transaction.OrderID, transaction.Amount, claims.Name, "", claims.Email)

	payment, err := h.paymentService.CreateTransaction(ctx.Request().Context(), paymentRequest)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"url_pembayaran": payment})
}

func (h *TransactionHandler) WebHookTransaction(ctx echo.Context) error {
	var input entity.MidtransRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// cari transaction by order id kita dapat user id
	transaction, err := h.transactionService.FindByOrderID(ctx.Request().Context(), input.OrderID)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	status := "unpaid"

	if input.TransactionStatus == "settlement" {
		status = "paid"
	}

	err = h.transactionService.UpdateStatus(ctx.Request().Context(), transaction.OrderID, status)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "success"})
}
