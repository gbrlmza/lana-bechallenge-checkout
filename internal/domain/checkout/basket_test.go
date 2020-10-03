package checkout

import (
	"context"
	"errors"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type serviceTest struct {
	Ctx       context.Context
	Locker    *FakeLocker
	Storage   *FakeStorage
	Container *Container
	Service   Service
}

func buildTestDependencies() serviceTest {
	st := serviceTest{
		Ctx:     context.Background(),
		Locker:  &FakeLocker{},
		Storage: &FakeStorage{},
	}
	st.Container = &Container{
		Locker:  st.Locker,
		Storage: st.Storage,
	}
	st.Service = NewService(st.Container)
	return st
}

func Test_service_BasketCreate_Error(t *testing.T) {
	// Given
	st := buildTestDependencies()
	st.Storage.On("BasketSave", st.Ctx, mock.Anything).Return(errors.New("save-error"))

	// When
	basket, err := st.Service.BasketCreate(st.Ctx)

	// Then
	assert.EqualError(t, err, "save-error")
	assert.Nil(t, basket)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketCreate_Success(t *testing.T) {
	// Given
	st := buildTestDependencies()
	st.Storage.On("BasketSave", st.Ctx, mock.Anything).Return(nil)

	// When
	basket, err := st.Service.BasketCreate(st.Ctx)

	// Then
	assert.NotNil(t, basket)
	assert.Nil(t, err)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketGet_Error(t *testing.T) {
	// Given
	st := buildTestDependencies()
	var b *entities.Basket
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(b, errors.New("get-error"))

	// When
	basket, err := st.Service.BasketGet(st.Ctx, basketID)

	// Then
	assert.EqualError(t, err, "get-error")
	assert.Nil(t, basket)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketGet_Success(t *testing.T) {
	// Given
	st := buildTestDependencies()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{}, nil)

	// When
	basket, err := st.Service.BasketGet(st.Ctx, basketID)

	// Then
	assert.NotNil(t, basket)
	assert.Nil(t, err)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketDelete_LockError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, "basket-1680cd34-931e-4b0c-b7e3-ab314d688398").
		Return(errors.New("lock-error"))

	// When
	err := st.Service.BasketDelete(st.Ctx, basketID)

	// Then
	assert.EqualError(t, err, "lock-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketDelete_StorageError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketDelete", st.Ctx, basketID).Return(errors.New("delete-error"))

	// When
	err := st.Service.BasketDelete(st.Ctx, basketID)

	// Then
	assert.EqualError(t, err, "delete-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketDelete_Success(t *testing.T) {
	// Given
	st := buildTestDependencies()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketDelete", st.Ctx, basketID).Return(nil)

	// When
	err := st.Service.BasketDelete(st.Ctx, basketID)

	// Then
	assert.Nil(t, err)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketAddItem_LockError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, "basket-1680cd34-931e-4b0c-b7e3-ab314d688398").
		Return(errors.New("lock-error"))

	// When
	err := st.Service.BasketAddItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "lock-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketAddItem_GetBasketError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).
		Return(&entities.Basket{}, errors.New("get-basket-error"))

	// When
	err := st.Service.BasketAddItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "get-basket-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketAddItem_GetProductError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{ID: basketID}, nil)
	st.Storage.On("ProductGet", st.Ctx, item.ProductID).
		Return(&entities.Product{}, errors.New("get-product-error"))

	// When
	err := st.Service.BasketAddItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "get-product-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketAddItem_GetPromotionError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	promotion := entities.Promotion{ID: "2X1"}
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{ID: basketID}, nil)
	st.Storage.On("ProductGet", st.Ctx, item.ProductID).Return(&entities.Product{
		ID:          "PEN",
		PromotionID: &promotion.ID,
	}, nil)
	st.Storage.On("PromotionGet", st.Ctx, promotion.ID).
		Return(&entities.Promotion{}, errors.New("get-promotion-error"))

	// When
	err := st.Service.BasketAddItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "get-promotion-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketAddItem_SaveBasketError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	promotion := entities.Promotion{ID: "2X1"}
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{
		ID:    basketID,
		Items: make(map[string]entities.BasketItem),
	}, nil)
	st.Storage.On("ProductGet", st.Ctx, item.ProductID).Return(&entities.Product{
		ID:          "PEN",
		PromotionID: &promotion.ID,
	}, nil)
	st.Storage.On("PromotionGet", st.Ctx, promotion.ID).Return(&entities.Promotion{}, nil)
	st.Storage.On("BasketSave", st.Ctx, mock.Anything).Return(errors.New("save-basket-error"))

	// When
	err := st.Service.BasketAddItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "save-basket-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketAddItem_Success(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	promotion := entities.Promotion{ID: "2X1"}
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{
		ID:    basketID,
		Items: make(map[string]entities.BasketItem),
	}, nil)
	st.Storage.On("ProductGet", st.Ctx, item.ProductID).Return(&entities.Product{
		ID:          "PEN",
		PromotionID: &promotion.ID,
	}, nil)
	st.Storage.On("PromotionGet", st.Ctx, promotion.ID).Return(&entities.Promotion{}, nil)
	st.Storage.On("BasketSave", st.Ctx, mock.Anything).Return(nil)

	// When
	err := st.Service.BasketAddItem(st.Ctx, basketID, item)

	// Then
	assert.Nil(t, err)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketRemoveItem_LockError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, "basket-1680cd34-931e-4b0c-b7e3-ab314d688398").
		Return(errors.New("lock-error"))

	// When
	err := st.Service.BasketRemoveItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "lock-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketRemoveItem_GetBasketError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).
		Return(&entities.Basket{}, errors.New("get-basket-error"))

	// When
	err := st.Service.BasketRemoveItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "get-basket-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketRemoveItem_ProductNotInBasketError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{ID: basketID}, nil)

	// When
	err := st.Service.BasketRemoveItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "item PEN not found in basket 1680cd34-931e-4b0c-b7e3-ab314d688398")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketRemoveItem_QuantityError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 10}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{
		ID: basketID,
		Items: map[string]entities.BasketItem{
			"PEN": {Product: entities.Product{ID: "PEN"}, Quantity: 1},
		},
	}, nil)

	// When
	err := st.Service.BasketRemoveItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "can't remove 10 PEN. item quantity: 1")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketRemoveItem_SaveBasketError(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 1}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{
		ID: basketID,
		Items: map[string]entities.BasketItem{
			"PEN": {Product: entities.Product{ID: "PEN"}, Quantity: 1},
		},
	}, nil)
	st.Storage.On("BasketSave", st.Ctx, mock.Anything).Return(errors.New("save-basket-error"))

	// When
	err := st.Service.BasketRemoveItem(st.Ctx, basketID, item)

	// Then
	assert.EqualError(t, err, "save-basket-error")
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_BasketRemoveItem_Success(t *testing.T) {
	// Given
	st := buildTestDependencies()
	item := entities.ItemDetail{ProductID: "PEN", Quantity: 1}
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Locker.On("Lock", st.Ctx, mock.Anything).Return(nil)
	st.Locker.On("Unlock", st.Ctx, mock.Anything).Return(nil)
	st.Storage.On("BasketGet", st.Ctx, basketID).Return(&entities.Basket{
		ID: basketID,
		Items: map[string]entities.BasketItem{
			"PEN": {Product: entities.Product{ID: "PEN"}, Quantity: 1},
		},
	}, nil)
	st.Storage.On("BasketSave", st.Ctx, mock.Anything).Return(nil)

	// When
	err := st.Service.BasketRemoveItem(st.Ctx, basketID, item)

	// Then
	assert.Nil(t, err)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}
