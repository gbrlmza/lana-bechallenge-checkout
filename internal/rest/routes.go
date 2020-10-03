package rest

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (h *Handler) RouterInit() http.Handler {
	r := chi.NewRouter()

	// Health check endpoint for infrastructure monitoring & load balancers instances management
	r.Get("/ping", h.Ping)

	// API evolves over time and we need some kind of versioning system.
	// I'm using the common URI versioning approach, but could be by header version,
	// query param, accept header, domain, etc.
	r.Route("/v1/", func(r chi.Router) {

		// Basket endpoints
		r.Route("/baskets", func(r chi.Router) {

			// Create basket
			r.Post("/", h.BasketCreate)

			// Get basket details
			r.Get("/{basketID}", h.BasketGet)

			// Delete basket
			r.Delete("/{basketID}", h.BasketDelete)

			// Add product to basket
			r.Post("/{basketID}/items", h.BasketAddItems)

			// Remove product from basket
			r.Delete("/{basketID}/items/{productID}", h.BasketRemoveItem)

		})

		// Product endpoints
		r.Route("/products", func(r chi.Router) {

			// Get product list
			r.Get("/", h.ProductList)

			// Get product information
			r.Get("/{productID}", h.ProductGet)

		})
	})

	return r
}
