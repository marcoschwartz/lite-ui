package main

import (
	"demo-app/templates"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static files
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Home().Render(r.Context(), w)
	})

	r.Get("/inputs", func(w http.ResponseWriter, r *http.Request) {
		templates.InputsPage().Render(r.Context(), w)
	})

	r.Get("/feedback", func(w http.ResponseWriter, r *http.Request) {
		templates.FeedbackPage().Render(r.Context(), w)
	})

	r.Get("/navigation", func(w http.ResponseWriter, r *http.Request) {
		templates.NavigationPage().Render(r.Context(), w)
	})

	r.Get("/layout", func(w http.ResponseWriter, r *http.Request) {
		templates.LayoutComponentsPage().Render(r.Context(), w)
	})

	r.Get("/data-display", func(w http.ResponseWriter, r *http.Request) {
		templates.DataDisplayPage().Render(r.Context(), w)
	})

	r.Get("/datetime", func(w http.ResponseWriter, r *http.Request) {
		templates.DateTimePage().Render(r.Context(), w)
	})

	r.Get("/icons", func(w http.ResponseWriter, r *http.Request) {
		templates.IconsPage().Render(r.Context(), w)
	})

	port := ":3058"
	log.Printf("Starting server on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
