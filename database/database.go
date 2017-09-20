package database

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lucasgomide/menyoo-api/store"
	"github.com/lucasgomide/menyoo-api/types"
)

func Connect(url string) *sql.DB {
	if url == "" {
		url = "postgres://localhost:5432/menyoo?sslmode=disable"
	}
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("Cannot open database connection with %#v, got error: %s", url, err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database ping has been failed, erro: %s", err)
	}

	return db
}

func NewStore(db *sql.DB) types.Store {
	newDb := sqlx.NewDb(db, "postgres")
	return struct {
		*store.ProductStore
	}{
		store.NewProductStore(newDb),
	}
}
