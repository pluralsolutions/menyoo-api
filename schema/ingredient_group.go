package schema

type IngredientGroup struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	Basic     bool   `db:"basic"`
	ProductID int    `db:"product_id"`
	// Ingredients []Ingredient
}
