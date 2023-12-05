package handler

import (
	"errors"
	"net/http"

	"github.com/fidya02/Capstone-Project/entity"
	"github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/fidya02/Capstone-Project/internal/service" //

	// mengimpor paket jwt yang berisi fungsi-fungsi untuk membuat dan mengelola token JWT
	"github.com/labstack/echo/v4" // mengimpor paket echo yang berisi fungsi-fungsi untuk membuat dan mengelola web server
)

type OrderHandler struct {
	OrderService service.OrderUsecase // mendefinisikan field OrderService yang bertipe interface OrderUsecase dari paket service
}

func NewOrderHandler(OrderService service.OrderUsecase) *OrderHandler {
	return &OrderHandler{OrderService} // mengembalikan pointer ke struct OrderHandler yang berisi field OrderService
}

func (h *OrderHandler) CreateOrder(ctx echo.Context) error { // mendefinisikan fungsi CreateOrder yang menerima parameter ctx yang bertipe echo.Context dan mengembalikan error
	var input struct { // mendefinisikan struct anonim untuk menampung input dari user
		TicketID int64  `json:"ticket_id" validate:"required"` // mendefinisikan field TicketID yang bertipe int64 dan memiliki tag json dan validate
		Quantity int64  `json:"quantity" validate:"required"`  // mendefinisikan field Quantity yang bertipe int64 dan memiliki tag json dan validate
		UserID   int64  `json:"user_id" validate:"required"`   // mendefinisikan field UserID yang bertipe int64 dan memiliki tag json dan validate
		Status   string `json:"status" validate:"required"`    // mendefinisikan field Status yang bertipe string dan memiliki tag json dan validate
	}

	if err := ctx.Bind(&input); err != nil { // memanggil fungsi Bind dari ctx untuk mengisi struct input dengan data dari permintaan user dan mengecek apakah ada error
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err)) // jika ada error, mengembalikan respons JSON dengan status bad request dan pesan error dari fungsi ValidatorErrors
	}

	userBalance, err := h.OrderService.GetUserBalance(ctx.Request().Context(), input.UserID) // memanggil fungsi GetUserBalance dari h.OrderService untuk mendapatkan saldo user berdasarkan input.UserID dan mengecek apakah ada error
	if err != nil {                                                                          // jika ada error
		return ctx.JSON(http.StatusUnprocessableEntity, err) // mengembalikan respons JSON dengan status unprocessable entity dan pesan error
	}

	ticketPrice, err := h.OrderService.GetTicketPrice(ctx.Request().Context(), input.TicketID) // memanggil fungsi GetTicketPrice dari h.OrderService untuk mendapatkan harga tiket berdasarkan input.TicketID dan mengecek apakah ada error
	if err != nil {                                                                            // jika ada error
		return ctx.JSON(http.StatusUnprocessableEntity, err) // mengembalikan respons JSON dengan status unprocessable entity dan pesan error
	}

	if userBalance < (input.Quantity * ticketPrice) { // jika saldo user kurang dari jumlah harga tiket yang dibeli
		return ctx.JSON(http.StatusUnprocessableEntity, errors.New("insufficient balance")) // mengembalikan respons JSON dengan status unprocessable entity dan pesan error "insufficient balance"
	}

	order := entity.NewOrder(input.TicketID, input.Quantity, input.UserID, input.Status) // memanggil fungsi NewOrder dari paket entity untuk membuat struct order baru berdasarkan input dari user
	err = h.OrderService.CreateOrder(ctx.Request().Context(), order)                     // memanggil fungsi CreateOrder dari h.OrderService untuk menyimpan order ke database dan mengecek apakah ada error
	if err != nil {                                                                      // jika ada error
		return ctx.JSON(http.StatusUnprocessableEntity, err) // mengembalikan respons JSON dengan status unprocessable entity dan pesan error
	}

	err = h.OrderService.UpdateUserBalance(ctx.Request().Context(), input.UserID, input.Quantity*ticketPrice) // memanggil fungsi UpdateUserBalance dari h.OrderService untuk mengurangi saldo user sesuai dengan jumlah harga tiket yang dibeli dan mengecek apakah ada error
	if err != nil {                                                                                           // jika ada error
		return ctx.JSON(http.StatusUnprocessableEntity, err) // mengembalikan respons JSON dengan status unprocessable entity dan pesan error
	}

	return ctx.JSON(http.StatusCreated, "Order created successfully") // jika tidak ada error, mengembalikan respons JSON dengan status created dan pesan "Order created successfully"
}

func (h *OrderHandler) GetAllOrders(ctx echo.Context) error { // mendefinisikan fungsi GetAllOrders yang menerima parameter ctx yang bertipe echo.Context dan mengembalikan error
	orders, err := h.OrderService.GetOrders(ctx.Request().Context()) // memanggil fungsi GetOrders dari h.OrderService untuk mendapatkan semua order dari database dan mengecek apakah ada error
	if err != nil {                                                  // jika ada error
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error())) // mengembalikan respons JSON dengan status bad request dan pesan error dari fungsi NewHTTPError
	}

	var orderDetails []map[string]interface{} // mendefinisikan slice orderDetails yang berisi map dengan key bertipe string dan value bertipe interface{}
	for _, order := range orders {            // melakukan iterasi untuk setiap order di dalam slice orders
		ticket, err := h.OrderService.GetTicketByID(ctx.Request().Context(), order.TicketID) // memanggil fungsi GetTicketByID dari h.OrderService untuk mendapatkan tiket berdasarkan order.TicketID dan mengecek apakah ada error
		if err != nil {                                                                      // jika ada error
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error())) // mengembalikan respons JSON dengan status internal server error dan pesan error dari fungsi NewHTTPError
		}

		orderDetail := map[string]interface{}{ // mendefinisikan map orderDetail yang berisi informasi tentang order dan tiket
			"user_id": order.UserID, // menambahkan key "user_id" dengan value order.UserID
			"ticket":  ticket,       // menambahkan key "ticket" dengan value ticket
		}
		orderDetails = append(orderDetails, orderDetail) // menambahkan map orderDetail ke dalam slice orderDetails
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{ // mengembalikan respons JSON dengan status ok dan data berikut
		"message":       "Get all orders success", // menambahkan key "message" dengan value "Get all orders success"
		"order_details": orderDetails,             // menambahkan key "order_details" dengan value orderDetails
	})
}

func (h *OrderHandler) GetOrderByUserID(ctx echo.Context) error { // mendefinisikan fungsi GetOrderByUserID yang menerima parameter ctx yang bertipe echo.Context dan mengembalikan error
	userID := ctx.Param("user_id")                                                  // mendapatkan parameter user_id dari ctx
	orders, err := h.OrderService.GetOrderByUserID(ctx.Request().Context(), userID) // memanggil fungsi GetOrderByUserID dari h.OrderService untuk mendapatkan order berdasarkan userID dan mengecek apakah ada error
	if err != nil {                                                                 // jika ada error
		return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error())) // mengembalikan respons JSON dengan status bad request dan pesan error dari fungsi NewHTTPError
	}

	var orderDetails []map[string]interface{} // mendefinisikan slice orderDetails yang berisi map dengan key bertipe string dan value bertipe interface{}
	for _, order := range orders {            // melakukan iterasi untuk setiap order di dalam slice orders
		ticket, err := h.OrderService.GetTicketByID(ctx.Request().Context(), order.TicketID) // memanggil fungsi GetTicketByID dari h.OrderService untuk mendapatkan tiket berdasarkan order.TicketID dan mengecek apakah ada error
		if err != nil {                                                                      // jika ada error
			return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error())) // mengembalikan respons JSON dengan status internal server error dan pesan error dari fungsi NewHTTPError
		}

		orderDetail := map[string]interface{}{ // mendefinisikan map orderDetail yang berisi informasi tentang order dan tiket
			"user_id": order.UserID, // menambahkan key "user_id" dengan value order.UserID
			"ticket":  ticket,       // menambahkan key "ticket" dengan value ticket
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Get all orders success",
		"order_details": orderDetails,
	})
}
