package repository

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

// GetTicket implements service.OrderRepository.
func (*OrderRepository) GetTicket(ctx context.Context, ticketID int64) (*entity.Ticket, error) {
	panic("unimplemented")
}

// UpdateTicket implements service.OrderRepository.
func (*OrderRepository) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	panic("unimplemented")
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *entity.Order) error {
	err := r.db.WithContext(ctx).Create(&order).Error
	return err
}

// GetTicketByID mengambil tiket berdasarkan ID yang diberikan.
func (r *OrderRepository) GetTicketByID(ctx context.Context, id int64) (*entity.Ticket, error) {
	ticket := new(entity.Ticket)
	result := r.db.WithContext(ctx).First(&ticket, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ticket, nil
}

// GetOrders mengambil semua pesanan dari database.
func (r *OrderRepository) GetOrders(ctx context.Context) ([]*entity.Order, error) {
	orders := make([]*entity.Order, 0)
	err := r.db.WithContext(ctx).Preload("Ticket").Find(&orders).Error
	return orders, err
}

// GetOrderByUserID mengambil pesanan untuk ID pengguna yang diberikan.
func (r *OrderRepository) GetOrderByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	orders := make([]*entity.Order, 0)
	err := r.db.WithContext(ctx).Preload("Ticket").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

// GetTicketPrice mendapatkan harga tiket berdasarkan ID tiket.
func (r *OrderRepository) GetTicketPrice(ctx context.Context, ticketID int64) (int64, error) {
	ticket := new(entity.Ticket)
	if err := r.db.WithContext(ctx).Where("id = ?", ticketID).First(ticket).Error; err != nil {
		return 0, err
	}
	return int64(ticket.Price), nil
}

// UserCreateOrder membuat pesanan oleh pengguna.
func (r *OrderRepository) UserCreateOrder(ctx context.Context, order *entity.Order) error {
	err := r.db.WithContext(ctx).Create(&order).Error
	return err
}

// GetOrderHistory mengambil riwayat pesanan berdasarkan ID pengguna.
func (r *OrderRepository) GetOrderHistory(ctx context.Context, userID int64) ([]*entity.Order, error) {
	orders := make([]*entity.Order, 0)
	err := r.db.WithContext(ctx).Preload("Ticket").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

// GetUserBalance mendapatkan saldo pengguna berdasarkan ID pengguna.
func (r *OrderRepository) GetUserBalance(ctx context.Context, userID int64) (int64, error) {
	var userBalance int64
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Select("wallet").
		Where("id = ?", userID).
		Find(&userBalance).Error; err != nil {
		return 0, err
	}
	return userBalance, nil
}

// UpdateUserBalance memperbarui saldo pengguna berdasarkan ID pengguna dan jumlah yang diberikan.
func (r *OrderRepository) UpdateUserBalance(ctx context.Context, userID int64, amount int64) error {
	// Ambil saldo pengguna saat ini
	currentBalance, err := r.GetUserBalance(ctx, userID)
	if err != nil {
		return err
	}

	// Hitung saldo baru setelah penambahan atau pengurangan
	newBalance := currentBalance + amount

	// Perbarui saldo pengguna ke database
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", userID).
		Update("wallet", newBalance).Error; err != nil {
		return err
	}

	return nil
}
