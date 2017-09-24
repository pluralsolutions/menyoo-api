package cmd

import (
	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
)

type CmdProduct struct {
	types.Store
}

func (cmd CmdProduct) ProductsByRestaurant(restaurantID int) (
	products []schema.Product,
	err error,
) {

	products, err = cmd.Store.FindProductsBy(
		schema.Product{RestaurantID: restaurantID},
	)
	return products, err
}

func (cmd CmdProduct) ProductByRestaurantAndID(restaurantID int, productID int) (product schema.Product, err error) {
	product, err = cmd.Store.FindFullProductBy(
		schema.Product{RestaurantID: restaurantID, ID: productID},
	)

	return product, err
}
