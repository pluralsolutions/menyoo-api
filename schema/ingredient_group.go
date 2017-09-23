package schema

type IngredientGroup struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Basic       bool         `json:"basic"`
	ProductID   int          `json:"product_id"`
	Ingredients []Ingredient `json:"ingredients"`
}
