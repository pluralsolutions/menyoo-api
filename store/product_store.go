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

func (d *ProductStore) FindProductsBy(
	filter schema.Product,
) (
	products []schema.Product,
	err error,
) {

	err = d.Find(
		&products,
		&filter,
	).Error

	return products, err
}

func (d *ProductStore) FindFullProductBy(
	filter schema.Product,
) (
	product schema.Product,
	err error,
) {

	err = d.
		Preload("IngredientGroups.Ingredients").
		Find(
			&product,
			&filter,
		).Error

	return product, err
}
