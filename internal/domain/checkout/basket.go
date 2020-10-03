package checkout

import (
	"context"
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
)

func (s *service) BasketCreate(ctx context.Context) (*entities.Basket, error) {
	// In a more realistic scenario the, a user may have only one basket or at least one basket
	// per session and the creation of basket may be idempotent and some other constraints.
	// This kind of logic is business logic and should be placed here in the domain service.

	basket := entities.NewBasket()
	if err := s.Storage.BasketSave(ctx, basket); err != nil {
		return nil, err
	}

	return basket, nil
}

func (s *service) BasketGet(ctx context.Context, basketID string) (*entities.Basket, error) {
	basket, err := s.Storage.BasketGet(ctx, basketID)
	if err != nil {
		return nil, err
	}

	return basket, nil
}

func (s *service) BasketDelete(ctx context.Context, basketID string) error {
	// Lock basket
	lockKey := s.getBasketLockKey(basketID)
	if err := s.Locker.Lock(ctx, lockKey); err != nil {
		return err
	}
	defer s.Locker.Unlock(ctx, lockKey)

	// Delete basket
	return s.Storage.BasketDelete(ctx, basketID)
}

func (s *service) BasketAddItems(ctx context.Context, basketID string, items []entities.ItemDetail) error {
	// Lock basket
	lockKey := s.getBasketLockKey(basketID)
	if err := s.Locker.Lock(ctx, lockKey); err != nil {
		return err
	}
	defer s.Locker.Unlock(ctx, lockKey)

	// Note: In this basic case there is no need for a lock keep alive function but in a real case
	// we will probably need a lock keep alive to ensure the resource remains locked before confirming
	// the changes because some of the operation can take longer than expected(network problems, lag,
	// slow repository, load, etc...). With and ACID storage this means to be sure the resource is
	// locked before committing the changes.

	// Get Basket
	basket, err := s.Storage.BasketGet(ctx, basketID)
	if err != nil {
		return err
	}

	for _, item := range items {
		// Obtain product
		product, err := s.Storage.ProductGet(ctx, item.ProductID)
		if err != nil {
			return err
		}

		// Add product to basket
		if err := basket.AddItem(product, item.Quantity); err != nil {
			return err
		}
	}

	// Update basket
	if err := s.Storage.BasketSave(ctx, basket); err != nil {
		return err
	}

	// Done
	return nil
}

func (s *service) BasketRemoveItem(ctx context.Context, basketID string, productID string, quantity int) error {
	// Lock basket
	lockKey := s.getBasketLockKey(basketID)
	if err := s.Locker.Lock(ctx, lockKey); err != nil {
		return err
	}
	defer s.Locker.Unlock(ctx, lockKey)

	// Get Basket
	basket, err := s.Storage.BasketGet(ctx, basketID)
	if err != nil {
		return err
	}

	// Obtain product
	product, err := s.Storage.ProductGet(ctx, productID)
	if err != nil {
		return err
	}

	// Remove product from basket
	if err := basket.RemoveItem(product, quantity); err != nil {
		return err
	}

	// Update basket
	if err := s.Storage.BasketSave(ctx, basket); err != nil {
		return err
	}

	// Done
	return nil
}

func (s *service) getBasketLockKey(basketID string) string {
	return fmt.Sprintf("backet-%s", basketID)
}
