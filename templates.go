package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/csrf"
)

var DefaultFiles = []string{"templates/index.html", "templates/_nav.html"}

type Template struct {
	Title string
	Data     map[string]interface{}
	layout string
	template *template.Template
}

type templateData struct {
	Title   string
	Data    map[string]interface{}
	Session ActiveUser
	CsrfField  template.HTML
}

func (t *Template) Execute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "html")
	w.WriteHeader(http.StatusOK)

	u, _ := GetUser(r)
	data := templateData{t.Title, t.Data, u, csrf.TemplateField(r)}
	t.template.ExecuteTemplate(w, t.layout, data)
}

func NewTemplate(title string, layout string, files ...string) (*Template, error) {
	fs := append(DefaultFiles, files...)
	t, err := template.ParseFiles(fs...)
	if err != nil {
		return nil, err
	}
	return &Template{Title: title, layout: layout, template: t}, nil
}
