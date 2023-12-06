package repository

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

// GetUserBalance implements service.OrderRepository.
func (*OrderRepository) GetUserBalance(ctx context.Context, userID int64) (int64, error) {
	panic("unimplemented")
}

// UpdateUserBalance implements service.OrderRepository.
func (*OrderRepository) UpdateUserBalance(ctx context.Context, userID int64, amount int64) error {
	panic("unimplemented")
}

// NewOrderRepository creates a new instance of OrderRepository.
// It takes a *gorm.DB parameter and returns a pointer to OrderRepository.
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

// CreateOrder creates a new order in the database.
//
// It takes a context and an order entity as parameters.
// It returns an error if there was a problem creating the order.
func (r *OrderRepository) CreateOrder(ctx context.Context, order *entity.Order) error {
	err := r.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

// GetTicket retrieves a ticket from the database based on the given ticket ID.
// It returns the ticket if found, otherwise it returns an error.
func (r *OrderRepository) GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error) {
	// Create a new ticket object
	ticket := new(entity.Ticket)

	// Retrieve the ticket from the database using the provided ticket ID
	if err := r.db.WithContext(ctx).Where("id = ?", ticketID).First(&ticket).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

// UpdateTicket updates a ticket in the order repository.
func (r *OrderRepository) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	// Update the ticket in the database
	if err := r.db.WithContext(ctx).
		Model(&entity.Ticket{}).
		Where("id = ?", ticket.ID).
		Updates(&ticket).Error; err != nil {
		return err
	}
	return nil
}

// Add the following method to implement the missing GetTicketByID
func (r *OrderRepository) GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	ticket := new(entity.Ticket)
	result := r.db.WithContext(ctx).First(&ticket, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ticket, nil
}

// GetOrders retrieves all orders from the database.
// It returns a slice of order entities and an error, if any.
func (r *OrderRepository) GetOrders(ctx context.Context) ([]*entity.Order, error) {
	// Initialize an empty slice to store the orders
	orders := make([]*entity.Order, 0)

	// Query the database and preload the associated ticket entity
	err := r.db.WithContext(ctx).Preload("Ticket").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	// Return the orders and a nil error
	return orders, nil
}

// GetOrderByUserID fetches the orders for a given user ID.
func (r *OrderRepository) GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	// Initialize an empty slice to store the orders
	orders := make([]*entity.Order, 0)

	// Fetch the orders from the database, including the associated ticket
	err := r.db.WithContext(ctx).Preload("Ticket").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	// Return the orders
	return orders, nil
}

// GetTicketPrice
func (r *OrderRepository) GetTicketPrice(ctx context.Context, ticketID int64) (int64, error) {
	ticket := new(entity.Ticket)
	if err := r.db.WithContext(ctx).Where("id = ?", ticketID).First(ticket).Error; err != nil {
		return 0, err
	}

	return int64(ticket.Price), nil
}

// UserCreateOrder
func (r *OrderRepository) UserCreateOrder(ctx context.Context, order *entity.Order) error {
	err := r.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

// GetOrderHistory
func (r *OrderRepository) GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error) {
	orders := make([]*entity.Order, 0)
	err := r.db.WithContext(ctx).Preload("Ticket").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
