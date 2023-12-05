package repository

import (
	"context"
	"errors"

	"github.com/fidya02/Capstone-Project/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

// CreateOrder creates a new order in the database.
// It takes a context and an order entity as parameters.
// It returns an error if there was a problem creating the order.
func (r *OrderRepository) CreateOrder(ctx context.Context, order *entity.Order) error {
	err := r.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

// GetTicket retrieves a ticket by its ID.
//
// Parameters:
//
//	ctx: The context.Context for the request.
//	ticketID: The ID of the ticket to retrieve.
//
// Returns:
//
//	*entity.Ticket: The retrieved ticket.
//	error: Any error that occurred during the retrieval.
func (r *OrderRepository) GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error) {
	// Create a new instance of the Ticket struct.
	ticket := new(entity.Ticket)
	// Query the database for the ticket with the given ID.
	if err := r.db.WithContext(ctx).Where("id = ?", ticketID).First(&ticket).Error; err != nil {
		// If an error occurred, return nil for the ticket and the error.
		return nil, err
	}
	// Return the retrieved ticket.
	return ticket, nil
}

// UpdateTicket updates the ticket in the order repository.
// It takes the context and the ticket to be updated as input.
// It returns an error if the update operation fails.
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

// GetTicketByID retrieves a ticket from the database based on its ID.
func (r *OrderRepository) GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	// Create a new ticket instance
	ticket := new(entity.Ticket)

	// Query the database to retrieve the ticket with the given ID
	result := r.db.WithContext(ctx).First(&ticket, id)

	// If there was an error retrieving the ticket, return the error
	if result.Error != nil {
		return nil, result.Error
	}

	// Return the retrieved ticket
	return ticket, nil
}

func (r *OrderRepository) GetOrders(ctx context.Context) ([]*entity.Order, error) {
	orders := make([]*entity.Order, 0)
	err := r.db.WithContext(ctx).Preload("Ticket").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrderByUserID retrieves all orders associated with a given user ID.
// It returns a slice of order entities and an error, if any.
func (r *OrderRepository) GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	// Initialize an empty slice to store the orders
	orders := make([]*entity.Order, 0)

	// Query the database to retrieve orders associated with the given user ID
	err := r.db.WithContext(ctx).Preload("Ticket").Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	// Return the retrieved orders
	return orders, nil
}

// UpdateUserBalance updates the balance of a user in the database.
// It checks if the user has sufficient balance before deducting the specified amount.
// If the user has insufficient balance, it returns an error.
// If the update is successful, it returns nil.
func (r *OrderRepository) UpdateUserBalance(ctx context.Context, userID int64, total int64) error {
	// Retrieve the user from the database
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(user).Error; err != nil {
		return err
	}

	// Check if the user has sufficient balance
	if user.Saldo < total {
		return errors.New("insufficient balance")
	}

	// Deduct the specified amount from the user's balance
	user.Saldo -= total

	// Update the user in the database
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserBalance retrieves the balance of a user from the database.
// It takes a context and a userID as input and returns the user's balance as an int64.
// If there is an error retrieving the balance, it returns 0 and the error.
func (r *OrderRepository) GetUserBalance(ctx context.Context, userID int64) (int64, error) {
	// Create a new instance of the User struct
	user := new(entity.User)

	// Query the database for the user with the given ID
	// and assign the result to the user struct
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(user).Error; err != nil {
		return 0, err
	}

	// Return the user's balance
	return user.Saldo, nil
}

// GetTicketPrice retrieves the price of a ticket with the given ticketID.
// It returns the ticket price as an int64 and any error encountered.
func (r *OrderRepository) GetTicketPrice(ctx context.Context, ticketID int64) (int64, error) {
	// Create a new instance of the Ticket struct
	ticket := new(entity.Ticket)

	// Query the database for the ticket with the given ticketID
	// and store the result in the ticket variable
	if err := r.db.WithContext(ctx).Where("id = ?", ticketID).First(ticket).Error; err != nil {
		// Return the error if there was an issue retrieving the ticket
		return 0, err
	}

	// Return the price of the ticket as an int64
	return int64(ticket.Price), nil
}

// UserCreateOrder creates a new order for a user.
func (r *OrderRepository) UserCreateOrder(ctx context.Context, order *entity.Order) error {
	err := r.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

// GetOrderHistory retrieves the order history for a given user.
// It returns a slice of order entities and an error if any.
func (r *OrderRepository) GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error) {
	// Initialize an empty slice to store the orders
	orders := make([]*entity.Order, 0)

	// Retrieve the orders from the database
	err := r.db.WithContext(ctx).
		Preload("Ticket").
		Where("user_id = ?", userID).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}
