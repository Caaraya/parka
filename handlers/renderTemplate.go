package handlers

import (
	"html/template"
	"net/http"
)

// Parse template files
var templates = template.Must(template.ParseFiles("templates/index.html"))

// Declaration of struct needed for the template
type Page struct {
	PageTitle string
	Heading   string
	Name      string
	Country   string
}

// A custom render function which takes the filename of template html file
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
