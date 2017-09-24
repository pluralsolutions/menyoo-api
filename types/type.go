package types

import (
	"github.com/lucasgomide/menyoo-api/schema"
)

type Store interface {
	ProductStore
	OrderStore
	ProductOrderStore
}

type ProductStore interface {
	ProductsByRestaurant(restaurantID int) ([]schema.Product, error)
	ProductByRestaurantAndID(restaurantID int, productID int) (schema.Product, error)
}

type OrderStore interface {
	CreateOrder(order *schema.Order) error
	OrderByRestaurantAndUserAndID(restaurantID int, userID string, orderID int) (schema.Order, error)
	ShowOrder(order schema.Order) (schema.Order, error)
}

type ProductOrderStore interface {
	UpdateProductOrder(*schema.ProductOrder, schema.ProductOrder) error
	DeleteProductOrder(*schema.ProductOrder) error
	ProductOrderByID(ID int) (schema.ProductOrder, error)
}

type ProductCmd interface {
	ProductsByRestaurant(restaurantID int) ([]schema.Product, error)
	ProductByRestaurantAndID(restaurantID int, productID int) (schema.Product, error)
}

type OrderCmd interface {
	CreateOrder(order schema.Order) (schema.Order, error)
	ShowOrder(order schema.Order) (schema.Order, error)
}

type ProductOrderCmd interface {
	UpdateProductOrderQuantity(params interface{}) (schema.ProductOrder, error)
}
