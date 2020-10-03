package checkout

import (
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_service_ProductList_Success(t *testing.T) {
	// Given
	st := buildTestDependencies()
	st.Storage.On("ProductList", st.Ctx).Return([]entities.Product{}, nil)

	// When
	products, err := st.Service.ProductList(st.Ctx)

	// Then
	assert.NotNil(t, products)
	assert.Nil(t, err)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}

func Test_service_ProductGet_Success(t *testing.T) {
	// Given
	st := buildTestDependencies()
	productID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	st.Storage.On("ProductGet", st.Ctx, productID).Return(&entities.Product{}, nil)

	// When
	product, err := st.Service.ProductGet(st.Ctx, productID)

	// Then
	assert.NotNil(t, product)
	assert.Nil(t, err)
	st.Storage.AssertExpectations(t)
	st.Locker.AssertExpectations(t)
}
