package service

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"
)

type NotificationUsecase interface {
	GetAllNotifications(ctx context.Context) ([]*entity.Notification, error)
	CreateNotification(ctx context.Context, Notification *entity.Notification) error
	UserGetNotification(ctx context.Context) ([]*entity.Notification, error)
}

type NotificationRepository interface {
	GetAllNotifications(ctx context.Context) ([]*entity.Notification, error)
	CreateNotification(ctx context.Context, Notification *entity.Notification) error
	UserGetNotification(ctx context.Context) ([]*entity.Notification, error)
}

type NotificationService struct {
	Repository NotificationRepository
}

func NewNotificationService(Repository NotificationRepository) *NotificationService {
	return &NotificationService{Repository: Repository}
}

// Get All Notifications
func (s *NotificationService) GetAllNotifications(ctx context.Context) ([]*entity.Notification, error) {
	return s.Repository.GetAllNotifications(ctx)
}

// func to creates a notification
func (s *NotificationService) CreateNotification(ctx context.Context, Notification *entity.Notification) error {
	return s.Repository.CreateNotification(ctx, Notification)
}

// get notification
func (s *NotificationService) UserGetNotifications(ctx context.Context) ([]*entity.Notification, error) {
	return s.Repository.UserGetNotification(ctx)
}
