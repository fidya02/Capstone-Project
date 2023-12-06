package repository

import (
	"GOLANG/entity"
	"context"
	"database/sql"
)

// SQLTicketRepository represents an implementation of TicketRepository for SQL databases.
type SQLTicketRepository struct {
	DB *sql.DB // SQL database connection
	// Other fields if needed for the database connection
}

// NewSQLTicketRepository creates a new instance of SQLTicketRepository.
func NewSQLTicketRepository(db *sql.DB) *SQLTicketRepository {
	return &SQLTicketRepository{
		DB: db,
	}
}

// GetAllTickets retrieves all tickets from the database.
func (repo *SQLTicketRepository) GetAllTickets(ctx context.Context) ([]entity.Ticket, error) {
	// Query to get all tickets from the database
	rows, err := repo.DB.QueryContext(ctx, "SELECT * FROM tickets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []entity.Ticket
	for rows.Next() {
		var ticket entity.Ticket
		// Scan rows and construct Ticket objects
		if err := rows.Scan(&ticket.ID, &ticket.Nama, &ticket.Tanggal, &ticket.Venue, &ticket.Harga, &ticket.Keterangan, &ticket.CreatedAt, &ticket.UpdatedAt, &ticket.DeleteAt); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

// CreateTicket creates a new ticket in the database.
func (repo *SQLTicketRepository) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	// Implement logic to insert a new ticket into the database
	stmt, err := repo.DB.PrepareContext(ctx, "INSERT INTO tickets(nama, tanggal, venue, harga, keterangan, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, ticket.Nama, ticket.Tanggal, ticket.Venue, ticket.Harga, ticket.Keterangan, ticket.CreatedAt, ticket.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTicket updates an existing ticket in the database.
// ... (Implementation of UpdateTicket, GetTicketByID, DeleteTicket methods remains similar to the earlier examples)
// (Omitted for brevity)
