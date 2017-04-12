package main

import "net/http"

func LoggedInUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := GetUser(r)
		if err != nil {
			//v.Alerts.New("Warning", "alert-warning", "You need to login to access that page.")
			//http.Redirect(w, r, RequireLoginRedirectTo, http.StatusTemporaryRedirect)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
