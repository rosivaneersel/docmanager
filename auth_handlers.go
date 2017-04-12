package main

import (
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("Password")

		u, err := Users.GetUserByAuthentication(email, password)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		CreateSession(u.ID.String(), u.Username, w)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	DestroySession(w)
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
