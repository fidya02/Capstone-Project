package handler

import (
	"errors"
	"net/http"

	"github.com/fidya02/Capstone-Project/common"
	"github.com/fidya02/Capstone-Project/entity"
	"github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/fidya02/Capstone-Project/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderService service.OrderUsecase
}

// NewOrderHandler creates a new instance of OrderHandler with the given OrderService.
// It takes an OrderService as a parameter and returns a pointer to an OrderHandler.
func NewOrderHandler(OrderService service.OrderUsecase) *OrderHandler {
	return &OrderHandler{OrderService}
}

// CreateOrder handles the creation of an order.
// It takes a context and returns an error.
func (h *OrderHandler) CreateOrder(ctx echo.Context) error {
	// Define the input struct with the required fields
	type createOrderInput struct {
		TicketID int64  `json:"ticket_id" validate:"required"`
		Quantity int64  `json:"quantity" validate:"required"`
		UserID   int64  `json:"user_id" validate:"required"`
		Status   string `json:"status" validate:"required"`
	}

	// Bind the request body to the input struct
	var input createOrderInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Get the user balance before creating the order
	userBalance, err := h.OrderService.GetUserBalance(ctx.Request().Context(), input.UserID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Get the ticket price from the TicketService
	ticketPrice, err := h.OrderService.GetTicketPrice(ctx.Request().Context(), input.TicketID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Check if the balance is sufficient to create the order
	if userBalance < (input.Quantity * ticketPrice) {
		return ctx.JSON(http.StatusUnprocessableEntity, errors.New("insufficient balance"))
	}

	// Create a new order entity
	order := entity.NewOrder(input.TicketID, input.Quantity, input.UserID, input.Status)

	// Create the order using the OrderService
	err = h.OrderService.CreateOrder(ctx.Request().Context(), order)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Update the user balance after creating the order
	err = h.OrderService.UpdateUserBalance(ctx.Request().Context(), input.UserID, input.Quantity*ticketPrice)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusCreated, "Order created successfully")
}

// GetAllOrders retrieves all orders from the usecase and returns them as a JSON response.
func (h *OrderHandler) GetAllOrders(ctx echo.Context) error {
	// Get all orders from the usecase
	orders, err := h.OrderService.GetOrders(ctx.Request().Context())
	if err != nil {
		// Return a JSON response with a 400 status code and an error message if there was an error getting the orders
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	var orderDetails []map[string]interface{}
	for _, order := range orders {
		// Get the ticket details by ID from the usecase
		ticket, err := h.OrderService.GetTicketByID(ctx.Request().Context(), order.TicketID)
		if err != nil {
			// Return a JSON response with a 500 status code and an error message if there was an error getting the ticket details
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
		}

		// Create a map with the order details
		orderDetail := map[string]interface{}{
			"user_id": order.UserID,
			"ticket":  ticket,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	// Return a JSON response with a 200 status code and the order details
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}

// GetOrderByUserID retrieves all orders for a given user ID.
func (h *OrderHandler) GetOrderByUserID(ctx echo.Context) error {
	// Get all orders from the OrderService.
	orders, err := h.OrderService.GetOrders(ctx.Request().Context())
	if err != nil {
		// Return a JSON response with the error message if there was an error.
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	var orderDetails []map[string]interface{}
	for _, order := range orders {
		// Get the ticket details for each order.
		ticket, err := h.OrderService.GetTicketByID(ctx.Request().Context(), order.TicketID)
		if err != nil {
			// Return a JSON response with the error message if there was an error.
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
		}

		// Create a map of order details with user ID and ticket information.
		orderDetail := map[string]interface{}{
			"user_id": order.UserID,
			"ticket":  ticket,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	// Return a JSON response with the order details.
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}

// UserCreateOrder creates a new order for a user
func (h *OrderHandler) UserCreateOrder(ctx echo.Context) error {
	// Define the input struct
	var input struct {
		UserID   int64
		TicketID int64  `json:"ticket_id" validate:"required"`
		Quantity int64  `json:"quantity" validate:"required"`
		Status   string `json:"status" default:"success"`
	}

	// Get JWT token from the context
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, errors.New("missing or invalid token"))
	}

	// Extract claims from the JWT token
	claims, ok := token.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, errors.New("invalid token claims"))
	}

	// Assign UserID from the JWT claims to the input struct
	input.UserID = claims.ID

	// Bind the request body to the input struct
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Get user balance before creating the order
	userBalance, err := h.OrderService.GetUserBalance(ctx.Request().Context(), input.UserID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Get ticket price from TicketService
	ticketPrice, err := h.OrderService.GetTicketPrice(ctx.Request().Context(), input.TicketID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Check if the user has sufficient balance to create the order
	if userBalance < (input.Quantity * ticketPrice) {
		return ctx.JSON(http.StatusUnprocessableEntity, errors.New("insufficient balance"))
	}

	// Create a new order entity
	order := entity.NewOrder(input.TicketID, input.Quantity, input.UserID, input.Status)

	// Create the order
	err = h.OrderService.CreateOrder(ctx.Request().Context(), order)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Update user balance after creating the order
	err = h.OrderService.UpdateUserBalance(ctx.Request().Context(), input.UserID, input.Quantity*ticketPrice)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusCreated, "Order created successfully")
}

// GetOrderHistory retrieves the order history of a user using JWT authentication.
func (h *OrderHandler) GetOrderHistory(ctx echo.Context) error {
	// Get JWT token from the context
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, errors.New("missing or invalid token"))
	}

	// Extract claims from the JWT token
	claims, ok := token.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, errors.New("invalid token claims"))
	}

	// Get all orders by user ID
	orders, err := h.OrderService.GetOrderByUserID(ctx.Request().Context(), claims.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	// Prepare order details
	var orderDetails []map[string]interface{}
	for _, order := range orders {
		ticket, err := h.OrderService.GetTicketByID(ctx.Request().Context(), order.TicketID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
		}

		orderDetail := map[string]interface{}{
			"user_id": order.UserID,
			"ticket":  ticket,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	// Return order history as JSON response
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}
