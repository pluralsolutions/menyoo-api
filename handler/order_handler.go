package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
)

type OrderHandler struct {
	types.OrderCmd
}

func (cmd OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var order schema.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		fmt.Print("Error ", err)
	}

	result, err := cmd.CreateOrder(order)

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusCreated, result)
}

func (cmd OrderHandler) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	restaurantID, err := strconv.Atoi(params["restaurant_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	orderID, err := strconv.Atoi(params["order_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	if restaurantID <= 0 || orderID <= 0 {
		badRequest(w, missingParamsError())
		return
	}

	uID := r.Header.Get("uid")
	if uID == "" {
		UnauthorizedRequest(w)
		return
	}

	order, err := cmd.ShowOrder(schema.Order{UserID: uID, RestaurantID: restaurantID, ID: orderID})

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusOK, order)
}

func (cmd OrderHandler) Place(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	restaurantID, err := strconv.Atoi(params["restaurant_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	orderID, err := strconv.Atoi(params["order_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	if restaurantID <= 0 || orderID <= 0 {
		badRequest(w, missingParamsError())
		return
	}

	uID := r.Header.Get("uid")
	if uID == "" {
		UnauthorizedRequest(w)
		return
	}

	order, err := cmd.PlaceOrder(
		schema.Order{UserID: uID, RestaurantID: restaurantID, ID: orderID},
	)

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusOK, order)
}

func (cmd OrderHandler) CurrentOrder(w http.ResponseWriter, r *http.Request) {
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

	result, err := cmd.OrderCmd.CurrentOrder(
		schema.Order{RestaurantID: restaurantID, UserID: uID},
	)

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusOK, result)
}
