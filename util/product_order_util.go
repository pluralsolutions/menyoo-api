package util

import "github.com/lucasgomide/menyoo-api/schema"

func CalculatesProductOrderPrice(pd schema.ProductOrder) int {
	var additionalPrice int
	for _, ingredient := range pd.Ingredients {
		additionalPrice += ingredient.PriceCents
	}
	return (pd.Quantity * pd.Product.PriceCents) + additionalPrice
}
