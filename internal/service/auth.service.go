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
	Register(ctx context.Context, user *entity.User) error
}

type RegisterRepository interface {
	Register(ctx context.Context, user *entity.User) error
}

type registerService struct {
	repository RegisterRepository
}

func NewRegisterService(repository RegisterRepository) *registerService {
	return &registerService{
		repository: repository,
	}
}

func (s *registerService) Register(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repository.Register(ctx, user)
}

// Buyer
type BuyerCreateAccountUseCase interface {
	BuyerCreateAccount(ctx context.Context, user *entity.User) error
}

type BuyerCreateAccountRepository interface {
	BuyerCreateAccount(ctx context.Context, user *entity.User) error
}

type buyercreateaccountService struct {
	repository BuyerCreateAccountRepository
}

func NewBuyerCreateAccountService(repository BuyerCreateAccountRepository) *buyercreateaccountService {
	return &buyercreateaccountService{
		repository: repository,
	}
}
