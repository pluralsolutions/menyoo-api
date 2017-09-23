package schema

type IngredientGroup struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Basic       bool         `json:"basic"`
	ProductID   int          `json:"product_id"`
	Ingredients []Ingredient `json:"ingredients"`
}

// ListIngredientsID bblbla
func (ig IngredientGroup) ListIngredientsID() []int {
	ingredientsID := make([]int, len(ig.Ingredients))

	for _, i := range ig.Ingredients {
		ingredientsID = append(ingredientsID, i.ID)
	}
	return ingredientsID
}
