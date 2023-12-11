package repository

import (
	"context"
	"errors"

	"github.com/fidya02/Capstone-Project/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID)
	if user.Name != "" {
		query = query.Update("name", user.Name)
	}
	if user.Password != "" {
		query = query.Update("password", user.Password)
	}
	if user.Role != "" {
		query = query.Update("role", user.Role)
	}
	if user.Wallet != 0 {
		query = query.Update("wallet", user.Wallet)
	}
	if user.Email != "" {
		query = query.Update("email", user.Email)
	}
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user with that email not found")
	}
	return user, nil
}

// Update User Self
func (r *UserRepository) UpdateProfile(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// update user balance by id
func (r *UserRepository) UpdateUserBalance(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetProfile
func (r *UserRepository) GetProfile(ctx context.Context, userID int64) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UserBalance
func (r *UserRepository) GetUserBalance(ctx context.Context, userID int64) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteAccount
func (r *UserRepository) DeleteAccount(ctx context.Context, email string) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, email).Error; err != nil {
		return err
	}
	return nil
}

// upgrade wallet
func (r *UserRepository) UpgradeWallet(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// logout
func (r *UserRepository) UserLogout(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// wallet updates by ID
func (r *UserRepository) UpdateWallet(ctx context.Context, userID int64, updatedWallet int64) error {
	user := &entity.User{ID: userID, Wallet: updatedWallet}

	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", userID).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.WithContext(ctx).First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
