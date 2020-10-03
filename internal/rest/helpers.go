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

func (h Handler) GetQueryParamIntValue(r *http.Request, name string, defaultValue int) (int, error) {
	value := r.URL.Query().Get(name)
	if value == "" {
		return defaultValue, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, lanaerr.New(fmt.Errorf("invalid %s value: %s", name, value), http.StatusBadRequest)
	}

	return intValue, nil
}
