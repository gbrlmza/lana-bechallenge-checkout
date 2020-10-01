package entities

/*
	The idea of a discount entity is to have a configurable abstraction of promotions
	that can be associate to products. This way we can update promotions to all related
	products or create new promotions. We can even create new types of promotions. Of
	course more params would be needed to support other kind of promotions. In real
	scenarios promotions apply to certain stock units and are valid for certain time span
	and change often so it's a good idea having some configurable parameters.
*/

type Discount struct {
	Code          string       `json:"code"`
	Type          DiscountType `json:"type"`
	RequiredItems int          `json:"required_items"`
	FreeItems     int          `json:"free_items"`
	Reduction     float64      `json:"reduction"`
}

type DiscountType string

const (
	DiscountTypeBuyXGetN DiscountType = "BuyXGetN"
	DiscountTypeBulk     DiscountType = "Bulk"
)
