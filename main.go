package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lucasgomide/menyoo-api/cmd"
	"github.com/lucasgomide/menyoo-api/database"
	"github.com/lucasgomide/menyoo-api/handler"
)

func main() {
	log.Println("Starting database connection")
	var store = database.NewStore(
		database.Connect(os.Getenv("DATABASE_URL")),
	)

	router := mux.NewRouter()

	router.HandleFunc("/restaurants/{restaurant_id}/products", handler.NewProductsHandler(cmd.NewCmdProduct(store)).Handler)
	router.HandleFunc("/restaurants/{restaurant_id}/products/{product_id}", handler.NewProductsHandler(cmd.NewCmdProduct(store)).Show)

	// func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	title := vars["title"]
	// 	page := vars["page"]

	// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	// )

	// router.HandlerFunc(
	// 	"GET",
	// 	"/products",
	// 	handler.NewProductsHandler(cmd.NewCmdProduct(store)).Handler,
	// )

	// router.HandlerFunc(
	// 	"POST",
	// 	"/orders",
	// 	handler.NewOrdersHandler(cmd.NewCmdOrder(store)).Handler,
	// )

	log.Println("Starting server..")
	log.Fatal(http.ListenAndServe(":8080", router))
}
