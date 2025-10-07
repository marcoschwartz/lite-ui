package main

import (
	"demo-app/templates"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

	r.Get("/ssr-test", func(w http.ResponseWriter, r *http.Request) {
		templates.SSRTestPage().Render(r.Context(), w)
	})

	// Mock API endpoints with 1-second delay
	r.Get("/api/users", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		users := []templates.User{
			{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", Avatar: "https://ui-avatars.com/api/?name=Alice+Johnson&background=6366f1&color=fff"},
			{ID: 2, Name: "Bob Smith", Email: "bob@example.com", Avatar: "https://ui-avatars.com/api/?name=Bob+Smith&background=8b5cf6&color=fff"},
			{ID: 3, Name: "Carol Williams", Email: "carol@example.com", Avatar: "https://ui-avatars.com/api/?name=Carol+Williams&background=ec4899&color=fff"},
			{ID: 4, Name: "David Brown", Email: "david@example.com", Avatar: "https://ui-avatars.com/api/?name=David+Brown&background=10b981&color=fff"},
			{ID: 5, Name: "Eve Davis", Email: "eve@example.com", Avatar: "https://ui-avatars.com/api/?name=Eve+Davis&background=f59e0b&color=fff"},
		}
		templates.UsersResponse(users).Render(r.Context(), w)
	})

	r.Get("/api/products", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		products := []templates.Product{
			{ID: 1, Name: "Laptop Pro 15\"", Price: "$1,299", Category: "Electronics", IconColor: "blue"},
			{ID: 2, Name: "Wireless Mouse", Price: "$29", Category: "Accessories", IconColor: "purple"},
			{ID: 3, Name: "Mechanical Keyboard", Price: "$149", Category: "Accessories", IconColor: "indigo"},
			{ID: 4, Name: "USB-C Hub", Price: "$79", Category: "Accessories", IconColor: "green"},
			{ID: 5, Name: "4K Monitor", Price: "$499", Category: "Electronics", IconColor: "blue"},
			{ID: 6, Name: "Webcam HD", Price: "$89", Category: "Electronics", IconColor: "red"},
		}
		templates.ProductsResponse(products).Render(r.Context(), w)
	})

	r.Get("/api/items", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		pageStr := r.URL.Query().Get("page")
		page, _ := strconv.Atoi(pageStr)
		if page == 0 {
			page = 1
		}

		// Generate mock items for this page
		itemsPerPage := 5
		totalPages := 4
		start := (page - 1) * itemsPerPage
		items := []templates.Item{}

		for i := 0; i < itemsPerPage; i++ {
			itemNum := start + i + 1
			items = append(items, templates.Item{
				ID:          itemNum,
				Title:       fmt.Sprintf("Item #%d", itemNum),
				Description: fmt.Sprintf("This is item number %d loaded from page %d", itemNum, page),
			})
		}

		hasMore := page < totalPages
		templates.ItemsResponse(items, page, hasMore).Render(r.Context(), w)
	})

	port := ":3058"
	log.Printf("Starting server on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
