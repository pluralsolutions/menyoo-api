package schema

type Order struct {
	ID            int
	UserID        string         `json:"user_id"`
	RestaurantID  int            `json:"restaurant_id"`
	Status        string         `json:"status"`
	ProductsOrder []ProductOrder `json:"products_order"`
}
