package main

import (
	"fmt"
	"log"
	"net/http"

	"rest-api-golang/database"
	"rest-api-golang/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Load configurations
	LoadAppConfig()

	// Initialize database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router (new instance of MUX router)
	router := mux.NewRouter().StrictSlash(true)

	// Registers Route Handlers into the MUX router
	routes.RegisterProductRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router)) // Starts the REST API Server
}
