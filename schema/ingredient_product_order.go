package schema

type IngredientProductOrders struct {
	ProducOrdertID int    `db:"product_order_id"`
	IngredientID   string `db:"ingredient_id"`
	InsertedAt     string `db:"inserted_at"`
}
