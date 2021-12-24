package handlers

import (
	"fmt"
	"net/http"

	"github.com/caaraya/parka-server/app"
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
		Heading:   "something else", // Change this to something else and see the results in html page
		Name:      "NepCodex",
		Country:   "Nepal",
	}
	renderTemplate(w, "index", p)
}

func ShapeGenHandler(w http.ResponseWriter, r *http.Request) {
	Shape := &app.Shape{
		StrokeThickness: 1.0,
		Points:          6,
		Fill:            app.Color{Hex: "#005533", Opacity: 1.0},
		Stroke:          app.Color{Hex: "#ffffff", Opacity: 1.0},
		Path:            "STRAIGHT",
		MinRad:          2.0,
		SizeCon:         app.SizeConstraint{Width: 4, Height: 4, PixelScale: 100},
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, Shape.Generate())
}

// Similarly, for '/users'
func UserHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello User From Users Page, NepCodex.com")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
