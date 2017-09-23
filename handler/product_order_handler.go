package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lucasgomide/menyoo-api/types"
)

type ProductOrderHandler struct {
	types.ProductOrderCmd
}

func (cmd ProductOrderHandler) UpdateQuantity(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Quantity       int    `json:"quantity"`
		UserID         string `json:"user_id"`
		RestaurantID   int
		OrderID        int
		ProductOrderID int
	}
	muxParams := mux.Vars(r)

	restaurantID, err := strconv.Atoi(muxParams["restaurant_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	orderID, err := strconv.Atoi(muxParams["order_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	productOrderID, err := strconv.Atoi(muxParams["product_order_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		badRequest(w, err)
		return
	}

	if restaurantID <= 0 || orderID <= 0 || productOrderID <= 0 || params.UserID == "" {
		badRequest(w, missingParamsError())
		return
	}

	params.RestaurantID = restaurantID
	params.OrderID = orderID
	params.ProductOrderID = productOrderID

	result, err := cmd.UpdateProductOrderQuantity(params)

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusOK, result)
}

func missingParamsError() error {
	return errors.New("You missing some parameters")
}
