package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
)

type OrderHandler struct {
	types.OrderCmd
}

func (d OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var m schema.Order
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Print("Error", err)
	}

	fmt.Print(m)
}
