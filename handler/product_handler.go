package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lucasgomide/menyoo-api/types"
)

type ProductHandler struct {
	types.ProductCmd
}

func (d ProductHandler) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	restaurantID, err := strconv.Atoi(params["restaurant_id"])

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

func (d ProductHandler) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productID, err := strconv.Atoi(params["product_id"])
	restaurantID, err := strconv.Atoi(params["restaurant_id"])

	if err != nil {
		badRequest(w, err)
		return
	}

	result, err := d.ProductByRestaurantAndID(restaurantID, productID)

	if err != nil {
		badRequest(w, err)
		return
	}

	w.WriteHeader(200)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
