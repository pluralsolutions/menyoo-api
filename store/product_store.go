package store

import (
	"github.com/jinzhu/gorm"
	"github.com/lucasgomide/menyoo-api/schema"
)

type ProductStore struct {
	*gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{db}
}

func (d *ProductStore) ProductsByRestaurant(restaurantID int) (products []schema.Product, err error) {
	err = d.Find(
		&products,
		&schema.Product{RestaurantID: restaurantID},
	).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (d *ProductStore) ProductByRestaurantAndID(restaurantID int, productID int) (product schema.Product, err error) {
	err = d.
		Preload("IngredientGroups.Ingredients").
		Find(
			&product,
			&schema.Product{ID: productID, RestaurantID: restaurantID},
		).Error

	if err != nil {
		return product, err
	}

	return product, nil
}
