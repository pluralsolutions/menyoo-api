package schema

import "time"

type Product struct {
	ID               int               `json:"id"`
	RestaurantID     int               `json:"restaurant_id"`
	Title            string            `json:"title"`
	Description      string            `json:"description"`
	Image            string            `json:"image"`
	PriceCents       int               `json:"price_cents"`
	IngredientGroups []IngredientGroup `json:"ingredient_groups"`
	DeletedAt        *time.Time        `json:"-"`
	UpdatedAt        *time.Time        `json:"-"`
}
