package main

import (
	"net/http"
	"fmt"
	a "github.com/arjanvaneersel/docmanager/alerts"
	"github.com/gorilla/mux"
	"strconv"
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
	//Todo: Redirect after change, empty form
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
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
		return
	}
	t.Data["Group"] = group
	t.Execute(w,r)
}

func GroupCreateUpdateDocumentType(w http.ResponseWriter, r *http.Request) {
	var idx int

	t, err := NewTemplate("Document Manager", "base", "templates/group_document_type.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
		return
	}

	vars := mux.Vars(r)
	groupID := vars["gid"]
	if groupID == "" {
		http.NotFound(w, r)
		return
	}

	group, err := Groups.GetByID(groupID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if r.Method == "POST" {
		code := r.FormValue("Code")
		name := r.FormValue("Name")
		fi := r.FormValue("DocumentTypeIndex")
		if fi == "" || fi == "-1"{
			idx = -1
		} else {
			idx, _ = strconv.Atoi(fi)
			t.Data["GID"] = groupID
			t.Data["DocumentTypeIndex"] = idx
			t.Data["DocumentType"] = group.DocumentTypes[idx]
		}

		group.CreateOrUpdateDocumentType(idx, DocumentType{Code: code, Name: name})
		err = Groups.Update(group)
		if err != nil {
			a.Alerts.New("Error", "alert-danger", err.Error())
			t.Data["DocumentTypeIndex"] = idx
			t.Data["DocumentType"] = group.DocumentTypes[idx]
			t.Execute(w, r)
			return
		}
		a.Alerts.New("Success", "alert-danger", "Successfully added document type")
		http.Redirect(w, r, "/group/" + group.ID.Hex(), http.StatusFound)
	}

	i := vars["idx"]
	if i == "" {
		idx = -1
	} else {
		idx, _ = strconv.Atoi(i)
		t.Data["DocumentTypeIndex"] = idx
		t.Data["DocumentType"] = group.DocumentTypes[idx]
	}
	t.Data["GID"] = groupID
	t.Execute(w, r)
}

func GroupDeleteDocumentType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["gid"]
	if groupID == "" {
		http.NotFound(w, r)
		return
	}

	fi := r.FormValue("DocumentTypeIndex")
	if fi == "" {
		http.NotFound(w, r)
		return
	}

	group, err := Groups.GetByID(groupID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	i, _ := strconv.Atoi(fi)
	group.DocumentTypes = append(group.DocumentTypes[:i], group.DocumentTypes[i+1:]...)

	a.Alerts.New("Success", "alert-danger", "Successfully deleted document type")
	http.Redirect(w, r, "/group/" + group.ID.Hex(), http.StatusFound)
}

func GroupCreateUpdateBatch(w http.ResponseWriter, r *http.Request) {
	var idx int

	t, err := NewTemplate("Document Manager", "base", "templates/group_document_type.html")
	if err != nil {
		fmt.Fprintf(w, "Template error: %s", err)
		return
	}

	vars := mux.Vars(r)
	groupID := vars["gid"]
	if groupID == "" {
		http.NotFound(w, r)
		return
	}

	group, err := Groups.GetByID(groupID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if r.Method == "POST" {
		code := r.FormValue("Code")
		name := r.FormValue("Name")
		fi := r.FormValue("DocumentTypeIndex")
		if fi == "" || fi == "-1"{
			idx = -1
		} else {
			idx, _ = strconv.Atoi(fi)
			t.Data["GID"] = groupID
			t.Data["DocumentTypeIndex"] = idx
			t.Data["DocumentType"] = group.DocumentTypes[idx]
		}

		group.CreateOrUpdateDocumentType(idx, DocumentType{Code: code, Name: name})
		err = Groups.Update(group)
		if err != nil {
			a.Alerts.New("Error", "alert-danger", err.Error())
			t.Data["DocumentTypeIndex"] = idx
			t.Data["DocumentType"] = group.DocumentTypes[idx]
			t.Execute(w, r)
			return
		}
		a.Alerts.New("Success", "alert-danger", "Successfully added document type")
		http.Redirect(w, r, "/group/" + group.ID.Hex(), http.StatusFound)
	}

	i := vars["idx"]
	if i == "" {
		idx = -1
	} else {
		idx, _ = strconv.Atoi(i)
		t.Data["DocumentTypeIndex"] = idx
		t.Data["DocumentType"] = group.DocumentTypes[idx]
	}
	t.Data["GID"] = groupID
	t.Execute(w, r)
}