package store

import (
	"github.com/jinzhu/gorm"
	"github.com/lucasgomide/menyoo-api/schema"
)

type ProductOrderStore struct {
	*gorm.DB
}

func NewProductOrderStore(db *gorm.DB) *ProductOrderStore {
	return &ProductOrderStore{db}
}

func (d *ProductOrderStore) UpdateProductOrder(
	productOrder *schema.ProductOrder,
	updateAttributes schema.ProductOrder,
) error {

	return d.Model(&productOrder).UpdateColumns(updateAttributes).Error
}

func (d *ProductOrderStore) DeleteProductOrder(
	productOrder *schema.ProductOrder,
) error {

	return d.Delete(&productOrder).Error
}

func (d *ProductOrderStore) FindFullProductOrderBy(
	filter schema.ProductOrder,
) (
	pd schema.ProductOrder,
	err error,
) {

	err = d.
		Preload("Product").
		Preload("Ingredients").
		Find(
			&pd,
			&filter,
		).Error

	return pd, err
}
