package schema

type Ingredient struct {
	ID              int
	IngredientGroup IngredientGroup
	Name            string `json:"name"`
	PriceCents      int    `json:"price_cents"`
}
