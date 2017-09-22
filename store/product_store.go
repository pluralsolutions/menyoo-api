package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/lucasgomide/menyoo-api/schema"
)

type ProductStore struct {
	*sqlx.DB
}

func NewProductStore(db *sqlx.DB) *ProductStore {
	return &ProductStore{db}
}

func (d *ProductStore) ProductsByRestaurant(restaurantID int) (result []schema.Product, err error) {
	err = d.Select(
		&result,
		`
			SELECT
				id, restaurant_id, title, description,
				image, price_cents
			FROM products
			WHERE restaurant_id = $1
		`,
		restaurantID,
	)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *ProductStore) ProductByRestaurantAndID(restaurantID int, productID int) (result []schema.Product, err error) {
	err = d.Select(
		&result,
		`
			SELECT
				p.id, p.restaurant_id, p.title, p.description,
				p.image, p.price_cents,
				ig.title as "ingredient_groups.title",
				ig.product_id as "ingredient_groups.product_id",
				ig.basic as "ingredient_groups.basic"
			FROM products p
			JOIN ingredient_groups ig on ig.product_id = p.id
			WHERE p.restaurant_id = $1 and p.id = $2
		`,
		restaurantID,
		productID,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}
