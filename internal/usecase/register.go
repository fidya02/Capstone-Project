package usecase

import (
	"fmt"
	"ticketing_app/internal/domain"
)

// RegisterUsecase adalah use case untuk proses registrasi pengguna
type RegisterUsecase struct{}

// RegisterUser mendaftarkan pengguna baru
func (u *RegisterUsecase) RegisterUser(newUser domain.User) {
	// Implementasi logika registrasi di sini
	// Misalnya, menyimpan data pengguna ke dalam database
	fmt.Println("Mendaftarkan pengguna baru:", newUser.NamaLengkap)
}
