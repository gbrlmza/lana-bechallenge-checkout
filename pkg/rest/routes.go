package rest

import (
	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	// Health check endpoint for infrastructure monitoring & load balancers
	r.Get("/ping", Ping)

	return r
}
