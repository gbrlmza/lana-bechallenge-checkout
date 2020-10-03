package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasket_SaveItem(t *testing.T) {
	// Given
	b := NewBasket()
	bi := NewBasketItem(Product{ID: "PEN", Price: 50}, nil)
	bi.AddQuantity(1)

	// When
	b.SaveItem(bi)

	// Then
	assert.Equal(t, 1, len(b.Items))
	assert.Equal(t, 50.0, b.Subtotal)
	assert.Equal(t, 0.0, b.Discount)
	assert.Equal(t, 50.0, b.Total)
}

func TestBasket_SaveItem_ZeroQuantity(t *testing.T) {
	// Given
	b := NewBasket()
	bi := NewBasketItem(Product{ID: "PEN", Price: 50}, nil)
	bi.AddQuantity(10)
	b.SaveItem(bi)
	bi.RemoveQuantity(10)

	// When
	b.SaveItem(bi)

	// Then
	assert.Equal(t, 0, len(b.Items))
	assert.Equal(t, 0.0, b.Subtotal)
	assert.Equal(t, 0.0, b.Discount)
	assert.Equal(t, 0.0, b.Total)
}

func TestBasket_GetItem(t *testing.T) {
	// Given
	b := NewBasket()
	bi := NewBasketItem(Product{ID: "PEN", Price: 50}, nil)
	bi.AddQuantity(1)
	b.SaveItem(bi)

	// When
	basketItem := b.GetItem("PEN")

	// Then
	assert.NotNil(t, basketItem)
}

func TestBasket_GetItem_Nil(t *testing.T) {
	// Given
	b := NewBasket()

	// When
	basketItem := b.GetItem("PEN")

	// Then
	assert.Nil(t, basketItem)
}
