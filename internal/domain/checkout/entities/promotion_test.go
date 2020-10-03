package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPromotion_Apply_WithoutRequiredItems(t *testing.T) {
	// Given
	p := &Promotion{RequiredItems: 10}
	bi := &BasketItem{Quantity: 5}

	// When
	amount := p.Apply(bi)

	// Then
	assert.Zero(t, amount)
}

func TestPromotion_Apply_2X1(t *testing.T) {
	// Given
	promotion := &Promotion{
		RequiredItems: 2,
		FreeItems:     1,
		Reduction:     0,
	}
	product := Product{
		ID:    "PEN",
		Name:  "PEN",
		Price: 15.5,
	}
	bi := NewBasketItem(product, promotion)
	bi.AddQuantity(3)

	// When
	amount := promotion.Apply(bi)

	// Then
	assert.Equal(t, 15.5, amount)
}

func TestPromotion_Apply_25OFF3ORMORE(t *testing.T) {
	// Given
	promotion := &Promotion{
		RequiredItems: 3,
		FreeItems:     0,
		Reduction:     25,
	}
	product := Product{
		ID:    "PEN",
		Name:  "PEN",
		Price: 50,
	}
	bi := NewBasketItem(product, promotion)
	bi.AddQuantity(3)

	// When
	amount := promotion.Apply(bi)

	// Then
	assert.Equal(t, 37.5, amount)
}
