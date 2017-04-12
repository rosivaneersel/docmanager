package main

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"fmt"
)

var Users *users

func main() {
	db, err := NewDatabase("localhost", "docmanager")
	if err != nil {
		panic(err)
	}
	Users = NewUserController(db)

	r := mux.NewRouter()
	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/p", LoggedInUser(handleProtected))

	http.ListenAndServe(":8000", r)
}

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

func handleRoot(w http.ResponseWriter, r *http.Request) {
	t, err := NewTemplate("Document Manager", "base", "root.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
	}
	t.Execute(w,r)
}

func handleProtected(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from this protected page!")
}