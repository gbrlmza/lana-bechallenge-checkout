package storage

import (
	"context"
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/utils/lanaerr"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

type storage struct {
	data struct {
		products  map[string]entities.Product
		baskets   map[string]entities.Basket
		discounts map[string]entities.Discount
	}
	mutex struct {
		product  sync.Mutex
		basket   sync.Mutex
		discount sync.Mutex
	}
}

func NewStorage(ctx context.Context) *storage {
	s := &storage{}
	s.initializeData()
	return s
}

func (s *storage) BasketSave(ctx context.Context, basket *entities.Basket) error {
	// Lock basket map
	s.mutex.basket.Lock()
	defer s.mutex.basket.Unlock()

	// Generate ID if needed
	if basket.ID == "" {
		basket.ID = uuid.New().String()
		basket.CreatedAt = time.Now()
	}

	// Save basket
	s.data.baskets[basket.ID] = *basket

	return nil
}

func (s *storage) BasketGet(ctx context.Context, basketID string) (*entities.Basket, error) {
	// Lock basket map
	s.mutex.basket.Lock()
	defer s.mutex.basket.Unlock()

	// Get basket from storage data
	if basket, ok := s.data.baskets[basketID]; ok {
		return &basket, nil
	}

	// Basket not found
	return nil, lanaerr.New(fmt.Errorf("basket %s not found", basketID), http.StatusNotFound)
}

func (s *storage) BasketDelete(ctx context.Context, basketID string) error {
	// Lock basket map
	s.mutex.basket.Lock()
	defer s.mutex.basket.Unlock()

	// Delete
	delete(s.data.baskets, basketID)

	return nil
}

func (s *storage) ProductGet(ctx context.Context, productID string) (*entities.Product, error) {
	// Lock product map
	s.mutex.product.Lock()
	defer s.mutex.product.Unlock()

	// Get product from storage data
	if product, ok := s.data.products[productID]; ok {
		return &product, nil
	}

	// Product not found
	return nil, lanaerr.New(fmt.Errorf("product %s not found", productID), http.StatusNotFound)
}

func (s *storage) ProductList(ctx context.Context) ([]entities.Product, error) {
	// Lock product map
	s.mutex.product.Lock()
	defer s.mutex.product.Unlock()

	// TODO: Explain search & pagination

	// Get all products
	products := make([]entities.Product, 0)
	for _, p := range s.data.products {
		products = append(products, p)
	}

	return products, nil
}

func (s *storage) DiscountGet(ctx context.Context, discountID string) (*entities.Discount, error) {
	// Lock discount map
	s.mutex.discount.Lock()
	defer s.mutex.discount.Unlock()

	// Get discount from storage data
	if discount, ok := s.data.discounts[discountID]; ok {
		return &discount, nil
	}

	// Discount not found
	return nil, lanaerr.New(fmt.Errorf("discount %s not found", discountID), http.StatusNotFound)
}
