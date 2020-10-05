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
		products   map[string]entities.Product
		baskets    map[string]entities.Basket
		promotions map[string]entities.Promotion
	}
	mutex struct {
		product   sync.Mutex
		basket    sync.Mutex
		promotion sync.Mutex
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

	// Get all products
	products := make([]entities.Product, 0)
	for _, p := range s.data.products {
		products = append(products, p)
	}

	return products, nil
}

func (s *storage) PromotionGet(ctx context.Context, promotionID string) (*entities.Promotion, error) {
	// Lock promotion map
	s.mutex.promotion.Lock()
	defer s.mutex.promotion.Unlock()

	// Get promotion from storage data
	if discount, ok := s.data.promotions[promotionID]; ok {
		return &discount, nil
	}

	// Promotion not found
	return nil, lanaerr.New(fmt.Errorf("promotion %s not found", promotionID), http.StatusNotFound)
}
