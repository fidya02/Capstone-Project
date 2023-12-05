package repository

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"

	"gorm.io/gorm"
)

type RegisterRepository struct {
	db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) *RegisterRepository {
	return &RegisterRepository{
		db: db,
	}
}

func (r *RegisterRepository) Register(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
