package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitializeUserRoutes(r *chi.Mux) {

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
}
