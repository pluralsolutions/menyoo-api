package schema

import "github.com/jinzhu/gorm"

type IngredientGroup struct {
	gorm.Model

	Title       string       `json:"title"`
	Basic       bool         `json:"basic"`
	ProductID   int          `json:"product_id"`
	Ingredients []Ingredient `json:"ingredients"`
}
