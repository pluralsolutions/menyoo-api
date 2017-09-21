package schema

type Ingredient struct {
	ID                int    `db:"id"`
	IngredientGroupID int    `db:"ingredient_group_id"`
	Name              string `db:"basci"`
	PriceCents        int    `db:"price_cents"`
}
