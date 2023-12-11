package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/fidya02/Capstone-Project/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(ctx context.Context, payment *entity.Transaction) error {
	if payment == nil {
		return errors.New("payment is nil")
	}

	// Validasi untuk nilai-nilai yang diperlukan tidak kosong atau tidak negatif
	if payment.Amount <= 0 {
		return errors.New("invalid amount")
	}

	if err := r.db.WithContext(ctx).Create(payment).Error; err != nil {
		log.Printf("Error creating payment: %s", err)
		return err
	}
	return nil
}

func (r *PaymentRepository) FindByOrderID(ctx context.Context, orderID int64) (*entity.Transaction, error) {
	payment := &entity.Transaction{}
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepository) FindByUserID(ctx context.Context, userID int64) ([]*entity.Transaction, error) {
	var payments []*entity.Transaction
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Preload(clause.Associations).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *PaymentRepository) UpdateStatus(ctx context.Context, orderID int64, status string) error {
	if err := r.db.WithContext(ctx).Model(&entity.Transaction{}).Where("order_id = ?", orderID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepository) SoftDelete(ctx context.Context, orderID int64) error {
	payment := &entity.Transaction{}
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("payment not found")
		}
		return err
	}

	if err := r.db.WithContext(ctx).Model(payment).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
