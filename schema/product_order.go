package schema

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProductOrder struct {
	ID              int          `json:"id"`
	Product         Product      `json:"product"`
	ProductID       int          `json:"product_id"`
	OrderID         int          `json:"order_id"`
	Quantity        int          `json:"quantity"`
	TotalPriceCents int          `json:"total_price_cents"`
	Ingredients     []Ingredient `json:"ingredients" gorm:"many2many:ingredient_product_orders;"`
	DeletedAt       *time.Time   `json:"-"`
	UpdatedAt       *time.Time   `json:"-"`
}

func (po *ProductOrder) AfterDelete(tx *gorm.DB) (err error) {
	i := IngredientProductOrder{}
	return tx.Where(
		IngredientProductOrder{ProductOrderID: po.ID},
	).Delete(&i).Error
}
