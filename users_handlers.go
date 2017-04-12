package main

import (
	"net/http"
	"fmt"
	a "github.com/arjanvaneersel/docmanager/alerts"
)

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("Username")
		password := r.FormValue("Password")
		password2 := r.FormValue("Password2")
		email := r.FormValue("Email")

		if password != password2 {
			a.Alerts.New("Error", "alert-danger", "Passwords don't match")
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		newUser := &User{Username: username, Email: email}
		newUser.SetPassword(password)

		err := Users.Create(newUser)
		if err != nil {
			a.Alerts.New("Error", "alert-danger", err.Error())
			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		a.Alerts.New("Success", "alert-info", "You have successfully registered an account. Please check your email for activation instructions.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	t, err := NewTemplate("Document Manager", "base", "register.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
	}
	t.Execute(w,r)
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
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

		CreateSession(u.ID.String(), u.Username, w)

		a.Alerts.New("Success", "alert-success", "You have succesfully logged in.")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	t, err := NewTemplate("Document Manager", "base", "login.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
	}
	t.Execute(w,r)
}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
	DestroySession(w)
	a.Alerts.New("Success", "alert-success", "You have succesfully logged out.")
	http.Redirect(w, r, "/", http.StatusFound)
	return
}