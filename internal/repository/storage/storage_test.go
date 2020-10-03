package storage_test

import (
	"context"
	"encoding/json"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/repository/storage"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_storage_BasketGet_NotFound(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)

	// When
	b, err := s.BasketGet(ctx, "78235217-43fe-4e7a-8f18-e5f83df01ca6")

	// Then
	assert.EqualError(t, err, "basket 78235217-43fe-4e7a-8f18-e5f83df01ca6 not found")
	assert.Nil(t, b)
}

func Test_storage_BasketSave_NewBasket_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)
	basket := &entities.Basket{}

	// When
	err := s.BasketSave(ctx, basket)
	sBasket, _ := s.BasketGet(ctx, basket.ID)

	// Then
	assert.Nil(t, err)
	assert.NotEmpty(t, basket.ID)
	assert.Equal(t, *basket, *sBasket)
}

func Test_storage_BasketSave_ExistingBasket_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)
	basket := &entities.Basket{
		ID:        "cf31bf2b-42a3-4cb5-ae51-34fbe30d163f",
		CreatedAt: time.Date(2020, 10, 1, 0, 0, 0, 0, time.UTC),
		Total:     100,
	}

	// When
	err := s.BasketSave(ctx, basket)
	sBasket, _ := s.BasketGet(ctx, basket.ID)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, *basket, *sBasket)
}

func Test_storage_BasketDelete_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)
	basket := &entities.Basket{
		ID:        "cf31bf2b-42a3-4cb5-ae51-34fbe30d163f",
		CreatedAt: time.Date(2020, 10, 1, 0, 0, 0, 0, time.UTC),
		Total:     100,
	}
	err := s.BasketSave(ctx, basket)

	// When
	err = s.BasketDelete(ctx, basket.ID)
	sBasket, _ := s.BasketGet(ctx, basket.ID)

	// Then
	assert.Nil(t, err)
	assert.Nil(t, sBasket)
}

func Test_storage_ProductGet_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)

	// When
	p, err := s.ProductGet(ctx, "PEN")
	jsonP, _ := json.Marshal(p)

	// Then
	expectedProduct := `{"id":"PEN","name":"Lana Pen","price":5,"discount_id":"BUY2GET1FREE"}`
	assert.Equal(t, expectedProduct, string(jsonP))
	assert.Nil(t, err)
}

func Test_storage_ProductGet_NotFound(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)

	// When
	p, err := s.ProductGet(ctx, "BOOK")

	// Then
	assert.EqualError(t, err, "product BOOK not found")
	assert.Nil(t, p)
}

func Test_storage_ProductList_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)

	// When
	list, err := s.ProductList(ctx)
	jsonList, _ := json.Marshal(list)

	// Then
	expectedList := `[{"id":"PEN","name":"Lana Pen","price":5,"discount_id":"BUY2GET1FREE"},{"id":"TSHIRT","name":"Lana T-Shirt","price":20,"discount_id":"BUY3GET25OFF"},{"id":"MUG","name":"Lana Coffee Mug","price":7.5,"discount_id":null}]`
	assert.Equal(t, expectedList, string(jsonList))
	assert.Nil(t, err)
}

func Test_storage_DiscountGet_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)

	// When
	d, err := s.DiscountGet(ctx, "BUY2GET1FREE")
	jsonD, _ := json.Marshal(d)

	// Then
	expectedDiscount := `{"code":"BUY2GET1FREE","type":"BuyXGetN","required_items":2,"free_items":1,"reduction":0}`
	assert.Equal(t, expectedDiscount, string(jsonD))
	assert.Nil(t, err)
}

func Test_storage_DiscountGet_NotFound(t *testing.T) {
	// Given
	ctx := context.Background()
	s := storage.NewStorage(ctx)

	// When
	d, err := s.DiscountGet(ctx, "3X2")

	// Then
	assert.EqualError(t, err, "discount 3X2 not found")
	assert.Nil(t, d)
}
