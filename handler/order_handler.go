package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
