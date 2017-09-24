package cmd

import (
	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
	"github.com/lucasgomide/menyoo-api/util"
)

type CmdProductOrder struct {
	types.Store
}

func (cmd CmdProductOrder) UpdateProductOrderQuantity(
	params interface{},
) (
	pd schema.ProductOrder,
	err error,
) {

	vStruct := params.(struct {
		Quantity       int    `json:"quantity"`
		UserID         string `json:"user_id"`
		RestaurantID   int
		OrderID        int
		ProductOrderID int
	})

	_, err = cmd.Store.OrderByRestaurantAndUserAndID(
		vStruct.RestaurantID,
		vStruct.UserID,
		vStruct.OrderID,
	)

	if err != nil {
		return pd, err
	}

	currentProductOrder, err := cmd.Store.ProductOrderByID(vStruct.ProductOrderID)

	if err != nil {
		return pd, err
	}

	if currentProductOrder.Quantity == vStruct.Quantity {
		return currentProductOrder, err
	}

	if vStruct.Quantity <= 0 {
		err = cmd.Store.DeleteProductOrder(&currentProductOrder)
		if err != nil {
			return pd, err
		}
		return pd, err
	}

	currentProductOrder.Quantity = vStruct.Quantity

	attributesUpdates := schema.ProductOrder{
		Quantity:        currentProductOrder.Quantity,
		TotalPriceCents: util.CalculatesProductOrderPrice(currentProductOrder),
	}

	err = cmd.Store.UpdateProductOrder(&currentProductOrder, attributesUpdates)

	if err != nil {
		return pd, err
	}

	return currentProductOrder, err
}
