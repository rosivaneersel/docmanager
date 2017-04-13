package main

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/gorilla/csrf"
	"time"
	"log"
)

var Users *users
var Groups *groups

func main() {
	db, err := NewDatabase("localhost, 192.168.100.105", "docmanager")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	Users = NewUserController(db)
	Groups = NewGroupController(db)

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

	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/login", AuthLoginHandler)
	r.HandleFunc("/logout", LoggedInUser(AuthLogoutHandler))

	r.HandleFunc("/register", UserCreateHandler)

	r.HandleFunc("/group/create", LoggedInUser(GroupCreateHandler))
	r.HandleFunc("/group/{id}", LoggedInUser(GroupShowHandler))

	server.ListenAndServe()
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := NewTemplate("Document Manager", "base", "templates/root.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
		return
	}

	u, err := GetUser(r)
	if err == nil {
		user, err := Users.GetUserByID(u.ID)
		if err != nil {
			fmt.Fprintf(w, "Corrupt session. User ID is invalid.\n")
			return
		}
		log.Println("Active user, getting groups")
		groups, err := Groups.GetByIDs(user.Groups)
		if err != nil {
			log.Printf("Couldn't get user groups. %x", err)
		}
		t.Data["Groups"] = groups
	}
	log.Println("Executing template")
	t.Execute(w,r)
}
