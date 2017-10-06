package util

import (
	"github.com/plural-solutions/menyoo-api/schema"
)

func CalculatesProductOrderPrice(pd schema.ProductOrder) int {
	var additionalPrice int
	for _, ingredient := range pd.Ingredients {
		additionalPrice += ingredient.PriceCents
	}
	return pd.Quantity * (pd.Product.PriceCents + additionalPrice)
}

func IsSameProductOrder(po1 schema.ProductOrder, po2 schema.ProductOrder) bool {
	return po1.ProductID == po2.ProductID &&
		IsSameIngredients(po1.Ingredients, po2.Ingredients)
}

func IsSameIngredients(ig1 []schema.Ingredient, ig2 []schema.Ingredient) bool {
	if len(ig1) != len(ig2) {
		return false
	}

	for i, ig := range ig1 {
		if ig.ID != ig2[i].ID {
			return false
		}
	}
	return true
}
