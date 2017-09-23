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
	err = d.Find(
		&product,
		&schema.Product{ID: productID, RestaurantID: restaurantID},
	).Related(&product.IngredientGroups).Error

	if err != nil {
		return product, err
	}

	for i, ig := range product.IngredientGroups {
		d.Where("ingredient_group_id = ?", &ig.ID).Find(&ig.Ingredients)
		product.IngredientGroups[i].Ingredients = ig.Ingredients
	}

	if err != nil {
		return product, err
	}

	return product, nil
}
