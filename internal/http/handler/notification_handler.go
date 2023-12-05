package handler

imoport (
	"Ticketing/entity"
	"Ticketing/internal/service"
	"net/http"

	"Ticketing/internal/http/validator"
	"time"

	"github.com/labstack/echo/v4"
)

	type NotificationHandler struct {
	notificationService service.NotificationUsecase
}
	func NewNotificationHandler(notificationService service.NotificationUsecase) *NotificationHandler {
		return &NotificationHandler{notificationService}
	}
	//GetAllNotifications
	func (h *NotificationHandler) GetAllNotifications(c echo.Context) error {
		Notifications, err := h.notificationService.GetAllNotifications(c.Request().Context)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": Notifications,		
	})
}

		//func to create notification
		func (h *NotificationHandler) CreateNotification(c echo.Context) error {
			var input struct {
				Type 		string 		`json:"type"`
				Message 	string 		`json:"message"`
				IsRead 		bool 		`json:"is_read"`
				Create_at 	time.Time 	`json:"create_at
	}
	//Input validation
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	//Create a notification object
	Notificaton := entity.Notification{
		Type:      input.Type,
		Message:   input.Message,
		IsRead:    input.IsRead,
		CreatedAt: time.Now(),
	}

	err := h.notificationService.CreateNotification(c.Request().Context(), &Notification)
	if err != nil {
		return c.JSON(http.StatusCreated, Notification)
	}

	//get notifications
	func (h *NotificationHandler) UserGetNotification(c echo.Context) error {
		Notifications, err := h.notificationService.UserGetNotification(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": Notifications,
		})
	}