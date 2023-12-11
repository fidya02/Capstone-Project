package entity

import "time"

type Transaction struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	OrderID   string     `json:"order_id"`
	Amount    int64      `json:"amount"`
	Method    string     `json:"method"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewTransaction(orderID string, userID int64, amount int64, status string) *Transaction {
	return &Transaction{
		UserID:  userID,
		OrderID: orderID,
		Amount:  amount,
		Status:  status,
		// Method:    method,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
