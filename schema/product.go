package schema

import "time"

type Product struct {
	ID               int
	RestaurantID     int               `json:"restaurant_id" db:"restaurant_id"`
	Title            string            `json:"title" db:"title"`
	Description      string            `json:"description" db:"description"`
	Image            string            `json:"image" db:"image"`
	PriceCents       string            `json:"price_cents" db:"price_cents"`
	IngredientGroups []IngredientGroup `json:"ingredient_groups" db:"ingredient_groups"`
	InsertedAt       *time.Time        `json:"inserted_at" db:"inserted_at"`
}
