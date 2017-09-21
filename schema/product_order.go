package schema

type ProductOrder struct {
	ID                      int `db:"id"`
	ProductID               int `db:"product_id"`
	OrderID                 int `db:"order_id"`
	Quantity                int `db:"quantity"`
	TotalPriceCents         int `db:"total_price_cents"`
	IngredientProductOrders IngredientProductOrders
}
