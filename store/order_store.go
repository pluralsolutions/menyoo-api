package store

import (
	"github.com/jinzhu/gorm"
	"github.com/lucasgomide/menyoo-api/schema"
)

type OrderStore struct {
	*gorm.DB
}

func NewOrderStore(db *gorm.DB) *OrderStore {
	return &OrderStore{db}
}

func (d *OrderStore) OrderByRestaurantAndUserAndID(
	restaurantID int,
	userID string,
	orderID int,
) (order schema.Order, err error) {

	err = d.Find(
		&order,
		&schema.Order{UserID: userID, ID: orderID, RestaurantID: restaurantID},
	).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (d *OrderStore) CreateOrder(order *schema.Order) error {
	return d.Create(order).Error
}

func (d *OrderStore) ShowOrder(order schema.Order) (rOrder schema.Order, err error) {
	err = d.
		Preload("Products.Ingredients").
		Find(&rOrder,
			&schema.Order{
				UserID:       order.UserID,
				ID:           order.ID,
				RestaurantID: order.RestaurantID,
			}).Error

	return rOrder, err
}
