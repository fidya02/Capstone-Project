package entity

import (
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewUser(name, email, password, role string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		Role:      role,
		CreatedAt: time.Now(),
	}
}

func UpdateUser(id int64, name, email, password, role string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		Role:      role,
		UpdatedAt: time.Now(),
	}
}
func Regist(email, password, role string) *User {
	return &User{
		Email:    email,
		Password: password,
		Role:     role,
	}
}

func UpdateProfile(id int64, name, email, password string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		UpdatedAt: time.Now(),
	}
}

// func DeleteUserSelfByEmail(email string) *User {
// 	return &User{
// 		Email:     email,
// 		DeletedAt: nil,
// 	}
// }
