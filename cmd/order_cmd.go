package cmd

import (
	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
)

type CmdOrder struct {
	types.Store
}

func (cmd CmdOrder) AddOrder(restaurantID int, productsOrder schema.ProductOrder) (schema.Order, error) {
	return schema.Order{}, nil
}
