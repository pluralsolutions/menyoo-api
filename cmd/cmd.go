package cmd

import (
	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
)

type CmdProduct struct {
	types.Store
}

func NewCmdProduct(store types.Store) *CmdProduct {
	return &CmdProduct{store}
}

func (cmd CmdProduct) AddProduct(product *schema.Product) (string, error) {
	return "Okay, " + product.Title + " have to pay ", nil
}

func (cmd CmdProduct) ProductsByRestaurant(restaurantID int) ([]schema.Product, error) {
	result, err := cmd.Store.ProductsByRestaurant(restaurantID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
