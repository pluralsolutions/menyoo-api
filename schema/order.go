package schema

type Order struct {
	ID           int            `json:"id"`
	UserID       string         `json:"user_id"`
	RestaurantID int            `json:"restaurant_id"`
	Status       string         `json:"status"`
	Products     []ProductOrder `json:"products"`
}
