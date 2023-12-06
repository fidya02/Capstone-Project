package entity

import (
	"context"
	"fmt"
	"time"
)

type Ticket struct {
	ID         int64     `json:"id"`
	Nama       string    `json:"nama"`
	Tanggal    time.Time `json:"tanggal"`
	Venue      string    `json:"venue"`
	Harga      float64   `json:"harga"`
	Keterangan string    `json:"keterangan"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeleteAt   time.Time `json:"-"`
}

type TicketRepository interface {
	GetAllTickets(ctx context.Context) ([]Ticket, error)
	CreateTicket(ctx context.Context, ticket *Ticket) error
	UpdateTicket(ctx context.Context, ticket *Ticket) error
	GetTicketByID(ctx context.Context, id int64) (*Ticket, error)
	DeleteTicket(ctx context.Context, id int64) error
}

// req untuk create ticket
func NewTicket(nama, venue string, tanggal time.Time, harga float64, keterangan string) *Ticket {
	return &Ticket{
		Nama:       nama,
		Tanggal:    tanggal,
		Venue:      venue,
		Harga:      harga,
		Keterangan: keterangan,
		CreatedAt:  time.Now(),
	}
}

// req untuk update ticket
func UpdateTicket(id int64, nama, venue string, tanggal time.Time, harga float64, keterangan string) *Ticket {
	return &Ticket{
		ID:         id,
		Nama:       nama,
		Tanggal:    tanggal,
		Venue:      venue,
		Harga:      harga,
		Keterangan: keterangan,
		UpdatedAt:  time.Now(),
	}
}

// Membuat metode untuk mencetak detail event
func (e Ticket) PrintEventDetails() {
	fmt.Printf("Nama Ticket: %s\n", e.Nama)
	fmt.Printf("Hari dan Tanggal: %s\n", e.Tanggal.Format("Monday, 2 January 06"))
	fmt.Printf("Venue: %s\n", e.Venue)
	fmt.Printf("Harga: IDR %.0f\n", e.Harga)
	fmt.Printf("Keterangan: %s\n", e.Keterangan)
	fmt.Println("---------------------------------------------------")
}

// req untuk login
// func Login(email, password string) *User {
// 	return &User{
// 		Email:    email,
// 		Password: password,
// 	}
// }

//note : ketika type data untuk ID hanya int, maka akan error ketika dijalankan. karena ID tidak bisa di tambahkan otmatis oleh database
// namun ketika type data untuk ID diubah menjadi int64, maka tidak akan error ketika dijalankan. karena ID bisa di tambahkan otmatis oleh database melalui postman.
