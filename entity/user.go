package entity

import (
	"time"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Number   string `json:"number"`
	// Wallet    int        `json:"Wallet" gorm:"column:Wallet"`
	Wallet    int64      `json:"wallet"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewUser(name string, email string, number string, password string, role string, wallet int64) *User {
	return &User{
		Name:      name,
		Email:     email,
		Number:    number,
		Password:  password,
		Role:      role,
		Wallet:    wallet,
		CreatedAt: time.Now(),
	}
}

func UpdateUser(id int64, wallet int64, number, name string, email, password, role string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Number:    number,
		Email:     email,
		Password:  password,
		Role:      role,
		Wallet:    wallet,
		UpdatedAt: time.Now(),
	}
}
func Regist(email, number, password, role string) *User {
	return &User{
		Email:    email,
		Number:   number,
		Password: password,
		Role:     role,
	}
}

func UpdateProfile(id int64, name, email, number, password string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Number:    number,
		Password:  password,
		UpdatedAt: time.Now(),
	}
}

// Update the return type to be *User
func DeleteUserSelfByEmail(email string) *User {
	return &User{
		Email:     email,
		DeletedAt: nil,
	}
}

func UpgradeWallet(id int64, wallet int64) *User {
	return &User{
		ID:     id,
		Wallet: wallet,
	}
}

// user logout
func UserLogout(id int64) *User {
	return &User{
		ID: id,
	}
}

// updateWallet
func UpdateWallet(id int64, wallet int64) *User {
	return &User{
		ID:     id,
		Wallet: wallet,
	}
}
