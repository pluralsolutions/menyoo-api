package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"../store"
	"../types"
)

func Connect(url string) *gorm.DB {
	if url == "" {
		url = "postgres://localhost:5432/menyoo?sslmode=disable"
	}
	db, err := gorm.Open("postgres", url)

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)
	return db
}

func NewStore(db *gorm.DB) types.Store {
	return struct {
		*store.ProductStore
		*store.OrderStore
		*store.ProductOrderStore
		*store.EvaluationStore
	}{
		store.NewProductStore(db),
		store.NewOrderStore(db),
		store.NewProductOrderStore(db),
		store.NewEvaluationStore(db),
	}
}
