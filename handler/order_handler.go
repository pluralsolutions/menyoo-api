package handler

import (
	"net/http"

	"github.com/lucasgomide/menyoo-api/types"
)

type OrderHandler struct {
	types.OrderCmd
}

func (d OrderHandler) Handler(w http.ResponseWriter, r *http.Request) {
}
