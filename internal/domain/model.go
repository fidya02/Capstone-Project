package domain

// User adalah model yang merepresentasikan informasi pengguna
type User struct {
	NamaLengkap  string
	TempatLahir  string
	TanggalLahir string
	Gender       string
	Email        string
	Password     string
}
