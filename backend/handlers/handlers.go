package handlers

import (
	"encoding/json"
	"errors"
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
	//anim := app.GetDefault()

	p := &Page{
		Selected: app.Shape{
			StrokeThickness: 1.0,
			Points:          4,
			Fill: app.Color{
				Hex:     "#ccccccc",
				Opacity: 1.0,
			},
			Stroke: app.Color{
				Hex:     "#772233",
				Opacity: 0.7,
			},
			Path:   "ARKED",
			MinRad: 1.0,
			SizeCon: app.SizeConstraint{
				Width:      4.0,
				Height:     4.0,
				PixelScale: 100,
			},
		},
		//Animations: []app.Animation{anim},
	}

	renderTemplate(w, "index", p)
}

func ShapesGenHandler(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e []app.Shape
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	for v := range e {
		fmt.Fprintln(w, e[v].Generate())
	}
}

func ShapeGenHandler(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e app.Shape
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, e.Generate())
}

func LoadAnimations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		anim := []app.Animation{app.GetDefault()}
		jsonResp, _ := json.Marshal(anim)
		w.Write(jsonResp)
	case "POST":
		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}
		var e app.Animation
		var unmarshalErr *json.UnmarshalTypeError

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&e)
		if err != nil {
			if errors.As(err, &unmarshalErr) {
				errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
			} else {
				errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
			}
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, e.GenerateGIF())
		fmt.Fprint(w, e.GenerateSpriteSheet())
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
