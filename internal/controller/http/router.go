package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func NewRouter(c Controller) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/swagger", httpSwagger.WrapHandler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/api/todo-list/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c.Create(r.Context(), w, r)
	})

	r.Put("/api/todo-list/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")
		c.Update(r.Context(), w, r, id)
	})

	r.Delete("/api/todo-list/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")
		c.Delete(r.Context(), w, id)
	})

	r.Put("/api/todo-list/tasks/{id}/done", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")
		c.Done(r.Context(), w, id)
	})

	r.Get("/api/todo-list/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c.List(r.Context(), w, r)
	})

	return r
}
