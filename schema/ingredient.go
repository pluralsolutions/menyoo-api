package schema

type Ingredient struct {
	ID              int             `json:"id"`
	IngredientGroup IngredientGroup `json:"-"`
	Name            string          `json:"name"`
	PriceCents      int             `json:"price_cents"`
}
