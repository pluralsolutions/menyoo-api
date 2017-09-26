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

func (d *ProductStore) FindProductByUser(
	productID int,
	userID string,
) (productOrder schema.ProductOrder, err error) {

	err = d.
		Joins("JOIN orders ON orders.id = product_orders.order_id").
		Where("product_orders.id = ? AND orders.user_id = ?", productID, userID).
		Find(&productOrder).Error

	return productOrder, err
}

func (d *ProductStore) FindProductOrderWithEvaluations(
	restaurantID int,
	userID string,
) (productOrder []schema.ProductOrder, err error) {

	err = d.
		Preload("Product").
		Preload("Evaluation").
		Joins("JOIN orders ON orders.id = product_orders.order_id").
		Where("orders.user_id = ? AND orders.restaurant_id = ?", userID, restaurantID).
		Find(&productOrder).Error

	return productOrder, err
}
