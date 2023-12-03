package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	NamaLengkap  string
	TempatLahir  string
	TanggalLahir string
	Gender       string
	Email        string
	Password     string
}

func isValidDateFormat(dateStr string) bool {
	const layout = "02/01/2006" // Format DD/MM/YYYY
	_, err := time.Parse(layout, dateStr)
	return err == nil
}

func register() User {
	var newUser User

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Registrasi Akun ===")
	fmt.Print("Nama Lengkap (max 50 karakter): ")
	inputNamaLengkap, _ := reader.ReadString('\n')
	inputNamaLengkap = strings.TrimSpace(inputNamaLengkap)
	newUser.NamaLengkap = inputNamaLengkap

	if len(newUser.NamaLengkap) > 50 {
		newUser.NamaLengkap = newUser.NamaLengkap[:50]
	}

	fmt.Print("Tempat Lahir (max 20 karakter): ")
	tempatLahir, _ := reader.ReadString('\n')
	tempatLahir = strings.TrimSpace(tempatLahir)
	newUser.TempatLahir = tempatLahir

	if len(newUser.TempatLahir) > 20 {
		newUser.TempatLahir = newUser.TempatLahir[:20]
	}

	fmt.Scanln(&newUser.TanggalLahir)
	for !isValidDateFormat(newUser.TanggalLahir) {
		fmt.Print("Tanggal Lahir (DD/MM/YYYY): ")
		fmt.Scanln(&newUser.TanggalLahir)
	}

	fmt.Print("Jenis Kelamin (L = Laki-laki, P = Perempuan): ")
	fmt.Scanln(&newUser.Gender)
	newUser.Gender = strings.ToUpper(newUser.Gender[:1])

	fmt.Print("Email (max 20 karakter): ")
	fmt.Scanln(&newUser.Email)
	newUser.Email = newUser.Email[:min(len(newUser.Email), 20)]
	fmt.Printf("Email: %s\n", newUser.Email)

	fmt.Print("Password (max 10 karakter): ")
	fmt.Scanln(&newUser.Password)
	newUser.Password = newUser.Password[:min(len(newUser.Password), 10)]

	return newUser
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	newUser := register()

	// Simulasi registrasi berhasil, kemudian pengguna diarahkan ke proses login
	fmt.Println("\nRegistrasi Berhasil! Silakan login untuk melanjutkan.")

	// Gunakan newUser untuk menampilkan informasi pengguna yang terdaftar
	fmt.Println("Informasi Pengguna:")
	fmt.Println("Nama Lengkap:", newUser.NamaLengkap)
	fmt.Println("Tempat Lahir:", newUser.TempatLahir)
	fmt.Println("Tanggal Lahir:", newUser.TanggalLahir)
	fmt.Println("Gender:", newUser.Gender)
	fmt.Println("Email:", newUser.Email)

	// Di sini, Anda dapat melanjutkan dengan proses login setelah registrasi berhasil
	// Anda dapat menambahkan logika login atau mengarahkan pengguna ke fungsi login yang sesuai.
}
