package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/plural-solutions/menyoo-api/types"
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

	if restaurantID <= 0 {
		badRequest(w, missingParamsError())
		return
	}

	result, err := d.ProductsByRestaurant(restaurantID)
	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusOK, result)
}

func (d ProductHandler) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productID, err := strconv.Atoi(params["product_id"])

	if err != nil {
		badRequest(w, err)
		return
	}

	restaurantID, err := strconv.Atoi(params["restaurant_id"])

	if err != nil {
		badRequest(w, err)
		return
	}

	if restaurantID <= 0 || productID <= 0 {
		badRequest(w, missingParamsError())
		return
	}

	result, err := d.ProductByRestaurantAndID(restaurantID, productID)

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusOK, result)
}
