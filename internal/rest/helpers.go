package rest

import (
	"fmt"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/utils/lanaerr"
	"net/http"
	"strconv"
)

type BasketProduct struct {
	Quantity int `json:"quantity"`
}

func (h Handler) HandleError(w http.ResponseWriter, err error) {
	lErr := lanaerr.FromErr(err)

	w.WriteHeader(lErr.GetStatusCode())
	w.Write([]byte(lErr.Error()))
}

func (h Handler) GetQueryParamUintValue(r *http.Request, name string, defaultValue uint) (uint, error) {
	value := r.URL.Query().Get(name)
	if value == "" {
		return defaultValue, nil
	}

	intValue, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return 0, lanaerr.New(fmt.Errorf("invalid %s value: %s", name, value), http.StatusBadRequest)
	}

	return uint(intValue), nil
}
