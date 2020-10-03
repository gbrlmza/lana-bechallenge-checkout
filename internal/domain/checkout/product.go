package checkout

import (
	"context"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
)

func (s *service) ProductList(ctx context.Context) ([]entities.Product, error) {
	return s.Storage.ProductList(ctx)
}

func (s *service) ProductGet(ctx context.Context, productID string) (*entities.Product, error) {
	return s.Storage.ProductGet(ctx, productID)
}
