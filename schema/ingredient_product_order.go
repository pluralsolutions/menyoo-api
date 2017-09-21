package schema

type IngredientProductOrders struct {
	ProducOrdertID int `db:"product_order_id"`
	IngredientID   int `db:"ingredient_id"`
}
