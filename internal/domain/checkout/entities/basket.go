package entities

import "time"

type Basket struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Products  []Product `json:"products"`
	Discount  float64   `json:"-"`
	Total     float64   `json:"total"`
}
