package service

import (
	"context"
	"errors"

	"github.com/fidya02/Capstone-Project/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
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
