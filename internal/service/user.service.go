package service

import (
	"context"
	"errors"

	"github.com/fidya02/Capstone-Project/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
	UpdateProfile(ctx context.Context, user *entity.User) error
	UpdateUserBalance(ctx context.Context, user *entity.User) error
	GetProfile(ctx context.Context, userID int64) (*entity.User, error)
	GetUserBalance(ctx context.Context, userID int64) (*entity.User, error)
	DeleteAccount(ctx context.Context, email string) error
	UpgradeSaldo(ctx context.Context, user *entity.User) error
	UserLogout(ctx context.Context, user *entity.User) error
	UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
}

type UserService struct {
	repository UserRepository
}

// CreateUser implements UserUseCase.
func (*UserService) CreateUser(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}

// DeleteAccount implements UserUseCase.
func (*UserService) DeleteAccount(ctx context.Context, email string) error {
	panic("unimplemented")
}

// GetAll implements UserUseCase.
func (*UserService) GetAll(ctx context.Context) ([]*entity.User, error) {
	panic("unimplemented")
}

// GetProfile implements UserUseCase.
func (*UserService) GetProfile(ctx context.Context, userID int64) (*entity.User, error) {
	panic("unimplemented")
}

// GetUserBalance implements UserUseCase.
func (*UserService) GetUserBalance(ctx context.Context, userID int64) (*entity.User, error) {
	panic("unimplemented")
}

// GetUserByID implements UserUseCase.
func (*UserService) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	panic("unimplemented")
}

// UpdateProfile implements UserUseCase.
func (*UserService) UpdateProfile(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}

// UpdateSaldo implements UserUseCase.
func (*UserService) UpdateSaldo(ctx context.Context, userID int64, updatedSaldo int64) error {
	panic("unimplemented")
}

// UpdateUser implements UserUseCase.
func (*UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}

// UpdateUserBalance implements UserUseCase.
func (*UserService) UpdateUserBalance(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}

// UpgradeSaldo implements UserUseCase.
func (*UserService) UpgradeSaldo(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}

// UserLogout implements UserUseCase.
func (*UserService) UserLogout(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository}
}

func (s *UserService) FindAll(ctx context.Context) ([]*entity.User, error) {
	return s.repository.FindAll(ctx)
}

func (s *UserService) Create(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repository.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user *entity.User) error {
	if user.Role != "" {
		if user.Role != "Administrator" && user.Role != "Buyer" {
			return errors.New("role harus di isi Administrator / Buyer")
		}
	}
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return s.repository.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

func (s *UserService) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.FindByID(ctx, id)
}
