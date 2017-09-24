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

func (d *OrderStore) FindOrderBy(filter schema.Order) (order schema.Order, err error) {
	err = d.Find(
		&order,
		&filter,
	).Error

	return order, err
}

func (d *OrderStore) FindOrderByExcludeBy(
	filter schema.Order,
	exclude schema.Order,
) (order schema.Order, err error) {

	err = d.
		Not(&exclude).
		Find(
			&order,
			&filter,
		).Error

	return order, err
}

func (d *OrderStore) CountOrdersBy(filter schema.Order) (order schema.Order, count int, err error) {
	err = d.
		Preload("Products.Ingredients").
		Preload("Products.Product").
		Where(filter).
		Find(
			&order,
		).Count(&count).
		Error

	return order, count, err
}

func (d *OrderStore) FindFullOrderByExcludeBy(
	filter schema.Order,
	exclude schema.Order,
) (order schema.Order, err error) {

	err = d.
		Preload("Products.Ingredients").
		Not(exclude).
		Find(
			&order,
			&filter,
		).Error

	return order, err
}

func (d *OrderStore) CreateOrder(order *schema.Order) error {
	return d.Create(order).Error
}

func (d *OrderStore) SaveOrder(order *schema.Order) error {
	return d.Save(order).Error
}
