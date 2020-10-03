package entities

import "math"

/*
	The idea of a promotion entity is to have a configurable abstraction of promotions
	that can be associate to products. This way we can update promotions to all related
	products or create new promotions. We can even create new types of promotions. Of
	course more params would be needed to support other kind of promotions. In real
	scenarios promotions apply to certain stock units and are valid for certain time span
	and change often so it's a good idea having some configurable parameters.
*/

type Promotion struct {
	ID            string  `json:"id"`
	RequiredItems uint    `json:"required_items"`
	FreeItems     uint    `json:"free_items"`
	Reduction     float64 `json:"reduction"`
}

type PromotionType string

func (d Promotion) Apply(item *BasketItem) float64 {
	amount := 0.0

	// Check required items
	if item.Quantity < d.RequiredItems {
		return amount
	}

	// Apply free items discount
	amount += float64(d.FreeItems) * math.Floor(float64(item.Quantity)/float64(d.RequiredItems)) * item.Product.Price

	// Apply reduction
	amount += (d.Reduction / 100) * item.Total

	// Round to two decimals
	amount = math.Round(amount*100) / 100

	return amount
}
