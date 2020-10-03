package checkout

import (
	"context"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
)

// The container has all the external functionality required by the business logic(domain).
// Provides an abstraction to the domain, because the business logic doesn't(and shouldn't)
// known where the data is stored to or retrieved from or witch external service is accessed.
// This allow us to change databases, services or any external provider without affecting
// the business model.
//
// Also allows us to test the domain logic regardless of the specific repository implementations.

type Container struct {
	Storage Storage
	Locker  Locker
}

type Storage interface {
	// Basket
	BasketSave(ctx context.Context, basket *entities.Basket) error
	BasketGet(ctx context.Context, basketID string) (*entities.Basket, error)
	BasketDelete(ctx context.Context, basketID string) error

	// Product
	ProductGet(ctx context.Context, productID string) (*entities.Product, error)
	ProductList(ctx context.Context) ([]entities.Product, error)

	// Promotion
	PromotionGet(ctx context.Context, promotionID string) (*entities.Promotion, error)
}

type Locker interface {
	Lock(ctx context.Context, resource string) error
	Unlock(ctx context.Context, resource string) error
}
