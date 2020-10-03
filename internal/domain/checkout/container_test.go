package checkout

import (
	"context"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/stretchr/testify/mock"
)

//==================================================================================================
// Fake Storage
//==================================================================================================
type FakeStorage struct {
	mock.Mock
}

func (f FakeStorage) BasketSave(ctx context.Context, basket *entities.Basket) error {
	args := f.Called(ctx, basket)
	return args.Error(0)
}

func (f FakeStorage) BasketGet(ctx context.Context, basketID string) (*entities.Basket, error) {
	args := f.Called(ctx, basketID)
	return args.Get(0).(*entities.Basket), args.Error(1)
}

func (f FakeStorage) BasketDelete(ctx context.Context, basketID string) error {
	args := f.Called(ctx, basketID)
	return args.Error(0)
}

func (f FakeStorage) ProductGet(ctx context.Context, productID string) (*entities.Product, error) {
	args := f.Called(ctx, productID)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (f FakeStorage) ProductList(ctx context.Context) ([]entities.Product, error) {
	args := f.Called(ctx)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (f FakeStorage) PromotionGet(ctx context.Context, promotionID string) (*entities.Promotion, error) {
	args := f.Called(ctx, promotionID)
	return args.Get(0).(*entities.Promotion), args.Error(1)
}

//==================================================================================================
// Fake Locker
//==================================================================================================
type FakeLocker struct {
	mock.Mock
}

func (f FakeLocker) Lock(ctx context.Context, resource string) error {
	args := f.Called(ctx, resource)
	return args.Error(0)
}

func (f FakeLocker) Unlock(ctx context.Context, resource string) error {
	args := f.Called(ctx, resource)
	return args.Error(0)
}
