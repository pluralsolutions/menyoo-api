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
	if err := d.Create(order).Error; err != nil {
		return err
	}

	return nil
}
