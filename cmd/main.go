package main

import (
	"fmt"
	"ticketing_app/internal/delivery/cli"
)

func main() {
	fmt.Println("Ticketing App CLI")

	// Inisialisasi handler untuk registrasi
	registerHandler := cli.NewRegisterHandler()

	// Tampilkan prompt registrasi
	registerHandler.HandleRegister()
}
