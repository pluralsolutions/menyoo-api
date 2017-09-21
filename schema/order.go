package schema

type Order struct {
	ID     int
	UserID int    `db:"user_id"`
	Status string `db:"status"`
}
