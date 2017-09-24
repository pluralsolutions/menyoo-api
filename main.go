package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

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
	log.Println("Database connected successfuly")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

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

	router.HandleFunc(
		"/users/me/restaurants/{restaurant_id}/orders/{order_id}/place",
		handler.NewOrdersHandler(cmd.NewCmdOrder(store)).Place,
	).Methods("POST")

	router.HandleFunc(
		"/users/me/restaurants/{restaurant_id}/products/{product_id}/evaluations",
		handler.NewEvaluationsHandler(cmd.NewCmdEvaluation(store)).Create,
	).Methods("POST")

	log.Println("Starting server at port " + port)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(
		handlers.CompressHandler(loggedRouter),
	)))
}
