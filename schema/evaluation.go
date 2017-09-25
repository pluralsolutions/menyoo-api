package schema

import "time"

type Evaluation struct {
	ID             int        `json:"id"`
	UserID         string     `json:"user_id"`
	ProductOrderID int        `json:"product_order_id"`
	Score          int        `json:"score"`
	UpdatedAt      *time.Time `json:"-"`
}
