package entities

import (
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/utils/lanaerr"
	"math"
	"net/http"
)

type BasketItem struct {
	Product   Product    `json:"product"`
	Promotion *Promotion `json:"-"`
	Quantity  uint       `json:"quantity"`
	Total     float64    `json:"total"`
	Discount  float64    `json:"discount"`
}

func NewBasketItem(product Product, promotion *Promotion) *BasketItem {
	return &BasketItem{
		Product:   product,
		Promotion: promotion,
		Quantity:  0,
		Total:     0,
		Discount:  0,
	}
}

func (b *BasketItem) AddQuantity(quantity uint) {
	b.Quantity += quantity
	b.updateTotals()
}

func (b *BasketItem) RemoveQuantity(quantity uint) error {
	if quantity > b.Quantity {
		err := fmt.Errorf("can't remove %d %s. item quantity: %d", quantity, b.Product.ID, b.Quantity)
		return lanaerr.New(err, http.StatusBadRequest)
	}
	b.Quantity -= quantity
	b.updateTotals()
	return nil
}

func (b *BasketItem) updateTotals() {
	// Calculate total of new quantity
	b.Total = float64(b.Quantity) * b.Product.Price

	// Apply promotions if needed
	if b.Promotion != nil {
		b.Discount = b.Promotion.Apply(b)
	}

	b.Total = math.Round(b.Total*100) / 100
	b.Discount = math.Round(b.Discount*100) / 100
}
