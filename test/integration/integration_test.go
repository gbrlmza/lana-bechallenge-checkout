package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/cmd/container"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func buildTestDependencies() http.Handler {
	ctx := context.Background()

	// Container & service initialization
	container := container.NewContainer(ctx)
	service := checkout.NewService(container)

	// Handler
	handler := rest.NewHandler(service)
	router := handler.RouterInit()

	return router
}

func TestHandler_Integration_Success(t *testing.T) {
	//==================================================================================================
	// App initialization
	//==================================================================================================
	var w *httptest.ResponseRecorder
	var r *http.Request
	var url string
	var b *strings.Reader
	router := buildTestDependencies()
	basket := &entities.Basket{}

	// ### Integration test steps:
	// 1-Create basket
	// 2-Add items to basket
	// 3-Add more items to basket
	// 4-Remove items from basket
	// 5-Get basket
	// 6-Delete basket
	// 7-Get basket

	//==================================================================================================
	// 1-Create basket
	//==================================================================================================
	w = httptest.NewRecorder()
	r, _ = http.NewRequest(http.MethodPost, "/v1/baskets", nil)
	router.ServeHTTP(w, r)

	// Get basket data from response
	json.Unmarshal(w.Body.Bytes(), basket)

	//==================================================================================================
	// 2-Add item to basket: 2 PEN
	//==================================================================================================
	w = httptest.NewRecorder()
	b = strings.NewReader(`{"id":"PEN","quantity":2}`)
	url = fmt.Sprintf("/v1/baskets/%s/items", basket.ID)
	r, _ = http.NewRequest(http.MethodPost, url, b)
	router.ServeHTTP(w, r)

	//==================================================================================================
	// 3-Add more items to basket: 30 TSHIRT
	//==================================================================================================
	w = httptest.NewRecorder()
	b = strings.NewReader(`{"id":"TSHIRT","quantity":30}`)
	url = fmt.Sprintf("/v1/baskets/%s/items", basket.ID)
	r, _ = http.NewRequest(http.MethodPost, url, b)
	router.ServeHTTP(w, r)

	//==================================================================================================
	// 4-Remove items from basket: 27 TSHIRT
	//==================================================================================================
	w = httptest.NewRecorder()
	url = fmt.Sprintf("/v1/baskets/%s/items/%s?quantity=27", basket.ID, "TSHIRT")
	r, _ = http.NewRequest(http.MethodDelete, url, nil)
	router.ServeHTTP(w, r)

	//==================================================================================================
	// 5-Get basket
	//==================================================================================================
	w = httptest.NewRecorder()
	url = fmt.Sprintf("/v1/baskets/%s", basket.ID)
	r, _ = http.NewRequest(http.MethodGet, url, nil)
	router.ServeHTTP(w, r)

	// Get basket data from response
	newBasket := &entities.Basket{}
	json.Unmarshal(w.Body.Bytes(), newBasket)

	// Assert expected basket values
	assert.Equal(t, 60.0, newBasket.Subtotal)
	assert.Equal(t, 15.0, newBasket.Discount)
	assert.Equal(t, 45.0, newBasket.Total)
	assert.Equal(t, 2, len(newBasket.Items))

	// PEN: 2 units, $10 total, $5 discount
	assert.Equal(t, uint(2), newBasket.Items["PEN"].Quantity)
	assert.Equal(t, 10.0, newBasket.Items["PEN"].Total)
	assert.Equal(t, 5.0, newBasket.Items["PEN"].Discount)

	// TSHIRT: 3 units, $60 total, $15 discount
	assert.Equal(t, uint(3), newBasket.Items["TSHIRT"].Quantity)
	assert.Equal(t, 60.0, newBasket.Items["TSHIRT"].Total)
	assert.Equal(t, 15.0, newBasket.Items["TSHIRT"].Discount)

	//==================================================================================================
	// 6-Delete basket
	//==================================================================================================
	w = httptest.NewRecorder()
	url = fmt.Sprintf("/v1/baskets/%s", basket.ID)
	r, _ = http.NewRequest(http.MethodDelete, url, nil)
	router.ServeHTTP(w, r)

	//==================================================================================================
	// 6-Get basket
	//==================================================================================================
	w = httptest.NewRecorder()
	url = fmt.Sprintf("/v1/baskets/%s", basket.ID)
	r, _ = http.NewRequest(http.MethodGet, url, nil)
	router.ServeHTTP(w, r)

	// Basket shouldn't exists
	assert.Equal(t, http.StatusNotFound, w.Code)
}
