package handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	// "github.com/lucasgomide/zenon/schema"
	"net/http"

	"github.com/lucasgomide/menyoo-api/types"
)

type ProductHandler struct {
	types.Cmd
}

func NewProductsHandler(cmd types.Cmd) *ProductHandler {
	return &ProductHandler{cmd}
}

func (d ProductHandler) Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	restaurantID, err := strconv.Atoi(query.Get("restaurant_id"))

	if err != nil {
		badRequest(w, err)
		return
	}

	result, err := d.ProductsByRestaurant(restaurantID)
	fmt.Print(result)
	if err != nil {
		badRequest(w, err)
		return
	}

	w.WriteHeader(200)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	fmt.Fprintf(w, "%s", err)
}
