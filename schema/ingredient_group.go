package schema

type IngredientGroup struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Basic       bool   `db:"basci"`
	ProductID   int    `db:"product_id"`
	Ingredients []Ingredient
}
