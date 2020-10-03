package rest

import (
	"encoding/json"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout/entities"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/utils/lanaerr"
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
	UrlParamBasketID   = "basketID"
	UrlParamProductID  = "productID"
	QueryParamQuantity = "quantity"
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
		h.HandleError(w, err)
		return
	}

	// Success
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, basket)
}

func (h Handler) BasketGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	basketID := chi.URLParam(r, UrlParamBasketID)

	// Service call
	basket, err := h.srv.BasketGet(ctx, basketID)
	if err != nil {
		h.HandleError(w, err)
		return
	}

	// Success
	render.JSON(w, r, basket)
}

func (h Handler) BasketDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	basketID := chi.URLParam(r, UrlParamBasketID)

	// Service call
	if err := h.srv.BasketDelete(ctx, basketID); err != nil {
		h.HandleError(w, err)
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
}

func (h Handler) BasketAddItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	basketID := chi.URLParam(r, UrlParamBasketID)

	// Item from payload
	item := entities.ItemDetail{}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		err = lanaerr.New(err, http.StatusBadRequest)
		h.HandleError(w, err)
		return
	}

	// Service call
	if err := h.srv.BasketAddItem(ctx, basketID, item); err != nil {
		h.HandleError(w, err)
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
}

func (h Handler) BasketRemoveItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	basketID := chi.URLParam(r, UrlParamBasketID)
	productID := chi.URLParam(r, UrlParamProductID)
	quantity, err := h.GetQueryParamUintValue(r, QueryParamQuantity, 1)
	if err != nil {
		h.HandleError(w, err)
		return
	}
	itemDetail := entities.ItemDetail{
		ProductID: productID,
		Quantity:  quantity,
	}

	// Service call
	err = h.srv.BasketRemoveItem(ctx, basketID, itemDetail)
	if err != nil {
		h.HandleError(w, err)
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
		h.HandleError(w, err)
		return
	}

	// Success
	render.JSON(w, r, products)
}

func (h Handler) ProductGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Request params
	productID := chi.URLParam(r, UrlParamProductID)

	// Service call
	product, err := h.srv.ProductGet(ctx, productID)
	if err != nil {
		h.HandleError(w, err)
		return
	}

	// Success
	render.JSON(w, r, product)
}
