package handlers

import (
	"fmt"
	"net/http"
)

// We need ResponseWriter for sending response, and Request for handling request
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// If following code is emitted, the url other than '/' will be handled as same
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	p := &Page{
		PageTitle: "Demo Home Page",
		Heading:   "", // Change this to something else and see the results in html page
		Name:      "NepCodex",
		Country:   "Nepal",
	}
	renderTemplate(w, "index", p)
}

// Similarly, for '/users'
func UserHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello User From Users Page, NepCodex.com")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
