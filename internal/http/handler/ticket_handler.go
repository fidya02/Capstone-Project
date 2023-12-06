package handler

import (
	"GOLANG/entity"
	"GOLANG/internal/http/validator"
	"GOLANG/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	ticketService service.TicketUsecase
}

func NewTicketHandler(ticketService service.TicketUsecase) *TicketHandler {
	return &TicketHandler{ticketService}
}

func (h *TicketHandler) GetAllTickets(ctx echo.Context) error {
	tickets, err := h.ticketService.GetAll(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}

func (h *TicketHandler) CreateTicket(ctx echo.Context) error {
	var input struct {
		Nama       string    `json:"nama" validate:"required"`
		Tanggal    time.Time `json:"tanggal" validate:"required"`
		Venue      string    `json:"venue" validate:"required"`
		Harga      float64   `json:"harga" validate:"required"`
		Keterangan string    `json:"keterangan"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	ticket := entity.Ticket{
		Nama:       input.Nama,
		Tanggal:    input.Tanggal,
		Venue:      input.Venue,
		Harga:      input.Harga,
		Keterangan: input.Keterangan,
		CreatedAt:  time.Now(),
	}

	err := h.ticketService.CreateTicket(ctx.Request().Context(), &ticket)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusCreated, "Ticket created successfully")
}

func (h *TicketHandler) UpdateTicket(ctx echo.Context) error {
	var input struct {
		ID         int64     `param:"id" validate:"required"`
		Nama       string    `json:"nama" validate:"required"`
		Tanggal    time.Time `json:"tanggal" validate:"required"`
		Venue      string    `json:"venue" validate:"required"`
		Harga      float64   `json:"harga" validate:"required"`
		Keterangan string    `json:"keterangan"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	ticket := entity.Ticket{
		ID:         input.ID,
		Nama:       input.Nama,
		Tanggal:    input.Tanggal,
		Venue:      input.Venue,
		Harga:      input.Harga,
		Keterangan: input.Keterangan,
		UpdatedAt:  time.Now(),
	}

	err := h.ticketService.UpdateTicket(ctx.Request().Context(), &ticket)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Ticket updated successfully",
		"ticket":  ticket,
	})
}

func (h *TicketHandler) GetTicketByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	ticket, err := h.ticketService.GetTicketByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":         ticket.ID,
			"nama":       ticket.Nama,
			"tanggal":    ticket.Tanggal,
			"venue":      ticket.Venue,
			"harga":      ticket.Harga,
			"keterangan": ticket.Keterangan,
			"created":    ticket.CreatedAt,
			"updated":    ticket.UpdatedAt,
		},
	})
}

func (h *TicketHandler) DeleteTicket(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.ticketService.Delete(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Ticket deleted successfully",
	})
}
