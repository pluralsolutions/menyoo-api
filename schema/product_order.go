package schema

type ProductOrder struct {
	ProductID               int                      `json:"product_id"`
	OrderID                 int                      `json:"order_id"`
	Quantity                int                      `json:"quantity"`
	TotalPriceCents         int                      `json:"total_price_cents"`
	IngredientProductOrders []IngredientProductOrder `json:"ingredient_product_orders"`
}
