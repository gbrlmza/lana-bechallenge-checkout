package entities

import (
	"math"
	"time"
)

type Basket struct {
	ID        string                `json:"id"`
	CreatedAt time.Time             `json:"created_at"`
	Items     map[string]BasketItem `json:"items"`
	Subtotal  float64               `json:"subtotal"`
	Discount  float64               `json:"discount"`
	Total     float64               `json:"total"`
}

func NewBasket() *Basket {
	b := &Basket{}
	b.Items = make(map[string]BasketItem, 0)
	return b
}

func (b *Basket) GetItem(productID string) *BasketItem {
	if item, ok := b.Items[productID]; ok {
		return &item
	}
	return nil
}

func (b *Basket) SaveItem(item *BasketItem) {
	if item.Quantity == 0 {
		// If the item has zero quantity, remove it from basket
		delete(b.Items, item.Product.ID)
	} else {
		// Save
		b.Items[item.Product.ID] = *item
	}

	// Recalculate totals amounts
	b.updateTotals()
}

func (b *Basket) updateTotals() {
	b.Subtotal = 0
	b.Discount = 0
	for _, i := range b.Items {
		b.Subtotal = i.Total
		b.Discount = i.Discount
	}
	b.Total = b.Subtotal - b.Discount

	b.Subtotal = math.Round(b.Subtotal*100) / 100
	b.Discount = math.Round(b.Discount*100) / 100
	b.Total = math.Round(b.Total*100) / 100
}
