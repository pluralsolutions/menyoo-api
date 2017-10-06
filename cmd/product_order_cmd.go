package cmd

import (
	"github.com/plural-solutions/menyoo-api/schema"
	"github.com/plural-solutions/menyoo-api/types"
	"github.com/plural-solutions/menyoo-api/util"
)

type CmdProductOrder struct {
	types.Store
}

func (cmd CmdProductOrder) UpdateProductOrderQuantity(
	params interface{},
) (
	order schema.Order,
	err error,
) {

	vStruct := params.(struct {
		Quantity       int `json:"quantity"`
		UserID         string
		RestaurantID   int
		OrderID        int
		ProductOrderID int
	})

	_, err = cmd.Store.FindOrderBy(
		schema.Order{
			RestaurantID: vStruct.RestaurantID,
			UserID:       vStruct.UserID,
			ID:           vStruct.OrderID,
			Status:       "requested",
		},
	)

	if err != nil {
		return order, err
	}

	currentProductOrder, err := cmd.Store.FindFullProductOrderBy(
		schema.ProductOrder{OrderID: vStruct.OrderID, ID: vStruct.ProductOrderID},
	)

	if err != nil {
		return order, err
	}

	if currentProductOrder.Quantity == vStruct.Quantity {
		return cmd.Store.FindFullOrderBy(schema.Order{ID: currentProductOrder.OrderID})
	}

	if vStruct.Quantity <= 0 {
		err = cmd.Store.DeleteProductOrder(&currentProductOrder)
		if err != nil {
			return order, err
		}
		return cmd.Store.FindFullOrderBy(schema.Order{ID: currentProductOrder.OrderID})
	}

	currentProductOrder.Quantity = vStruct.Quantity

	attributesUpdates := schema.ProductOrder{
		Quantity:        currentProductOrder.Quantity,
		TotalPriceCents: util.CalculatesProductOrderPrice(currentProductOrder),
	}

	err = cmd.Store.UpdateProductOrder(&currentProductOrder, attributesUpdates)

	if err != nil {
		return order, err
	}

	return cmd.Store.FindFullOrderBy(schema.Order{ID: currentProductOrder.OrderID})
}

func (cmd CmdProductOrder) ProductsByUser(
	restaurantID int,
	userID string,
) (
	po []schema.ProductOrder,
	err error,
) {
	po, err = cmd.Store.FindProductOrderWithEvaluations(restaurantID, userID)
	return po, err
}
