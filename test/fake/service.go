package fake

import (
	"context"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/stretchr/testify/mock"
)

type FakeService struct {
	mock.Mock
}

func (f FakeService) BasketCreate(ctx context.Context) (*entities.Basket, error) {
	args := f.Called(ctx)
	return args.Get(0).(*entities.Basket), args.Error(1)
}

func (f FakeService) BasketGet(ctx context.Context, basketID string) (*entities.Basket, error) {
	args := f.Called(ctx, basketID)
	return args.Get(0).(*entities.Basket), args.Error(1)
}

func (f FakeService) BasketDelete(ctx context.Context, basketID string) error {
	args := f.Called(ctx, basketID)
	return args.Error(0)
}

func (f FakeService) BasketAddItem(ctx context.Context, basketID string, itemDetail entities.ItemDetail) error {
	args := f.Called(ctx, basketID, itemDetail)
	return args.Error(0)
}

func (f FakeService) BasketRemoveItem(ctx context.Context, basketID string, itemDetail entities.ItemDetail) error {
	args := f.Called(ctx, basketID, itemDetail)
	return args.Error(0)
}

func (f FakeService) ProductList(ctx context.Context) ([]entities.Product, error) {
	args := f.Called(ctx)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (f FakeService) ProductGet(ctx context.Context, productID string) (*entities.Product, error) {
	args := f.Called(ctx, productID)
	return args.Get(0).(*entities.Product), args.Error(1)
}
