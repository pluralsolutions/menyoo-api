package schema

type Order struct {
	UserID        string         `json:"user_id"`
	Status        string         `json:"status"`
	ProductsOrder []ProductOrder `json:"products_order"`
}
