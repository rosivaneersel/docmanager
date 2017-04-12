package main

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/gorilla/csrf"
	"time"
)

var Users *users

func main() {
	db, err := NewDatabase("localhost, 192.168.100.105", "docmanager")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	Users = NewUserController(db)

	secure := false
	csrf := csrf.Protect([]byte("dfgiort8u54u3498t9tu53yerer450rw44rt"), csrf.Secure(secure))

	r := mux.NewRouter()
	server := &http.Server {
		Addr: ":8000",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		MaxHeaderBytes: 1,
		Handler: csrf(r),
	}

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/login", UserLoginHandler)
	r.HandleFunc("/logout", LoggedInUser(UserLogoutHandler))
	r.HandleFunc("/register", UserCreateHandler)

	r.HandleFunc("/p", LoggedInUser(handleProtected))

	server.ListenAndServe()
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	t, err := NewTemplate("Document Manager", "base", "templates/root.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
		return
	}
	t.Execute(w,r)
}

func handleProtected(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from this protected page!")
}