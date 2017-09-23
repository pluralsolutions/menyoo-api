package schema

type ProductOrder struct {
	ID              int          `json:"id"`
	ProductID       int          `json:"product_id"`
	OrderID         int          `json:"order_id"`
	Quantity        int          `json:"quantity"`
	TotalPriceCents int          `json:"total_price_cents"`
	Ingredients     []Ingredient `json:"ingredients" gorm:"many2many:ingredient_product_orders;"`
}
