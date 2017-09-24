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

	router.HandleFunc(
		"/restaurants/{restaurant_id}/products", handler.NewProductsHandler(cmd.NewCmdProduct(store)).Handler,
	).Methods("GET")

	router.HandleFunc(
		"/restaurants/{restaurant_id}/products/{product_id}", handler.NewProductsHandler(cmd.NewCmdProduct(store)).Show,
	).Methods("GET")

	router.HandleFunc(
		"/orders", handler.NewOrdersHandler(cmd.NewCmdOrder(store)).Create,
	).Methods("POST")

	router.HandleFunc(
		"/restaurants/{restaurant_id}/orders/{order_id}/products/{product_order_id}/quantity",
		handler.NewProductOrdersHandler(
			cmd.NewCmdProductOrder(store),
		).UpdateQuantity,
	).Methods("PUT")

	router.HandleFunc(
		"/users/me/restaurants/{restaurant_id}/orders/{order_id}",
		handler.NewOrdersHandler(cmd.NewCmdOrder(store)).Show,
	).Methods("GET")

	log.Println("Starting server..")
	log.Fatal(http.ListenAndServe(":8080", router))
}
