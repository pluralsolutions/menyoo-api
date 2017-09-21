package schema

type ProductOrder struct {
	ID                      int    `db:"id"`
	ProductID               int    `db:"product_id"`
	OrderID                 string `db:"order_id"`
	Quantity                string `db:"quantity"`
	TotalPriceCents         string `db:"total_price_cents"`
	IngredientProductOrders IngredientProductOrders
	InsertedAt              string `db:"inserted_at"`
}
