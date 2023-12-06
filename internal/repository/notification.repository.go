package repository

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

// Get all notification
func (r *NotificationRepository) GetAllNotifications(ctx context.Context) ([]*entity.Notification, error) {
	Notifications := make([]*entity.Notification, 0)
	result := r.db.WithContext(ctx).Find(&Notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return Notifications, nil
}

// Create a notification
func (r *NotificationRepository) CreateNotification(ctx context.Context, Notification *entity.Notification) error {
	result := r.db.WithContext(ctx).Create(&Notification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get Notifications
func (r *NotificationRepository) UserGetNotification(ctx context.Context) ([]*entity.Notification, error) {
	Notifications := make([]*entity.Notification, 0)

	//Retrieve Notifications
	result := r.db.WithContext(ctx).Where("is_read = ?", false).Find(&Notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	//Mark retrieved Notifications as read
	for _, notification := range Notifications {
		//Suppose you have a method a update the is_read field
		err := r.MarkNotificationAsRead(ctx, notification.ID)
		if err != nil {
			return nil, err
		}
	}
	return Notifications, nil
}
func (r *NotificationRepository) MarkNotificationAsRead(ctx context.Context, notificationID int) error {
	result := r.db.WithContext(ctx).Model(&entity.Notification{}).Where("id = ?", notificationID).Update("is_read", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
