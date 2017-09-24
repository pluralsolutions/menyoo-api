package schema

import "time"

type Evaluation struct {
	ID        int        `json:"id"`
	UserID    string     `json:"user_id"`
	ProductID int        `json:"product_id"`
	Score     int        `json:"score"`
	UpdatedAt *time.Time `json:"-"`
}
