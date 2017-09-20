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
				image, price_cents, inserted_at
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
