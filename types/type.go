package types

import (
	"github.com/lucasgomide/menyoo-api/schema"
)

type Store interface {
	ProductStore
	OrderStore
	ProductOrderStore
	EvaluationStore
}

type ProductStore interface {
	FindProductsBy(filter schema.Product) ([]schema.Product, error)
	FindFullProductBy(filter schema.Product) (schema.Product, error)
}

type OrderStore interface {
	SaveOrder(order *schema.Order) error
	CreateOrder(order *schema.Order) error
	FindOrderBy(filter schema.Order, exclude ...schema.Order) (schema.Order, error)
	FindFullOrderBy(filter schema.Order, exclude ...schema.Order) (schema.Order, error)
	CountOrdersBy(filter schema.Order) (schema.Order, int, error)
}

type ProductOrderStore interface {
	UpdateProductOrder(*schema.ProductOrder, schema.ProductOrder) error
	DeleteProductOrder(*schema.ProductOrder) error
	FindFullProductOrderBy(filter schema.ProductOrder) (schema.ProductOrder, error)
	FindProductByUser(productID int, userID string) (schema.ProductOrder, error)
	FindProductOrderWithEvaluations(restaurantID int, userID string) ([]schema.ProductOrder, error)
}

type EvaluationStore interface {
	CreateEvaluation(order *schema.Evaluation) error
}

type ProductCmd interface {
	ProductsByRestaurant(restaurantID int) ([]schema.Product, error)
	ProductByRestaurantAndID(restaurantID int, productID int) (schema.Product, error)
}

type OrderCmd interface {
	CreateOrder(order schema.Order) (schema.Order, error)
	ShowOrder(order schema.Order) (schema.Order, error)
	PlaceOrder(order schema.Order) (schema.Order, error)
}

type ProductOrderCmd interface {
	UpdateProductOrderQuantity(params interface{}) (schema.Order, error)
	ProductsByUser(restaurantID int, userID string) ([]schema.ProductOrder, error)
}

type EvaluationCmd interface {
	CreateEvaluation(evaluation schema.Evaluation) (schema.Evaluation, error)
}
