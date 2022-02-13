package main // For executable package, we use main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caaraya/parka-server/handlers"
)

// Created own package for handling routes

// Main program starts here
func main() {
	// Handling Routes
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/shape", handlers.ShapeGenHandler)
	http.HandleFunc("/shapes", handlers.ShapesGenHandler)
	http.HandleFunc("/animations", handlers.LoadAnimations)
	// Handle Static Content
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Use environment variable "PORT" or otherwise assign
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		log.Printf("The port is %s", port)
	}

	// Logging
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)

	// Create server or exit
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
