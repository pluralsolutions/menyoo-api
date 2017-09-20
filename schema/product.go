package schema

// Debit model represent the debit's schema
type Product struct {
	ID           int
	RestaurantID int    `db:"restaurant_id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	Image        string `db:"image"`
	PriceCents   string `db:"price_cents"`
	InsertedAt   string `db:"inserted_at"`
}
