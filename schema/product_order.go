package schema

type ProductOrder struct {
	ID                      int
	ProductID               int          `json:"product_id"`
	OrderID                 int          `json:"order_id"`
	Quantity                int          `json:"quantity"`
	TotalPriceCents         int          `json:"total_price_cents"`
	IngredientProductOrders []Ingredient `json:"ingredient_product_orders" gorm:"many2many:ingredient_product_orders;"`
}
