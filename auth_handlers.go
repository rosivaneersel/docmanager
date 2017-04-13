package main

import (
	"net/http"
	"fmt"
	a "github.com/arjanvaneersel/docmanager/alerts"
)

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		remember := r.FormValue("Remember")

		_ = remember // Ignore remember for now

		u, err := Users.GetUserByAuthentication(email, password)
		if err != nil {
			a.Alerts.New("Error", "alert-danger", "Invalid login.")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		CreateSession(u.ID.Hex(), u.Username, w)

		a.Alerts.New("Success", "alert-success", "You have succesfully logged in.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	t, err := NewTemplate("Document Manager", "base", "templates/login.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
		return
	}
	t.Execute(w,r)
}

func AuthLogoutHandler(w http.ResponseWriter, r *http.Request) {
	DestroySession(w)
	a.Alerts.New("Success", "alert-success", "You have succesfully logged out.")
	http.Redirect(w, r, "/", http.StatusFound)
	return
}