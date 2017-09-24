package cmd

import (
	"errors"

	"github.com/lucasgomide/menyoo-api/schema"
	"github.com/lucasgomide/menyoo-api/types"
	"github.com/lucasgomide/menyoo-api/util"
)

type CmdOrder struct {
	types.Store
}

func (cmd CmdOrder) CreateOrder(order schema.Order) (result schema.Order, err error) {
	if order.Products == nil {
		err = errors.New("Params products is required")
		return result, err
	}

	for _, pd := range order.Products {
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

		pd.Ingredients = validateSelectedIngredients(product, pd.Ingredients)

		pd.TotalPriceCents = util.CalculatesProductOrderPrice(pd)

		result.Products = append(result.Products, pd)
	}

	if result.Products == nil {
		err = errors.New("No products added")
		return result, err
	}

	result.UserID = order.UserID
	result.Status = "requested"
	result.RestaurantID = order.RestaurantID

	if err := cmd.Store.CreateOrder(&result); err != nil {
		return result, err
	}
	return result, err
}

func (cmd CmdOrder) ShowOrder(order schema.Order) (rOrder schema.Order, err error) {
	rOrder, err = cmd.Store.FindFullOrderByExcludeBy(order, schema.Order{Status: "paid"})
	return rOrder, err
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
