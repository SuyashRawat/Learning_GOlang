package main

import (
	"fmt"
	mongodb "g_o/MongoDb/Database"
	routes "g_o/MongoDb/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Crud in MongoDb")
	// Connection to MongoDB
	mongodb.ConnectMongo()
	// Routes setup
	r := mux.NewRouter()
	routes.SetupRoutes(r)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
