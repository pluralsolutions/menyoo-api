package schema

import "time"

type Ingredient struct {
	ID              int             `json:"id"`
	IngredientGroup IngredientGroup `json:"-"`
	Name            string          `json:"name"`
	PriceCents      int             `json:"price_cents"`
	DeletedAt       *time.Time      `json:"-"`
	UpdatedAt       *time.Time      `json:"-"`
}
