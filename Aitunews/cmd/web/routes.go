package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/category", app.category)
	mux.HandleFunc("/contact", app.contact)

	// Static files example (replace "/static/" with your actual static files path)
	staticDir := http.Dir("./static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	return mux
}
