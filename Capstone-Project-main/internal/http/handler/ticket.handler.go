package handler

import (
	"net/http"
	"time"

	"github.com/fidya02/Capstone-Project/entity"
	"github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/fidya02/Capstone-Project/internal/service"
	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	ticketService service.TicketUseCase
}

func NewTicketHandler(ticketService service.TicketUseCase) *TicketHandler {
	return &TicketHandler{ticketService}
}

func (h *TicketHandler) FindAllTickets(ctx echo.Context) error {
	tickets, err := h.ticketService.FindAllTickets(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) FindTicketByID(ctx echo.Context) error {
	id := ctx.Param("id")
	ticket, err := h.ticketService.FindTicketByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"ticket": ticket,
	})
}

func (h *TicketHandler) CreateTicket(ctx echo.Context) error {
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Image       string `json:"image"`
		Price       int64  `json:"price" validate:"required"`
		Date        string `json:"date" validate:"required"`
		Location    string `json:"location" validate:"required"`
		Status      string `json:"status" validate:"required"`
		Quantity    int    `json:"quantity" validate:"required"`
		Category    string `json:"category" validate:"required"`
		Sold        int64  `json:"sold"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	timestr := input.Date.Format("2006-01-02T15:04:05")

	ticket := entity.Ticket{
		Name:        input.Name,
		Description: input.Description,
		Image:       input.Image,
		Price:       input.Price,
		Date:        timestr,
		Location:    input.Location,
		Status:      input.Status,
		Quantity:    input.Quantity,
		Category:    input.Category,
		Sold:        input.Sold,
		CreateAt:    time.Now(),
	}

	err := h.ticketService.CreateTicket(ctx.Request().Context(), &ticket)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create ticket",
	})
}

func (h *TicketHandler) UpdateTicket(ctx echo.Context) error {
	var input struct {
		ID          int64  `param:"id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Image       string `json:"image"`
		Price       int64  `json:"price" validate:"required"`
		Date        string `json:"date" validate:"required"`
		Location    string `json:"location" validate:"required"`
		Status      string `json:"status" validate:"required"`
		Quantity    int    `json:"quantity" validate:"required"`
		Category    string `json:"category" validate:"required"`
		Sold        int64  `json:"sold"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	timestr := input.Date.Format("2006-01-02T15:04:05")

	ticket := entity.Ticket{
		ID:          int(input.ID),
		Name:        input.Name,
		Description: input.Description,
		Image:       input.Image,
		Price:       input.Price,
		Date:        timestr,
		Location:    input.Location,
		Quantity:    input.Quantity,
		Category:    input.Category,
	}

	err := h.ticketService.UpdateTicket(ctx.Request().Context(), &ticket)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update ticket",
		"ticket": map[string]interface{}{
			"id":          ticket.ID,
			"name":        ticket.Name,
			"description": ticket.Description,
			"image":       ticket.Image,
			"price":       ticket.Price,
			"date":        ticket.Date,
			"location":    ticket.Location,
			"status":      ticket.Status,
			"quantity":    ticket.Quantity,
			"category":    ticket.Category,
			"sold":        ticket.Sold,
		},
	})
}

func (h *TicketHandler) DeleteTicket(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.ticketService.DeleteTicket(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete ticket",
	})
}

func (h *TicketHandler) SearTicket(ctx echo.Context) error {
	var input struct {
		Keyword string `param:"keyword" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.SearchTicket(ctx.Request().Context(), input.Keyword)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) FilterTicketByCategory(ctx echo.Context) error {
	var input struct {
		Category string `param:"category" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.FilterTicketByCategory(ctx.Request().Context(), input.Category)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) FilterTicketByLocation(ctx echo.Context) error {
	var input struct {
		Location string `param:"location" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.FilterTicketByLocation(ctx.Request().Context(), input.Location)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) FilterTicketByRangeTime(ctx echo.Context) error {
	var input struct {
		Date string `param:"date" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.FilterTicketByRangeTime(ctx.Request().Context(), input.Date)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) FilterTicketByPrice(ctx echo.Context) error {
	var input struct {
		Price int64 `param:"price" validate:"required"`
	}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tickets, err := h.ticketService.FilterTicketByPrice(ctx.Request().Context(), input.Price)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) SortTicketByNewest(ctx echo.Context) error {
	sortTicket := ctx.Param("sort")

	if sortTicket == "newest" {
		tickets, err := h.ticketService.SortTicketByNewest(ctx.Request().Context())
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, err)
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"tickets": tickets,
		})
	}
}

func (h *TicketHandler) SortTicketByOldest(ctx echo.Context) error {
	sortTicket := ctx.Param("sort")

	if sortTicket == "oldest" {
		tickets, err := h.ticketService.SortTicketByOldest(ctx.Request().Context())
		if err != nil {
			return ctx.JSON(http.StatusUnprocessableEntity, err)
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"tickets": tickets,
		})
	}
}

func (h *TicketHandler) SortTicketByExpensive(ctx echo.Context) error {
	sortTicket := ctx.Param("sort")

	if sortTicket == "expensive" {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "invalid sort",
		})
	}

	tickets, err := h.ticketService.SortTicketByExpensive(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) SortTicketByCheap(ctx echo.Context) error {
	sortTicket := ctx.Param("sort")

	if sortTicket == "cheap" {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "invalid sort",
		})
	}

	tickets, err := h.ticketService.SortTicketByCheap(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) SortTicketBySold(ctx echo.Context) error {
	sortTicket := ctx.Param("sort")

	if sortTicket == "sold" {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "invalid sort",
		})
	}

	tickets, err := h.ticketService.SortTicketBySold(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}

func (h *TicketHandler) SortTicketByAvailable(ctx echo.Context) error {
	sortTicket := ctx.Param("sort")

	if sortTicket == "available" {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "invalid sort",
		})
	}

	tickets, err := h.ticketService.SortTicketByAvailable(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"tickets": tickets,
	})
}
