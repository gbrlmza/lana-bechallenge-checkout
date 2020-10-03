package checkout

import (
	"context"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
)

/*
	The domain(aka core, services or usecases) are the core of the application. It is a technology agnostic
	component that contains all the business logic. For the domain how the application is served or accessed
	and where the data is stored doesn't matter. We can change from REST to gRPC, or from one storage type to
	another but the domain logic remains the same.
*/

type Service interface {
	// Basket
	BasketCreate(ctx context.Context) (*entities.Basket, error)
	BasketGet(ctx context.Context, basketID string) (*entities.Basket, error)
	BasketDelete(ctx context.Context, basketID string) error
	BasketAddItems(ctx context.Context, basketID string, items []entities.ItemDetail) error
	BasketRemoveItem(ctx context.Context, basketID string, productID string, quantity int) error

	// Product
	ProductList(ctx context.Context) ([]entities.Product, error)
	ProductGet(ctx context.Context, productCode string) (*entities.Product, error)
}

type service struct {
	*Container
}

func NewService(cont *Container) Service {
	return &service{
		Container: cont,
	}
}
