package main

import (
	"net/http"
	a "github.com/arjanvaneersel/docmanager/alerts"
)

func LoggedInUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := GetUser(r)
		if err != nil {
			a.Alerts.New("Warning", "alert-warning", "You need to login to access that page.")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next(w, r)
	}
}
