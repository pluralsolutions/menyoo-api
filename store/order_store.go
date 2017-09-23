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

func (d *OrderStore) CreateOrder(order schema.Order) error {
	if err := d.Create(&order).Error; err != nil {
		return err
	}

	return nil
}
