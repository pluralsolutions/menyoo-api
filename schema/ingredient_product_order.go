package schema

import "time"

type IngredientProductOrder struct {
	ProductOrderID int        `json:"product_order_id"`
	IngredientID   int        `json:"ingredient_id"`
	DeletedAt      *time.Time `json:"-"`
	UpdatedAt      *time.Time `json:"-"`
}
