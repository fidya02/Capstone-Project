package service

import (
	"GOLANG/entity"
	"context"
	"errors"
)

// TicketUsecase represents the use case methods for ticket-related operations.
type TicketUsecase interface {
	GetAll(ctx context.Context) ([]entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error)
	Delete(ctx context.Context, id int64) error
}

// TicketService represents a struct that implements the TicketUsecase interface.
type TicketService struct {
	// You might have other dependencies injected here, such as a repository.
	ticketRepo entity.TicketRepository // Assuming this is the repository for Ticket
}

// NewTicketService creates a new TicketService with the provided repository.
func NewTicketService(ticketRepo entity.TicketRepository) *TicketService {
	return &TicketService{ticketRepo}
}

// GetAll implements the method to retrieve all tickets.
func (s *TicketService) GetAll(ctx context.Context) ([]entity.Ticket, error) {
	tickets, err := s.ticketRepo.GetAllTickets(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// CreateTicket implements the method to create a new ticket.
func (s *TicketService) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	err := s.ticketRepo.CreateTicket(ctx, ticket)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTicket implements the method to update an existing ticket.
func (s *TicketService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	// Check if the ticket exists before updating
	_, err := s.GetTicketByID(ctx, ticket.ID)
	if err != nil {
		return errors.New("ticket not found")
	}

	err = s.ticketRepo.UpdateTicket(ctx, ticket)
	if err != nil {
		return err
	}
	return nil
}

// GetTicketByID implements the method to retrieve a ticket by its ID.
func (s *TicketService) GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	ticket, err := s.ticketRepo.GetTicketByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

// Delete implements the method to delete a ticket by its ID.
func (s *TicketService) Delete(ctx context.Context, id int64) error {
	// Check if the ticket exists before deleting
	_, err := s.GetTicketByID(ctx, id)
	if err != nil {
		return errors.New("ticket not found")
	}

	err = s.ticketRepo.DeleteTicket(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
