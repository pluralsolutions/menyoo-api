package schema

type Ingredient struct {
	ID                int    `json:"id"`
	IngredientGroupID int    `json:"ingredient_group_id"`
	Name              string `json:"name"`
	PriceCents        int    `json:"price_cents"`
}
