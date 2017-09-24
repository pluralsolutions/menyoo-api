package schema

import "time"

type IngredientGroup struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Basic       bool         `json:"basic"`
	ProductID   int          `json:"product_id"`
	Ingredients []Ingredient `json:"ingredients" gorm:"ForeignKey:ID"`
	DeletedAt   *time.Time   `json:"-"`
	UpdatedAt   *time.Time   `json:"-"`
}
