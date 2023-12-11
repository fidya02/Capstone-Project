package entity

import "time"

type Payment struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	OrderID   int64      `json:"order_id"`
	Amount    float64    `json:"amount"`
	Method    string     `json:"method"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewPayment(userID, orderID int64, amount float64, method, status string) *Payment {
	return &Payment{
		UserID:    userID,
		OrderID:   orderID,
		Amount:    amount,
		Method:    method,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
