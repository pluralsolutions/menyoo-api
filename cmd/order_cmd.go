package cmd

import (
	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
)

type CmdOrder struct {
	types.Store
}

func (cmd CmdOrder) CreateOrder(order schema.Order) (result schema.Order, err error) {
	order.Status = "requested"
	for _, pd := range order.ProductsOrder {
		// productPrice = pd.ProductID
		pd.TotalPriceCents = pd.Quantity
	}
	return result, err
}
