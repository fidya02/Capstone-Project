package service

import (
	"context"
	"errors"

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

// NewOrderService creates a new instance of OrderService.
// It takes an OrderRepository parameter and returns a pointer to OrderService.
func NewOrderService(repository OrderRepository) *OrderService {
	return &OrderService{repository}
}

// CreateOrder creates a new order in the system.
// It checks if the ticket is available and deducts the quantity from the ticket quota.
// It also updates the user's balance.
func (s *OrderService) CreateOrder(ctx context.Context, order *entity.Order) error {
	// Get the ticket from the repository
	ticket, err := s.repository.GetTicket(ctx, order.TicketID)
	if err != nil {
		return err
	}

	// Check if the ticket quota is sufficient for the order
	if int64(ticket.Quota) < order.Quantity {
		return errors.New("ticket is not available")
	}

	// Calculate the total price of the order
	order.Total = ticket.Price * int64(order.Quantity)

	// Create the order in the repository
	if err := s.repository.CreateOrder(ctx, order); err != nil {
		return err
	}

	// Deduct the order quantity from the ticket quota
	ticket.Quota -= order.Quantity
	if err := s.repository.UpdateTicket(ctx, ticket); err != nil {
		return err
	}

	// Update the user's balance
	if err := s.repository.UpdateUserBalance(ctx, order.UserID, order.Total); err != nil {
		return err
	}

	return nil
}

// GetTicket retrieves a ticket with the given ID from the repository.
// It returns the ticket if found, or an error if the ticket does not exist or an error occurred during retrieval.
func (s *OrderService) GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error) {
	return s.repository.GetTicket(ctx, ticketID)
}

// UpdateTicket updates the ticket in the order service.
func (s *OrderService) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	return s.repository.UpdateTicket(ctx, ticket)
}

// UpdateUserBalance updates the balance of a user.
//
// It takes the user ID and the new balance as parameters.
// It returns an error if there was a problem updating the balance.
func (s *OrderService) UpdateUserBalance(ctx context.Context, userID int64, saldo int64) error {
	return s.repository.UpdateUserBalance(ctx, userID, saldo)
}

// GetOrders retrieves a list of orders from the repository.
// It returns a slice of Order pointers and an error if any.
func (s *OrderService) GetOrders(ctx context.Context) ([]*entity.Order, error) {
	// Call the GetOrders method of the repository to fetch the orders.
	orders, err := s.repository.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// GetTicketByID retrieves a ticket by its ID.
func (s *OrderService) GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	return s.repository.GetTicketByID(ctx, id)
}

// GetOrderByUserID retrieves the orders associated with a specific user.
// It takes a context and a user ID as input parameters.
// It returns a slice of order entities and an error, if any.
func (s *OrderService) GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	return s.repository.GetOrderByUserID(ctx, userID)
}

// GetUserBalance returns the balance of a user.
func (s *OrderService) GetUserBalance(ctx context.Context, userID int64) (int64, error) {
	// Call the repository method to get the user balance
	balance, err := s.repository.GetUserBalance(ctx, userID)
	// Return the balance and any error that occurred
	return balance, err
}

// GetTicketPrice retrieves the price of a ticket with the given ID.
// It returns the ticket price as an int64.
// If the ticket does not exist, it returns an error.
func (s *OrderService) GetTicketPrice(ctx context.Context, ticketID int64) (int64, error) {
	// Get the ticket from the repository
	ticket, err := s.repository.GetTicket(ctx, ticketID)
	if err != nil {
		return 0, err
	}
	// Return the ticket price
	return int64(ticket.Price), nil
}

// UserCreateOrder creates a new order for a user.
func (s *OrderService) UserCreateOrder(ctx context.Context, order *entity.Order) error {
	// Get the ticket associated with the order.
	ticket, err := s.repository.GetTicket(ctx, order.TicketID)
	if err != nil {
		return err
	}

	// Check if the ticket has enough quota.
	if int64(ticket.Quota) < order.Quantity {
		return errors.New("ticket is not available")
	}

	// Calculate the total price of the order.
	order.Total = ticket.Price * int64(order.Quantity)

	// Create the order in the repository.
	if err := s.repository.CreateOrder(ctx, order); err != nil {
		return err
	}

	// Update the quota of the ticket.
	ticket.Quota -= order.Quantity
	if err := s.repository.UpdateTicket(ctx, ticket); err != nil {
		return err
	}

	// Update the user's balance.
	if err := s.repository.UpdateUserBalance(ctx, order.UserID, order.Total); err != nil {
		return err
	}

	return nil
}

// GetOrderHistory retrieves the order history for a specific user.
// It takes the user ID as input and returns a slice of order entities and an error.
func (s *OrderService) GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error) {
	// Call the GetOrderByUserID method of the repository to fetch the order history.
	orders, err := s.repository.GetOrderByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
