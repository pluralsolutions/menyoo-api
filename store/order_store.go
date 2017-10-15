package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/plural-solutions/menyoo-api/schema"
)

type OrderStore struct {
	*gorm.DB
}

func NewOrderStore(db *gorm.DB) *OrderStore {
	return &OrderStore{db}
}

func (d *OrderStore) FindOrderBy(
	filter schema.Order,
	exclude ...schema.Order,
) (order schema.Order, err error) {

	fmt.Print("Exclude", exclude)
	err = d.Find(
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

func (d *OrderStore) FindFullOrderBy(
	filter schema.Order,
	exclude ...schema.Order,
) (order schema.Order, err error) {

	err = d.
		Preload("Products.Ingredients").
		Preload("Products.Product").
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

func (d *OrderStore) CurrentOrder(
	filter schema.Order,
) (order schema.Order, err error) {

	result := d.
		Not(schema.Order{Status: "paid"}).
		Preload("Products").
		Preload("Products.Ingredients").
		Preload("Products.Product").
		Find(&order, &filter)

	if result.RecordNotFound() {
		return order, nil
	} else if err = result.Error; err != nil {
		return order, err
	} else {
		return order, nil
	}
}

// AllOrders show all orders, without any details of products
func (d *OrderStore) AllOrders(
	filter schema.Order,
) (order schema.Order, err error) {

	result := d.
		Find(&order, &filter).
		Omit("Products").
		Order("InsertedAt DESC")

	if result.RecordNotFound() {
		return order, nil
	} else if err = result.Error; err != nil {
		return order, err
	} else {
		return order, nil
	}
}
