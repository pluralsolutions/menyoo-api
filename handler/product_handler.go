package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lucasgomide/menyoo-api/types"
)

type ProductHandler struct {
	types.ProductCmd
}

func (d ProductHandler) Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	restaurantID, err := strconv.Atoi(query.Get("restaurant_id"))

	if err != nil {
		badRequest(w, err)
		return
	}

	result, err := d.ProductsByRestaurant(restaurantID)
	if err != nil {
		badRequest(w, err)
		return
	}

	w.WriteHeader(200)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
