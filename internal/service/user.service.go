package service

import (
	"context"
	"errors"

	"github.com/fidya02/Capstone-Project/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	// FindAll(ctx context.Context) ([]*entity.User, error)

	// Update(ctx context.Context, user *entity.User)
	// Delete(ctx context.Context, id int64) error

	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	// UpdateProfile(ctx context.Context, user *entity.User) error
	// UpdateUserBalance(ctx context.Context, user *entity.User) error
	// GetProfile(ctx context.Context, userID int64) (*entity.User, error)
	// GetUserBalance(ctx context.Context, userID int64) (*entity.User, error)
	// DeleteAccount(ctx context.Context, email string) error
	// UpgradeWallet(ctx context.Context, user *entity.User) error
	// UserLogout(ctx context.Context, user *entity.User) error
	// UpdateWallet(ctx context.Context, userID int64, updatedWallet int64) error
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
	UpgradeWallet(ctx context.Context, user *entity.User) error
	UserLogout(ctx context.Context, user *entity.User) error
	UpdateWallet(ctx context.Context, userID int64, updatedWallet int64) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
}

type UserRepository interface {
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
	UpgradeWallet(ctx context.Context, user *entity.User) error
	UserLogout(ctx context.Context, user *entity.User) error
	UpdateWallet(ctx context.Context, userID int64, updatedWallet int64) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository}
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repository.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
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
	return s.repository.UpdateUser(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

func (s *UserService) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *UserService) DeleteAccount(ctx context.Context, email string) error {
	return s.DeleteAccount(ctx, email)
}

// GetAll implements UserUseCase.
func (s *UserService) GetAll(ctx context.Context) ([]*entity.User, error) {
	users, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetProfile implements UserUseCase.
func (s *UserService) GetProfile(ctx context.Context, userID int64) (*entity.User, error) {
	return s.repository.GetProfile(ctx, userID)
}

// GetUserBalance implements UserUseCase.
func (s *UserService) GetUserBalance(ctx context.Context, userID int64) (*entity.User, error) {
	return s.repository.GetUserBalance(ctx, userID)
}

// GetUserByID implements UserUseCase.
func (s *UserService) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return s.repository.GetUserByID(ctx, id)
}

// UpdateProfile implements UserUseCase.
func (s *UserService) UpdateProfile(ctx context.Context, user *entity.User) error {
	if user.Name == "" {
		return errors.New("Nama harus diisi")
	}
	if user.Email == "" {
		return errors.New("Email harus diisi")
	}
	return s.repository.UpdateProfile(ctx, user)
}

// UpdateWallet implements UserUseCase.
func (s *UserService) UpdateWallet(ctx context.Context, userID int64, updatedWallet int64) error {
	return s.repository.UpdateWallet(ctx, userID, updatedWallet)
}

// UpdateUserBalance implements UserUseCase.
func (s *UserService) UpdateUserBalance(ctx context.Context, user *entity.User) error {
	if user.Wallet < 0 {
		return errors.New("Saldo tidak boleh bernilai negatif")
	}
	return s.repository.UpdateUserBalance(ctx, user)
}

// UpgradeWallet implements UserUseCase.
func (s *UserService) UpgradeWallet(ctx context.Context, user *entity.User) error {
	return s.repository.UpgradeWallet(ctx, user)
}

// UserLogout implements UserUseCase.
func (s *UserService) UserLogout(ctx context.Context, user *entity.User) error {
	return s.repository.UserLogout(ctx, user)

}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.repository.FindByEmail(ctx, email)
}
