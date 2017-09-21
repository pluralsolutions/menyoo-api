package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/lucasgomide/menyoo-api/cmd"
	"github.com/lucasgomide/menyoo-api/database"
	"github.com/lucasgomide/menyoo-api/handler"
)

func main() {
	log.Println("Starting database connection")
	var store = database.NewStore(
		database.Connect(os.Getenv("DATABASE_URL")),
	)

	router := httprouter.New()

	router.HandlerFunc(
		"GET",
		"/products",
		handler.NewProductsHandler(cmd.NewCmdProduct(store)).Handler,
	)

	router.HandlerFunc(
		"POST",
		"/orders",
		handler.NewOrdersHandler(cmd.NewCmdOrder(store)).Handler,
	)

	log.Println("Starting server..")
	log.Fatal(http.ListenAndServe(":8080", router))
}
