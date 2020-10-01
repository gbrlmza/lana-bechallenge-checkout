package checkout

import (
	"context"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
)

func (s *service) BasketCreate(ctx context.Context) (*entities.Basket, error) {
	return nil, nil
}

func (s *service) BasketGet(ctx context.Context, basketID string) (*entities.Basket, error) {
	return nil, nil
}

func (s *service) BasketDelete(ctx context.Context, basketID string) error {
	return nil
}

func (s *service) BasketAddProduct(ctx context.Context, basketID string, productCode string, quantity int) error {
	return nil
}

func (s *service) BasketRemoveProduct(ctx context.Context, basketID string, productCode string, quantity int) error {
	return nil
}
