package entity

import (
	"time"
)

// Order merepresentasikan entitas pesanan dalam sistem.
type Order struct {
	ID        int64      `json:"id"`
	TicketID  int64      `json:"ticket_id"`
	Ticket    Ticket     `json:"ticket"`
	UserID    int64      `json:"user_id"`
	User      User       `json:"user"`
	Quantity  int64      `json:"quantity"`
	Total     int64      `json:"total"`
	Price     int64      `json:"price"`
	Status    string     `json:"status"`
	OrderAt   time.Time  `json:"order_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	OrderBy   string     `json:"order_by"`
	UpdateBy  string     `json:"-"`
	DeletedBy string     `json:"-"`
}

// NewOrder adalah fungsi pembuat (constructor) yang membuat instance baru dari pesanan dengan detail yang diberikan.
func NewOrder(ticketID, userID, quantity, price int64, status string) *Order {
	return &Order{
		TicketID: ticketID,
		UserID:   userID,
		Quantity: quantity,
		Price:    price,
		OrderAt:  time.Now(),
		Status:   status,
	}
}

// OrderDetail berisi informasi terperinci tentang suatu pesanan.
type OrderDetail struct {
	UserID    int64        `json:"user_id"`
	Quantity  int64        `json:"quantity"`
	Total     int64        `json:"total"`
	Status    string       `json:"status"`
	OrderAt   time.Time    `json:"order_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Ticket    TicketDetail `json:"ticket"`
}
