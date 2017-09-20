package handler

import (
	"fmt"
	"net/http"

	"github.com/lucasgomide/menyoo-api/types"
)

func NewProductsHandler(cmd types.ProductCmd) *ProductHandler {
	return &ProductHandler{cmd}
}

func NewOrdersHandler(cmd types.OrderCmd) *OrderHandler {
	return &OrderHandler{cmd}
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	fmt.Fprintf(w, "%s", err)
}
