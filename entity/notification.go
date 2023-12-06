package entity

import (
	"time"
)

type Notification struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewNotification(id int, typ, massage string, isRead bool, createAt, updatedAt, deletedAt time.Time) *Notification {
	return &Notification{
		ID:        id,
		Type:      typ,
		Message:   massage,
		IsRead:    isRead,
		CreatedAt: createAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
