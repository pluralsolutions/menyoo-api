package schema

import "time"

type Order struct {
	ID           int            `json:"id"`
	UserID       string         `json:"user_id"`
	RestaurantID int            `json:"restaurant_id"`
	Status       string         `json:"status"`
	Products     []ProductOrder `json:"products"`
	DeletedAt    *time.Time     `json:"-"`
	UpdatedAt    *time.Time     `json:"-"`
}
