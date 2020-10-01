package rest

import (
	"github.com/go-chi/render"
	"net/http"
)

type BasketProduct struct {
	Quantity int `json:"quantity"`
}

func getBasketProductQuantity(r *http.Request) (int, error) {
	// Bind payload
	basketProduct := &BasketProduct{}
	if err := render.DecodeJSON(r.Body, basketProduct); err != nil {
		return 0, err
	}

	// Set default quantity
	if basketProduct.Quantity == 0 {
		basketProduct.Quantity = 1
	}

	return basketProduct.Quantity, nil
}
