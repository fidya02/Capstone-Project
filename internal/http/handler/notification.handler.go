package handler

import (
	"net/http"

	"github.com/fidya02/Capstone-Project/entity"
	"github.com/fidya02/Capstone-Project/internal/service"

	"time"

	"github.com/fidya02/Capstone-Project/internal/http/validator"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	notificationService service.NotificationUsecase
}

func NewNotificationHandler(notificationService service.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{notificationService}
}

// GetAllNotifications
func (h *NotificationHandler) GetAllNotifications(ctx echo.Context) error {
	Notifications, err := h.notificationService.GetAllNotifications(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": Notifications,
	})
}

// func to create notification
func (h *NotificationHandler) CreateNotification(ctx echo.Context) error {
	var input struct {
		Type       string    `json:"type"`
		Message    string    `json:"message"`
		IsRead     bool      `json:"is_read"`
		Created_at time.Time `json:"created_at"`
	}
	//Input validation
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	//Create a notification object
	notification := entity.Notification{
		Type:      input.Type,
		IsRead:    input.IsRead,
		Message:   input.Message,
		CreatedAt: time.Now(),
	}

	err := h.notificationService.CreateNotification(ctx.Request().Context(), &notification)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"data": "success create notification",
	})
}

// get notifications
func (h *NotificationHandler) UserGetNotification(ctx echo.Context) error {
	Notifications, err := h.notificationService.UserGetNotification(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": Notifications,
	})
}
