package main

import (
	"net/http"
	"fmt"
	a "github.com/arjanvaneersel/docmanager/alerts"
	"github.com/gorilla/mux"
)

func GroupCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("Name")
		email := r.FormValue("Email")

		u, _ := GetUser(r)
		user, err := Users.GetUserByID(u.ID)
		if err != nil {
			a.Alerts.New("Error", "alert-danger", err.Error())
			http.Redirect(w, r, "/group/create", http.StatusFound)
			return
		}

		newGroup := &Group{Name: name, Email: email}
		err = Groups.Create(newGroup)
		if err != nil {
			a.Alerts.New("Error", "alert-danger", err.Error())
			http.Redirect(w, r, "/group/create", http.StatusFound)
			return
		}

		user.Groups = append(user.Groups, newGroup.ID)
		err = Users.Update(user)
		if err != nil {
			a.Alerts.New("Error", "alert-danger", err.Error())
			http.Redirect(w, r, "/group/create", http.StatusFound)
			return
		}
		
		http.Redirect(w, r, "/group/" + newGroup.ID.Hex(), http.StatusFound)
		return
	}
	t, err := NewTemplate("Document Manager", "base", "templates/group_create.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
		return
	}
	t.Execute(w,r)
}

func GroupShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	group, err := Groups.GetByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	t, err := NewTemplate("Document Manager", "base", "templates/group_show.html")
	t.Data["Group"] = group
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
		return
	}
	t.Execute(w,r)
}