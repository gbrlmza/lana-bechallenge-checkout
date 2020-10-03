package rest_test

import (
	"errors"
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/rest"
	"github.com/gbrlmza/lana-bechallenge-checkout/test/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Ping_Success(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()

	// When
	r, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketCreate_Error(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()

	srv.On("BasketCreate", mock.Anything).Return(&entities.Basket{}, errors.New("create-error"))

	// When
	r, _ := http.NewRequest(http.MethodPost, "/v1/baskets", nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "create-error", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketCreate_Success(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()

	srv.On("BasketCreate", mock.Anything).Return(&entities.Basket{
		ID: "1680cd34-931e-4b0c-b7e3-ab314d688398",
	}, nil)

	// When
	r, _ := http.NewRequest(http.MethodPost, "/v1/baskets", nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusCreated, w.Code)
	expectedBody := `{"id":"1680cd34-931e-4b0c-b7e3-ab314d688398","created_at":"0001-01-01T00:00:00Z","items":null,"subtotal":0,"discount":0,"total":0}`
	assert.Equal(t, expectedBody, strings.TrimSpace(w.Body.String()))
	srv.AssertExpectations(t)
}

func TestHandler_BasketGet_Error(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	srv.On("BasketGet", mock.Anything, basketID).Return(&entities.Basket{}, errors.New("get-error"))

	// When
	r, _ := http.NewRequest(http.MethodGet, "/v1/baskets/"+basketID, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "get-error", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketGet_Success(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	srv.On("BasketGet", mock.Anything, basketID).Return(&entities.Basket{
		ID: "1680cd34-931e-4b0c-b7e3-ab314d688398",
	}, nil)

	// When
	r, _ := http.NewRequest(http.MethodGet, "/v1/baskets/"+basketID, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"id":"1680cd34-931e-4b0c-b7e3-ab314d688398","created_at":"0001-01-01T00:00:00Z","items":null,"subtotal":0,"discount":0,"total":0}`
	assert.Equal(t, expectedBody, strings.TrimSpace(w.Body.String()))
	srv.AssertExpectations(t)
}

func TestHandler_BasketDelete_Error(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	srv.On("BasketDelete", mock.Anything, basketID).Return(errors.New("delete-error"))

	// When
	r, _ := http.NewRequest(http.MethodDelete, "/v1/baskets/"+basketID, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "delete-error", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketDelete_Success(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	srv.On("BasketDelete", mock.Anything, basketID).Return(nil)

	// When
	r, _ := http.NewRequest(http.MethodDelete, "/v1/baskets/"+basketID, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketAddItem_PayloadError(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"

	// When
	b := strings.NewReader(`{"quantity":"TEXT"}`)
	r, _ := http.NewRequest(http.MethodPost, "/v1/baskets/"+basketID+"/items", b)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "payload error", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketAddItem_ServiceError(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	srv.On("BasketAddItem", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("add-error"))

	// When
	b := strings.NewReader(`{"id": "PEN","quantity": 1}`)
	r, _ := http.NewRequest(http.MethodPost, "/v1/baskets/"+basketID+"/items", b)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "add-error", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketAddItem_Success(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	srv.On("BasketAddItem", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// When
	b := strings.NewReader(`{"id": "PEN","quantity": 1}`)
	r, _ := http.NewRequest(http.MethodPost, "/v1/baskets/"+basketID+"/items", b)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketRemoveItem_QueryError(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	productID := "PEN"

	// When
	url := fmt.Sprintf("/v1/baskets/%s/items/%s?quantity=TEXT", basketID, productID)
	r, _ := http.NewRequest(http.MethodDelete, url, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "invalid quantity value: TEXT", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketRemoveItem_ServiceError(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	productID := "PEN"
	srv.On("BasketRemoveItem", mock.Anything, basketID, entities.ItemDetail{
		ProductID: productID,
		Quantity:  1,
	}).Return(errors.New("remove-error"))

	// When
	url := fmt.Sprintf("/v1/baskets/%s/items/%s?quantity=1", basketID, productID)
	r, _ := http.NewRequest(http.MethodDelete, url, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "remove-error", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_BasketRemoveItem_Success(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	basketID := "1680cd34-931e-4b0c-b7e3-ab314d688398"
	productID := "PEN"
	srv.On("BasketRemoveItem", mock.Anything, basketID, entities.ItemDetail{
		ProductID: productID,
		Quantity:  1,
	}).Return(nil)

	// When
	url := fmt.Sprintf("/v1/baskets/%s/items/%s", basketID, productID)
	r, _ := http.NewRequest(http.MethodDelete, url, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_ProductList_ServiceError(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	srv.On("ProductList", mock.Anything).Return([]entities.Product{}, errors.New("list-error"))

	// When
	r, _ := http.NewRequest(http.MethodGet, "/v1/products", nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "list-error", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_ProductList_Success(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	srv.On("ProductList", mock.Anything).Return([]entities.Product{}, nil)

	// When
	r, _ := http.NewRequest(http.MethodGet, "/v1/products", nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", strings.TrimSpace(w.Body.String()))
	srv.AssertExpectations(t)
}

func TestHandler_ProductGet_ServiceError(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	productID := "PEN"
	srv.On("ProductGet", mock.Anything, productID).Return(&entities.Product{}, errors.New("get-error"))

	// When
	r, _ := http.NewRequest(http.MethodGet, "/v1/products/"+productID, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "get-error", w.Body.String())
	srv.AssertExpectations(t)
}

func TestHandler_ProductGet_Success(t *testing.T) {
	// Given
	srv := &fake.FakeService{}
	handler := rest.NewHandler(srv)
	router := handler.RouterInit()
	w := httptest.NewRecorder()
	productID := "PEN"
	srv.On("ProductGet", mock.Anything, productID).Return(&entities.Product{
		ID:    productID,
		Name:  "Lana Pen",
		Price: 10.50,
	}, nil)

	// When
	r, _ := http.NewRequest(http.MethodGet, "/v1/products/"+productID, nil)
	router.ServeHTTP(w, r)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"id":"PEN","name":"Lana Pen","price":10.5,"promotion_id":null}`
	assert.Equal(t, expectedBody, strings.TrimSpace(w.Body.String()))
	srv.AssertExpectations(t)
}
