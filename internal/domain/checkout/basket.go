package checkout

import (
	"context"
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/repository/metrics"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/utils/lanaerr"
	"net/http"
)

func (s *service) BasketCreate(ctx context.Context) (*entities.Basket, error) {
	basket := entities.NewBasket()
	if err := s.Storage.BasketSave(ctx, basket); err != nil {
		return nil, err
	}

	// Metric
	metrics.Counter(ctx, "basket_created", 1)

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

func (s *service) BasketAddItem(ctx context.Context, basketID string, itemDetail entities.ItemDetail) error {
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
	product, err := s.Storage.ProductGet(ctx, itemDetail.ProductID)
	if err != nil {
		return err
	}

	// Check if product is already in the basket
	basketItem := basket.GetItem(itemDetail.ProductID)
	if basketItem == nil { // If not, create new basket item
		var promotion *entities.Promotion
		if product.PromotionID != nil {
			if promotion, err = s.Storage.PromotionGet(ctx, *product.PromotionID); err != nil {
				return err
			}
		}
		basketItem = entities.NewBasketItem(*product, promotion)
	}

	// Add quantity
	basketItem.AddQuantity(itemDetail.Quantity)

	// Save item in basket
	basket.SaveItem(basketItem)

	// Save basket
	if err := s.Storage.BasketSave(ctx, basket); err != nil {
		return err
	}

	// Metric
	metrics.Counter(ctx, "basket_items_added", float64(itemDetail.Quantity))

	// Done
	return nil
}

func (s *service) BasketRemoveItem(ctx context.Context, basketID string, itemDetail entities.ItemDetail) error {
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

	// Check if product is in the basket
	basketItem := basket.GetItem(itemDetail.ProductID)
	if basketItem == nil {
		err := fmt.Errorf("item %s not found in basket %s", itemDetail.ProductID, basketID)
		return lanaerr.New(err, http.StatusNotFound)
	}

	// Remove quantity
	if err := basketItem.RemoveQuantity(itemDetail.Quantity); err != nil {
		return err
	}

	// Update item in basket
	basket.SaveItem(basketItem)

	// Save basket
	if err := s.Storage.BasketSave(ctx, basket); err != nil {
		return err
	}

	// Done
	return nil
}

func (s *service) getBasketLockKey(basketID string) string {
	return fmt.Sprintf("basket-%s", basketID)
}
