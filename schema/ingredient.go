package schema

import "github.com/jinzhu/gorm"

type Ingredient struct {
	gorm.Model
	IngredientGroupID int    `json:"ingredient_group_id"`
	Name              string `json:"name"`
	PriceCents        int    `json:"price_cents"`
}
