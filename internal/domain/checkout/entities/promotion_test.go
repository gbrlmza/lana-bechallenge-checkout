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
