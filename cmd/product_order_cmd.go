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
		return pd, err
	}

	currentProductOrder, err := cmd.Store.FindFullProductOrderBy(
		schema.ProductOrder{OrderID: vStruct.OrderID, ID: vStruct.ProductOrderID},
	)

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
