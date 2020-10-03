package entities

import (
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/utils/lanaerr"
	"net/http"
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

type BasketItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
	Discount float64 `json:"discount"`
}

func NewBasket() *Basket {
	b := &Basket{}
	b.Items = make(map[string]BasketItem, 0)
	return b
}

func (b *Basket) AddItem(item *Product, quantity int) error {
	// Check if item is already in basket
	if basketItem, ok := b.Items[item.ID]; ok {
		// Only update quantity
		basketItem.Quantity += quantity
		b.Items[item.ID] = basketItem
		return nil
	}

	// Initialize items map if nil
	if b.Items == nil {
		b.Items = make(map[string]BasketItem, 0)
	}

	// Add item
	b.Items[item.ID] = BasketItem{
		Product:  *item,
		Quantity: quantity,
	}

	return nil
}

func (b *Basket) RemoveItem(item *Product, quantity int) error {
	// Check if item is in basket
	basketItem, ok := b.Items[item.ID]
	if !ok {
		return lanaerr.New(fmt.Errorf("item %s not found in basket", item.ID), http.StatusBadRequest)
	}

	if quantity > basketItem.Quantity { // Check if the quantity can be deleted
		return lanaerr.New(fmt.Errorf("item quantity is %d, can't delete %d",
			basketItem.Quantity, quantity), http.StatusBadRequest)
	} else if quantity == basketItem.Quantity { // Remove the item completely
		delete(b.Items, item.ID)
	} else {
		basketItem.Quantity -= quantity
		b.Items[item.ID] = basketItem
	}

	return nil
}
