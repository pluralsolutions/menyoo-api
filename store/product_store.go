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

	query := d.Find(
		&products,
		&filter,
	)

	for idx, product := range products {
		product.CalculateEvaluation(d.DB)
		products[idx] = product
	}

	return products, query.Error
}

func (d *ProductStore) FindFullProductBy(
	filter schema.Product,
) (
	product schema.Product,
	err error,
) {
	query := d.
		Preload("IngredientGroups.Ingredients").
		Find(
			&product,
			&filter,
		)

	product.CalculateEvaluation(d.DB)

	return product, query.Error
}
