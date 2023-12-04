package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Username string
	Password string
}

var users = map[string]User{
	"user1": {"user1", "password1"},
	"user2": {"user2", "password2"},
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Silakan login</h1><form method='post' action='/login'><input type='text' name='username' placeholder='Username'><br><input type='password' name='password' placeholder='Password'><br><input type='submit' value='Login'></form>")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		user, found := users[username]
		if !found || user.Password != password {
			http.Error(w, "Username atau password salah", http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "<h1>Login berhasil</h1><p>Selamat datang, %s!</p>", username)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/", loginPage)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
