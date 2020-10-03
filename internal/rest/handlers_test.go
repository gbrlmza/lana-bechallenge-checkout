package rest_test

import (
	"context"
	"errors"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/rest"
	"github.com/gbrlmza/lana-bechallenge-checkout/test/fake"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandler_Ping_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	s := &fake.FakeService{}
	h := rest.NewHandler(s)
	//b := ioutil.NopCloser(strings.NewReader(`{"flag_type_id":"standard"}`))
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	w := httptest.NewRecorder()

	// When
	h.Ping(w, r)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestHandler_BasketCreate_Error(t *testing.T) {
	// Given
	ctx := context.Background()
	s := &fake.FakeService{}
	h := rest.NewHandler(s)
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "", nil)
	w := httptest.NewRecorder()

	s.On("BasketCreate", ctx).Return(&entities.Basket{}, errors.New("basket-create-error"))

	// When
	h.BasketCreate(w, r)

	// Then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "basket-create-error", w.Body.String())
}

func TestHandler_BasketCreate_Success(t *testing.T) {
	// Given
	ctx := context.Background()
	s := &fake.FakeService{}
	h := rest.NewHandler(s)
	r, _ := http.NewRequestWithContext(ctx, http.MethodPost, "", nil)
	w := httptest.NewRecorder()

	s.On("BasketCreate", ctx).Return(&entities.Basket{
		ID:        "1680cd34-931e-4b0c-b7e3-ab314d688398",
		CreatedAt: time.Date(2020, 10, 02, 0, 0, 0, 0, time.UTC),
		Items:     make(map[string]entities.BasketItem, 0),
	}, nil)

	// When
	h.BasketCreate(w, r)

	// Then
	expectedBasket := `{"id":"1680cd34-931e-4b0c-b7e3-ab314d688398","created_at":"2020-10-02T00:00:00Z","items":{},"subtotal":0,"discount":0,"total":0}`
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedBasket, strings.TrimSpace(w.Body.String()))
}
