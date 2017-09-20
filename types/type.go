package types

import (
	"github.com/lucasgomide/menyoo-api/schema"
)

type Store interface {
	ProductStore
}

type ProductStore interface {
	ProductsByRestaurant(restaurantID int) ([]schema.Product, error)
}

type Cmd interface {
	ProductsByRestaurant(restaurantID int) ([]schema.Product, error)
}
