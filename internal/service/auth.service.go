package service

import (
	"context"
	"errors"

	"github.com/fidya02/Capstone-Project/entity"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Login(ctx context.Context, email, password string) (*entity.User, error)
}

type LoginRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type LoginService struct {
	repo LoginRepository
}

func NewLoginService(repo LoginRepository) *LoginService {
	return &LoginService{
		repo: repo,
	}
}

func (s *LoginService) Login(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user with that email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect login credentials")
	}

	return user, nil
}

type RegisterUseCase interface {
	Registration(ctx context.Context, user *entity.User) error
}

type RegisterRepository interface {
	Registration(ctx context.Context, user *entity.User) error
}

type registrationService struct {
	repository RegisterRepository
}

func NewRegistrationService(repository RegisterRepository) *registrationService {
	return &registrationService{
		repository: repository,
	}
}

func (s *registrationService) Registration(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repository.Registration(ctx, user)
}
