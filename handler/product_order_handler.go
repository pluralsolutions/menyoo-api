package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lucasgomide/menyoo-api/schema"

	"github.com/gorilla/mux"
	"github.com/lucasgomide/menyoo-api/types"
)

type ProductOrderHandler struct {
	types.ProductOrderCmd
}

func (cmd ProductOrderHandler) UpdateQuantity(w http.ResponseWriter, r *http.Request) {
	uID := r.Header.Get("uid")
	if uID == "" {
		UnauthorizedRequest(w)
		return
	}

	var params struct {
		Quantity       int `json:"quantity"`
		UserID         string
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

	if restaurantID <= 0 || orderID <= 0 || productOrderID <= 0 {
		badRequest(w, missingParamsError())
		return
	}

	params.RestaurantID = restaurantID
	params.OrderID = orderID
	params.ProductOrderID = productOrderID
	params.UserID = uID

	result, err := cmd.UpdateProductOrderQuantity(params)

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusOK, result)
}

func (d ProductOrderHandler) ByUser(w http.ResponseWriter, r *http.Request) {
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

	uID := r.Header.Get("uid")
	if uID == "" {
		UnauthorizedRequest(w)
		return
	}

	result, err := d.ProductsByUser(restaurantID, uID)

	type S struct {
		Products []schema.ProductOrder `json:"products"`
	}

	serializer := S{Products: result}

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusOK, serializer)
}
