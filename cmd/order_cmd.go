package cmd

import (
	"errors"

	"../schema"
	"../types"
	"../util"
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

		product, err := cmd.Store.FindFullProductBy(
			schema.Product{RestaurantID: order.RestaurantID, ID: pd.ProductID},
		)

		if err != nil {
			continue
		}

		pd.Ingredients = validateSelectedIngredients(product, pd.Ingredients)
		pd.Product = product
		pd.TotalPriceCents = util.CalculatesProductOrderPrice(pd)

		result.Products = append(result.Products, pd)
	}

	if result.Products == nil {
		err = errors.New("No products added")
		return result, err
	}

	selectedOrder, count, _ := cmd.Store.CountOrdersBy(
		schema.Order{
			UserID:       order.UserID,
			RestaurantID: order.RestaurantID,
			Status:       "requested",
		},
	)

	result.UserID = order.UserID
	result.Status = "requested"
	result.RestaurantID = order.RestaurantID

	if count > 0 {
		result.ID = selectedOrder.ID

		for _, product := range selectedOrder.Products {
			for i, resultProduct := range result.Products {
				if util.IsSameProductOrder(product, resultProduct) {
					result.Products[i].ID = product.ID
					result.Products[i].Quantity = product.Quantity + 1

					result.Products[i].TotalPriceCents = util.CalculatesProductOrderPrice(
						schema.ProductOrder{
							Quantity:    result.Products[i].Quantity,
							Product:     product.Product,
							Ingredients: result.Products[i].Ingredients,
						},
					)
				}
			}
		}
		if err := cmd.Store.SaveOrder(&result); err != nil {
			return result, err
		}
	} else {

		if err := cmd.Store.CreateOrder(&result); err != nil {
			return result, err
		}
	}
	result, err = cmd.Store.FindFullOrderBy(schema.Order{ID: result.ID})
	return result, err
}

func (cmd CmdOrder) ShowOrder(order schema.Order) (rOrder schema.Order, err error) {
	rOrder, err = cmd.Store.FindFullOrderBy(order)
	return rOrder, err
}

func (cmd CmdOrder) PlaceOrder(order schema.Order) (rOrder schema.Order, err error) {
	rOrder, err = cmd.Store.FindOrderBy(order, schema.Order{Status: "paid"})
	if err != nil {
		return rOrder, err
	}

	rOrder.Status = "paid"
	err = cmd.Store.SaveOrder(&rOrder)

	return rOrder, err
}

func (cmd CmdOrder) CurrentOrder(
	order schema.Order,
) (
	o schema.Order,
	err error,
) {

	return cmd.Store.CurrentOrder(order)
}

func (cmd CmdOrder) AllOrders(
	order schema.Order,
) (
	o schema.Order,
	err error,
) {
	return cmd.Store.AllOrders(order)
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
