package cmd

import (
	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
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
		TotalPriceCents: CalculatesProductOrderPrice(currentProductOrder),
	}

	err = cmd.Store.UpdateProductOrder(&currentProductOrder, attributesUpdates)

	if err != nil {
		return pd, err
	}

	return currentProductOrder, err
}

func CalculatesProductOrderPrice(pd schema.ProductOrder) int {
	var additionalPrice int
	for _, ingredient := range pd.Ingredients {
		additionalPrice += ingredient.PriceCents
	}
	return (pd.Quantity * pd.Product.PriceCents) + additionalPrice
}
