package schema

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProductEvaluation struct {
	Score    int `json:"score"`
	Quantity int `json:"quantity"`
}

type Product struct {
	ID               int               `json:"id"`
	RestaurantID     int               `json:"restaurant_id"`
	Title            string            `json:"title"`
	Description      string            `json:"description"`
	Image            string            `json:"image"`
	PriceCents       int               `json:"price_cents"`
	IngredientGroups []IngredientGroup `json:"ingredient_groups"`
	DeletedAt        *time.Time        `json:"-"`
	UpdatedAt        *time.Time        `json:"-"`
	Evaluation       ProductEvaluation `json:"evaluation"`
}

func (p *Product) CalculateEvaluation(tx *gorm.DB) (err error) {
	evaluations := []Evaluation{}
	query := tx.
		Joins("JOIN product_orders on product_orders.id = evaluations.product_order_id").
		Joins("JOIN products on products.id = product_orders.product_id").
		Where("products.id = ?", p.ID).
		Find(&evaluations)

	if err = query.Error; err != nil {
		return err
	}
	score := 0
	for _, evaluation := range evaluations {
		score += evaluation.Score
	}
	quantityEvaluations := len(evaluations)
	if quantityEvaluations > 0 {
		p.Evaluation = ProductEvaluation{
			Score:    score / quantityEvaluations,
			Quantity: quantityEvaluations,
		}
	}

	return nil
}
