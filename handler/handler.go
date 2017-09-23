package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasgomide/menyoo-api/types"
)

type errorBadRequest struct {
	Err string `json:"error"`
}

func NewProductsHandler(cmd types.ProductCmd) *ProductHandler {
	return &ProductHandler{cmd}
}

func NewOrdersHandler(cmd types.OrderCmd) *OrderHandler {
	return &OrderHandler{cmd}
}

func badRequest(w http.ResponseWriter, err error) {
	body := errorBadRequest{err.Error()}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&body)
}
