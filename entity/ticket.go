package entity

import "time"

type Ticket struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       int64      `json:"price"`
	Status      string     `json:"status"`
	Image       string     `json:"image"`
	Location    string     `json:"location"`
	Quantity    int64      `json:"quantity"`
	Category    string     `json:"category"`
	Date        string     `json:"date"` // Menggunakan string custom
	Sold        int64      `json:"sold"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

func NewTicket(name string, description string, price int64, status string, image string, location string, quantity int64, category string, date string, sold int64) *Ticket {
	return &Ticket{
		Name:        name,
		Description: description,
		Price:       price,
		Status:      status,
		Image:       image,
		Location:    location,
		Quantity:    quantity,
		Category:    category,
		Date:        date,
		Sold:        sold,
	}
}

func UpdateTicket(id int64, name string, description string, price int64, status string, image string, location string, quantity int64, category string, date string, sold int64) *Ticket {
	return &Ticket{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Status:      status,
		Image:       image,
		Location:    location,
		Quantity:    quantity,
		Category:    category,
		Date:        date,
		Sold:        sold,
	}
}

type TicketDetails struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Image       string    `json:"image"`
}
