package service

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *entity.Order) error
	GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetOrders(ctx context.Context) ([]*entity.Order, error)
	GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error)
	GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error)
	UpdateUserBalance(ctx context.Context, userID int64, amount int64) error
	GetUserBalance(ctx context.Context, userID int64) (int64, error)
	GetTicketPrice(ctx context.Context, ticketID int64) (int64, error)
	UserCreateOrder(ctx context.Context, order *entity.Order) error
	GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entity.Order) error
	GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetOrders(ctx context.Context) ([]*entity.Order, error)
	GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error)
	GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error)
	UpdateUserBalance(ctx context.Context, userID int64, amount int64) error
	GetUserBalance(ctx context.Context, userID int64) (int64, error)
	GetTicketPrice(ctx context.Context, ticketID int64) (int64, error)
	UserCreateOrder(ctx context.Context, order *entity.Order) error
	GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error)
}

type OrderService struct {
	repository OrderRepository
}

func NewOrderService(repository OrderRepository) *OrderService {
	return &OrderService{repository}
}

// CreateOrder creates a new order in the OrderService.
// It retrieves information about the ticket based on the ticket ID in the order.
// It calculates the total price of the order.
// It creates the order in the repository.
// It updates the user's balance by deducting the total price of the order.
// If any error occurs, it returns the error.
func (s *OrderService) CreateOrder(ctx context.Context, order *entity.Order) error {
	// Retrieve information about the ticket based on the ticket ID
	ticket, err := s.repository.GetTicket(ctx, order.TicketID)
	if err != nil {
		return err
	}

	// Calculate the total price of the order
	order.Total = ticket.Price * int64(order.Quantity)

	// Create the order in the repository
	if err := s.repository.CreateOrder(ctx, order); err != nil {
		return err
	}

	// Deduct the total price from the user's balance
	if err := s.repository.UpdateUserBalance(ctx, order.UserID, order.Total); err != nil {
		return err
	}

	return nil
}

// Implementasi fungsi GetTicket
func (s *OrderService) GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error) {
	return s.repository.GetTicket(ctx, ticketID)
}

// Implementasi fungsi UpdateTicket
func (s *OrderService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.repository.UpdateTicket(ctx, ticket)
}

// implementasi fungsi update user balance
func (s *OrderService) UpdateUserBalance(ctx context.Context, userID int64, saldo int64) error {
	return s.repository.UpdateUserBalance(ctx, userID, saldo)
}

func (s *OrderService) GetOrders(ctx context.Context) ([]*entity.Order, error) {
	return s.repository.GetOrders(ctx)
}

func (s *OrderService) GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	return s.repository.GetTicketByID(ctx, id)
}

// get order by user_id
func (s *OrderService) GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	return s.repository.GetOrderByUserID(ctx, userID)
}

// get user balance
func (s *OrderService) GetUserBalance(ctx context.Context, userID int64) (int64, error) {
	return s.repository.GetUserBalance(ctx, userID)
}

// GetTicketPrice
func (s *OrderService) GetTicketPrice(ctx context.Context, ticketID int64) (int64, error) {
	ticket, err := s.repository.GetTicket(ctx, ticketID)
	if err != nil {
		return 0, err
	}
	return int64(ticket.Price), nil
}

// UserCreateOrder creates a new order for a user.
func (s *OrderService) UserCreateOrder(ctx context.Context, order *entity.Order) error {
	// Get ticket information based on the ticket ID in the order.
	ticket, err := s.repository.GetTicket(ctx, order.TicketID)
	if err != nil {
		return err
	}

	// Calculate the total price of the order.
	order.Total = ticket.Price * int64(order.Quantity)

	// Create the order.
	if err := s.repository.CreateOrder(ctx, order); err != nil {
		return err
	}

	// Deduct user's balance.
	if err := s.repository.UpdateUserBalance(ctx, order.UserID, order.Total); err != nil {
		return err
	}

	return nil
}

// GetOrderHistory returns the order history for a given user.
func (s *OrderService) GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error) {
	// Call the repository's GetOrderByUserID method to fetch the order history for the user.
	return s.repository.GetOrderByUserID(ctx, userID)
}
