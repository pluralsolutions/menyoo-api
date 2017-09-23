package cmd

import (
	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
)

type CmdOrder struct {
	types.Store
}

func (cmd CmdOrder) CreateOrder(order schema.Order) (result schema.Order, err error) {
	for _, pd := range order.ProductsOrder {
		if pd.Quantity <= 0 {
			continue
		}
		product, err := cmd.Store.ProductByRestaurantAndID(
			order.RestaurantID,
			pd.ProductID,
		)

		if err != nil {
			continue
		}

		validSelectedIngredients := validateSelectedIngredients(product, pd.Ingredients)
		pd.Ingredients = validSelectedIngredients
		var additionalPrice int
		for _, vsi := range validSelectedIngredients {
			additionalPrice += vsi.PriceCents
		}

		pd.TotalPriceCents = (pd.Quantity * product.PriceCents) + additionalPrice

		result.ProductsOrder = append(result.ProductsOrder, pd)
	}
	result.UserID = order.UserID
	result.Status = "requested"
	result.RestaurantID = order.RestaurantID

	if err := cmd.Store.CreateOrder(&result); err != nil {
		return result, err
	}
	return result, err
}

func validateSelectedIngredients(p schema.Product, selectedIngredients []schema.Ingredient) (ingredients []schema.Ingredient) {
	for _, ig := range p.IngredientGroups {
		for _, i := range ig.Ingredients {
			for _, ipo := range selectedIngredients {
				if ipo.ID == i.ID {
					ingredients = append(ingredients, i)
				}
			}
		}
	}
	return ingredients
}
