package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter(c Controller) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Возвращаем ответ в json
	//r.Use(func(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		w.Header().Set("Content-Type", "application/json")
	//		next.ServeHTTP(w, r)
	//	})
	//})

	//r.Route("/api/todo-list", func(r chi.Router) {
	//	r.Post("/tasks/", func(w http.ResponseWriter, r *http.Request) {
	//		log.Println("eba")
	//		c.Create(w, r)
	//	})
	//})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/api/todo-list/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c.Create(r.Context(), w, r)
	})

	return r
}
