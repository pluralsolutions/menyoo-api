package cmd

import (
	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
)

type CmdProduct struct {
	types.Store
}

func (cmd CmdProduct) ProductsByRestaurant(restaurantID int) ([]schema.Product, error) {
	result, err := cmd.Store.ProductsByRestaurant(restaurantID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cmd CmdProduct) ProductByRestaurantAndID(restaurantID int, productID int) (result []schema.Product, err error) {
	result, err = cmd.Store.ProductByRestaurantAndID(restaurantID, productID)

	if err != nil {
		return result, err
	}
	return result, nil
}
