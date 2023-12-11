package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Role      string     `json:"role"`
	Number    string     `json:"number"`
	Wallet    int64      `json:"wallet"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewUser(name, email, number, password, role string, wallet int64) *User {
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

func UpdateUser(id int64, wallet int64, number, name, email, password, role string) *User {
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

func UserLogout(id int64) *User {
	return &User{
		ID: id,
	}
}

func UpdateWallet(id int64, wallet int64) *User {
	return &User{
		ID:     id,
		Wallet: wallet,
	}
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetUserByID(db *gorm.DB, userID int64) (*User, error) {
	user := &User{}
	if err := db.First(user, userID).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserWalletByID(db *gorm.DB, userID int64, newWalletValue int64) error {
	return db.Model(&User{}).Where("id = ?", userID).Update("wallet", newWalletValue).Error
}

func DeleteUserByID(db *gorm.DB, userID int64) error {
	return db.Delete(&User{}, userID).Error
}
