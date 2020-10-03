package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBasketItem_NewBasketItem(t *testing.T) {
	// Given
	product := Product{
		ID:    "PEN",
		Name:  "PEN",
		Price: 50,
	}

	// When
	bi := NewBasketItem(product, nil)

	// Then
	assert.NotNil(t, bi)
}

func TestBasketItem_AddQuantity(t *testing.T) {
	// Given
	product := Product{
		ID:    "PEN",
		Name:  "PEN",
		Price: 50,
	}

	// When
	bi := NewBasketItem(product, nil)
	bi.AddQuantity(10)

	// Then
	assert.NotNil(t, bi)
	assert.Equal(t, uint(10), bi.Quantity)
	assert.Equal(t, 500.0, bi.Total)
	assert.Equal(t, 0.0, bi.Discount)
}

func TestBasketItem_RemoveQuantity_Error(t *testing.T) {
	// Given
	product := Product{
		ID:    "PEN",
		Name:  "PEN",
		Price: 50,
	}

	// When
	bi := NewBasketItem(product, nil)
	bi.AddQuantity(10)
	err := bi.RemoveQuantity(20)

	// Then
	assert.NotNil(t, bi)
	assert.EqualError(t, err, "can't remove 20 PEN. item quantity: 10")
}

func TestBasketItem_RemoveQuantity_Success(t *testing.T) {
	// Given
	product := Product{
		ID:    "PEN",
		Name:  "PEN",
		Price: 50,
	}

	// When
	bi := NewBasketItem(product, nil)
	bi.AddQuantity(10)
	err := bi.RemoveQuantity(5)

	// Then
	assert.NotNil(t, bi)
	assert.Nil(t, err)
	assert.Equal(t, uint(5), bi.Quantity)
	assert.Equal(t, 250.0, bi.Total)
	assert.Equal(t, 0.0, bi.Discount)
}
