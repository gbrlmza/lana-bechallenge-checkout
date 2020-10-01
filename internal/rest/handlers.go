package rest

import (
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type Handler struct {
	srv checkout.Service
}

func NewHandler(srv checkout.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

const (
	UrlParamBasketID  = "basketID"
	UrlParamProductID = "productID"
)

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (h Handler) BasketCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Service call
	basket, err := h.srv.BasketCreate(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Success
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, basket)
}

func (h Handler) BasketGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	basketID := chi.URLParam(r, UrlParamBasketID)

	// Service call
	basket, err := h.srv.BasketGet(ctx, basketID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, basket)
}

func (h Handler) BasketDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	basketID := chi.URLParam(r, UrlParamBasketID)

	// Service call
	err := h.srv.BasketDelete(ctx, basketID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
}

func (h Handler) BasketAddProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	basketID := chi.URLParam(r, UrlParamBasketID)
	productID := chi.URLParam(r, UrlParamProductID)
	quantity, err := getBasketProductQuantity(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Service call
	err = h.srv.BasketAddProduct(ctx, basketID, productID, quantity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	// Success
	w.WriteHeader(http.StatusOK)
}

func (h Handler) BasketRemoveProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	basketID := chi.URLParam(r, UrlParamBasketID)
	productID := chi.URLParam(r, UrlParamProductID)
	quantity, err := getBasketProductQuantity(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Service call
	err = h.srv.BasketRemoveProduct(ctx, basketID, productID, quantity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
}

func (h Handler) ProductList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Service call
	products, err := h.srv.ProductList(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, products)
}

func (h Handler) ProductGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	productID := chi.URLParam(r, UrlParamProductID)

	// Service call
	product, err := h.srv.ProductGet(ctx, productID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, product)
}
